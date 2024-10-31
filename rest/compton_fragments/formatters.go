package compton_fragments

import "fmt"

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
		return fmt.Sprintf("%2d:%2d", minutes, seconds)
	} else {
		return fmt.Sprintf("%d:%2d:%2d", hours, minutes, seconds)
	}
}
