package cloth

import (
	"bytes"
	"encoding/binary"
	"reflect"

	"github.com/fatih/structs"
	"github.com/osamingo/boolconv"
	"google.golang.org/cloud/bigtable"
)

// ReadItems converts Mutation into Struct.
func ReadItems(ris []*bigtable.ReadItem, s interface{}) (err error) {

	if len(ris) == 0 || s == nil {
		return
	}

	fs := structs.New(s).Fields()
	if len(fs) == 0 {
		return
	}

	for _, ri := range ris {

		for _, f := range fs {

			t := f.Tag(tag)
			if t == "" {
				continue
			}

			if ri.Column == getTagInfo(t).Column {

				switch f.Kind() {

				case reflect.String:
					err = f.Set(string(ri.Value))

				case reflect.Bool:
					err = f.Set(boolconv.BtoB(ri.Value).Tob())

				case reflect.Int:
					var n int64
					binary.Read(bytes.NewReader(ri.Value), binary.BigEndian, &n)
					err = f.Set(int(n))

				case reflect.Uint:
					var n uint64
					binary.Read(bytes.NewReader(ri.Value), binary.BigEndian, &n)
					err = f.Set(uint(n))

				case reflect.Int8:
					var n int8
					binary.Read(bytes.NewReader(ri.Value), binary.BigEndian, &n)
					err = f.Set(n)

				case reflect.Uint8:
					var n uint8
					binary.Read(bytes.NewReader(ri.Value), binary.BigEndian, &n)
					err = f.Set(n)

				case reflect.Int16:
					var n int16
					binary.Read(bytes.NewReader(ri.Value), binary.BigEndian, &n)
					err = f.Set(n)

				case reflect.Uint16:
					var n uint16
					binary.Read(bytes.NewReader(ri.Value), binary.BigEndian, &n)
					err = f.Set(n)

				case reflect.Int32:
					var n int32
					binary.Read(bytes.NewReader(ri.Value), binary.BigEndian, &n)
					err = f.Set(n)

				case reflect.Uint32:
					var n uint32
					binary.Read(bytes.NewReader(ri.Value), binary.BigEndian, &n)
					err = f.Set(n)

				case reflect.Int64:
					var n int64
					binary.Read(bytes.NewReader(ri.Value), binary.BigEndian, &n)
					err = f.Set(n)

				case reflect.Uint64:
					var n uint64
					binary.Read(bytes.NewReader(ri.Value), binary.BigEndian, &n)
					err = f.Set(n)

				case reflect.Float32:
					var n float32
					binary.Read(bytes.NewReader(ri.Value), binary.BigEndian, &n)
					err = f.Set(n)

				case reflect.Float64:
					var n float64
					binary.Read(bytes.NewReader(ri.Value), binary.BigEndian, &n)
					err = f.Set(n)
				}

				if err != nil {
					return
				}

				continue

			}

		}

	}

	return
}
