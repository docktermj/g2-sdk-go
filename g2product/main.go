// The Senzing G2product Package is a Go wrapper over
// Senzing's G2product C binding.
//
// The purpose of a G2product object is:
//   - ...
//   - ...
//   - ...
//
// To use G2product, the LD_LIBRARY_PATH environment variable must include
// a path to Senzing's libraries.  Example:
//
//	export LD_LIBRARY_PATH=/opt/senzing/g2/lib
package g2product

import (
	"context"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type G2productImpl struct{}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const MessageIdFormat = "senzing-6002%04d"

// ----------------------------------------------------------------------------
// Interfaces
// ----------------------------------------------------------------------------

type G2product interface {
	ClearLastException(ctx context.Context) error
	Destroy(ctx context.Context) error
	GetLastException(ctx context.Context) (string, error)
	GetLastExceptionCode(ctx context.Context) (int, error)
	Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error
	License(ctx context.Context) (string, error)
	ValidateLicenseFile(ctx context.Context, licenseFilePath string) (string, error)
	ValidateLicenseStringBase64(ctx context.Context, licenseString string) (string, error)
	Version(ctx context.Context) (string, error)
}
