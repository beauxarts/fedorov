package data

// TODO: move to scrinium
type DataType int

const (
	Unknown DataType = iota
	LitResMyBooksFresh
	LitResMyBooksDetails
	LiveLibDetails
)

var dataTypeStrings = map[DataType]string{
	LitResMyBooksFresh:   "litres_my_books_fresh",
	LitResMyBooksDetails: "litres_my_books_details",
	LiveLibDetails:       "livelib_details",
}

func (dt DataType) String() string {
	if dts, ok := dataTypeStrings[dt]; ok {
		return dts
	}
	return ""
}

func ParseDataType(dts string) DataType {
	for dt, str := range dataTypeStrings {
		if str == dts {
			return dt
		}
	}
	return Unknown
}
