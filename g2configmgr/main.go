/*
Package G2configmgr is a Go wrapper over Senzing's G2Configmgr C binding.

To use G2configmgr, the LD_LIBRARY_PATH environment variable must include
a path to Senzing's libraries.  Example:

	export LD_LIBRARY_PATH=/opt/senzing/g2/lib
*/
package g2configmgr

import (
	"context"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type G2configmgr interface {
	AddConfig(ctx context.Context, configStr string, configComments string) (int64, error)
	ClearLastException(ctx context.Context) error
	Destroy(ctx context.Context) error
	GetConfig(ctx context.Context, configID int64) (string, error)
	GetConfigList(ctx context.Context) (string, error)
	GetDefaultConfigID(ctx context.Context) (int64, error)
	GetLastException(ctx context.Context) (string, error)
	GetLastExceptionCode(ctx context.Context) (int, error)
	Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error
	ReplaceDefaultConfigID(ctx context.Context, oldConfigID int64, newConfigID int64) error
	SetDefaultConfigID(ctx context.Context, configID int64) error
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const MessageIdFormat = "senzing-6002%04d"

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var Messages = map[int]string{
	1:    "Call to G2ConfigMgr_addConfig(%s, %s) failed. Return code: %d",
	2:    "Call to G2Config_destroy() failed. Return code: %d",
	3:    "Call to G2ConfigMgr_getConfig(%d) failed. Return code: %d",
	4:    "Call to G2ConfigMgr_getConfigList() failed. Return code: %d",
	5:    "Call to G2ConfigMgr_getDefaultConfigID() failed. Return code: %d",
	6:    "Call to G2Config_init(%s, %s, %d) failed. Return code: %d",
	7:    "Call to G2ConfigMgr_replaceDefaultConfigID(%d, %d) failed. Return code: %d",
	8:    "Call to G2ConfigMgr_setDefaultConfigID(%d) failed. Return code: %d",
	2999: "Cannot retrieve last error message.",
}
