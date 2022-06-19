/*
Package g2engine ...
*/
package g2engine

//

/*
#include <stdlib.h>
#include <stdio.h>
#include "libg2.h"
#cgo CFLAGS: -g
#cgo LDFLAGS: -shared

typedef void* EntityListBySizeHandle;
typedef void*(*resize_buffer_type)(void *, size_t);

*/
import "C"
import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"unsafe"

	errormsg "github.com/docktermj/go-json-log-message/message"
)

const initialByteArraySize = 65535

// ----------------------------------------------------------------------------
// Internal methods - names begin with lower case
// ----------------------------------------------------------------------------

// Get space for an array of bytes of a given size.
func (g2engine *G2engineImpl) getByteArrayC(size int) *C.char {
	bytes := C.malloc(C.size_t(size))
	return (*C.char)(bytes)
}

func (g2engine *G2engineImpl) getByteArray(size int) []byte {
	return make([]byte, size)
}

func (g2engine *G2engineImpl) getErrorLevel(errorNumber int) string {

	// Create a map of the different levels. Map will be unsorted.

	errorLevelsMap := map[int]string{
		1000:  "I", // Informational
		2000:  "W", // Warning
		3000:  "E", // Error
		4000:  "D", // Debug
		5000:  "T", // Trace
		9000:  "R", // Reserved
		10000: "F", // Fatal
	}

	// Create a list of sorted keys.

	errorLevelsKeys := make([]int, 0, len(errorLevelsMap))
	for key := range errorLevelsMap {
		errorLevelsKeys = append(errorLevelsKeys, key)
	}
	sort.Ints(errorLevelsKeys)

	// Using the sorted key, find the level.

	for _, errorLevelsKey := range errorLevelsKeys {
		if errorNumber < errorLevelsKey {
			return errorLevelsMap[errorLevelsKey]
		}
	}
	return "" // Unknown
}

func (g2engine *G2engineImpl) getMessageId(errorNumber int) string {
	return fmt.Sprintf(
		"%s%s",
		fmt.Sprintf(MessageIdFormat, errorNumber),
		g2engine.getErrorLevel(errorNumber))
}

func (g2engine *G2engineImpl) getError(ctx context.Context, errorNumber int, details ...string) error {
	lastException, err := g2engine.GetLastException(ctx)
	defer g2engine.ClearLastException(ctx)

	var result error

	if err != nil {
		errorMessage := errormsg.BuildMessage(
			g2engine.getMessageId(errorNumber),
			err.Error(),
			details...,
		)
		result = fmt.Errorf(errorMessage)
	} else {
		errorMessage := errormsg.BuildMessage(
			g2engine.getMessageId(errorNumber),
			lastException,
			details...,
		)
		result = errors.New(errorMessage)
	}
	return result
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// ClearLastException returns the available memory, in bytes, on the host system.
func (g2engine *G2engineImpl) ClearLastException(ctx context.Context) error {
	// _DLEXPORT void G2_clearLastException();
	var err error = nil
	C.G2_clearLastException()
	return err
}

// GetLastException returns the last exception encountered in the Xyzzy Engine.
func (g2engine *G2engineImpl) GetLastException(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2_getLastException(char *buffer, const size_t bufSize);
	var err error = nil
	stringBuffer := g2engine.getByteArray(initialByteArraySize)
	C.G2_getLastException((*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
	stringBuffer = bytes.Trim(stringBuffer, "\x00")
	if len(stringBuffer) == 0 {
		errorMessage := errormsg.BuildMessage(
			g2engine.getMessageId(2999),
			"Cannot retrieve last error message",
		)
		err = errors.New(errorMessage)
	}
	return string(stringBuffer), err
}

// Init initializes the Xyzzy G2diagnosis.
func (g2engine *G2engineImpl) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	// _DLEXPORT int G2_init(const char *moduleName, const char *iniParams, const int verboseLogging);
	var err error = nil

	// Transform Go datatypes to C datatypes.

	moduleNameForC := C.CString(moduleName)
	defer C.free(unsafe.Pointer(moduleNameForC))

	iniParamsForC := C.CString(iniParams)
	defer C.free(unsafe.Pointer(iniParamsForC))

	// Call Xyzzy.

	result := C.G2_init(moduleNameForC, iniParamsForC, C.int(verboseLogging))

	// Handle result.

	if result != 0 {
		err = g2engine.getError(ctx, 18, moduleName, iniParams, strconv.Itoa(verboseLogging))
	}
	return err
}
