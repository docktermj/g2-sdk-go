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

void* resizeStringBuffer(void *ptr, size_t size) {
  //deallocate old buffer
  if (ptr != 0)
    free(ptr);
  //allocate new buffer
  void* buffer = malloc(size);
  return buffer;
}

typedef void*(*resize_buffer_type)(void *, size_t);

char* G2Diagnostic_checkDBPerf_local(int secondsToRun) {
  size_t bufferSize = 1;
  char *charBuff = (char *)malloc(1);
  resize_buffer_type resizeFuncPointer = &resizeStringBuffer;
  G2Diagnostic_checkDBPerf(secondsToRun, &charBuff, &bufferSize, resizeFuncPointer);
  return charBuff;
}

char* G2Diagnostic_getDBInfo_local() {
  size_t bufferSize = 1;
  char *charBuff = (char *)malloc(1);
  resize_buffer_type resizeFuncPointer = &resizeStringBuffer;
  G2Diagnostic_getDBInfo(&charBuff, &bufferSize, resizeFuncPointer);
  return charBuff;
}

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
func (g2diagnostic *G2diagnosticImpl) getByteArrayC(size int) *C.char {
	bytes := C.malloc(C.size_t(size))
	return (*C.char)(bytes)
}

func (g2diagnostic *G2diagnosticImpl) getByteArray(size int) []byte {
	return make([]byte, size)
}

// Change the pointer to an array of bytes to a larger array of a given size.
// TODO: not sure this works.
//  FIXME: export resizeStringBuffer
func xxxResizeStringBuffer(stringBuffer unsafe.Pointer, size C.size_t) {
	fmt.Println(">>> Requesting larger buffer of", size, "bytes")

	newByteBuffer := make([]byte, size)
	stringBuffer = unsafe.Pointer(&newByteBuffer)
}

// ----------------------------------------------------------------------------
// Work in progress
// ----------------------------------------------------------------------------

func (g2diagnostic *G2diagnosticImpl) CheckDBPerf(ctx context.Context, secondsToRun int) (string, error) {
	var err error = nil
	stringBuffer := C.GoString(C.G2Diagnostic_checkDBPerf_local(C.int(secondsToRun)))
	return stringBuffer, err
}

// ----------------------------------------------------------------------------
// TODO:Interface methods
// ----------------------------------------------------------------------------

func (g2diagnostic *G2diagnosticImpl) CloseEntityListBySize(ctx context.Context, entityListBySizeHandle int) error {
	//  _DLEXPORT int G2Diagnostic_closeEntityListBySize(EntityListBySizeHandle entityListBySizeHandle);
	return nil
}

func (g2diagnostic *G2diagnosticImpl) Destroy(ctx context.Context) error {
	//  _DLEXPORT int G2Diagnostic_destroy();
	return nil
}

func (g2diagnostic *G2diagnosticImpl) FetchNextEntityBySize(ctx context.Context, entityListBySizeHandle int) (string, error) {
	//  _DLEXPORT int G2Diagnostic_fetchNextEntityBySize(EntityListBySizeHandle entityListBySizeHandle, char *responseBuf, const size_t bufSize);
	return "", nil
}

func (g2diagnostic *G2diagnosticImpl) FindEntitiesByFeatureIDs(ctx context.Context, features string) (string, error) {
	//  _DLEXPORT int G2Diagnostic_findEntitiesByFeatureIDs(const char *features, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	return "", nil
}

func (g2diagnostic *G2diagnosticImpl) GetDataSourceCounts(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getDataSourceCounts(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	return "", nil
}

func (g2diagnostic *G2diagnosticImpl) GetEntityDetails(ctx context.Context, entityID int64, includeInternalFeatures int) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getEntityDetails(const long long entityID, const int includeInternalFeatures, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	return "", nil
}

func (g2diagnostic *G2diagnosticImpl) GetEntityListBySize(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getEntityListBySize(const size_t entitySize,EntityListBySizeHandle* entityListBySizeHandle);
	return "", nil
}

func (g2diagnostic *G2diagnosticImpl) GetEntityResume(ctx context.Context, entityID int64) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getEntityResume(const long long entityID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	return "", nil
}

func (g2diagnostic *G2diagnosticImpl) GetEntitySizeBreakdown(ctx context.Context, minimumEntitySize int, includeInternalFeatures int) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getEntitySizeBreakdown(const size_t minimumEntitySize, const int includeInternalFeatures, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	return "", nil
}

func (g2diagnostic *G2diagnosticImpl) GetFeature(ctx context.Context, libFeatID int64) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getFeature(const long long libFeatID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	return "", nil
}

func (g2diagnostic *G2diagnosticImpl) GetGenericFeatures(ctx context.Context, featureType string, maximumEstimatedCount string) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getGenericFeatures(const char* featureType, const size_t maximumEstimatedCount, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	return "", nil
}

func (g2diagnostic *G2diagnosticImpl) GetLastExceptionCode(ctx context.Context) (int, error) {
	//  _DLEXPORT int G2Diagnostic_getLastException(char *buffer, const size_t bufSize);
	return 0, nil
}

func (g2diagnostic *G2diagnosticImpl) GetMappingStatistics(ctx context.Context, includeInternalFeatures int) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getMappingStatistics(const int includeInternalFeatures, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	return "", nil
}

func (g2diagnostic *G2diagnosticImpl) GetRelationshipDetails(ctx context.Context, relationshipID int64, includeInternalFeatures int) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getRelationshipDetails(const long long relationshipID, const int includeInternalFeatures, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	return "", nil
}

func (g2diagnostic *G2diagnosticImpl) GetResolutionStatistics(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getResolutionStatistics(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	return "", nil
}

func (g2diagnostic *G2diagnosticImpl) InitWithConfigID(ctx context.Context, moduleName string, iniParams string, initConfigID int64, verboseLogging int) error {
	//  _DLEXPORT int G2Diagnostic_initWithConfigID(const char *moduleName, const char *iniParams, const long long initConfigID, const int verboseLogging);
	return nil
}

func (g2diagnostic *G2diagnosticImpl) Reinit(ctx context.Context, initConfigID int64) error {
	//  _DLEXPORT int G2Diagnostic_reinit(const long long initConfigID);
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

// GetDBInfo returns information about the database connection.
func (g2diagnostic *G2diagnosticImpl) GetDBInfo(ctx context.Context) (string, error) {
	var err error = nil
	stringBuffer := C.GoString(C.G2Diagnostic_getDBInfo_local())
	return stringBuffer, err
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
