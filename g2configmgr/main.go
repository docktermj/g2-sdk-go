// The Senzing G2configmgr Package is a Go wrapper over
// Senzing's G2Configmgr C binding.
//
// The purpose of a g2configmgr object is:
//   - ...
//   - ...
//   - ...
//
// To use g2configmgr, the LD_LIBRARY_PATH environment variable must include
// a path to Senzing's libraries.  Example:
//
//	export LD_LIBRARY_PATH=/opt/senzing/g2/lib
package g2configmgr

import (
	"context"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type G2configmgrImpl struct{}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const MessageIdFormat = "senzing-6002%04d"

// ----------------------------------------------------------------------------
// Interfaces
// ----------------------------------------------------------------------------

type G2configmgr interface {
	AddConfig(ctx context.Context, configStr string, configComments string) (int64, error)
	ClearLastException(ctx context.Context) error
	Destroy(ctx context.Context) error
	// GetConfig(ctx context.Context, configID int64) (string, error)
	// GetConfigList(ctx context.Context) (string, error)
	// GetDefaultConfigID(ctx context.Context) (int64, error)
	GetLastException(ctx context.Context) (string, error)
	GetLastExceptionCode(ctx context.Context) (int, error)
	Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error
	// ReplaceDefaultConfigID(ctx context.Context, oldConfigID int64, newConfigID int64) error
	// SetDefaultConfigID(ctx context.Context, configID int64) error
}
