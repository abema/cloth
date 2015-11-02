package cloth

import (
	"strings"
	"unicode"
)

// TagInfo is a field tag information.
type TagInfo struct {
	Ignore    bool
	Omitempty bool
	RowKey    bool
	Qualifier bool
	Column    string
}

const (
	// BigtableTagName is a "bigtable"
	BigtableTagName = "bigtable"
	// ColumnQualifierDelimiter is a ":"
	ColumnQualifierDelimiter = ":"
)

// GetBigtableTagInfo gets TagInfo by a field tag.
func GetBigtableTagInfo(tag string) (ti TagInfo) {

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
		if ss[i] == "qualifier" && len(ss) == 1 {
			ti.Qualifier = true
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
