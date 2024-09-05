package tool

import (
	"strconv"
	"unsafe"
)

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StringToBytes(s string) []byte {
	sh := (*[2]uintptr)(unsafe.Pointer(&s))
	bh := [3]uintptr{sh[0], sh[1], sh[1]}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func ConvertInt64SliceToStringSlice(int64Slice []int64) []string {
	stringSlice := make([]string, len(int64Slice))
	for i, id := range int64Slice {
		stringSlice[i] = strconv.FormatInt(id, 10)
	}
	return stringSlice
}
