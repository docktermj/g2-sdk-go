package g2diagnostic

/*
g2diagnostic

The purpose of a g2diagnostic object is:

 1. ...
*/

import (
  "context"
)

// ----------------------------------------------------------------------------
// G2diagnostic Interface Definition
// ----------------------------------------------------------------------------

type G2diagnostic interface {
  CheckDBPerf(ctx context.Context, secondsToRun int)(string, error)
  ClearLastException(ctx context.Context)
  CloseEntityListBySize(ctx context.Context, entityListBySizeHandle EntityListBySizeHandle) (error)
  Destroy(ctx context.Context) (error)
  FetchNextEntityBySize(ctx context.Context, entityListBySizeHandle EntityListBySizeHandle)(string, error)
  FindEntitiesByFeatureIDs(ctx context.Context, features string )(string, error)
  GetAvailableMemory(ctx context.Context)(int64, error)
  GetDBInfo(ctx context.Context)(string, error)
  GetDataSourceCounts(ctx context.Context)(string, error)
  GetEntityDetails(ctx context.Context, entityID int64, includeInternalFeatures int)(string, error)
  GetEntityListBySize(ctx context.Context)(string, error)
  GetEntityResume(ctx context.Context, entityID int64)(string, error)
  GetEntitySizeBreakdown(ctx context.Context, minimumEntitySize int, includeInternalFeatures int)(string, error)
  GetFeature(ctx context.Context, libFeatID int64)(string, error)
  GetGenericFeatures(ctx context.Context, featureType string, maximumEstimatedCount string)(string, error)
  GetLastException(ctx context.Context)(string, error)
  GetLastExceptionCode(ctx context.Context)(int, error)
  GetLogicalCores(ctx context.Context)(int, error)
  GetMappingStatistics(ctx context.Context,t includeInternalFeatures int)(string, error)
  GetPhysicalCores(ctx context.Context)(int, error)
  GetRelationshipDetails(ctx context.Context, relationshipID int64, includeInternalFeatures int)(string, error)
  GetResolutionStatistics(ctx context.Context)(string, error)
  GetTotalSystemMemory(ctx context.Context)(int64, error)
  Init(cctx context.Context, moduleName string, iniParams string, verboseLogging int)(error)
  InitWithConfigID(ctx context.Context, moduleName string, iniParams string, initConfigID int64, verboseLogging int)(error)
  Reinit(ctx context.Context, initConfigID int64)(error)
}
