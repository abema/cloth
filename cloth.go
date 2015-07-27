package cloth

import (
	"strings"
	"unicode"
)

type tagInfo struct {
	Ignore    bool
	Omitempty bool
	RowKey    bool
	Column    string
}

const (
	tagName   = "bigtable"
	delimiter = ":"
)

// ColumnQualifierPrefix is a prefix of column qualifier.
var ColumnQualifierPrefix string

func getTagInfo(tag string) (ti tagInfo) {

	ss := strings.FieldsFunc(tag, func(c rune) bool {
		return c == ',' || unicode.IsSpace(c)
	})

	for i := range ss {
		if ss[i] == "-" {
			ti.Ignore = true
			return
		}
		if ss[i] == "rowkey" && len(ss) == 1 {
			ti.RowKey = true
			continue
		}
		if ss[i] == "omitempty" && len(ss) > 1 {
			ti.Omitempty = true
			continue
		}
		ti.Column = ss[i]
	}

	return
}
