/*
Package g2config ...
*/
package g2config

/*
#include "g2config.h"
#cgo CFLAGS: -g
#cgo LDFLAGS: -shared
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
func (g2config *G2configImpl) getByteArrayC(size int) *C.char {
	bytes := C.malloc(C.size_t(size))
	return (*C.char)(bytes)
}

// TODO: Document.
func (g2config *G2configImpl) getByteArray(size int) []byte {
	return make([]byte, size)
}

// TODO: Document.
func (g2config *G2configImpl) getError(ctx context.Context, errorNumber int, details ...string) error {
	lastException, err := g2config.GetLastException(ctx)
	defer g2config.ClearLastException(ctx)
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
func (g2config *G2configImpl) AddDataSource(ctx context.Context, configHandle uintptr, inputJson string) (string, error) {
	// _DLEXPORT int G2Config_addDataSource(ConfigHandle configHandle, const char *inputJson, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	inputJsonForC := C.CString(inputJson)
	defer C.free(unsafe.Pointer(inputJsonForC))
	result := C.G2Config_addDataSource_helper(C.uintptr_t(configHandle), inputJsonForC)
	response := C.GoString(result.response)
	returnCode := result.returnCode
	if returnCode != 0 {
		err = g2config.getError(ctx, 1, inputJson)
	}
	return response, err
}

// ClearLastException returns the available memory, in bytes, on the host system.
func (g2config *G2configImpl) ClearLastException(ctx context.Context) error {
	// _DLEXPORT void G2Config_clearLastException();
	var err error = nil
	C.G2Config_clearLastException()
	return err
}

// TODO: Document.
func (g2config *G2configImpl) Close(ctx context.Context, configHandle uintptr) error {
	// _DLEXPORT int G2Config_close(ConfigHandle configHandle);
	var err error = nil
	result := C.G2config_close_helper(C.uintptr_t(configHandle))
	if result != 0 {
		err = g2config.getError(ctx, 2)
	}
	return err
}

// TODO: Document.
func (g2config *G2configImpl) Create(ctx context.Context) (uintptr, error) {
	// _DLEXPORT int G2Config_create(ConfigHandle* configHandle);
	var err error = nil
	result := C.G2config_create_helper()
	if result == nil {
		err = g2config.getError(ctx, 3)
	}
	return (uintptr)(result), err
}

// TODO: Document.
func (g2config *G2configImpl) DeleteDataSource(ctx context.Context, configHandle uintptr, inputJson string) error {
	// _DLEXPORT int G2Config_deleteDataSource(ConfigHandle configHandle, const char *inputJson);
	var err error = nil
	inputJsonForC := C.CString(inputJson)
	defer C.free(unsafe.Pointer(inputJsonForC))
	returnCode := C.G2Config_deleteDataSource_helper(C.uintptr_t(configHandle), inputJsonForC)
	if returnCode != 0 {
		err = g2config.getError(ctx, 4, inputJson)
	}
	return err
}

// TODO: Document.
func (g2config *G2configImpl) Destroy(ctx context.Context) error {
	// _DLEXPORT int G2Config_destroy();
	var err error = nil
	result := C.G2Config_destroy()
	if result != 0 {
		err = g2config.getError(ctx, 5)
	}
	return err
}

// GetLastException returns the last exception encountered in the Senzing Engine.
func (g2config *G2configImpl) GetLastException(ctx context.Context) (string, error) {
	// _DLEXPORT int G2Config_getLastException(char *buffer, const size_t bufSize);
	var err error = nil
	stringBuffer := g2config.getByteArray(initialByteArraySize)
	C.G2Config_getLastException((*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
	stringBuffer = bytes.Trim(stringBuffer, "\x00")
	if len(stringBuffer) == 0 {
		err = logger.BuildError(MessageIdFormat, 2999, "Cannot retrieve last error message.")
	}
	return string(stringBuffer), err
}

// TODO: Document.
func (g2config *G2configImpl) GetLastExceptionCode(ctx context.Context) (int, error) {
	//  _DLEXPORT int G2Config_getLastExceptionCode();
	var err error = nil
	result := C.G2Config_getLastExceptionCode()
	return int(result), err
}

// TODO: Document.
func (g2config *G2configImpl) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	// _DLEXPORT int G2Config_init(const char *moduleName, const char *iniParams, const int verboseLogging);
	var err error = nil
	moduleNameForC := C.CString(moduleName)
	defer C.free(unsafe.Pointer(moduleNameForC))
	iniParamsForC := C.CString(iniParams)
	defer C.free(unsafe.Pointer(iniParamsForC))
	result := C.G2Config_init(moduleNameForC, iniParamsForC, C.int(verboseLogging))
	if result != 0 {
		err = g2config.getError(ctx, 6, moduleName, iniParams, strconv.Itoa(verboseLogging))
	}
	return err
}

// TODO: Document.
func (g2config *G2configImpl) ListDataSources(ctx context.Context, configHandle uintptr) (string, error) {
	// _DLEXPORT int G2Config_listDataSources(ConfigHandle configHandle, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	result := C.G2Config_listDataSources_helper(C.uintptr_t(configHandle))
	response := C.GoString(result.response)
	returnCode := result.returnCode
	if returnCode != 0 {
		err = g2config.getError(ctx, 7)
	}
	return response, err
}

// TODO: Document.
func (g2config *G2configImpl) Load(ctx context.Context, configHandle uintptr, jsonConfig string) error {
	// _DLEXPORT int G2Config_load(const char *jsonConfig,ConfigHandle* configHandle);
	var err error = nil
	jsonConfigForC := C.CString(jsonConfig)
	defer C.free(unsafe.Pointer(jsonConfigForC))
	returnCode := C.G2Config_load_helper(C.uintptr_t(configHandle), jsonConfigForC)
	if returnCode != 0 {
		err = g2config.getError(ctx, 8, jsonConfig)
	}
	return err
}

// TODO: Document.
func (g2config *G2configImpl) Save(ctx context.Context, configHandle uintptr) (string, error) {
	// _DLEXPORT int G2Config_save(ConfigHandle configHandle, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	result := C.G2Config_save_helper(C.uintptr_t(configHandle))
	response := C.GoString(result.response)
	returnCode := result.returnCode
	if returnCode != 0 {
		err = g2config.getError(ctx, 9)
	}
	return response, err
}
