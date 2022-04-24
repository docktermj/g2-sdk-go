/*
Package g2diagnostic ...
*/
package g2diagnostic

/*
#include <stdlib.h>
#include <stdio.h>
#include "libg2diagnostic.h"
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -shared
*/
import "C"
import (
	"context"
	"errors"
	"unsafe"
)

// ----------------------------------------------------------------------------
// Structure
// ----------------------------------------------------------------------------

type G2diagnosticImpl struct{}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// GetAvailableMemory returns the available memory, in bytes, on the host system.
func (g2diagnostic *G2diagnosticImpl) GetAvailableMemory(ctx context.Context) (int64, error) {
	result := C.G2Diagnostic_getAvailableMemory()
	return int64(result), nil
}

// ClearLastException returns the available memory, in bytes, on the host system.
func (g2diagnostic *G2diagnosticImpl) ClearLastException(ctx context.Context) error {
	C.G2Diagnostic_clearLastException()
	return nil
}

// GetLastException returns the last exception encountered in the Xyzzy Engine.
func (g2diagnostic *G2diagnosticImpl) GetLastException(ctx context.Context) (string, error) {
	stringBuffer := make([]byte, 65535)
	C.G2Diagnostic_getLastException((*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
	return string(stringBuffer), nil
}

// GetLogicalCores returns the number of logical cores on the host system.
func (g2diagnostic *G2diagnosticImpl) GetLogicalCores(ctx context.Context) (int, error) {
	result := C.G2Diagnostic_getLogicalCores()
	return int(result), nil
}

// GetPhysicalCores returns the number of physical cores on the host system.
func (g2diagnostic *G2diagnosticImpl) GetPhysicalCores(ctx context.Context) (int, error) {
	result := C.G2Diagnostic_getPhysicalCores()
	return int(result), nil
}

// GetTotalSystemMemory returns the total memory, in bytes, on the host system.
func (g2diagnostic *G2diagnosticImpl) GetTotalSystemMemory(ctx context.Context) (int64, error) {
	result := C.G2Diagnostic_getTotalSystemMemory()
	return int64(result), nil
}

// Init initializes the Xyzzy G2diagnosis.
func (g2diagnostic *G2diagnosticImpl) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {

	// Transform Go datatypes to C datatypes.

	moduleNameForC := C.CString(moduleName)
	defer C.free(unsafe.Pointer(moduleNameForC))

	iniParamsForC := C.CString(iniParams)
	defer C.free(unsafe.Pointer(iniParamsForC))

	// Call Xyzzy.

	result := C.G2Diagnostic_init(moduleNameForC, iniParamsForC, C.int(verboseLogging))

	// Handle result.

	var err error = nil
	if result != 0 {
		err = errors.New("xyzzy-6000nnnn Unable to ...")
	}
	return err
}

// Null shows how to report a BUG inline.
func (g2diagnostic *G2diagnosticImpl) Null(ctx context.Context) (int64, error) {
	// BUG(mjd): Just an example of how to show bugs in GoDoc.
	return 1, nil
}
