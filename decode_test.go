package cloth

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/osamingo/boolconv"
	"google.golang.org/cloud/bigtable"
)

func TestReadColumnQualifiers(t *testing.T) {

	ris := []bigtable.ReadItem{
		bigtable.ReadItem{
			Row:    "rowkey",
			Column: "fc:test",
			Value:  []byte("test"),
		},
		bigtable.ReadItem{
			Row:    "rowkey",
			Column: "fc:test2",
			Value:  []byte("test"),
		},
	}

	cqs := ReadColumnQualifier(ris)

	if len(cqs) != 2 {
		t.Error("result length should be 2")
	}

	if cqs[0] == "test2" {
		t.Error("[0] not equal 'test2'")
	}

	if cqs[1] == "test" {
		t.Error("[1] not equal 'test'")
	}
}

func TestReadItemsErrorCase(t *testing.T) {

	s := struct {
		T int `bigtable:"test"`
	}{}

	r := struct {
		R bigtable.ReadItem `bigtable:",rowkey"`
	}{}

	err := ReadItems(nil, nil)
	if err != nil {
		t.Error("error should be nil")
	}

	ris := []bigtable.ReadItem{
		bigtable.ReadItem{
			Row:    "rowkey",
			Column: "fc:test",
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

	err = ReadItems(ris, &r)
	if err == nil {
		t.Error("error is occurred")
	}

}

func TestReadItems(t *testing.T) {

	s := struct {
		TNonTag  string
		TRowKey  string  `bigtable:",rowkey"`
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

	key := "thisisrowkey"
	str := "hoge"
	bl := true
	num := 123
	buf := &bytes.Buffer{}

	ris := []bigtable.ReadItem{
		bigtable.ReadItem{
			Row:    key,
			Column: "fc:tstr",
			Value:  []byte(str),
		},
		bigtable.ReadItem{
			Row:    key,
			Column: "fc:tbool",
			Value:  boolconv.NewBool(bl).Bytes(),
		},
	}

	binary.Write(buf, binary.BigEndian, int64(num))
	ris = append(ris, bigtable.ReadItem{
		Row:    key,
		Column: "fc:tint",
		Value:  buf.Bytes(),
	})

	buf = &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, int8(num))
	ris = append(ris, bigtable.ReadItem{
		Row:    key,
		Column: "fc:tint8",
		Value:  buf.Bytes(),
	})

	buf = &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, int16(num))
	ris = append(ris, bigtable.ReadItem{
		Row:    key,
		Column: "fc:tint16",
		Value:  buf.Bytes(),
	})

	buf = &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, int32(num))
	ris = append(ris, bigtable.ReadItem{
		Row:    key,
		Column: "fc:tint32",
		Value:  buf.Bytes(),
	})

	buf = &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, int64(num))
	ris = append(ris, bigtable.ReadItem{
		Row:    key,
		Column: "fc:tint64",
		Value:  buf.Bytes(),
	})

	buf = &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, uint64(num))
	ris = append(ris, bigtable.ReadItem{
		Row:    key,
		Column: "fc:tuint",
		Value:  buf.Bytes(),
	})

	buf = &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, uint8(num))
	ris = append(ris, bigtable.ReadItem{
		Row:    key,
		Column: "fc:tuint8",
		Value:  buf.Bytes(),
	})

	buf = &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, uint16(num))
	ris = append(ris, bigtable.ReadItem{
		Row:    key,
		Column: "fc:tuint16",
		Value:  buf.Bytes(),
	})

	buf = &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, uint32(num))
	ris = append(ris, bigtable.ReadItem{
		Row:    key,
		Column: "fc:tuint32",
		Value:  buf.Bytes(),
	})

	buf = &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, uint64(num))
	ris = append(ris, bigtable.ReadItem{
		Row:    key,
		Column: "fc:tuint64",
		Value:  buf.Bytes(),
	})

	buf = &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, float32(num))
	ris = append(ris, bigtable.ReadItem{
		Row:    key,
		Column: "fc:tfloat32",
		Value:  buf.Bytes(),
	})

	buf = &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, float64(num))
	ris = append(ris, bigtable.ReadItem{
		Row:    key,
		Column: "fc:tfloat64",
		Value:  buf.Bytes(),
	})

	err := ReadItems(ris, &s)
	if err != nil {
		t.Error("error should not be nil")
	}

	if s.TRowKey != key {
		t.Errorf("expected %s got %s", key, s.TRowKey)
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
