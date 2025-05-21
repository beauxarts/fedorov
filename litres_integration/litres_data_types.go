package litres_integration

type LitResDataType int

const (
	LitResUnknownData LitResDataType = iota
	LitResHistoryLog
	LitResOperations
)

var litresDataTypeStrings = map[LitResDataType]string{
	LitResUnknownData: "litres-unknown-data",
	LitResHistoryLog:  "litres-history-log",
	LitResOperations:  "litres-operations",
}

func AllLitResDataTypes() []LitResDataType {
	alrdt := make([]LitResDataType, 0, len(litresDataTypeStrings)-1)
	for lrdt := range litresDataTypeStrings {
		if lrdt == LitResUnknownData {
			continue
		}
		alrdt = append(alrdt, lrdt)
	}

	return alrdt
}

func (lrdt LitResDataType) String() string {
	if dts, ok := litresDataTypeStrings[lrdt]; ok {
		return dts
	}
	return ""
}

func ParseLitResDataType(lrdts string) LitResDataType {
	for dt, str := range litresDataTypeStrings {
		if str == lrdts {
			return dt
		}
	}
	return LitResUnknownData
}
