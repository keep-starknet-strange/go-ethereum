package poseidon

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -l:libposeidon.a
#include "poseidon.h"
*/
import "C"

import (
	"unsafe"
)


func Hash(input []byte) ([]byte) {
    output := make([]byte, 64)
    count := C.c_hash_sw2(
	    (*C.uint8_t)(unsafe.Pointer(&input[0])), C.size_t(len(input)),
	    (*C.uint8_t)(unsafe.Pointer(&output[0])), C.size_t(len(output)),
    )
    return output[:count]
}
