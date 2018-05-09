package util

import (
	"fmt"
	"strconv"
	"time"
)

var AddDays, _ = time.ParseDuration("-8760h")
var DefaultUntil = strconv.FormatInt(time.Now().Unix(), 10)
var DefaultFrom = strconv.FormatInt(time.Now().Add(AddDays).Unix(), 10)

func GetNow() int64 {
	return time.Now().Unix()
}

func GetTimestampStartOfDay() int64 {
	return time.Now().Unix()/(3600*24)*(3600*24) - 8*3600
}

func GetTimestampStartOfHour() int64 {
	return time.Now().Unix() / 3600 * 3600
}

func GetTimestampStartOfWeek() int64 {
	today := GetTimestampStartOfDay()
	return today - ((int64(time.Now().Weekday())+6)%7)*24*3600
}

func GetLastNDayDateString(n int64) string {
	now := time.Now().Add(-time.Hour * 24 * time.Duration(n))
	return fmt.Sprintf("%4d%02d%02d", now.Year(), now.Month(), now.Day())
}

func GetTimestamp(timeString string, accuracy string) (int64, error) {
	now := GetNow()

	if timeString == "now" {
		return now, nil
	}

	num, err := strconv.Atoi(timeString[:len(timeString)-1])
	if err != nil {
		return now, err
	}

	switch timeString[len(timeString)-1] {
	case 'd':
		// 确保获取的是当日 0 时 的时间
		return GetTimestampStartOfDay() + int64(num)*24*3600 + 24*3600, nil
	case 'h':
		return GetTimestampStartOfHour() + int64(num)*3600 + 3600, nil
	case 'w':
		return GetTimestampStartOfWeek() + int64(num)*7*24*3600 + 7*24*3600, nil

	default:
		if num, err := strconv.Atoi(timeString); err == nil {
			return int64(num), nil
		}
		return now, nil
	}
}
