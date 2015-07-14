package cloth

import (
	"testing"

	"bytes"
	"encoding/binary"

	"github.com/osamingo/boolconv"
	"google.golang.org/cloud/bigtable"
)

func TestReadItems(t *testing.T) {

	s := struct {
		TString  string `bigtable:"tstr"`
		TInt     int64  `bigtable:"tint"`
		TBool    bool   `bigtable:"tbool"`
		TOmitStr string `bigtable:"tomitstr, omitempty"`
		TOmitInt int    `bigtable:"tomitint, omitempty"`
		TIgnore  int64  `bigtable:"-"`
	}{}

	str := "hoge"
	num := int64(1000)

	b := &bytes.Buffer{}
	binary.Write(b, binary.BigEndian, num)

	ris := []*bigtable.ReadItem{
		&bigtable.ReadItem{
			Column: "tstr",
			Value:  []byte(str),
		},
		&bigtable.ReadItem{
			Column: "tbool",
			Value:  boolconv.True.Bytes(),
		},
		&bigtable.ReadItem{
			Column: "tint",
			Value:  b.Bytes(),
		},
		&bigtable.ReadItem{
			Column: "nothing",
			Value:  []byte("fuga"),
		},
	}

	err := ReadItems(ris, &s)
	if err != nil {
		t.Error("error should not be nil")
	}

	if s.TString != str {
		t.Errorf("expected %s got %s", str, s.TString)
	}

	if !s.TBool {
		t.Errorf("expected %v got %v", true, s.TBool)
	}

	if s.TInt != num {
		t.Errorf("expected %d got %d", num, s.TInt)
	}

}
