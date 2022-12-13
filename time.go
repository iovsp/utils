package utils

import (
	"time"
)

var (
	cstzone       *time.Location = time.FixedZone("CST", 0)
	timeForamtStr string         = "2006-01-02 15:04:05"
)

func TimeOnlyTime() string {
	s := time.Now().Format(timeForamtStr)
	return s[11:]
}

func TimeOfString(s string) time.Time {
	tm, _ := time.Parse(timeForamtStr, s)
	return tm
}

func UnixOfString(s string) int64 {
	tm, _ := time.Parse(timeForamtStr, s)
	return tm.Unix()
}

func UnixToString(sec int64) string {
	return time.Unix(sec, 0).In(cstzone).Format(timeForamtStr)
}

func TimeStringAdd(s string, t int) string {
	tm := TimeOfString(s)
	sec := tm.Unix() + int64(t)
	return UnixToString(sec)
}

func TimeStringSub(s string, t int) string {
	tm := TimeOfString(s)
	sec := tm.Unix() - int64(t)
	return UnixToString(sec)
}

func TimeOnly(t *time.Time) string {
	s := t.Format(timeForamtStr)
	return s[11:]
}

func DateOnly(t *time.Time) string {
	s := t.Format(timeForamtStr)
	return s[:10]
}

func DayNext(s string) string {
	nextDay := TimeOfString(s+" 00:00:00").AddDate(0, 0, 1)
	return DateOnly(&nextDay)
}

func TimeString(t *time.Time) string {
	if t == nil {
		return time.Now().Format(timeForamtStr)
	}
	return t.Format(timeForamtStr)
}

func Timeout(tm *time.Time, sec float64) bool {
	return time.Since(*tm).Seconds() > sec
}

