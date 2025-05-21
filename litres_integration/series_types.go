package litres_integration

type SeriesType int

const (
	SeriesUnknown SeriesType = iota
	SeriesDetails
	SeriesArts
)

var seriesTypesStrings = map[SeriesType]string{
	SeriesUnknown: "series-unknown",
	SeriesDetails: "series-details",
	SeriesArts:    "series-similar",
}

func AllSeriesTypes() []SeriesType {
	ast := make([]SeriesType, 0, len(seriesTypesStrings)-1)
	for st := range seriesTypesStrings {
		if st == SeriesUnknown {
			continue
		}
		ast = append(ast, st)
	}

	return ast
}

func (st SeriesType) String() string {
	if ats, ok := seriesTypesStrings[st]; ok {
		return ats
	}
	return seriesTypesStrings[SeriesUnknown]
}

func ParseSeriesType(str string) SeriesType {
	for st, sts := range seriesTypesStrings {
		if sts == str {
			return st
		}
	}
	return SeriesUnknown
}
