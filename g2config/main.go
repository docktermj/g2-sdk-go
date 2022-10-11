// The Senzing G2config Package is a Go wrapper over
// Senzing's G2Config C binding.
//
// The purpose of a g2config object is:
//   - ...
//   - ...
//   - ...
//
// To use g2config, the LD_LIBRARY_PATH environment variable must include
// a path to Senzing's libraries.  Example:
//
//	export LD_LIBRARY_PATH=/opt/senzing/g2/lib
package g2config

import (
    "context"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type G2configImpl struct{}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const MessageIdFormat = "senzing-6014%04d"

// ----------------------------------------------------------------------------
// Interfaces
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
