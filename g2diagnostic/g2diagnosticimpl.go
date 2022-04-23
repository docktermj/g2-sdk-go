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

// Null shows how to report a BUG inline.
func (g2diagnostic *G2diagnosticImpl) Null(ctx context.Context) (int64, error) {
	// BUG(mjd): Just an example of how to show bugs in GoDoc.
	return 1, nil
}
