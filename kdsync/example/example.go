package example

/*
#cgo LDFLAGS: -ldl -lm -Wl,-rpath,${SRCDIR}/bin ${SRCDIR}/bin/example.dylib
#include <stdlib.h>

extern int32_t merge_from(const char* data, int32_t length);
extern char* get_string(void);
*/
import "C"
import _ "unsafe"

func mergeFrom(b []byte) {
	C.merge_from(C.CString(string(b)), C.int32_t(len(b)))
}

func getString() string {
	return C.GoString(C.get_string())
}
