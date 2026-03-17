package main

/*
#cgo LDFLAGS: -ldl -lm -Wl,-rpath,${SRCDIR}/bin ${SRCDIR}/bin/example.dylib
#include <stdlib.h>

extern int32_t apply_sync(const char* data, int32_t length);
*/
import "C"

import (
	"github.com/iakud/knoll/kds/kdsc/example/kds"
)

func main() {
	all := kds.NewAll(0)
	// FIXME: init datas

	fullData, err := all.Marshal(nil)
	if err != nil {
		panic(err)
	}
	C.apply_sync(C.CString(string(fullData)), C.int32_t(len(fullData)))

	// FIXME: change datas
	dirtyData, err := all.Marshal(nil)
	if err != nil {
		panic(err)
	}
	C.apply_sync(C.CString(string(dirtyData)), C.int32_t(len(dirtyData)))
}
