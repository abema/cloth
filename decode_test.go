package cloth

import (
	"testing"

	"bytes"
	"encoding/binary"

	"github.com/osamingo/boolconv"
	"google.golang.org/cloud/bigtable"
)

func TestReadItemsErrorCase(t *testing.T) {

	s := struct {
		T int `bigtable:"test"`
	}{}

	err := ReadItems(nil, nil)
	if err != nil {
		t.Error("error should be nil")
	}

	ris := []*bigtable.ReadItem{
		&bigtable.ReadItem{
			Column: "test",
			Value:  []byte("test"),
		},
	}

	err = ReadItems(ris, struct{}{})
	if err != nil {
		t.Error("error should be nil")
	}

	err = ReadItems(ris, &s)
	if err == nil {
		t.Error("error is occurred")
	}

}

func TestReadItems(t *testing.T) {

	s := struct {
		TNonTag  string
		TString  string  `bigtable:"tstr"`
		TBool    bool    `bigtable:"tbool"`
		TInt     int     `bigtable:"tint"`
		TInt8    int8    `bigtable:"tint8"`
		TInt16   int16   `bigtable:"tint16"`
		TInt32   int32   `bigtable:"tint32"`
		TInt64   int64   `bigtable:"tint64"`
		TUint    uint    `bigtable:"tuint"`
		TUint8   uint8   `bigtable:"tuint8"`
		TUint16  uint16  `bigtable:"tuint16"`
		TUint32  uint32  `bigtable:"tuint32"`
		TUint64  uint64  `bigtable:"tuint64"`
		TFloat32 float32 `bigtable:"tfloat32"`
		TFloat64 float64 `bigtable:"tfloat64"`
	}{}

	str := "hoge"
	bl := true
	num := 123
	buf := &bytes.Buffer{}

	ris := []*bigtable.ReadItem{
		&bigtable.ReadItem{
			Column: "tstr",
			Value:  []byte(str),
		},
		&bigtable.ReadItem{
			Column: "tbool",
			Value:  boolconv.NewBool(bl).Bytes(),
		},
	}

	binary.Write(buf, binary.BigEndian, int64(num))
	ris = append(ris, &bigtable.ReadItem{
		Column: "tint",
		Value:  buf.Bytes(),
	})

	buf = &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, int8(num))
	ris = append(ris, &bigtable.ReadItem{
		Column: "tint8",
		Value:  buf.Bytes(),
	})

	buf = &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, int16(num))
	ris = append(ris, &bigtable.ReadItem{
		Column: "tint16",
		Value:  buf.Bytes(),
	})

	buf = &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, int32(num))
	ris = append(ris, &bigtable.ReadItem{
		Column: "tint32",
		Value:  buf.Bytes(),
	})

	buf = &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, int64(num))
	ris = append(ris, &bigtable.ReadItem{
		Column: "tint64",
		Value:  buf.Bytes(),
	})

	buf = &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, uint64(num))
	ris = append(ris, &bigtable.ReadItem{
		Column: "tuint",
		Value:  buf.Bytes(),
	})

	buf = &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, uint8(num))
	ris = append(ris, &bigtable.ReadItem{
		Column: "tuint8",
		Value:  buf.Bytes(),
	})

	buf = &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, uint16(num))
	ris = append(ris, &bigtable.ReadItem{
		Column: "tuint16",
		Value:  buf.Bytes(),
	})

	buf = &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, uint32(num))
	ris = append(ris, &bigtable.ReadItem{
		Column: "tuint32",
		Value:  buf.Bytes(),
	})

	buf = &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, uint64(num))
	ris = append(ris, &bigtable.ReadItem{
		Column: "tuint64",
		Value:  buf.Bytes(),
	})

	buf = &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, float32(num))
	ris = append(ris, &bigtable.ReadItem{
		Column: "tfloat32",
		Value:  buf.Bytes(),
	})

	buf = &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, float64(num))
	ris = append(ris, &bigtable.ReadItem{
		Column: "tfloat64",
		Value:  buf.Bytes(),
	})

	err := ReadItems(ris, &s)
	if err != nil {
		t.Error("error should not be nil")
	}

	if s.TString != str {
		t.Errorf("expected %s got %s", str, s.TString)
	}

	if !s.TBool {
		t.Errorf("expected %v got %v", bl, s.TBool)
	}

	if s.TInt != int(num) {
		t.Errorf("expected %d got %d", num, s.TInt)
	}

	if s.TInt8 != int8(num) {
		t.Errorf("expected %d got %d", num, s.TInt8)
	}

	if s.TInt16 != int16(num) {
		t.Errorf("expected %d got %d", num, s.TInt16)
	}

	if s.TInt32 != int32(num) {
		t.Errorf("expected %d got %d", num, s.TInt32)
	}

	if s.TInt64 != int64(num) {
		t.Errorf("expected %d got %d", num, s.TInt64)
	}

	if s.TUint != uint(num) {
		t.Errorf("expected %d got %d", num, s.TUint)
	}

	if s.TUint8 != uint8(num) {
		t.Errorf("expected %d got %d", num, s.TUint8)
	}

	if s.TUint16 != uint16(num) {
		t.Errorf("expected %d got %d", num, s.TUint16)
	}

	if s.TUint32 != uint32(num) {
		t.Errorf("expected %d got %d", num, s.TUint32)
	}

	if s.TUint64 != uint64(num) {
		t.Errorf("expected %d got %d", num, s.TUint64)
	}

	if s.TFloat32 != float32(num) {
		t.Errorf("expected %d got %v", num, s.TFloat32)
	}

	if s.TFloat64 != float64(num) {
		t.Errorf("expected %d got %v", num, s.TFloat64)
	}

}
