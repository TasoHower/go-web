package goid

import (
	"bytes"
	"runtime"
	"strconv"
)

/*
	go id 提供协程级别的单一协程 Id
*/
func GetID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
