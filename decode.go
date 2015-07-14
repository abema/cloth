package cloth

import (
	"reflect"

	"bytes"
	"encoding/binary"

	"github.com/fatih/structs"
	"github.com/osamingo/boolconv"
	"google.golang.org/cloud/bigtable"
)

// ReadItems converts Mutation into Struct.
func ReadItems(ris []*bigtable.ReadItem, s interface{}) (err error) {

	if len(ris) == 0 {
		return
	}

	fs := structs.New(s).Fields()
	if len(fs) == 0 {
		return
	}

	for _, ri := range ris {

		c := ri.Column

		for _, f := range fs {

			t := f.Tag(tag)
			if t == "" {
				continue
			}

			ti := getTagInfo(t)
			if ti.Column == c {

				switch f.Kind() {

				case reflect.String:

					err = f.Set(string(ri.Value))

				case reflect.Bool:

					err = f.Set(boolconv.BtoB(ri.Value).Tob())

				case reflect.Int8:
					var n int8
					if err = binary.Read(bytes.NewReader(ri.Value), binary.BigEndian, &n); err != nil {
						return
					}
					err = f.Set(n)

				case reflect.Uint8:
					var n uint8
					if err = binary.Read(bytes.NewReader(ri.Value), binary.BigEndian, &n); err != nil {
						return
					}
					err = f.Set(n)

				case reflect.Int16:
					var n int16
					if err = binary.Read(bytes.NewReader(ri.Value), binary.BigEndian, &n); err != nil {
						return
					}
					err = f.Set(n)

				case reflect.Uint16:
					var n uint16
					if err = binary.Read(bytes.NewReader(ri.Value), binary.BigEndian, &n); err != nil {
						return
					}
					err = f.Set(n)

				case reflect.Int32:
					var n int32
					if err = binary.Read(bytes.NewReader(ri.Value), binary.BigEndian, &n); err != nil {
						return
					}
					err = f.Set(n)

				case reflect.Uint32:
					var n uint32
					if err = binary.Read(bytes.NewReader(ri.Value), binary.BigEndian, &n); err != nil {
						return
					}
					err = f.Set(n)

				case reflect.Int64:
					var n int64
					if err = binary.Read(bytes.NewReader(ri.Value), binary.BigEndian, &n); err != nil {
						return
					}
					err = f.Set(n)

				case reflect.Uint64:
					var n uint64
					if err = binary.Read(bytes.NewReader(ri.Value), binary.BigEndian, &n); err != nil {
						return
					}
					err = f.Set(n)

				case reflect.Float32:
					var n float32
					if err = binary.Read(bytes.NewReader(ri.Value), binary.BigEndian, &n); err != nil {
						return
					}
					err = f.Set(n)

				case reflect.Float64:
					var n float64
					if err = binary.Read(bytes.NewReader(ri.Value), binary.BigEndian, &n); err != nil {
						return
					}
					err = f.Set(n)

				case reflect.Int:
					var n int
					if err = binary.Read(bytes.NewReader(ri.Value), binary.BigEndian, &n); err != nil {
						return
					}
					err = f.Set(n)

				case reflect.Uintptr:
					var n uintptr
					if err = binary.Read(bytes.NewReader(ri.Value), binary.BigEndian, &n); err != nil {
						return
					}
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
