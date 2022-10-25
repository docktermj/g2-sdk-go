/*
Package g2engine is a Go wrapper over Senzing's G2Engine C binding.

To use G2engine, the LD_LIBRARY_PATH environment variable must include
a path to Senzing's libraries.  Example:

	export LD_LIBRARY_PATH=/opt/senzing/g2/lib
*/
package g2engine

import (
	"context"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type FlagMask int64

type G2engine interface {
	AddRecord(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string) error
	AddRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string, flags int64) (string, error)
	AddRecordWithInfoWithReturnedRecordID(ctx context.Context, dataSourceCode string, jsonData string, loadID string, flags int64) (string, string, error)
	AddRecordWithReturnedRecordID(ctx context.Context, dataSourceCode string, jsonData string, loadID string) (string, error)
	CheckRecord(ctx context.Context, record string, recordQueryList string) (string, error)
	ClearLastException(ctx context.Context) error
	CloseExport(ctx context.Context, responseHandle uintptr) error
	CountRedoRecords(ctx context.Context) (int64, error)
	DeleteRecord(ctx context.Context, dataSourceCode string, recordID string, loadID string) error
	DeleteRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, loadID string, flags int64) (string, error)
	Destroy(ctx context.Context) error
	ExportConfig(ctx context.Context) (string, error)
	ExportConfigAndConfigID(ctx context.Context) (string, int64, error)
	ExportCSVEntityReport(ctx context.Context, csvColumnList string, flags int64) (uintptr, error)
	ExportJSONEntityReport(ctx context.Context, flags int64) (uintptr, error)
	FetchNext(ctx context.Context, responseHandle uintptr) (string, error)
	FindInterestingEntitiesByEntityID(ctx context.Context, entityID int64, flags int64) (string, error)
	FindInterestingEntitiesByRecordID(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error)
	FindNetworkByEntityID(ctx context.Context, entityList string, maxDegree int, buildOutDegree int, maxEntities int) (string, error)
	FindNetworkByEntityID_V2(ctx context.Context, entityList string, maxDegree int, buildOutDegree int, maxEntities int, flags int64) (string, error)
	FindNetworkByRecordID(ctx context.Context, recordList string, maxDegree int, buildOutDegree int, maxEntities int) (string, error)
	FindNetworkByRecordID_V2(ctx context.Context, recordList string, maxDegree int, buildOutDegree int, maxEntities int, flags int64) (string, error)
	FindPathByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int) (string, error)
	FindPathByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, flags int64) (string, error)
	FindPathByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int) (string, error)
	FindPathByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, flags int64) (string, error)
	FindPathExcludingByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string) (string, error)
	FindPathExcludingByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string, flags int64) (string, error)
	FindPathExcludingByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string) (string, error)
	FindPathExcludingByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string, flags int64) (string, error)
	FindPathIncludingSourceByEntityID(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string, requiredDsrcs string) (string, error)
	FindPathIncludingSourceByEntityID_V2(ctx context.Context, entityID1 int64, entityID2 int64, maxDegree int, excludedEntities string, requiredDsrcs string, flags int64) (string, error)
	FindPathIncludingSourceByRecordID(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string, requiredDsrcs string) (string, error)
	FindPathIncludingSourceByRecordID_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, maxDegree int, excludedRecords string, requiredDsrcs string, flags int64) (string, error)
	GetActiveConfigID(ctx context.Context) (int64, error)
	GetEntityByEntityID(ctx context.Context, entityID int64) (string, error)
	GetEntityByEntityID_V2(ctx context.Context, entityID int64, flags int64) (string, error)
	GetEntityByRecordID(ctx context.Context, dataSourceCode string, recordID string) (string, error)
	GetEntityByRecordID_V2(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error)
	GetLastException(ctx context.Context) (string, error)
	GetLastExceptionCode(ctx context.Context) (int, error)
	GetRecord(ctx context.Context, dataSourceCode string, recordID string) (string, error)
	GetRecord_V2(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error)
	GetRedoRecord(ctx context.Context) (string, error)
	GetRepositoryLastModifiedTime(ctx context.Context) (int64, error)
	GetVirtualEntityByRecordID(ctx context.Context, recordList string) (string, error)
	GetVirtualEntityByRecordID_V2(ctx context.Context, recordList string, flags int64) (string, error)
	HowEntityByEntityID(ctx context.Context, entityID int64) (string, error)
	HowEntityByEntityID_V2(ctx context.Context, entityID int64, flags int64) (string, error)
	Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error
	InitWithConfigID(ctx context.Context, moduleName string, iniParams string, initConfigID int64, verboseLogging int) error
	PrimeEngine(ctx context.Context) error
	Process(ctx context.Context, record string) error
	ProcessRedoRecord(ctx context.Context) (string, error)
	ProcessRedoRecordWithInfo(ctx context.Context, flags int64) (string, string, error)
	ProcessWithInfo(ctx context.Context, record string, flags int64) (string, error)
	ProcessWithResponse(ctx context.Context, record string) (string, error)
	ProcessWithResponseResize(ctx context.Context, record string) (string, error)
	PurgeRepository(ctx context.Context) error
	ReevaluateEntity(ctx context.Context, entityID int64, flags int64) error
	ReevaluateEntityWithInfo(ctx context.Context, entityID int64, flags int64) (string, error)
	ReevaluateRecord(ctx context.Context, dataSourceCode string, recordID string, flags int64) error
	ReevaluateRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error)
	Reinit(ctx context.Context, initConfigID int64) error
	ReplaceRecord(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string) error
	ReplaceRecordWithInfo(ctx context.Context, dataSourceCode string, recordID string, jsonData string, loadID string, flags int64) (string, error)
	SearchByAttributes(ctx context.Context, jsonData string) (string, error)
	SearchByAttributes_V2(ctx context.Context, jsonData string, flags int64) (string, error)
	Stats(ctx context.Context) (string, error)
	WhyEntities(ctx context.Context, entityID1 int64, entityID2 int64) (string, error)
	WhyEntities_V2(ctx context.Context, entityID1 int64, entityID2 int64, flags int64) (string, error)
	WhyEntityByEntityID(ctx context.Context, entityID int64) (string, error)
	WhyEntityByEntityID_V2(ctx context.Context, entityID int64, flags int64) (string, error)
	WhyEntityByRecordID(ctx context.Context, dataSourceCode string, recordID string) (string, error)
	WhyEntityByRecordID_V2(ctx context.Context, dataSourceCode string, recordID string, flags int64) (string, error)
	WhyRecords(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string) (string, error)
	WhyRecords_V2(ctx context.Context, dataSourceCode1 string, recordID1 string, dataSourceCode2 string, recordID2 string, flags int64) (string, error)
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const MessageIdFormat = "senzing-6004%04d"

const (
	G2_EXPORT_INCLUDE_RESOLVED         FlagMask = 0         // 0 we should include entities with "resolved" relationships
	G2_EXPORT_INCLUDE_POSSIBLY_SAME    FlagMask = 1 << iota // 1 we should include entities with "possibly same" relationships
	G2_EXPORT_INCLUDE_POSSIBLY_RELATED                      // 2 we should include entities with "possibly related" relationships
	G2_EXPORT_INCLUDE_NAME_ONLY                             // 3 we should include entities with "name only" relationships
	G2_EXPORT_INCLUDE_DISCLOSED                             // 4 we should include entities with "disclosed" relationships
	G2_EXPORT_INCLUDE_SINGLETONS                            // 5 we should include singleton entities

	/* flags for outputting entity relation data  */
	G2_ENTITY_INCLUDE_POSSIBLY_SAME_RELATIONS    // 6 include "possibly same" relationships on entities
	G2_ENTITY_INCLUDE_POSSIBLY_RELATED_RELATIONS // 7 include "possibly related" relationships on entities
	G2_ENTITY_INCLUDE_NAME_ONLY_RELATIONS        // 8 include "name only" relationships on entities
	G2_ENTITY_INCLUDE_DISCLOSED_RELATIONS        // 9 include "disclosed" relationships on entities

	/* flags for outputting entity feature data  */
	G2_ENTITY_INCLUDE_ALL_FEATURES            // 10 include all features for entities
	G2_ENTITY_INCLUDE_REPRESENTATIVE_FEATURES // 11 include only representative features on entities

	/* flags for getting extra information about an entity  */
	G2_ENTITY_INCLUDE_ENTITY_NAME            // 12 include the name of the entity
	G2_ENTITY_INCLUDE_RECORD_SUMMARY         // 13 include the record summary of the entity
	G2_ENTITY_INCLUDE_RECORD_DATA            // 14 include the basic record data for the entity
	G2_ENTITY_INCLUDE_RECORD_MATCHING_INFO   // 15 include the record matching info for the entity
	G2_ENTITY_INCLUDE_RECORD_JSON_DATA       // 16 include the record json data for the entity
	G2_ENTITY_INCLUDE_RECORD_FORMATTED_DATA  // 17 include the record formattted data for the entity
	G2_ENTITY_INCLUDE_RECORD_FEATURE_IDS     // 18 include the features identifiers for the records
	G2_ENTITY_INCLUDE_RELATED_ENTITY_NAME    // 19 include the name of the related entities
	G2_ENTITY_INCLUDE_RELATED_MATCHING_INFO  // 20 include the record matching info of the related entities
	G2_ENTITY_INCLUDE_RELATED_RECORD_SUMMARY // 21 include the record summary of the related entities
	G2_ENTITY_INCLUDE_RELATED_RECORD_DATA    // 22 include the basic record of the related entities

	/* flags for extra feature data  */
	G2_ENTITY_OPTION_INCLUDE_INTERNAL_FEATURES // 23 include internal features
	G2_ENTITY_OPTION_INCLUDE_FEATURE_STATS     // 24 include statistics on features

	/* flags for finding entity path data  */
	G2_FIND_PATH_PREFER_EXCLUDE // 25 excluded entities are still allowed, but not preferred

	/* flags for including search result feature scores  */
	G2_INCLUDE_FEATURE_SCORES // 26 include feature scores
	G2_SEARCH_INCLUDE_STATS   //  27 include statistics from search results
)

const (
	/* flags for exporting entity data  */
	G2_EXPORT_INCLUDE_ALL_ENTITIES      = G2_EXPORT_INCLUDE_RESOLVED | G2_EXPORT_INCLUDE_SINGLETONS
	G2_EXPORT_INCLUDE_ALL_RELATIONSHIPS = G2_EXPORT_INCLUDE_POSSIBLY_SAME | G2_EXPORT_INCLUDE_POSSIBLY_RELATED | G2_EXPORT_INCLUDE_NAME_ONLY | G2_EXPORT_INCLUDE_DISCLOSED

	/* flags for outputting entity relation data  */
	G2_ENTITY_INCLUDE_ALL_RELATIONS  = G2_ENTITY_INCLUDE_POSSIBLY_SAME_RELATIONS | G2_ENTITY_INCLUDE_POSSIBLY_RELATED_RELATIONS | G2_ENTITY_INCLUDE_NAME_ONLY_RELATIONS | G2_ENTITY_INCLUDE_DISCLOSED_RELATIONS
	G2_SEARCH_INCLUDE_FEATURE_SCORES = G2_INCLUDE_FEATURE_SCORES // include feature scores from search results

	/* flags for searching for entities  */
	G2_SEARCH_INCLUDE_RESOLVED         = G2_EXPORT_INCLUDE_RESOLVED
	G2_SEARCH_INCLUDE_POSSIBLY_SAME    = G2_EXPORT_INCLUDE_POSSIBLY_SAME
	G2_SEARCH_INCLUDE_POSSIBLY_RELATED = G2_EXPORT_INCLUDE_POSSIBLY_RELATED
	G2_SEARCH_INCLUDE_NAME_ONLY        = G2_EXPORT_INCLUDE_NAME_ONLY
	G2_SEARCH_INCLUDE_ALL_ENTITIES     = G2_SEARCH_INCLUDE_RESOLVED | G2_SEARCH_INCLUDE_POSSIBLY_SAME | G2_SEARCH_INCLUDE_POSSIBLY_RELATED | G2_SEARCH_INCLUDE_NAME_ONLY

	/* recommended settings */
	G2_RECORD_DEFAULT_FLAGS       = G2_ENTITY_INCLUDE_RECORD_JSON_DATA                                                                                                                                                                                                                                                                                                                   // the recommended default flag values for getting records
	G2_ENTITY_DEFAULT_FLAGS       = G2_ENTITY_INCLUDE_ALL_RELATIONS | G2_ENTITY_INCLUDE_REPRESENTATIVE_FEATURES | G2_ENTITY_INCLUDE_ENTITY_NAME | G2_ENTITY_INCLUDE_RECORD_SUMMARY | G2_ENTITY_INCLUDE_RECORD_DATA | G2_ENTITY_INCLUDE_RECORD_MATCHING_INFO | G2_ENTITY_INCLUDE_RELATED_ENTITY_NAME | G2_ENTITY_INCLUDE_RELATED_RECORD_SUMMARY | G2_ENTITY_INCLUDE_RELATED_MATCHING_INFO // the recommended default flag values for getting entities
	G2_ENTITY_BRIEF_DEFAULT_FLAGS = G2_ENTITY_INCLUDE_RECORD_MATCHING_INFO | G2_ENTITY_INCLUDE_ALL_RELATIONS | G2_ENTITY_INCLUDE_RELATED_MATCHING_INFO                                                                                                                                                                                                                                   // the recommended default flag values for a brief entity result
	G2_EXPORT_DEFAULT_FLAGS       = G2_EXPORT_INCLUDE_ALL_ENTITIES | G2_EXPORT_INCLUDE_ALL_RELATIONSHIPS | G2_ENTITY_INCLUDE_ALL_RELATIONS | G2_ENTITY_INCLUDE_REPRESENTATIVE_FEATURES | G2_ENTITY_INCLUDE_ENTITY_NAME | G2_ENTITY_INCLUDE_RECORD_DATA | G2_ENTITY_INCLUDE_RECORD_MATCHING_INFO | G2_ENTITY_INCLUDE_RELATED_MATCHING_INFO                                                // the recommended default flag values for exporting entities
	G2_FIND_PATH_DEFAULT_FLAGS    = G2_ENTITY_INCLUDE_ALL_RELATIONS | G2_ENTITY_INCLUDE_ENTITY_NAME | G2_ENTITY_INCLUDE_RECORD_SUMMARY | G2_ENTITY_INCLUDE_RELATED_MATCHING_INFO                                                                                                                                                                                                         // the recommended default flag values for finding entity paths
	G2_WHY_ENTITY_DEFAULT_FLAGS   = G2_ENTITY_DEFAULT_FLAGS | G2_ENTITY_INCLUDE_RECORD_FEATURE_IDS | G2_ENTITY_OPTION_INCLUDE_INTERNAL_FEATURES | G2_ENTITY_OPTION_INCLUDE_FEATURE_STATS | G2_INCLUDE_FEATURE_SCORES                                                                                                                                                                     // the recommended default flag values for why-analysis on entities
	G2_HOW_ENTITY_DEFAULT_FLAGS   = G2_ENTITY_DEFAULT_FLAGS | G2_ENTITY_INCLUDE_RECORD_FEATURE_IDS | G2_ENTITY_OPTION_INCLUDE_INTERNAL_FEATURES | G2_ENTITY_OPTION_INCLUDE_FEATURE_STATS | G2_INCLUDE_FEATURE_SCORES                                                                                                                                                                     // the recommended default flag values for how-analysis on entities

	G2_SEARCH_BY_ATTRIBUTES_ALL            = G2_SEARCH_INCLUDE_ALL_ENTITIES | G2_ENTITY_INCLUDE_REPRESENTATIVE_FEATURES | G2_ENTITY_INCLUDE_ENTITY_NAME | G2_ENTITY_INCLUDE_RECORD_SUMMARY | G2_SEARCH_INCLUDE_FEATURE_SCORES                               // the recommended flag values for searching by attributes, returning all matching entities
	G2_SEARCH_BY_ATTRIBUTES_STRONG         = G2_SEARCH_INCLUDE_RESOLVED | G2_SEARCH_INCLUDE_POSSIBLY_SAME | G2_ENTITY_INCLUDE_REPRESENTATIVE_FEATURES | G2_ENTITY_INCLUDE_ENTITY_NAME | G2_ENTITY_INCLUDE_RECORD_SUMMARY | G2_SEARCH_INCLUDE_FEATURE_SCORES // the recommended flag values for searching by attributes, returning only strongly matching entities
	G2_SEARCH_BY_ATTRIBUTES_MINIMAL_ALL    = G2_SEARCH_INCLUDE_ALL_ENTITIES                                                                                                                                                                                 // return minimal data with all matches
	G2_SEARCH_BY_ATTRIBUTES_MINIMAL_STRONG = G2_SEARCH_INCLUDE_RESOLVED | G2_SEARCH_INCLUDE_POSSIBLY_SAME                                                                                                                                                   // return minimal data with only the strongest matches
	G2_SEARCH_BY_ATTRIBUTES_DEFAULT_FLAGS  = G2_SEARCH_BY_ATTRIBUTES_ALL                                                                                                                                                                                    // the recommended default flag values for search-by-attributes
)

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var Messages = map[int]string{
	1:    "Call to G2_addRecord(%s, %s, %s, %s) failed. Return code: %d",
	2:    "Call to G2_addRecordWithInfo(%s, %s, %s, %s, %d) failed. Return code: %d",
	3:    "Call to G2_addRecordWithInfoWithReturnedRecordID(%s, %s, %s, %d) failed. Return values: %s; %s; Return code: %d",
	4:    "Call to G2_addRecordWithReturnedRecordID(%s, %s, %s) failed. Return code: %d",
	5:    "Call to G2_checkRecord(%s, %s) failed. Return code: %d",
	6:    "Call to G2_closeExport() failed. Return code: %d",
	7:    "Call to G2_deleteRecord(%s, %s, %s) failed. Return code: %d",
	8:    "Call to G2_deleteRecordWithInfo(%s, %s, %s, %d) failed. Return code: %d",
	9:    "Call to G2_destroy() failed. Return code: %d",
	10:   "Call to G2_exportConfigAndConfigID() failed. Return values: %s; %s; Return code: %d",
	11:   "Call to G2_exportConfig() failed. Return code: %d",
	12:   "Call to G2_exportCSVEntityReport(%s, %d) failed. Return code: %d",
	13:   "Call to G2_exportJSONEntityReport(%d) failed. Return code: %d",
	14:   "Call to G2_fetchNext() failed. Return code: %d",
	15:   "Call to G2_findInterestingEntitiesByEntityID(%d, %d) failed. Return code: %d",
	16:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	17:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	18:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	19:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	20:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	21:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	22:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	23:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	24:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	25:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	26:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	27:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	28:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	29:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	30:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	31:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	32:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	33:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	34:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	35:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	36:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	37:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	38:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	39:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	40:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	41:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	42:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	43:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	44:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	45:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	46:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	47:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	48:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	49:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	50:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	51:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	52:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	53:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	54:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	55:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	56:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	57:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	58:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	59:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	60:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	61:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	62:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	63:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	64:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	65:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	66:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	67:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	68:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	69:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	70:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	71:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	72:   "Call to G2Engine_xxxxx(%d) failed. Return code: %d",
	2999: "Cannot retrieve last error message.",
}
