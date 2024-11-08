package compton_fragments

import (
	"fmt"
	"github.com/beauxarts/scrinium/litres_integration"
	"strconv"
	"time"
)

func fmtBytes(b int) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

func fmtSeconds(ts int) string {
	if ts == 0 {
		return "unknown"
	}

	hours := ts / (60 * 60)
	minutes := (ts / 60) % 60
	seconds := ts % 60

	if hours == 0 {
		return fmt.Sprintf("%d:%02d", minutes, seconds)
	} else {
		return fmt.Sprintf("%d:%02d:%02d", hours, minutes, seconds)
	}
}

func fmtCurrentPagesOrSeconds(cpos string, at litres_integration.ArtType) string {
	switch at {
	case litres_integration.ArtTypeText:
		fallthrough
	case litres_integration.ArtTypePDF:
		cpos += " стр"
	case litres_integration.ArtTypeAudio:
		if vi, err := strconv.ParseInt(cpos, 10, 32); err == nil {
			cpos = fmtSeconds(int(vi))
		}
	}
	return cpos
}

func fmtYearWrittenAt(dwa string) string {
	yearWrittenAt := 0
	if dateWrittenAt, err := time.Parse("2006-01-02", dwa); err == nil {
		if dateWrittenAt.Month() == 1 && dateWrittenAt.Day() == 1 {
			yearWrittenAt = dateWrittenAt.Year() - 1
		} else {
			yearWrittenAt = dateWrittenAt.Year()
		}
	}
	return strconv.Itoa(yearWrittenAt)
}
