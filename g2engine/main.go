// The Xyzzy G2engine Package is a Go wrapper over
// Xyzzy's G2Engine C binding.
//
// The purpose of a g2engine object is:
//   • ...
//   • ...
//   • ...
// To use g2engine, the LD_LIBRARY_PATH environment variable must include
// a path to Xyzzy's libraries.  Example:
//  export LD_LIBRARY_PATH=/opt/xyzzy/g2/lib
package g2engine

import (
	"context"
)

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const MessageIdFormat = "xyzzy-6012%04d"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type G2engineImpl struct{}

// ----------------------------------------------------------------------------
// Interfaces
// ----------------------------------------------------------------------------

type G2engine interface {
	AddRecord(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string) error
	ClearLastException(ctx context.Context) error
	GetLastException(ctx context.Context) (string, error)
	Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error
}
