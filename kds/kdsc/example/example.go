package example

/*
#cgo LDFLAGS: -ldl -lm -Wl,-rpath,${SRCDIR}/bin ${SRCDIR}/bin/example.dylib
#include <stdlib.h>

extern int32_t apply_sync(const char* data, int32_t length);
extern char* dump(void);
*/
import "C"
import _ "unsafe"

func ApplySync(b []byte) {
	C.apply_sync(C.CString(string(b)), C.int32_t(len(b)))
}

func Dump() string {
	return C.GoString(C.dump())
}
