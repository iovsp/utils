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
