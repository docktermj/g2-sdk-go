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
	"context"

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
func (g2config *G2configImpl) AddDataSource(ctx context.Context, configHandle int64, inputJson string) (string, error) {
	// _DLEXPORT int G2Config_addDataSource(ConfigHandle configHandle, const char *inputJson, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2config *G2configImpl) ClearLastException(ctx context.Context) error {
	// _DLEXPORT void G2Config_clearLastException();
	var err error = nil
	return err
}

// TODO: Document.
func (g2config *G2configImpl) Close(ctx context.Context, configHandle int64) error {
	// _DLEXPORT int G2Config_close(ConfigHandle configHandle);
	var err error = nil
	return err
}

// TODO: Document.
func (g2config *G2configImpl) Create(ctx context.Context) (int64, error) {
	// _DLEXPORT int G2Config_create(ConfigHandle* configHandle);

	var err error = nil
	result := C.G2config_create_helper()
	return int64(result), err
}

// TODO: Document.
func (g2config *G2configImpl) DeleteDataSource(ctx context.Context, configHandle int64, inputJson string) error {
	// _DLEXPORT int G2Config_deleteDataSource(ConfigHandle configHandle, const char *inputJson);
	var err error = nil
	return err
}

// TODO: Document.
func (g2config *G2configImpl) Destroy(ctx context.Context) error {
	// _DLEXPORT int G2Config_destroy();
	var err error = nil
	return err
}

// TODO: Document.
func (g2config *G2configImpl) GetLastException(ctx context.Context) (string, error) {
	// _DLEXPORT int G2Config_getLastException(char *buffer, const size_t bufSize);
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2config *G2configImpl) GetLastExceptionCode(ctx context.Context) (int, error) {
	//  _DLEXPORT int G2Config_getLastExceptionCode();
	var err error = nil
	return 0, err
}

// TODO: Document.
func (g2config *G2configImpl) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	// _DLEXPORT int G2Config_init(const char *moduleName, const char *iniParams, const int verboseLogging);
	var err error = nil
	return err
}

// TODO: Document.
func (g2config *G2configImpl) ListDataSources(ctx context.Context, configHandle int64) (string, error) {
	// _DLEXPORT int G2Config_listDataSources(ConfigHandle configHandle, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2config *G2configImpl) Load(ctx context.Context, jsonConfig string) (string, error) {
	// _DLEXPORT int G2Config_load(const char *jsonConfig,ConfigHandle* configHandle);
	var err error = nil
	return "", err
}

// TODO: Document.
func (g2config *G2configImpl) Save(ctx context.Context, configHandle int64) (string, error) {
	// _DLEXPORT int G2Config_save(ConfigHandle configHandle, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
	var err error = nil
	return "", err
}
