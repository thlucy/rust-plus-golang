package main

// NOTE: There should be NO space between the comments and the `import "C"` line.
// The -ldl is sometimes necessary to fix linker errors about `dlsym`.

/*
#cgo LDFLAGS: ./lib/libhello.a -ldl
#include "./lib/hello.h"
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	for i := 0; ; i++ {
		func() {
			msg := C.CString(fmt.Sprintf("%d", i))
			defer C.free(unsafe.Pointer(msg))
			response := C.echo(msg)
			output := C.GoString(response)
			fmt.Println(output)
			C.dealloc(response)
		}()
	}
}
