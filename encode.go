package cloth

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"time"

	"github.com/fatih/structs"
	"github.com/osamingo/boolconv"
	"google.golang.org/cloud/bigtable"
)

// GenerateColumnsMutation generates Mutation from Struct.
func GenerateColumnsMutation(family string, t time.Time, i interface{}) (m *bigtable.Mutation, err error) {

	m = bigtable.NewMutation()
	err = SetColumns(family, t, m, i)

	return
}

// GenerateColumnQualifiersMutation generates Mutation from Slice.
func GenerateColumnQualifiersMutation(family string, t time.Time, slice interface{}) (m *bigtable.Mutation, err error) {

	m = bigtable.NewMutation()
	err = SetColumnQualifiers(family, t, m, slice)

	return
}

// SetColumns sets columns of Mutation by Struct.
func SetColumns(family string, t time.Time, m *bigtable.Mutation, i interface{}) (err error) {

	if family == "" {
		err = fmt.Errorf("cloth: family should not be empty")
		return
	}

	if i == nil {
		err = fmt.Errorf("cloth: struct should not be nil")
		return
	}

	fs := structs.New(i).Fields()
	if len(fs) == 0 {
		err = fmt.Errorf("cloth: fields are not found, %v", i)
		return
	}

	for _, f := range fs {

		tg := f.Tag(BigtableTagName)
		if tg == "" {
			continue
		}

		ti := GetBigtableTagInfo(tg)
		if ti.Ignore || ti.Column == "" || ti.Omitempty && f.IsZero() {
			continue
		}

		var b []byte
		b, err = getBytes(f)
		if err != nil {
			m = nil
			break
		}

		m.Set(family, ti.Column, bigtable.Time(t), b)
	}

	return
}

// SetColumnQualifiers sets column qualifiers of Mutation by Slice.
func SetColumnQualifiers(family string, t time.Time, m *bigtable.Mutation, slice interface{}) (err error) {

	if family == "" {
		err = fmt.Errorf("cloth: family should not be empty")
		return
	}

	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		err = fmt.Errorf("cloth: slice should be type slice")
		return
	}

	if s.Len() == 0 {
		err = fmt.Errorf("cloth: slice should not be empty")
		return
	}

	for i := 0; i < s.Len(); i++ {

		fs := structs.New(s.Index(i).Interface()).Fields()
		if len(fs) == 0 {
			err = fmt.Errorf("cloth: fields are not found, %v", i)
			return
		}

		for _, f := range fs {

			tg := f.Tag(BigtableTagName)
			if tg == "" {
				continue
			}

			ti := GetBigtableTagInfo(tg)
			if ti.Qualifier && !f.IsZero() {
				m.Set(family, fmt.Sprintf("%s", f.Value()), bigtable.Time(t), nil)
			}

		}
	}

	return
}

func getBytes(f *structs.Field) ([]byte, error) {

	var b *bytes.Buffer

	switch f.Kind() {

	case reflect.Slice:
		if reflect.ValueOf(f.Value()).Type().Elem().Kind() == reflect.Uint8 {
			// []byte
			return (f.Value()).([]byte), nil
		}

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
