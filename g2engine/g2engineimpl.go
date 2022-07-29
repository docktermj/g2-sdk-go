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

void* G2_resizeStringBuffer(void *ptr, size_t size) {
    //deallocate old buffer
    if (ptr != 0)
        free(ptr);
    //allocate new buffer
    void* buffer = malloc(size);
    return buffer;
}

char* G2_addRecordWithInfo_local(const char* dataSourceCode, const char* recordID, const char* jsonData, const char *loadID, const long long flags) {
    size_t bufferSize = 1;
    char *charBuff = (char *)malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    int returnCode = G2_addRecordWithInfo(dataSourceCode, recordID, jsonData, loadID, flags, &charBuff, &bufferSize, resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2_addRecordWithInfoWithReturnedRecordID_local(const char* dataSourceCode, const char* jsonData, const char *loadID, const long long flags) {
    size_t bufferSize = 1;
    size_t recordIDBufSize = 1;

    char recordIDBuf[100];
    char *charBuff = (char *)malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    int returnCode = G2_addRecordWithInfoWithReturnedRecordID(dataSourceCode, jsonData, loadID, flags, recordIDBuf, recordIDBufSize, &charBuff, &bufferSize, resizeFuncPointer);

    //  _DLEXPORT int G2_addRecordWithInfoWithReturnedRecordID(const char* dataSourceCode, const char* jsonData, const char *loadID, const long long flags, char *recordIDBuf, const size_t recordIDBufSize, char **responseBuf, size_t *responseBufSize, void *(*resizeFunc)(void *ptr, size_t newSize));

    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2_deleteRecordWithInfo_local(const char* dataSourceCode, const char* recordID, const char *loadID, const long long flags) {
    size_t bufferSize = 1;
    char *charBuff = (char *)malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    int returnCode = G2_deleteRecordWithInfo(dataSourceCode, recordID, loadID, flags, &charBuff, &bufferSize, resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}


char* G2_stats_local() {
    size_t bufferSize = 1;
    char *charBuff = (char *)malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    int returnCode = G2_stats(&charBuff, &bufferSize, resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

*/
import "C"
import (
	"bytes"
	"context"
	"strconv"
	"unsafe"

	"github.com/docktermj/go-xyzzy-helpers/logger"
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

// TODO: Document.
func (g2engine *G2engineImpl) getByteArray(size int) []byte {
	return make([]byte, size)
}

// TODO: Document.
func (g2engine *G2engineImpl) getError(ctx context.Context, errorNumber int, details ...string) error {
	lastException, err := g2engine.GetLastException(ctx)
	defer g2engine.ClearLastException(ctx)
	message := lastException
	if err != nil {
		message = err.Error()
	}
	return logger.BuildError(MessageIdFormat, errorNumber, message, details...)
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// TODO: Document.
func (g2engine *G2engineImpl) AddRecord(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string) error {
	//  _DLEXPORT int G2_addRecord(const char* dataSourceCode, const char* recordID, const char* jsonData, const char *loadID);
	var err error = nil

	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))

	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))

	jsonDataForC := C.CString(jsonData)
	defer C.free(unsafe.Pointer(jsonDataForC))

	loadIDForC := C.CString(loadID)
	defer C.free(unsafe.Pointer(loadIDForC))

	result := C.G2_addRecord(dataSourceCodeForC, recordIDForC, jsonDataForC, loadIDForC)

	// Handle result.

	if result != 0 {
		err = g2engine.getError(ctx, 1, dataSourceCode, recordID, jsonData, loadID)
	}
	return err
}

// TODO: Document.
func (g2engine *G2engineImpl) AddRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_addRecordWithInfo(const char* dataSourceCode, const char* recordID, const char* jsonData, const char *loadID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil

	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))

	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))

	jsonDataForC := C.CString(jsonData)
	defer C.free(unsafe.Pointer(jsonDataForC))

	loadIDForC := C.CString(loadID)
	defer C.free(unsafe.Pointer(loadIDForC))

	stringBuffer := C.GoString(C.G2_addRecordWithInfo_local(dataSourceCodeForC, recordIDForC, jsonDataForC, loadIDForC, C.longlong(flags)))

	// Handle result.

	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2, dataSourceCode, recordID, jsonData, loadID, strconv.FormatInt(flags, 2))
	}
	return stringBuffer, err
}

// TODO: Document.
func (g2engine *G2engineImpl) AddRecordWithInfoWithReturnedRecordID(ctx context.Context, dataSourceCode string, jsonData string, loadID string, flags int64) (string, string, error) {
	//  _DLEXPORT int G2_addRecordWithInfoWithReturnedRecordID(const char* dataSourceCode, const char* jsonData, const char *loadID, const long long flags, char *recordIDBuf, const size_t recordIDBufSize, char **responseBuf, size_t *responseBufSize, void *(*resizeFunc)(void *ptr, size_t newSize));

	var err error = nil

	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))

	jsonDataForC := C.CString(jsonData)
	defer C.free(unsafe.Pointer(jsonDataForC))

	loadIDForC := C.CString(loadID)
	defer C.free(unsafe.Pointer(loadIDForC))

	// FIXME:

	stringIdBuffer := ""

	// FIXME: figure out how to get 2 strings from a C function.
	stringBuffer := C.GoString(C.G2_addRecordWithInfoWithReturnedRecordID_local(dataSourceCodeForC, jsonDataForC, loadIDForC, C.longlong(flags)))

	// Handle result.

	if len(stringIdBuffer) == 0 || len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 2, dataSourceCode, stringIdBuffer, jsonData, loadID, strconv.FormatInt(flags, 2))
	}

	return stringIdBuffer, stringBuffer, err
}

// TODO: Document.
func (g2engine *G2engineImpl) AddRecordWithReturnedRecordID(ctx context.Context, dataSourceCode string, jsonData string, loadID string) (string, error) {
	//  _DLEXPORT int G2_addRecordWithReturnedRecordID(const char* dataSourceCode, const char* jsonData, const char *loadID, char *recordIDBuf, const size_t bufSize);
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) CheckRecord(ctx context.Context, record string, recordQueryList string) (string, error) {
	//  _DLEXPORT int G2_checkRecord(const char *record, const char* recordQueryList, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	return "", err
}

// ClearLastException returns the available memory, in bytes, on the host system.
func (g2engine *G2engineImpl) ClearLastException(ctx context.Context) error {
	// _DLEXPORT void G2_clearLastException();
	var err error = nil
	C.G2_clearLastException()
	return err
}

// TODO: Document.
func (g2engine *G2engineImpl) CloseExport(ctx context.Context, responseHandle interface{}) error {
	//  _DLEXPORT int G2_closeExport(ExportHandle responseHandle);
	var err error = nil
	return err
}

// TODO: Document.
func (g2engine *G2engineImpl) CountRedoRecords(ctx context.Context) (int64, error) {
	//  _DLEXPORT long long G2_countRedoRecords();
	var err error = nil
	return 0, err
}

// TODO: Document.
func (g2engine *G2engineImpl) DeleteRecord(ctx context.Context, dataSourceCode string, recordID string, loadID string) error {
	//  _DLEXPORT int G2_deleteRecord(const char* dataSourceCode, const char* recordID, const char* loadID);
	var err error = nil

	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))

	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))

	loadIDForC := C.CString(loadID)
	defer C.free(unsafe.Pointer(loadIDForC))

	result := C.G2_deleteRecord(dataSourceCodeForC, recordIDForC, loadIDForC)

	// Handle result.

	if result != 0 {
		err = g2engine.getError(ctx, 3, dataSourceCode, recordID, loadID)
	}

	return err
}

// TODO: Document.
func (g2engine *G2engineImpl) DeleteRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, loadID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_deleteRecordWithInfo(const char* dataSourceCode, const char* recordID, const char* loadID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	dataSourceCodeForC := C.CString(dataSourceCode)
	defer C.free(unsafe.Pointer(dataSourceCodeForC))

	recordIDForC := C.CString(recordID)
	defer C.free(unsafe.Pointer(recordIDForC))

	loadIDForC := C.CString(loadID)
	defer C.free(unsafe.Pointer(loadIDForC))

	stringBuffer := C.GoString(C.G2_deleteRecordWithInfo_local(dataSourceCodeForC, recordIDForC, loadIDForC, C.longlong(flags)))

	// Handle result.

	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 4, dataSourceCode, recordID, loadID, strconv.FormatInt(flags, 2))
	}
	return stringBuffer, err
}

// TODO: Document.
func (g2engine *G2engineImpl) Destroy(ctx context.Context) error {
	//  _DLEXPORT int G2_destroy();
	var err error = nil
	result := C.G2_destroy()
	if result != 0 {
		err = g2engine.getError(ctx, 5)
	}
	return err
}

// TODO: Document.
func (g2engine *G2engineImpl) ExportConfigAndConfigID(ctx context.Context) (string, int64, error) {
	//  _DLEXPORT int G2_exportConfigAndConfigID(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize), long long* configID );
	var err error = nil
	return "", 0, err
}

// TODO: Document.
func (g2engine *G2engineImpl) ExportConfig(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2_exportConfig(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) ExportCSVEntityReport(ctx context.Context, csvColumnList string, flags int64) (interface{}, error) {
	//  _DLEXPORT int G2_exportCSVEntityReport(const char* csvColumnList, const long long flags, ExportHandle* responseHandle);
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) ExportJSONEntityReport(ctx context.Context, flags int64) (interface{}, error) {
	//  _DLEXPORT int G2_exportJSONEntityReport(const long long flags, ExportHandle* responseHandle);
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) FetchNext(ctx context.Context, responseHandle interface{}) (string, error) {
	//  _DLEXPORT int G2_fetchNext(ExportHandle responseHandle, char *responseBuf, const size_t bufSize);
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) FindInterestingEntitiesByEntityID(ctx context.Context, entityID int64, flags int64) (string, error) {
	//  _DLEXPORT int G2_findInterestingEntitiesByEntityID(const long long entityID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) FindInterestingEntitiesByRecordID(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_findInterestingEntitiesByRecordID(const char* dataSourceCode, const char* recordID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) FindNetworkByEntityID(ctx context.Context, entityList string, maxDegree int, buildOutDegree int, maxEntities int) (string, error) {
	//  _DLEXPORT int G2_findNetworkByEntityID(const char* entityList, const int maxDegree, const int buildOutDegree, const int maxEntities, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) FindNetworkByEntityID_V2(ctx context.Context, entityList string, maxDegree int, buildOutDegree int, maxEntities int, flags int64) (string, error) {
	//  _DLEXPORT int G2_findNetworkByEntityID_V2(const char* entityList, const int maxDegree, const int buildOutDegree, const int maxEntities, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) FindNetworkByRecordID(ctx context.Context, recordList string, maxDegree int, buildOutDegree int, maxEntities int) (string, error) {
	//  _DLEXPORT int G2_findNetworkByRecordID(const char* recordList, const int maxDegree, const int buildOutDegree, const int maxEntities, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) FindNetworkByRecordID_V2(ctx context.Context, recordList string, maxDegree int, buildOutDegree int, maxEntities int, flags int64) (string, error) {
	//  _DLEXPORT int G2_findNetworkByRecordID_V2(const char* recordList, const int maxDegree, const int buildOutDegree, const int maxEntities, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) FindPathByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int) (string, error) {
	//  _DLEXPORT int G2_findPathByEntityID(const long long entityID1, const long long entityID2, const int maxDegree, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) FindPathByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, flags int64) (string, error) {
	//  _DLEXPORT int G2_findPathByEntityID_V2(const long long entityID1, const long long entityID2, const int maxDegree, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) FindPathByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int) (string, error) {
	//  _DLEXPORT int G2_findPathByRecordID(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) FindPathByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, flags int64) (string, error) {
	//  _DLEXPORT int G2_findPathByRecordID_V2(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) FindPathExcludingByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string) (string, error) {
	//  _DLEXPORT int G2_findPathExcludingByEntityID(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) FindPathExcludingByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string, flags int64) (string, error) {
	//  _DLEXPORT int G2_findPathExcludingByEntityID_V2(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) FindPathExcludingByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string) (string, error) {
	//  _DLEXPORT int G2_findPathExcludingByRecordID(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) FindPathExcludingByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string, flags int64) (string, error) {
	//  _DLEXPORT int G2_findPathExcludingByRecordID_V2(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) FindPathIncludingSourceByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string) (string, error) {
	//  _DLEXPORT int G2_findPathIncludingSourceByEntityID(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities, const char* requiredDsrcs, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) FindPathIncludingSourceByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string, requiredDsrcs string, flags int64) (string, error) {
	//  _DLEXPORT int G2_findPathIncludingSourceByEntityID_V2(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities, const char* requiredDsrcs, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) FindPathIncludingSourceByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string, requiredDsrcs string) (string, error) {
	//  _DLEXPORT int G2_findPathIncludingSourceByRecordID(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords, const char* requiredDsrcs, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) FindPathIncludingSourceByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string, requiredDsrcs string, flags int64) (string, error) {
	//  _DLEXPORT int G2_findPathIncludingSourceByRecordID_V2(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords, const char* requiredDsrcs, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) GetActiveConfigID(ctx context.Context) (int64, error) {
	//  _DLEXPORT int G2_getActiveConfigID(long long* configID);
	var err error = nil
	return 0, err
}

// TODO: Document.
func (g2engine *G2engineImpl) GetEntityByEntityID(ctx context.Context, entityID int64) (string, error) {
	//  _DLEXPORT int G2_getEntityByEntityID(const long long entityID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) GetEntityByEntityID_V2(ctx context.Context, entityID int64, flags int64) (string, error) {
	//  _DLEXPORT int G2_getEntityByEntityID_V2(const long long entityID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) GetEntityByRecordID(ctx context.Context, dataSourceCode string, recordID string) (string, error) {
	//  _DLEXPORT int G2_getEntityByRecordID(const char* dataSourceCode, const char* recordID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) GetEntityByRecordID_V2(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_getEntityByRecordID_V2(const char* dataSourceCode, const char* recordID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// GetLastException returns the last exception encountered in the Senzing Engine.
func (g2engine *G2engineImpl) GetLastException(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2_getLastException(char *buffer, const size_t bufSize);
	var err error = nil
	stringBuffer := g2engine.getByteArray(initialByteArraySize)
	C.G2_getLastException((*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
	stringBuffer = bytes.Trim(stringBuffer, "\x00")
	if len(stringBuffer) == 0 {
		err = logger.BuildError(MessageIdFormat, 2999, "Cannot retrieve last error message.")
	}
	return string(stringBuffer), err
}

// TODO: Document.
func (g2engine *G2engineImpl) GetLastExceptionCode(ctx context.Context) (int, error) {
	//  _DLEXPORT int G2_getLastExceptionCode();
	var err error = nil
	return 0, err
}

// TODO: Document.
func (g2engine *G2engineImpl) GetRecord(ctx context.Context, dataSourceCode string, recordID string) (string, error) {
	//  _DLEXPORT int G2_getRecord(const char* dataSourceCode, const char* recordID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) GetRecord_V2(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_getRecord_V2(const char* dataSourceCode, const char* recordID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) GetRedoRecord(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2_getRedoRecord(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) GetRepositoryLastModifiedTime(ctx context.Context) (int64, error) {
	//  _DLEXPORT int G2_getRepositoryLastModifiedTime(long long* lastModifiedTime);
	var err error = nil
	return 0, err
}

// TODO: Document.
func (g2engine *G2engineImpl) GetVirtualEntityByRecordID(ctx context.Context, recordList string) (int64, error) {
	//  _DLEXPORT int G2_getVirtualEntityByRecordID(const char* recordList, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return 0, err
}

// TODO: Document.
func (g2engine *G2engineImpl) GetVirtualEntityByRecordID_V2(ctx context.Context, recordList string, flags int64) (string, error) {
	//  _DLEXPORT int G2_getVirtualEntityByRecordID_V2(const char* recordList, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) HowEntityByEntityID(ctx context.Context, entityID int64) (string, error) {
	//  _DLEXPORT int G2_howEntityByEntityID(const long long entityID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) HowEntityByEntityID_V2(ctx context.Context, entityID int64, flags int64) (string, error) {
	//  _DLEXPORT int G2_howEntityByEntityID_V2(const long long entityID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// Init initializes the Senzing G2diagnosis.
func (g2engine *G2engineImpl) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	// _DLEXPORT int G2_init(const char *moduleName, const char *iniParams, const int verboseLogging);
	var err error = nil

	// Transform Go datatypes to C datatypes.

	moduleNameForC := C.CString(moduleName)
	defer C.free(unsafe.Pointer(moduleNameForC))

	iniParamsForC := C.CString(iniParams)
	defer C.free(unsafe.Pointer(iniParamsForC))

	// Call Senzing.

	result := C.G2_init(moduleNameForC, iniParamsForC, C.int(verboseLogging))

	// Handle result.

	if result != 0 {
		err = g2engine.getError(ctx, 6, moduleName, iniParams, strconv.Itoa(verboseLogging))
	}
	return err
}

// TODO: Document.
func (g2engine *G2engineImpl) InitWithConfigID(ctx context.Context, moduleName string, iniParams string, initConfigID int64, verboseLogging int) error {
	//  _DLEXPORT int G2_initWithConfigID(const char *moduleName, const char *iniParams, const long long initConfigID, const int verboseLogging);
	var err error = nil
	return err
}

// TODO: Document.
func (g2engine *G2engineImpl) PrimeEngine(ctx context.Context) error {
	//  _DLEXPORT int G2_primeEngine();
	var err error = nil
	return err
}

// TODO: Document.
func (g2engine *G2engineImpl) Process(ctx context.Context, record string) error {
	//  _DLEXPORT int G2_process(const char *record);
	var err error = nil
	return err
}

// TODO: Document.
func (g2engine *G2engineImpl) ProcessRedoRecord(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2_processRedoRecord(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) ProcessRedoRecordWithInfo(ctx context.Context, flags int64) (string, string, error) {
	//  _DLEXPORT int G2_processRedoRecordWithInfo(const long long flags, char **responseBuf, size_t *bufSize, char **infoBuf, size_t *infoBufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) ProcessWithInfo(ctx context.Context, record string, flags int64) (string, error) {
	//  _DLEXPORT int G2_processWithInfo(const char *record, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) ProcessWithResponse(ctx context.Context, record string) (string, error) {
	//  _DLEXPORT int G2_processWithResponse(const char *record, char *responseBuf, const size_t bufSize);
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) ProcessWithResponseResize(ctx context.Context, record string) (string, error) {
	//  _DLEXPORT int G2_processWithResponseResize(const char *record, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) PurgeRepository(ctx context.Context) error {
	//  _DLEXPORT int G2_purgeRepository();
	var err error = nil
	return err
}

// TODO: Document.
func (g2engine *G2engineImpl) ReevaluateEntity(ctx context.Context, entityID int64, flags int64) error {
	//  _DLEXPORT int G2_reevaluateEntity(const long long entityID, const long long flags);
	var err error = nil
	return err
}

// TODO: Document.
func (g2engine *G2engineImpl) ReevaluateEntityWithInfo(ctx context.Context, entityID int64, flags int64) (string, error) {
	//  _DLEXPORT int G2_reevaluateEntityWithInfo(const long long entityID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) ReevaluateRecord(ctx context.Context, dataSourceCode string, recordID string, flags int64) error {
	//  _DLEXPORT int G2_reevaluateRecord(const char* dataSourceCode, const char* recordID, const long long flags);
	var err error = nil
	return err
}

// TODO: Document.
func (g2engine *G2engineImpl) ReevaluateRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_reevaluateRecordWithInfo(const char* dataSourceCode, const char* recordID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) Reinit(ctx context.Context, initConfigID int64) error {
	//  _DLEXPORT int G2_reinit(const long long initConfigID);
	var err error = nil
	return err
}

// TODO: Document.
func (g2engine *G2engineImpl) ReplaceRecord(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string) error {
	//  _DLEXPORT int G2_replaceRecord(const char* dataSourceCode, const char* recordID, const char* jsonData, const char *loadID);
	var err error = nil
	return err
}

// TODO: Document.
func (g2engine *G2engineImpl) ReplaceRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_replaceRecordWithInfo(const char* dataSourceCode, const char* recordID, const char* jsonData, const char *loadID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) SearchByAttributes(ctx context.Context, jsonData string) (string, error) {
	//  _DLEXPORT int G2_searchByAttributes(const char* jsonData, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) SearchByAttributes_V2(ctx context.Context, jsonData string, flags int64) (string, error) {
	//  _DLEXPORT int G2_searchByAttributes_V2(const char* jsonData, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) Stats(ctx context.Context) (string, error) {
	//  _DLEXPORT int G2_stats(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	stringBuffer := C.GoString(C.G2_stats_local())
	if len(stringBuffer) == 0 {
		err = g2engine.getError(ctx, 7)
	}
	return stringBuffer, err
}

// TODO: Document.
func (g2engine *G2engineImpl) WhyEntities(ctx context.Context, entityID1 int64, entityID2 int64) (string, error) {
	//  _DLEXPORT int G2_whyEntities(const long long entityID1, const long long entityID2, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) WhyEntities_V2(ctx context.Context, entityID1 int64, entityID2 int64, flags int64) (string, error) {
	//  _DLEXPORT int G2_whyEntities_V2(const long long entityID1, const long long entityID2, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) WhyEntityByEntityID(ctx context.Context, entityID int64) (string, error) {
	//  _DLEXPORT int G2_whyEntityByEntityID(const long long entityID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) WhyEntityByEntityID_V2(ctx context.Context, entityID int64, flags int64) (string, error) {
	//  _DLEXPORT int G2_whyEntityByEntityID_V2(const long long entityID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) WhyEntityByRecordID(ctx context.Context, dataSourceCode string, recordID string) (string, error) {
	//  _DLEXPORT int G2_whyEntityByRecordID(const char* dataSourceCode, const char* recordID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) WhyEntityByRecordID_V2(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error) {
	//  _DLEXPORT int G2_whyEntityByRecordID_V2(const char* dataSourceCode, const char* recordID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) WhyRecords(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string) (string, error) {
	//  _DLEXPORT int G2_whyRecords(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2engine *G2engineImpl) WhyRecords_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, flags int64) (string, error) {
	//  _DLEXPORT int G2_whyRecords_V2(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}
