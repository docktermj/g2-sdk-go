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

extern void resizeStringBuffer(void*, size_t);
*/
import "C"
import (
	"context"
	"errors"
	"fmt"
	"unsafe"
)

const initialByteArraySize = 65535

// ----------------------------------------------------------------------------
// Structure
// ----------------------------------------------------------------------------

type G2diagnosticImpl struct{}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

// Get space for an array of bytes of a given size.
func (g2diagnostic *G2diagnosticImpl) getByteArray(size int) []byte {
	return make([]byte, size)
}

// Change the pointer to an array of bytes to a larger array of a given size.
// TODO: not sure this works.
//export resizeStringBuffer
func resizeStringBuffer(stringBuffer unsafe.Pointer, size C.size_t) {
	fmt.Println(">>> Requesting larger buffer of", size, "bytes")

	newByteBuffer := make([]byte, size)
	stringBuffer = unsafe.Pointer(&newByteBuffer)
}

// ----------------------------------------------------------------------------
// Work in progress
// ----------------------------------------------------------------------------

//   _DLEXPORT int G2Diagnostic_getDBInfo(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
func (g2diagnostic *G2diagnosticImpl) GetDBInfo(ctx context.Context) (string, error) {
	var err error = nil
	stringBuffer := g2diagnostic.getByteArray(initialByteArraySize)
	cStringBufferLength := C.ulong(initialByteArraySize)

	// TRIAL: In this version, cStringBufferPointerPointer is pointing to "0x0"
	// Log: NOTE: TRACE: G2Diagnostic_getDBInfo([0x0],[65535],[0x58ffa5])
	// cStringBufferPointer := (*C.char)(unsafe.Pointer(&stringBuffer[0]))
	// cStringBufferPointerPointer := (**C.char)(unsafe.Pointer(cStringBufferPointer))
	// C.G2Diagnostic_getDBInfo(cStringBufferPointerPointer, &cStringBufferLength, (*[0]byte)(C.resizeStringBuffer))

	// TRIAL: In this version the error is:
	// Log: panic: runtime error: cgo argument has Go pointer to Go pointer
	// cStringBufferPointer := (*C.char)(unsafe.Pointer(&stringBuffer[0]))
	// C.G2Diagnostic_getDBInfo(&cStringBufferPointer, &cStringBufferLength, (*[0]byte)(C.resizeStringBuffer))

	// TRIAL:
	// Log: panic: runtime error: cgo argument has Go pointer to Go pointer
	// stringBufferPointer := &stringBuffer
	// stringBufferPointerPointer := &stringBufferPointer
	// cStringBufferPointerPointer := (**C.char)(unsafe.Pointer(stringBufferPointerPointer))
	// C.G2Diagnostic_getDBInfo(cStringBufferPointerPointer, &cStringBufferLength, (*[0]byte)(C.resizeStringBuffer))

	// TRIAL:
	cStringBufferPointer := (*C.char)(&stringBuffer)
	// cStringBufferPointerPointer := (**C.char)(unsafe.Pointer(&cStringBufferPointer))
	C.G2Diagnostic_getDBInfo(&cStringBufferPointer, &cStringBufferLength, (*[0]byte)(C.resizeStringBuffer))

	return string(stringBuffer), err
}

// CheckDBPerf returns the available memory, in bytes, on the host system.
func (g2diagnostic *G2diagnosticImpl) CheckDBPerf(ctx context.Context, secondsToRun int) (string, error) {
	var err error = nil
	stringBuffer := g2diagnostic.getByteArray(initialByteArraySize)
	cSecondsToRun := C.int(secondsToRun)
	cStringBufferLength := C.ulong(initialByteArraySize)
	cStringBufferPointer := (*C.char)(unsafe.Pointer(&stringBuffer[0]))
	cStringBufferPointerPointer := (**C.char)(unsafe.Pointer(cStringBufferPointer))

	//	result := C.G2Diagnostic_checkDBPerf(cSecondsToRun, cStringBufferPointerPointer, &cStringBufferLength, unsafe.Pointer(C.resizeStringBuffer))
	C.G2Diagnostic_checkDBPerf(cSecondsToRun, cStringBufferPointerPointer, &cStringBufferLength, (*[0]byte)(C.resizeStringBuffer))
	return string(stringBuffer), err
}

// ----------------------------------------------------------------------------
// TODO:Interface methods
// ----------------------------------------------------------------------------

func (g2diagnostic *G2diagnosticImpl) CloseEntityListBySize(ctx context.Context, entityListBySizeHandle int) error {
	return nil
}

func (g2diagnostic *G2diagnosticImpl) Destroy(ctx context.Context) error {
	return nil
}

func (g2diagnostic *G2diagnosticImpl) FetchNextEntityBySize(ctx context.Context, entityListBySizeHandle int) (string, error) {
	return "", nil
}

func (g2diagnostic *G2diagnosticImpl) FindEntitiesByFeatureIDs(ctx context.Context, features string) (string, error) {
	return "", nil
}

func (g2diagnostic *G2diagnosticImpl) GetDataSourceCounts(ctx context.Context) (string, error) {
	return "", nil
}

func (g2diagnostic *G2diagnosticImpl) GetEntityDetails(ctx context.Context, entityID int64, includeInternalFeatures int) (string, error) {
	return "", nil
}

func (g2diagnostic *G2diagnosticImpl) GetEntityListBySize(ctx context.Context) (string, error) {
	return "", nil
}

func (g2diagnostic *G2diagnosticImpl) GetEntityResume(ctx context.Context, entityID int64) (string, error) {
	return "", nil
}

func (g2diagnostic *G2diagnosticImpl) GetEntitySizeBreakdown(ctx context.Context, minimumEntitySize int, includeInternalFeatures int) (string, error) {
	return "", nil
}

func (g2diagnostic *G2diagnosticImpl) GetFeature(ctx context.Context, libFeatID int64) (string, error) {
	return "", nil
}

func (g2diagnostic *G2diagnosticImpl) GetGenericFeatures(ctx context.Context, featureType string, maximumEstimatedCount string) (string, error) {
	return "", nil
}

func (g2diagnostic *G2diagnosticImpl) GetLastExceptionCode(ctx context.Context) (int, error) {
	return 0, nil
}

func (g2diagnostic *G2diagnosticImpl) GetMappingStatistics(ctx context.Context, includeInternalFeatures int) (string, error) {
	return "", nil
}

func (g2diagnostic *G2diagnosticImpl) GetRelationshipDetails(ctx context.Context, relationshipID int64, includeInternalFeatures int) (string, error) {
	return "", nil
}

func (g2diagnostic *G2diagnosticImpl) GetResolutionStatistics(ctx context.Context) (string, error) {
	return "", nil
}

func (g2diagnostic *G2diagnosticImpl) InitWithConfigID(ctx context.Context, moduleName string, iniParams string, initConfigID int64, verboseLogging int) error {
	return nil
}

func (g2diagnostic *G2diagnosticImpl) Reinit(ctx context.Context, initConfigID int64) error {
	return nil
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// ClearLastException returns the available memory, in bytes, on the host system.
func (g2diagnostic *G2diagnosticImpl) ClearLastException(ctx context.Context) error {
	var err error = nil
	C.G2Diagnostic_clearLastException()
	return err
}

// GetAvailableMemory returns the available memory, in bytes, on the host system.
func (g2diagnostic *G2diagnosticImpl) GetAvailableMemory(ctx context.Context) (int64, error) {
	var err error = nil
	result := C.G2Diagnostic_getAvailableMemory()
	return int64(result), err
}

// GetLastException returns the last exception encountered in the Xyzzy Engine.
func (g2diagnostic *G2diagnosticImpl) GetLastException(ctx context.Context) (string, error) {
	var err error = nil
	stringBuffer := g2diagnostic.getByteArray(initialByteArraySize)
	C.G2Diagnostic_getLastException((*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
	return string(stringBuffer), err
}

// GetLogicalCores returns the number of logical cores on the host system.
func (g2diagnostic *G2diagnosticImpl) GetLogicalCores(ctx context.Context) (int, error) {
	var err error = nil
	result := C.G2Diagnostic_getLogicalCores()
	return int(result), err
}

// GetPhysicalCores returns the number of physical cores on the host system.
func (g2diagnostic *G2diagnosticImpl) GetPhysicalCores(ctx context.Context) (int, error) {
	var err error = nil
	result := C.G2Diagnostic_getPhysicalCores()
	return int(result), err
}

// GetTotalSystemMemory returns the total memory, in bytes, on the host system.
func (g2diagnostic *G2diagnosticImpl) GetTotalSystemMemory(ctx context.Context) (int64, error) {
	var err error = nil
	result := C.G2Diagnostic_getTotalSystemMemory()
	return int64(result), err
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
