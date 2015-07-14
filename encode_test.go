package cloth

import (
	"bytes"
	"encoding/binary"
	"reflect"
	"testing"

	"github.com/osamingo/boolconv"
	"google.golang.org/cloud/bigtable"
)

func TestStom(t *testing.T) {

	s := struct {
		TString  string `bigtable:"tstr"`
		TInt     int64  `bigtable:"tint"`
		TBool    bool   `bigtable:"tbool"`
		TOmitStr string `bigtable:"tomitstr, omitempty"`
		TOmitInt int    `bigtable:"tomitint, omitempty"`
		TIgnore  int64  `bigtable:"-"`
	}{
		TString:  "test1",
		TInt:     100,
		TBool:    true,
		TOmitStr: "test2",
		TIgnore:  1000,
	}

	ret, err := Stom("fc", bigtable.Now(), &s)
	if err != nil {
		t.Error("failed to Stom. msg =", err)
	}

	ops := reflect.ValueOf(ret).Elem().FieldByName("ops")

	for i := 0; i < ops.Len(); i++ {

		c := ops.Index(i).Elem().FieldByName("SetCell").Elem().FieldByName("ColumnQualifier").Bytes()
		v := ops.Index(i).Elem().FieldByName("SetCell").Elem().FieldByName("Value").Bytes()

		switch string(c) {

		case "tstr":
			vv := string(v)
			if vv != "test1" {
				t.Errorf("expected %s got %s", "test1", vv)
			}

		case "tint":
			var vv int64
			binary.Read(bytes.NewBuffer(v), binary.BigEndian, &vv)
			if vv != 100 {
				t.Errorf("expected %d got %d", 100, vv)
			}

		case "tbool":
			vv := boolconv.BtoB(v)
			if vv != boolconv.True {
				t.Errorf("expected %v got %v", boolconv.True, vv)
			}

		case "tomitstr":
			vv := string(v)
			if vv != "test2" {
				t.Errorf("expected %s got %s", "test2", vv)
			}

		default:
			t.Error("undefined ccolumn", c)
		}

	}

}
