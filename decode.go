package cloth

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"

	"strings"

	"github.com/fatih/structs"
	"github.com/osamingo/boolconv"
	"google.golang.org/cloud/bigtable"
)

// ReadItems converts Mutation into Struct.
func ReadItems(ris []bigtable.ReadItem, s interface{}) (err error) {

	if len(ris) == 0 || s == nil {
		return
	}

	fs := structs.New(s).Fields()
	if len(fs) == 0 {
		return
	}

	for _, ri := range ris {

		for _, f := range fs {

			t := f.Tag(tagName)
			if t == "" {
				continue
			}

			ti := getTagInfo(t)
			if ti.RowKey {
				if err = setValue(f, []byte(ri.Row)); err != nil {
					return
				}

				continue
			}

			cs := strings.Split(ri.Column, delimiter)
			if cs[len(cs)-1] == ti.Column {
				if err = setValue(f, ri.Value); err != nil {
					return
				}

				continue
			}

		}

	}

	return
}

func setValue(f *structs.Field, val []byte) (err error) {

	switch f.Kind() {

	case reflect.String:
		f.Set(string(val))

	case reflect.Bool:
		f.Set(boolconv.BtoB(val).Tob())

	case reflect.Int:
		var n int64
		err = binary.Read(bytes.NewReader(val), binary.BigEndian, &n)
		f.Set(int(n))

	case reflect.Uint:
		var n uint64
		err = binary.Read(bytes.NewReader(val), binary.BigEndian, &n)
		f.Set(uint(n))

	case reflect.Int8:
		var n int8
		err = binary.Read(bytes.NewReader(val), binary.BigEndian, &n)
		f.Set(n)

	case reflect.Uint8:
		var n uint8
		err = binary.Read(bytes.NewReader(val), binary.BigEndian, &n)
		f.Set(n)

	case reflect.Int16:
		var n int16
		err = binary.Read(bytes.NewReader(val), binary.BigEndian, &n)
		f.Set(n)

	case reflect.Uint16:
		var n uint16
		err = binary.Read(bytes.NewReader(val), binary.BigEndian, &n)
		f.Set(n)

	case reflect.Int32:
		var n int32
		err = binary.Read(bytes.NewReader(val), binary.BigEndian, &n)
		f.Set(n)

	case reflect.Uint32:
		var n uint32
		err = binary.Read(bytes.NewReader(val), binary.BigEndian, &n)
		f.Set(n)

	case reflect.Int64:
		var n int64
		err = binary.Read(bytes.NewReader(val), binary.BigEndian, &n)
		f.Set(n)

	case reflect.Uint64:
		var n uint64
		err = binary.Read(bytes.NewReader(val), binary.BigEndian, &n)
		f.Set(n)

	case reflect.Float32:
		var n float32
		err = binary.Read(bytes.NewReader(val), binary.BigEndian, &n)
		f.Set(n)

	case reflect.Float64:
		var n float64
		err = binary.Read(bytes.NewReader(val), binary.BigEndian, &n)
		f.Set(n)

	default:
		err = fmt.Errorf("cloth: unsupported type. %v", f.Kind())

	}

	return
}
