/*
Package g2config is a Go wrapper over Senzing's G2Config C binding.

To use G2config, the LD_LIBRARY_PATH environment variable must include
a path to Senzing's libraries.  Example:

	export LD_LIBRARY_PATH=/opt/senzing/g2/lib
*/
package g2config

import (
	"context"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type G2config interface {
	AddDataSource(ctx context.Context, configHandle uintptr, inputJson string) (string, error)
	ClearLastException(ctx context.Context) error
	Close(ctx context.Context, configHandle uintptr) error
	Create(ctx context.Context) (uintptr, error)
	DeleteDataSource(ctx context.Context, configHandle uintptr, inputJson string) error
	Destroy(ctx context.Context) error
	GetLastException(ctx context.Context) (string, error)
	GetLastExceptionCode(ctx context.Context) (int, error)
	Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error
	ListDataSources(ctx context.Context, configHandle uintptr) (string, error)
	Load(ctx context.Context, configHandle uintptr, jsonConfig string) error
	Save(ctx context.Context, configHandle uintptr) (string, error)
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const MessageIdFormat = "senzing-6001%04d"

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var Messages = map[int]string{
	1:    "Call to G2Config_addDataSource(%s) failed. Return code: %d",
	2:    "Call to G2Config_close() failed. Return code: %d",
	3:    "Call to G2Config_create() failed. Return code: %d",
	4:    "Call to G2Config_deleteDataSource(%s) failed. Return code: %d",
	5:    "Call to G2Config_destroy() failed. Return code: %d",
	6:    "Call to G2Config_init(%s, %s, %d) failed. Return code: %d",
	7:    "Call to G2Config_listDataSources() failed. Return code: %d",
	8:    "Call to G2Config_load(%s) failed. Return code: %d",
	9:    "Call to G2Config_save() failed. Return code: %d",
	2999: "Cannot retrieve last error message.",
}
