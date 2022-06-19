// The Xyzzy G2Diagnostic Package is a Go wrapper over
// Xyzzy's G2Diagnostic C binding.
//
// The purpose of a g2diagnostic object is:
//   • ...
//   • ...
//   • ...
// To use g2diagnostic, the LD_LIBRARY_PATH environment variable must include
// a path to Xyzzy's libraries.  Example:
//  export LD_LIBRARY_PATH=/opt/xyzzy/g2/lib
package g2diagnostic

import (
	"context"
)

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const MessageIdFormat = "xyzzy-6013%04d"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type G2diagnosticImpl struct{}

// ----------------------------------------------------------------------------
// Interfaces
// ----------------------------------------------------------------------------

type G2diagnostic interface {
	CheckDBPerf(ctx context.Context, secondsToRun int) (string, error)
	ClearLastException(ctx context.Context) error
	CloseEntityListBySize(ctx context.Context, entityListBySizeHandle interface{}) error
	Destroy(ctx context.Context) error
	FetchNextEntityBySize(ctx context.Context, entityListBySizeHandle interface{}) (string, error)
	FindEntitiesByFeatureIDs(ctx context.Context, features string) (string, error)
	GetAvailableMemory(ctx context.Context) (int64, error)
	GetDataSourceCounts(ctx context.Context) (string, error)
	GetDBInfo(ctx context.Context) (string, error)
	GetEntityDetails(ctx context.Context, entityID int64, includeInternalFeatures int) (string, error)
	GetEntityListBySize(ctx context.Context, entitySize int) (interface{}, error)
	GetEntityResume(ctx context.Context, entityID int64) (string, error)
	GetEntitySizeBreakdown(ctx context.Context, minimumEntitySize int, includeInternalFeatures int) (string, error)
	GetFeature(ctx context.Context, libFeatID int64) (string, error)
	GetGenericFeatures(ctx context.Context, featureType string, maximumEstimatedCount int) (string, error)
	GetLastException(ctx context.Context) (string, error)
	GetLastExceptionCode(ctx context.Context) (string, error)
	GetLogicalCores(ctx context.Context) (int, error)
	GetMappingStatistics(ctx context.Context, includeInternalFeatures int) (string, error)
	GetPhysicalCores(ctx context.Context) (int, error)
	GetRelationshipDetails(ctx context.Context, relationshipID int64, includeInternalFeatures int) (string, error)
	GetResolutionStatistics(ctx context.Context) (string, error)
	GetTotalSystemMemory(ctx context.Context) (int64, error)
	Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error
	InitWithConfigID(ctx context.Context, moduleName string, iniParams string, initConfigID int64, verboseLogging int) error
	Reinit(ctx context.Context, initConfigID int64) error
}
