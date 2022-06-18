/*
Package g2diagnostic ...
*/
package g2diagnostic

//

/*
#include <stdlib.h>
#include <stdio.h>
#include "libg2diagnostic.h"
#cgo CFLAGS: -g
#cgo LDFLAGS: -shared

typedef void* EntityListBySizeHandle;
typedef void*(*resize_buffer_type)(void *, size_t);

void* resizeStringBuffer(void *ptr, size_t size) {
//deallocate old buffer
if (ptr != 0)
    free(ptr);
//allocate new buffer
void* buffer = malloc(size);
return buffer;
}

char* G2Diagnostic_checkDBPerf_local(int secondsToRun) {
size_t bufferSize = 1;
char *charBuff = (char *)malloc(1);
resize_buffer_type resizeFuncPointer = &resizeStringBuffer;
int returnCode = G2Diagnostic_checkDBPerf(secondsToRun, &charBuff, &bufferSize, resizeFuncPointer);
if (returnCode != 0) {
    return "";
}
return charBuff;
}

char* G2Diagnostic_findEntitiesByFeatureIDs_local(const char *features) {
size_t bufferSize = 1;
char *charBuff = (char *)malloc(1);
resize_buffer_type resizeFuncPointer = &resizeStringBuffer;
int returnCode = G2Diagnostic_findEntitiesByFeatureIDs(features, &charBuff, &bufferSize, resizeFuncPointer);
if (returnCode != 0) {
    return "";
}
return charBuff;
}

char* G2Diagnostic_getDataSourceCounts_local() {
size_t bufferSize = 1;
char *charBuff = (char *)malloc(1);
resize_buffer_type resizeFuncPointer = &resizeStringBuffer;
int returnCode = G2Diagnostic_getDataSourceCounts(&charBuff, &bufferSize, resizeFuncPointer);
if (returnCode != 0) {
    return "";
}
return charBuff;
}

char* G2Diagnostic_getDBInfo_local() {
size_t bufferSize = 1;
char *charBuff = (char *)malloc(1);
resize_buffer_type resizeFuncPointer = &resizeStringBuffer;
int returnCode = G2Diagnostic_getDBInfo(&charBuff, &bufferSize, resizeFuncPointer);
if (returnCode != 0) {
    return "";
}
return charBuff;
}

char* G2Diagnostic_getEntityDetails_local(const long long entityID, const int includeInternalFeatures) {
size_t bufferSize = 1;
char *charBuff = (char *)malloc(1);
resize_buffer_type resizeFuncPointer = &resizeStringBuffer;
int returnCode = G2Diagnostic_getEntityDetails(entityID, includeInternalFeatures, &charBuff, &bufferSize, resizeFuncPointer);
if (returnCode != 0) {
    return "";
}
return charBuff;
}

char* G2Diagnostic_getEntityResume_local(const long long entityID) {
size_t bufferSize = 1;
char *charBuff = (char *)malloc(1);
resize_buffer_type resizeFuncPointer = &resizeStringBuffer;
int returnCode = G2Diagnostic_getEntityResume(entityID, &charBuff, &bufferSize, resizeFuncPointer);
if (returnCode != 0) {
    return "";
}
return charBuff;
}

char* G2Diagnostic_getEntitySizeBreakdown_local(const size_t minimumEntitySize, const int includeInternalFeatures) {
size_t bufferSize = 1;
char *charBuff = (char *)malloc(1);
resize_buffer_type resizeFuncPointer = &resizeStringBuffer;
int returnCode = G2Diagnostic_getEntitySizeBreakdown(minimumEntitySize, includeInternalFeatures, &charBuff, &bufferSize, resizeFuncPointer);
if (returnCode != 0) {
    return "";
}
return charBuff;
}

char* G2Diagnostic_getFeature_local(const long long libFeatID) {
size_t bufferSize = 1;
char *charBuff = (char *)malloc(1);
resize_buffer_type resizeFuncPointer = &resizeStringBuffer;
int returnCode = G2Diagnostic_getFeature(libFeatID, &charBuff, &bufferSize, resizeFuncPointer);
if (returnCode != 0) {
    return "";
}
return charBuff;
}

char* G2Diagnostic_getGenericFeatures_local(const char *featureType, const size_t maximumEstimatedCount) {
size_t bufferSize = 1;
char *charBuff = (char *)malloc(1);
resize_buffer_type resizeFuncPointer = &resizeStringBuffer;
int returnCode = G2Diagnostic_getGenericFeatures(featureType, maximumEstimatedCount, &charBuff, &bufferSize, resizeFuncPointer);
if (returnCode != 0) {
    return "";
}
return charBuff;
}

char* G2Diagnostic_getMappingStatistics_local(const int includeInternalFeatures) {
size_t bufferSize = 1;
char *charBuff = (char *)malloc(1);
resize_buffer_type resizeFuncPointer = &resizeStringBuffer;
int returnCode = G2Diagnostic_getMappingStatistics(includeInternalFeatures, &charBuff, &bufferSize, resizeFuncPointer);
if (returnCode != 0) {
    return "";
}
return charBuff;
}

char* G2Diagnostic_getRelationshipDetails_local(const long long relationshipID, const int includeInternalFeatures) {
size_t bufferSize = 1;
char *charBuff = (char *)malloc(1);
resize_buffer_type resizeFuncPointer = &resizeStringBuffer;
int returnCode = G2Diagnostic_getRelationshipDetails(relationshipID, includeInternalFeatures, &charBuff, &bufferSize, resizeFuncPointer);
if (returnCode != 0) {
    return "";
}
return charBuff;
}

char* G2Diagnostic_getResolutionStatistics_local() {
size_t bufferSize = 1;
char *charBuff = (char *)malloc(1);
resize_buffer_type resizeFuncPointer = &resizeStringBuffer;
int returnCode = G2Diagnostic_getResolutionStatistics(&charBuff, &bufferSize, resizeFuncPointer);
if (returnCode != 0) {
    return "";
}
return charBuff;
}

*/
import "C"
import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"unsafe"
)

const initialByteArraySize = 65535

// ----------------------------------------------------------------------------
// Internal methods - names begin with lower case
// ----------------------------------------------------------------------------

// Get space for an array of bytes of a given size.
func (g2diagnostic *G2diagnosticImpl) getByteArrayC(size int) *C.char {
	bytes := C.malloc(C.size_t(size))
	return (*C.char)(bytes)
}

func (g2diagnostic *G2diagnosticImpl) getByteArray(size int) []byte {
	return make([]byte, size)
}

func (g2diagnostic *G2diagnosticImpl) getError(ctx context.Context, errorNumber int, details ...string) error {
	lastException, err := g2diagnostic.GetLastException(ctx)
	defer g2diagnostic.ClearLastException(ctx)
	var result error = nil
	if err != nil {
		result = fmt.Errorf("xyzzy-60119999: %w", err)
	} else {
		result = errors.New(fmt.Sprintf("xyzzy-6011%04d %s - %v", errorNumber, lastException, details))
	}
	return result
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// CheckDBPerf performs inserts to determine rate of insertion.
func (g2diagnostic *G2diagnosticImpl) CheckDBPerf(ctx context.Context, secondsToRun int) (string, error) {
	// _DLEXPORT int G2Diagnostic_checkDBPerf(int secondsToRun, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	stringBuffer := C.GoString(C.G2Diagnostic_checkDBPerf_local(C.int(secondsToRun)))
	if len(stringBuffer) == 0 {
		err = g2diagnostic.getError(ctx, 1, strconv.Itoa(secondsToRun))
	}
	return stringBuffer, err
}

// ClearLastException returns the available memory, in bytes, on the host system.
func (g2diagnostic *G2diagnosticImpl) ClearLastException(ctx context.Context) error {
	// _DLEXPORT void G2Diagnostic_clearLastException();
	var err error = nil
	C.G2Diagnostic_clearLastException()
	return err
}

func (g2diagnostic *G2diagnosticImpl) CloseEntityListBySize(ctx context.Context, entityListBySizeHandle interface{}) error {
	//  _DLEXPORT int G2Diagnostic_closeEntityListBySize(EntityListBySizeHandle entityListBySizeHandle);
	var err error = nil
	result := C.G2Diagnostic_closeEntityListBySize(C.EntityListBySizeHandle(&entityListBySizeHandle))
	if result != 0 {
		err = g2diagnostic.getError(ctx, 2)
	}
	return err
}

func (g2diagnostic *G2diagnosticImpl) Destroy(ctx context.Context) error {
	//  _DLEXPORT int G2Diagnostic_destroy();
	var err error = nil
	result := C.G2Diagnostic_destroy()
	if result != 0 {
		err = g2diagnostic.getError(ctx, 3)
	}
	return err
}

func (g2diagnostic *G2diagnosticImpl) FetchNextEntityBySize(ctx context.Context, entityListBySizeHandle interface{}) (string, error) {
	//  _DLEXPORT int G2Diagnostic_fetchNextEntityBySize(EntityListBySizeHandle entityListBySizeHandle, char *responseBuf, const size_t bufSize);
	var err error = nil
	stringBuffer := g2diagnostic.getByteArray(initialByteArraySize)
	result := C.G2Diagnostic_fetchNextEntityBySize(C.EntityListBySizeHandle(&entityListBySizeHandle), (*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
	if result != 0 {
		err = g2diagnostic.getError(ctx, 4)
	}
	return string(stringBuffer), err
}

func (g2diagnostic *G2diagnosticImpl) FindEntitiesByFeatureIDs(ctx context.Context, features string) (string, error) {
	//  _DLEXPORT int G2Diagnostic_findEntitiesByFeatureIDs(const char *features, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	featuresForC := C.CString(features)
	defer C.free(unsafe.Pointer(featuresForC))
	stringBuffer := C.GoString(C.G2Diagnostic_findEntitiesByFeatureIDs_local(featuresForC))
	if len(stringBuffer) == 0 {
		err = g2diagnostic.getError(ctx, 5, features)
	}
	return stringBuffer, err
}

// GetAvailableMemory returns the available memory, in bytes, on the host system.
func (g2diagnostic *G2diagnosticImpl) GetAvailableMemory(ctx context.Context) (int64, error) {
	// _DLEXPORT long long G2Diagnostic_getAvailableMemory();
	var err error = nil
	result := C.G2Diagnostic_getAvailableMemory()
	return int64(result), err
}

func (g2diagnostic *G2diagnosticImpl) GetDataSourceCounts(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getDataSourceCounts(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	stringBuffer := C.GoString(C.G2Diagnostic_getDataSourceCounts_local())
	if len(stringBuffer) == 0 {
		err = g2diagnostic.getError(ctx, 6)
	}
	return stringBuffer, err
}

// GetDBInfo returns information about the database connection.
func (g2diagnostic *G2diagnosticImpl) GetDBInfo(ctx context.Context) (string, error) {
	// _DLEXPORT int G2Diagnostic_getDBInfo(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	stringBuffer := C.GoString(C.G2Diagnostic_getDBInfo_local())
	if len(stringBuffer) == 0 {
		err = g2diagnostic.getError(ctx, 7)
	}
	return stringBuffer, err
}

func (g2diagnostic *G2diagnosticImpl) GetEntityDetails(ctx context.Context, entityID int64, includeInternalFeatures int) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getEntityDetails(const long long entityID, const int includeInternalFeatures, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	stringBuffer := C.GoString(C.G2Diagnostic_getEntityDetails_local(C.longlong(entityID), C.int(includeInternalFeatures)))
	if len(stringBuffer) == 0 {
		err = g2diagnostic.getError(ctx, 8, strconv.FormatInt(entityID, 10), strconv.Itoa(includeInternalFeatures))
	}
	return stringBuffer, err
}

func (g2diagnostic *G2diagnosticImpl) GetEntityListBySize(ctx context.Context, entitySize int) (interface{}, error) {
	//  _DLEXPORT int G2Diagnostic_getEntityListBySize(const size_t entitySize, EntityListBySizeHandle* entityListBySizeHandle);
	var err error = nil
	var entityListBySizeHandle unsafe.Pointer
	result := C.G2Diagnostic_getEntityListBySize(C.size_t(entitySize), (*C.EntityListBySizeHandle)(&entityListBySizeHandle))
	if result != 0 {
		err = g2diagnostic.getError(ctx, 9, strconv.Itoa(entitySize))
	}
	return entityListBySizeHandle, err
}

func (g2diagnostic *G2diagnosticImpl) GetEntityResume(ctx context.Context, entityID int64) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getEntityResume(const long long entityID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	stringBuffer := C.GoString(C.G2Diagnostic_getEntityResume_local(C.longlong(entityID)))
	if len(stringBuffer) == 0 {
		err = g2diagnostic.getError(ctx, 10, strconv.FormatInt(entityID, 10))
	}
	return stringBuffer, err
}

func (g2diagnostic *G2diagnosticImpl) GetEntitySizeBreakdown(ctx context.Context, minimumEntitySize int, includeInternalFeatures int) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getEntitySizeBreakdown(const size_t minimumEntitySize, const int includeInternalFeatures, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	stringBuffer := C.GoString(C.G2Diagnostic_getEntitySizeBreakdown_local(C.size_t(minimumEntitySize), C.int(includeInternalFeatures)))
	if len(stringBuffer) == 0 {
		err = g2diagnostic.getError(ctx, 11, strconv.Itoa(minimumEntitySize), strconv.Itoa(includeInternalFeatures))
	}
	return stringBuffer, err
}

func (g2diagnostic *G2diagnosticImpl) GetFeature(ctx context.Context, libFeatID int64) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getFeature(const long long libFeatID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	stringBuffer := C.GoString(C.G2Diagnostic_getFeature_local(C.longlong(libFeatID)))
	if len(stringBuffer) == 0 {
		err = g2diagnostic.getError(ctx, 12, strconv.FormatInt(libFeatID, 10))
	}
	return stringBuffer, err
}

func (g2diagnostic *G2diagnosticImpl) GetGenericFeatures(ctx context.Context, featureType string, maximumEstimatedCount int) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getGenericFeatures(const char* featureType, const size_t maximumEstimatedCount, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	featureTypeForC := C.CString(featureType)
	defer C.free(unsafe.Pointer(featureTypeForC))
	stringBuffer := C.GoString(C.G2Diagnostic_getGenericFeatures_local(featureTypeForC, C.size_t(maximumEstimatedCount)))
	if len(stringBuffer) == 0 {
		err = g2diagnostic.getError(ctx, 13, featureType, strconv.Itoa(maximumEstimatedCount))
	}
	return stringBuffer, err
}

// GetLastException returns the last exception encountered in the Xyzzy Engine.
func (g2diagnostic *G2diagnosticImpl) GetLastException(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getLastException(char *buffer, const size_t bufSize);
	var err error = nil
	stringBuffer := g2diagnostic.getByteArray(initialByteArraySize)
	C.G2Diagnostic_getLastException((*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
	if len(stringBuffer) == 0 {
		errorMessage := "xyzzy-60000001 Cannot retrieve last error message."
		err = errors.New(errorMessage)
	}
	return string(stringBuffer), err
}

func (g2diagnostic *G2diagnosticImpl) GetLastExceptionCode(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getLastException(char *buffer, const size_t bufSize);
	var err error = nil
	stringBuffer := g2diagnostic.getByteArray(initialByteArraySize)
	result := C.G2Diagnostic_getLastException((*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
	if result != 0 {
		err = g2diagnostic.getError(ctx, 14)
	}
	return string(stringBuffer), err
}

// GetLogicalCores returns the number of logical cores on the host system.
func (g2diagnostic *G2diagnosticImpl) GetLogicalCores(ctx context.Context) (int, error) {
	// _DLEXPORT int G2Diagnostic_getLogicalCores();
	var err error = nil
	result := C.G2Diagnostic_getLogicalCores()
	return int(result), err
}

func (g2diagnostic *G2diagnosticImpl) GetMappingStatistics(ctx context.Context, includeInternalFeatures int) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getMappingStatistics(const int includeInternalFeatures, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	stringBuffer := C.GoString(C.G2Diagnostic_getMappingStatistics_local(C.int(includeInternalFeatures)))
	if len(stringBuffer) == 0 {
		err = g2diagnostic.getError(ctx, 15, strconv.Itoa(includeInternalFeatures))
	}
	return stringBuffer, err
}

// GetPhysicalCores returns the number of physical cores on the host system.
func (g2diagnostic *G2diagnosticImpl) GetPhysicalCores(ctx context.Context) (int, error) {
	// _DLEXPORT int G2Diagnostic_getPhysicalCores();
	var err error = nil
	result := C.G2Diagnostic_getPhysicalCores()
	return int(result), err
}

func (g2diagnostic *G2diagnosticImpl) GetRelationshipDetails(ctx context.Context, relationshipID int64, includeInternalFeatures int) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getRelationshipDetails(const long long relationshipID, const int includeInternalFeatures, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	stringBuffer := C.GoString(C.G2Diagnostic_getRelationshipDetails_local(C.longlong(relationshipID), C.int(includeInternalFeatures)))
	if len(stringBuffer) == 0 {
		err = g2diagnostic.getError(ctx, 16, strconv.FormatInt(relationshipID, 10), strconv.Itoa(includeInternalFeatures))
	}
	return stringBuffer, err
}

func (g2diagnostic *G2diagnosticImpl) GetResolutionStatistics(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2Diagnostic_getResolutionStatistics(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	stringBuffer := C.GoString(C.G2Diagnostic_getResolutionStatistics_local())
	if len(stringBuffer) == 0 {
		err = g2diagnostic.getError(ctx, 17)
	}
	return stringBuffer, err
}

// GetTotalSystemMemory returns the total memory, in bytes, on the host system.
func (g2diagnostic *G2diagnosticImpl) GetTotalSystemMemory(ctx context.Context) (int64, error) {
	// _DLEXPORT long long G2Diagnostic_getTotalSystemMemory();
	var err error = nil
	result := C.G2Diagnostic_getTotalSystemMemory()
	return int64(result), err
}

// Init initializes the Xyzzy G2diagnosis.
func (g2diagnostic *G2diagnosticImpl) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	// _DLEXPORT int G2Diagnostic_init(const char *moduleName, const char *iniParams, const int verboseLogging);
	var err error = nil

	// Transform Go datatypes to C datatypes.

	moduleNameForC := C.CString(moduleName)
	defer C.free(unsafe.Pointer(moduleNameForC))

	iniParamsForC := C.CString(iniParams)
	defer C.free(unsafe.Pointer(iniParamsForC))

	// Call Xyzzy.

	result := C.G2Diagnostic_init(moduleNameForC, iniParamsForC, C.int(verboseLogging))

	// Handle result.

	if result != 0 {
		err = g2diagnostic.getError(ctx, 18, moduleName, iniParams, strconv.Itoa(verboseLogging))
	}
	return err
}

func (g2diagnostic *G2diagnosticImpl) InitWithConfigID(ctx context.Context, moduleName string, iniParams string, initConfigID int64, verboseLogging int) error {
	//  _DLEXPORT int G2Diagnostic_initWithConfigID(const char *moduleName, const char *iniParams, const long long initConfigID, const int verboseLogging);
	var err error = nil
	// Transform Go datatypes to C datatypes.

	moduleNameForC := C.CString(moduleName)
	defer C.free(unsafe.Pointer(moduleNameForC))

	iniParamsForC := C.CString(iniParams)
	defer C.free(unsafe.Pointer(iniParamsForC))

	// Call Xyzzy.

	result := C.G2Diagnostic_initWithConfigID(moduleNameForC, iniParamsForC, C.longlong(initConfigID), C.int(verboseLogging))

	// Handle result.

	if result != 0 {
		err = g2diagnostic.getError(ctx, 19, moduleName, iniParams, strconv.FormatInt(initConfigID, 10), strconv.Itoa(verboseLogging))
	}
	return err
}

// Null shows how to report a BUG inline.
func (g2diagnostic *G2diagnosticImpl) Null(ctx context.Context) (int64, error) {
	// BUG(mjd): Just an example of how to show bugs in GoDoc.
	var err error = nil
	return 1, err
}

func (g2diagnostic *G2diagnosticImpl) Reinit(ctx context.Context, initConfigID int64) error {
	//  _DLEXPORT int G2Diagnostic_reinit(const long long initConfigID);
	var err error = nil
	result := C.G2Diagnostic_reinit(C.longlong(initConfigID))
	if result != 0 {
		err = g2diagnostic.getError(ctx, 20, strconv.FormatInt(initConfigID, 10))
	}
	return err
}
