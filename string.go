package utils

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

// String2Bytes hacks string to []byte
func String2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// StringOfBytes hacks []byte to string
func StringOfBytes(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StringToFloat64(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
}

func StringToInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func StringSplitInts(s, sep string) (v []int) {
	ss := strings.Split(s, sep)
	for k := range ss {
		i, _ := strconv.Atoi(ss[k])
		v = append(v, i)
	}
	return
}

func ToString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func Contains(array []int, val int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			return true
		}
	}
	return false
}
