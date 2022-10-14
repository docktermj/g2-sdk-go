/*
Package g2configmgr ...
*/
package g2configmgr

/*
#include "g2configmgr.h"
#cgo CFLAGS: -g -I/opt/senzing/g2/sdk/c
#cgo LDFLAGS: -L/opt/senzing/g2/lib -lanalytics -ldb2plugin -lG2 -lg2AddressComp -lg2AddressHasher -lg2CloseNames -lg2CompJavaScoreSet -lg2ConfigParseAddr -lg2DateComp -lg2DistinctFeatJava -lg2DLComp -lg2EFeatJava -lg2EmailComp -lg2ExactDomainMatchComp -lg2ExactMatchComp -lg2FeatBuilder -lg2FormatSSN -lg2GenericHasher -lg2GEOLOCComp -lg2GNRNameComp -lg2GroupAssociationComp -lG2Hasher -lg2IDHasher -lg2JVMPlugin -lg2NameHasher -lg2ParseDOB -lg2ParseEmail -lg2ParseGEOLOC -lg2ParseID -lg2ParseName -lg2ParsePhone -lg2PartialAddresses -lg2PartialDates -lg2PartialNames -lg2PhoneComp -lg2PhoneHasher -lG2SSAdm -lg2SSNComp -lg2STBHasher -lg2StdCountry -lg2StdJava -lg2StdTokenizeName -lg2StrictSubsetFelems -lg2StrictSubsetNormalizedFelems -lg2StrictSubsetTokens -lg2StringComp -lmariadbplugin -lmssqlplugin -lNameDataObject -loracleplugin -lpostgresqlplugin -lscoring -lSpaceTimeBoxStandardizer -lsqliteplugin
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
func (g2configmgr *G2configmgrImpl) getByteArrayC(size int) *C.char {
	bytes := C.malloc(C.size_t(size))
	return (*C.char)(bytes)
}

// TODO: Document.
func (g2configmgr *G2configmgrImpl) getByteArray(size int) []byte {
	return make([]byte, size)
}

// TODO: Document.
func (g2configmgr *G2configmgrImpl) getError(ctx context.Context, errorNumber int, details ...interface{}) error {
	lastException, err := g2configmgr.GetLastException(ctx)
	defer g2configmgr.ClearLastException(ctx)
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
func (g2configmgr *G2configmgrImpl) AddConfig(ctx context.Context, configStr string, configComments string) (int64, error) {
	// _DLEXPORT int G2ConfigMgr_addConfig(const char* configStr, const char* configComments, long long* configID);
	var err error = nil
	configStrForC := C.CString(configStr)
	defer C.free(unsafe.Pointer(configStrForC))
	configCommentsForC := C.CString(configComments)
	defer C.free(unsafe.Pointer(configCommentsForC))
	result := C.G2ConfigMgr_addConfig_helper(configStrForC, configCommentsForC)
	configID := int64(C.longlong(result.configID))
	returnCode := result.returnCode
	if returnCode != 0 {
		err = g2configmgr.getError(ctx, 33, configStr, configComments, returnCode)
	}
	return configID, err
}

// ClearLastException returns the available memory, in bytes, on the host system.
func (g2configmgr *G2configmgrImpl) ClearLastException(ctx context.Context) error {
	// _DLEXPORT void G2Config_clearLastException();
	var err error = nil
	C.G2ConfigMgr_clearLastException()
	return err
}

// TODO: Document.
func (g2configmgr *G2configmgrImpl) Destroy(ctx context.Context) error {
	// _DLEXPORT int G2Config_destroy();
	var err error = nil
	result := C.G2ConfigMgr_destroy()
	if result != 0 {
		err = g2configmgr.getError(ctx, 5, result)
	}
	return err
}

// TODO:
func (g2configmgr *G2configmgrImpl) GetConfigList(ctx context.Context) (string, error) {
	// _DLEXPORT int G2ConfigMgr_getConfigList(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	result := C.G2ConfigMgr_getConfigList_helper()
	configList := C.GoString(result.configList)
	returnCode := result.returnCode
	if returnCode != 0 {
		err = g2configmgr.getError(ctx, 50, result)
	}
	return configList, err
}

// GetLastException returns the last exception encountered in the Senzing Engine.
func (g2configmgr *G2configmgrImpl) GetLastException(ctx context.Context) (string, error) {
	// _DLEXPORT int G2Config_getLastException(char *buffer, const size_t bufSize);
	var err error = nil
	stringBuffer := g2configmgr.getByteArray(initialByteArraySize)
	C.G2ConfigMgr_getLastException((*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
	stringBuffer = bytes.Trim(stringBuffer, "\x00")
	if len(stringBuffer) == 0 {
		err = logger.BuildError(MessageIdFormat, 2999, "Cannot retrieve last error message.")
	}
	return string(stringBuffer), err
}

// TODO: Document.
func (g2configmgr *G2configmgrImpl) GetLastExceptionCode(ctx context.Context) (int, error) {
	//  _DLEXPORT int G2Config_getLastExceptionCode();
	var err error = nil
	result := C.G2ConfigMgr_getLastExceptionCode()
	return int(result), err
}

// TODO: Document.
func (g2configmgr *G2configmgrImpl) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	// _DLEXPORT int G2Config_init(const char *moduleName, const char *iniParams, const int verboseLogging);
	var err error = nil
	moduleNameForC := C.CString(moduleName)
	defer C.free(unsafe.Pointer(moduleNameForC))
	iniParamsForC := C.CString(iniParams)
	defer C.free(unsafe.Pointer(iniParamsForC))
	result := C.G2ConfigMgr_init(moduleNameForC, iniParamsForC, C.int(verboseLogging))
	if result != 0 {
		err = g2configmgr.getError(ctx, 6, moduleName, iniParams, strconv.Itoa(verboseLogging), strconv.Itoa(int(result)))
	}
	return err
}
