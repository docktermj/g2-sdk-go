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
	"strconv"
	"unsafe"

	"github.com/docktermj/xyzzygoapi/g2helper"
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

func (g2engine *G2engineImpl) getError(ctx context.Context, errorNumber int, details ...string) error {
	lastException, err := g2engine.GetLastException(ctx)
	defer g2engine.ClearLastException(ctx)
	message := lastException
	if err != nil {
		message = err.Error()
	}
	return g2helper.BuildError(MessageIdFormat, errorNumber, message, details...)
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

func (g2engine *G2engineImpl) AddRecord() error {
	//  _DLEXPORT int G2_addRecord(const char* dataSourceCode, const char* recordID, const char* jsonData, const char *loadID);
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) AddRecordWithInfo() error {
	//  _DLEXPORT int G2_addRecordWithInfo(const char* dataSourceCode, const char* recordID, const char* jsonData, const char *loadID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) AddRecordWithInfoWithReturnedRecordID() error {
	//  _DLEXPORT int G2_addRecordWithInfoWithReturnedRecordID(const char* dataSourceCode, const char* jsonData, const char *loadID, const long long flags, char *recordIDBuf, const size_t recordIDBufSize, char **responseBuf, size_t *responseBufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) AddRecordWithReturnedRecordID() error {
	//  _DLEXPORT int G2_addRecordWithReturnedRecordID(const char* dataSourceCode, const char* jsonData, const char *loadID, char *recordIDBuf, const size_t bufSize);
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) CheckRecord() error {
	//  _DLEXPORT int G2_checkRecord(const char *record, const char* recordQueryList, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	return err
}

// ClearLastException returns the available memory, in bytes, on the host system.
func (g2engine *G2engineImpl) ClearLastException(ctx context.Context) error {
	// _DLEXPORT void G2_clearLastException();
	var err error = nil
	C.G2_clearLastException()
	return err
}

func (g2engine *G2engineImpl) CloseExport() error {
	//  _DLEXPORT int G2_closeExport(ExportHandle responseHandle);
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) CountRedoRecords() error {
	//  _DLEXPORT long long G2_countRedoRecords();
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) DeleteRecord() error {
	//  _DLEXPORT int G2_deleteRecord(const char* dataSourceCode, const char* recordID, const char* loadID);
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) DeleteRecordWithInfo() error {
	//  _DLEXPORT int G2_deleteRecordWithInfo(const char* dataSourceCode, const char* recordID, const char* loadID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) Destroy() error {
	//  _DLEXPORT int G2_destroy();
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) ExportConfigAndConfigID() error {
	//  _DLEXPORT int G2_exportConfigAndConfigID(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize), long long* configID );
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) ExportConfig() error {
	//  _DLEXPORT int G2_exportConfig(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) ExportCSVEntityReport() error {
	//  _DLEXPORT int G2_exportCSVEntityReport(const char* csvColumnList, const long long flags, ExportHandle* responseHandle);
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) ExportJSONEntityReport() error {
	//  _DLEXPORT int G2_exportJSONEntityReport(const long long flags, ExportHandle* responseHandle);
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) FetchNext() error {
	//  _DLEXPORT int G2_fetchNext(ExportHandle responseHandle, char *responseBuf, const size_t bufSize);
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) FindInterestingEntitiesByEntityID() error {
	//  _DLEXPORT int G2_findInterestingEntitiesByEntityID(const long long entityID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) FindInterestingEntitiesByRecordID() error {
	//  _DLEXPORT int G2_findInterestingEntitiesByRecordID(const char* dataSourceCode, const char* recordID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) FindNetworkByEntityID() error {
	//  _DLEXPORT int G2_findNetworkByEntityID(const char* entityList, const int maxDegree, const int buildOutDegree, const int maxEntities, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) FindNetworkByEntityID_V2() error {
	//  _DLEXPORT int G2_findNetworkByEntityID_V2(const char* entityList, const int maxDegree, const int buildOutDegree, const int maxEntities, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) FindNetworkByRecordID() error {
	//  _DLEXPORT int G2_findNetworkByRecordID(const char* recordList, const int maxDegree, const int buildOutDegree, const int maxEntities, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) FindNetworkByRecordID_V2() error {
	//  _DLEXPORT int G2_findNetworkByRecordID_V2(const char* recordList, const int maxDegree, const int buildOutDegree, const int maxEntities, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) FindPathByEntityID() error {
	//  _DLEXPORT int G2_findPathByEntityID(const long long entityID1, const long long entityID2, const int maxDegree, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) FindPathByEntityID_V2() error {
	//  _DLEXPORT int G2_findPathByEntityID_V2(const long long entityID1, const long long entityID2, const int maxDegree, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) FindPathByRecordID() error {
	//  _DLEXPORT int G2_findPathByRecordID(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) FindPathByRecordID_V2() error {
	//  _DLEXPORT int G2_findPathByRecordID_V2(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) FindPathExcludingByEntityID() error {
	//  _DLEXPORT int G2_findPathExcludingByEntityID(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) FindPathExcludingByEntityID_V2() error {
	//  _DLEXPORT int G2_findPathExcludingByEntityID_V2(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) FindPathExcludingByRecordID() error {
	//  _DLEXPORT int G2_findPathExcludingByRecordID(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) FindPathExcludingByRecordID_V2() error {
	//  _DLEXPORT int G2_findPathExcludingByRecordID_V2(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) FindPathIncludingSourceByEntityID() error {
	//  _DLEXPORT int G2_findPathIncludingSourceByEntityID(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities, const char* requiredDsrcs, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) FindPathIncludingSourceByEntityID_V2() error {
	//  _DLEXPORT int G2_findPathIncludingSourceByEntityID_V2(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities, const char* requiredDsrcs, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) FindPathIncludingSourceByRecordID() error {
	//  _DLEXPORT int G2_findPathIncludingSourceByRecordID(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords, const char* requiredDsrcs, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) FindPathIncludingSourceByRecordID_V2() error {
	//  _DLEXPORT int G2_findPathIncludingSourceByRecordID_V2(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords, const char* requiredDsrcs, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) GetActiveConfigID() error {
	//  _DLEXPORT int G2_getActiveConfigID(long long* configID);
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) GetEntityByEntityID() error {
	//  _DLEXPORT int G2_getEntityByEntityID(const long long entityID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) GetEntityByEntityID_V2() error {
	//  _DLEXPORT int G2_getEntityByEntityID_V2(const long long entityID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) GetEntityByRecordID() error {
	//  _DLEXPORT int G2_getEntityByRecordID(const char* dataSourceCode, const char* recordID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) GetEntityByRecordID_V2() error {
	//  _DLEXPORT int G2_getEntityByRecordID_V2(const char* dataSourceCode, const char* recordID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
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
		err = g2helper.BuildError(MessageIdFormat, 2999, "Cannot retrieve last error message.")
	}
	return string(stringBuffer), err
}

func (g2engine *G2engineImpl) GetLastExceptionCode() error {
	//  _DLEXPORT int G2_getLastExceptionCode();
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) GetRecord() error {
	//  _DLEXPORT int G2_getRecord(const char* dataSourceCode, const char* recordID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) GetRecord_V2() error {
	//  _DLEXPORT int G2_getRecord_V2(const char* dataSourceCode, const char* recordID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) GetRedoRecord() error {
	//  _DLEXPORT int G2_getRedoRecord(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) GetRepositoryLastModifiedTime() error {
	//  _DLEXPORT int G2_getRepositoryLastModifiedTime(long long* lastModifiedTime);
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) GetVirtualEntityByRecordID() error {
	//  _DLEXPORT int G2_getVirtualEntityByRecordID(const char* recordList, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) GetVirtualEntityByRecordID_V2() error {
	//  _DLEXPORT int G2_getVirtualEntityByRecordID_V2(const char* recordList, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) HowEntityByEntityID() error {
	//  _DLEXPORT int G2_howEntityByEntityID(const long long entityID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) HowEntityByEntityID_V2() error {
	//  _DLEXPORT int G2_howEntityByEntityID_V2(const long long entityID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
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

func (g2engine *G2engineImpl) InitWithConfigID() error {
	//  _DLEXPORT int G2_initWithConfigID(const char *moduleName, const char *iniParams, const long long initConfigID, const int verboseLogging);
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) PrimeEngine() error {
	//  _DLEXPORT int G2_primeEngine();
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) Process() error {
	//  _DLEXPORT int G2_process(const char *record);
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) ProcessRedoRecord() error {
	//  _DLEXPORT int G2_processRedoRecord(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) ProcessRedoRecordWithInfo() error {
	//  _DLEXPORT int G2_processRedoRecordWithInfo(const long long flags, char **responseBuf, size_t *bufSize, char **infoBuf, size_t *infoBufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) ProcessWithInfo() error {
	//  _DLEXPORT int G2_processWithInfo(const char *record, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) ProcessWithResponse() error {
	//  _DLEXPORT int G2_processWithResponse(const char *record, char *responseBuf, const size_t bufSize);
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) ProcessWithResponseResize() error {
	//  _DLEXPORT int G2_processWithResponseResize(const char *record, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) PurgeRepository() error {
	//  _DLEXPORT int G2_purgeRepository();
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) ReevaluateEntity() error {
	//  _DLEXPORT int G2_reevaluateEntity(const long long entityID, const long long flags);
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) ReevaluateEntityWithInfo() error {
	//  _DLEXPORT int G2_reevaluateEntityWithInfo(const long long entityID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) ReevaluateRecord() error {
	//  _DLEXPORT int G2_reevaluateRecord(const char* dataSourceCode, const char* recordID, const long long flags);
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) ReevaluateRecordWithInfo() error {
	//  _DLEXPORT int G2_reevaluateRecordWithInfo(const char* dataSourceCode, const char* recordID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) Reinit() error {
	//  _DLEXPORT int G2_reinit(const long long initConfigID);
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) ReplaceRecord() error {
	//  _DLEXPORT int G2_replaceRecord(const char* dataSourceCode, const char* recordID, const char* jsonData, const char *loadID);
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) ReplaceRecordWithInfo() error {
	//  _DLEXPORT int G2_replaceRecordWithInfo(const char* dataSourceCode, const char* recordID, const char* jsonData, const char *loadID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) SearchByAttributes() error {
	//  _DLEXPORT int G2_searchByAttributes(const char* jsonData, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) SearchByAttributes_V2() error {
	//  _DLEXPORT int G2_searchByAttributes_V2(const char* jsonData, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) Stats() error {
	//  _DLEXPORT int G2_stats(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) WhyEntities() error {
	//  _DLEXPORT int G2_whyEntities(const long long entityID1, const long long entityID2, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) WhyEntities_V2() error {
	//  _DLEXPORT int G2_whyEntities_V2(const long long entityID1, const long long entityID2, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) WhyEntityByEntityID() error {
	//  _DLEXPORT int G2_whyEntityByEntityID(const long long entityID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) WhyEntityByEntityID_V2() error {
	//  _DLEXPORT int G2_whyEntityByEntityID_V2(const long long entityID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) WhyEntityByRecordID() error {
	//  _DLEXPORT int G2_whyEntityByRecordID(const char* dataSourceCode, const char* recordID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) WhyEntityByRecordID_V2() error {
	//  _DLEXPORT int G2_whyEntityByRecordID_V2(const char* dataSourceCode, const char* recordID, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) WhyRecords() error {
	//  _DLEXPORT int G2_whyRecords(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}

func (g2engine *G2engineImpl) WhyRecords_V2() error {
	//  _DLEXPORT int G2_whyRecords_V2(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const long long flags, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return err
}
