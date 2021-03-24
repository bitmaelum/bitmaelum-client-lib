package main

//#include <stdint.h>
//#include <stdlib.h>
//typedef struct  { void* message; int size; char* error; } BytesReturn;
import "C"
import (
	"unsafe"

	bitmaelumClientBridge "github.com/bitmaelum/bitmaelum-client-lib/bridge"
)

//export BitMaelumClientBridgeCall
func BitMaelumClientBridgeCall(name *C.char, payload unsafe.Pointer, payloadSize C.int) *C.BytesReturn {
	output := (*C.BytesReturn)(C.malloc(C.size_t(C.sizeof_BytesReturn)))
	// we should free resources on dart side

	result, err := bitmaelumClientBridge.Call(C.GoString(name), C.GoBytes(payload, payloadSize))
	if err != nil {
		output.error = C.CString(err.Error())
		return output
	}
	output.error = nil
	output.message = C.CBytes(result)
	output.size = C.int(len(result))
	return output
}

func main() {}
