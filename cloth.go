package cloth

type tagInfo struct {
	Ignore    bool
	Omitempty bool
	Column    string
}

const (
	tag       = "bigtable"
	delimiter = ','
	ignore    = "-"
	omitempty = "omitempty"
)
