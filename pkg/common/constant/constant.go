package constant

const (
	HOUR  int64 = 3600
	DAY   int64 = 24 * 3600
	WEEK  int64 = 7 * 24 * 3600
	MONTH int64 = 30 * 24 * 3600
)

func GetAccuracyStep(accuracy string) int64 {
	switch accuracy {
	case "Hour":
		return HOUR
	case "Day":
		return DAY
	case "Week":
		return WEEK
	case "Month":
		return MONTH
	default:
		return DAY
	}
}
