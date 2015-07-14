package cloth

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"github.com/fatih/structs"
	"github.com/osamingo/boolconv"
	"google.golang.org/cloud/bigtable"
)

// Stom converts Struct into Mutation.
func Stom(family string, ts bigtable.Timestamp, i interface{}) (m *bigtable.Mutation, err error) {
	m = bigtable.NewMutation()
	err = SetColumns(family, ts, i, m)
	return
}

// SetColumns sets columns of Mutation by struct.
func SetColumns(family string, ts bigtable.Timestamp, i interface{}, m *bigtable.Mutation) (err error) {

	fs := structs.New(i).Fields()
	if len(fs) == 0 {
		err = fmt.Errorf("cloth: fields are not found, %v", i)
		return
	}

	for _, f := range fs {

		t := f.Tag(tag)
		if t == "" {
			continue
		}

		ti := getTagInfo(t)
		if ti.Ignore || ti.Column == "" || ti.Omitempty && f.IsZero() {
			continue
		}

		var b []byte
		b, err = getBytes(f)
		if err != nil {
			m = nil
			break
		}

		m.Set(family, ti.Column, ts, b)
	}

	return
}

func getTagInfo(tag string) (ti tagInfo) {

	ss := strings.FieldsFunc(tag, func(c rune) bool {
		return c == delimiter || unicode.IsSpace(c)
	})

	for _, s := range ss {
		if s == ignore {
			ti.Ignore = true
			return
		}
		if s == omitempty {
			ti.Omitempty = true
			continue
		}
		ti.Column = s
	}

	return
}

func getBytes(f *structs.Field) ([]byte, error) {

	var b *bytes.Buffer

	switch f.Kind() {

	case reflect.String:
		return []byte((f.Value()).(string)), nil

	case reflect.Bool:
		return boolconv.NewBool((f.Value()).(bool)).Bytes(), nil

	case reflect.Int8, reflect.Uint8:
		b = bytes.NewBuffer(make([]byte, 0, 1))

	case reflect.Int16, reflect.Uint16:
		b = bytes.NewBuffer(make([]byte, 0, 2))

	case reflect.Int32, reflect.Uint32, reflect.Float32:
		b = bytes.NewBuffer(make([]byte, 0, 4))

	case reflect.Int64, reflect.Uint64, reflect.Float64, reflect.Int, reflect.Uintptr:
		b = bytes.NewBuffer(make([]byte, 0, 8))

	}

	if b != nil {
		err := binary.Write(b, binary.BigEndian, f.Value())
		return b.Bytes(), err
	}

	return nil, fmt.Errorf("cloth: unsupported type. %v", f.Kind())
}
