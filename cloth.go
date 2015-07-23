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

const tagName = "bigtable"

func getTagInfo(tag string) (ti tagInfo) {

	ss := strings.FieldsFunc(tag, func(c rune) bool {
		return c == ',' || unicode.IsSpace(c)
	})

	for _, s := range ss {
		if s == "-" {
			ti.Ignore = true
			return
		}
		if s == "omitempty" && len(ss) > 1 {
			ti.Omitempty = true
			continue
		}
		if s == "rowkey" && len(ss) == 1 {
			ti.RowKey = true
			continue
		}
		ti.Column = s
	}

	return
}
