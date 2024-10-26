package main

/*
#include <sqlite3ext.h>
#include <stdint.h>
#include "sqlite_trace_callback.h"

*/
import "C"
import (
	"context"
	"time"
	"unsafe"
)

var collector *TraceCollector

func init() {
	collector = NewTraceCollector(context.Background())
}

//export sqlite3_extension_init
func sqlite3_extension_init(db *C.sqlite3, pzErrMsg **C.char, pApi *C.sqlite3_api_routines) C.int {
	return sqlite3_trace_init(db, pzErrMsg, pApi)
}

//export sqlite3_trace_init
func sqlite3_trace_init(db *C.sqlite3, _ **C.char, pApi *C.sqlite3_api_routines) C.int {
	C._SQLITE_EXTENSION_INIT2(pApi)
	C._sqlite3_trace_v2(db, C.SQLITE_TRACE_PROFILE, C.closure(C.go_trace_v2_callback), nil)
	return C.SQLITE_OK
}

//export go_trace_v2_callback
func go_trace_v2_callback(t C.uint, _ unsafe.Pointer, p unsafe.Pointer, x unsafe.Pointer) C.int {
	switch t {
	case C.SQLITE_TRACE_PROFILE:
		stmt := C.GoString(C._sqlite3_sql((*C.sqlite3_stmt)(p)))
		nanoseconds := *(*C.longlong)(x)
		collector.SaveProbe(stmt, time.Duration(nanoseconds))
	}
	return C.SQLITE_OK
}

func main() {}
