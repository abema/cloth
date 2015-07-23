package cloth

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"

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

	if family == "" {
		err = fmt.Errorf("cloth: family is empty")
		return
	}

	if i == nil {
		err = fmt.Errorf("cloth: struct is nil")
		return
	}

	fs := structs.New(i).Fields()
	if len(fs) == 0 {
		err = fmt.Errorf("cloth: fields are not found, %v", i)
		return
	}

	for _, f := range fs {

		t := f.Tag(tagName)
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

func getBytes(f *structs.Field) ([]byte, error) {

	var b *bytes.Buffer

	switch f.Kind() {

	case reflect.String:
		return []byte((f.Value()).(string)), nil

	case reflect.Bool:
		return boolconv.NewBool((f.Value()).(bool)).Bytes(), nil

	case reflect.Int8, reflect.Uint8:
		b = bytes.NewBuffer(make([]byte, 0, 2))

	case reflect.Int16, reflect.Uint16:
		b = bytes.NewBuffer(make([]byte, 0, binary.MaxVarintLen16))

	case reflect.Int32, reflect.Uint32:
		b = bytes.NewBuffer(make([]byte, 0, binary.MaxVarintLen32))

	case reflect.Int64, reflect.Uint64, reflect.Int, reflect.Uint, reflect.Float32, reflect.Float64:
		b = bytes.NewBuffer(make([]byte, 0, binary.MaxVarintLen64))

	}

	if b != nil {

		i := f.Value()
		if f.Kind() == reflect.Int {
			i = int64(i.(int))
		}
		if f.Kind() == reflect.Uint {
			i = uint64(i.(uint))
		}

		err := binary.Write(b, binary.BigEndian, i)
		return b.Bytes(), err
	}

	return nil, fmt.Errorf("cloth: unsupported type. %v", f.Kind())
}
