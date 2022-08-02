package g2engine

import (
	"context"
	"github.com/docktermj/go-xyzzy-helpers/g2configuration"
	"github.com/docktermj/go-xyzzy-helpers/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	g2engine G2engine
)

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context) G2engine {

	if g2engine == nil {
		g2engine = &G2engineImpl{}

		moduleName := "Test module name"
		verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
		iniParams, jsonErr := g2configuration.BuildSimpleSystemConfigurationJson("")
		if jsonErr != nil {
			logger.Fatalf("Cannot construct system configuration: %v", jsonErr)
		}

		initErr := g2engine.Init(ctx, moduleName, iniParams, verboseLogging)
		if initErr != nil {
			logger.Fatalf("Cannot Init: %v", initErr)
		}
	}
	return g2engine
}

func testError(test *testing.T, ctx context.Context, g2engine G2engine, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
		lastException, _ := g2engine.GetLastException(ctx)
		assert.FailNow(test, lastException)
	}
}

// ----------------------------------------------------------------------------
// Test harness
// ----------------------------------------------------------------------------

func TestGetObject(test *testing.T) {
	ctx := context.TODO()
	getTestObject(ctx)
}

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

func TestAddRecord(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx)

	dataSourceCode := "TEST"
	recordID := "987654321"
	jsonData := `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "987654321", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	loadID := "TEST"

	err := g2engine.AddRecord(ctx, dataSourceCode, recordID, jsonData, loadID)
	testError(test, ctx, g2engine, err)
}

func TestAddRecordWithInfo(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx)

	dataSourceCode := "TEST"
	recordID := "987654321"
	jsonData := `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "987654321", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	loadID := "TEST"
	var flags int64 = 0

	actual, err := g2engine.AddRecordWithInfo(ctx, dataSourceCode, recordID, jsonData, loadID, flags)
	testError(test, ctx, g2engine, err)
	test.Log("Actual:", actual)
}

func TestAddRecordWithInfoWithReturnedRecordID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx)

	dataSourceCode := "TEST"
	jsonData := `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	loadID := "TEST"
	var flags int64 = 0

	actual, actualRecordID, err := g2engine.AddRecordWithInfoWithReturnedRecordID(ctx, dataSourceCode, jsonData, loadID, flags)
	testError(test, ctx, g2engine, err)
	test.Log("Actual RecordID:", actualRecordID)
	test.Log("Actual:", actual)
}

func TestAddRecordWithReturnedRecordID(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx)

	dataSourceCode := "TEST"
	jsonData := `{"SOCIAL_HANDLE": "bobby", "DATE_OF_BIRTH": "1/2/1983", "ADDR_STATE": "WI", "ADDR_POSTAL_CODE": "54434", "SSN_NUMBER": "987-65-4321", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "Smith", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	loadID := "TEST"

	actual, err := g2engine.AddRecordWithReturnedRecordID(ctx, dataSourceCode, jsonData, loadID)
	testError(test, ctx, g2engine, err)
	test.Log("Actual:", actual)
}

func TestCheckRecord(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx)

	record := `{"DATA_SOURCE": "TEST", "NAMES": [{"NAME_TYPE": "PRIMARY", "NAME_LAST": "Smith", "NAME_MIDDLE": "M" }], "PASSPORT_NUMBER": "PP11111", "PASSPORT_COUNTRY": "US", "DRIVERS_LICENSE_NUMBER": "DL11111", "SSN_NUMBER": "111-11-1111"}`
	recordQueryList := `{"RECORDS": [{"DATA_SOURCE": "TEST","RECORD_ID": "987654321"},{"DATA_SOURCE": "TEST","RECORD_ID": "123456789"}]}`

	actual, err := g2engine.CheckRecord(ctx, record, recordQueryList)
	testError(test, ctx, g2engine, err)
	test.Log("Actual:", actual)
}

func TestClearLastException(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx)
	err := g2engine.ClearLastException(ctx)
	testError(test, ctx, g2engine, err)
}

// FAIL:
//func TestExport(test *testing.T) {
//	ctx := context.TODO()
//	g2diagnostic := getTestObject(ctx)
//	aSize := 10
//
//	aHandle, err := g2diagnostic.GetEntityListBySize(ctx, aSize)
//	testError(test, ctx, g2diagnostic, err)
//
//	anEntity, err := g2diagnostic.FetchNextEntityBySize(ctx, aHandle)
//	testError(test, ctx, g2diagnostic, err)
//	test.Log("Entity:", anEntity)
//
//	err = g2diagnostic.CloseExport(ctx, aHandle)
//	testError(test, ctx, g2diagnostic, err)
//}

func TestCountRedoRecords(test *testing.T) {
	ctx := context.TODO()
	g2engine := getTestObject(ctx)

	actual, err := g2engine.CountRedoRecords(ctx)
	testError(test, ctx, g2engine, err)
	test.Log("Actual:", actual)
}

//
//func TestDeleteRecord(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//
//	dataSourceCode := "TEST"
//	recordID := "987654321"
//	loadID := "TEST"
//
//	err := g2engine.DeleteRecord(ctx, dataSourceCode, recordID, loadID)
//	testError(test, ctx, g2engine, err)
//}
//
//func TestDeleteRecordWithInfo(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//
//	dataSourceCode := "TEST"
//	recordID := "987654321"
//	loadID := "TEST"
//	var flags int64 = 0
//
//	actual, err := g2engine.DeleteRecordWithInfo(ctx, dataSourceCode, recordID, loadID, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestDestroy(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	err := g2engine.Destroy(ctx)
//	testError(test, ctx, g2engine, err)
//}
//
//func TestExportConfigAndConfigID(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	actualConfig, actualConfigId, err := g2engine.ExportConfigAndConfigID(ctx)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual Config:", actualConfig)
//	test.Log("Actual Config ID:", actualConfigId)
//}
//
//func TestExportConfig(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	actual, err := g2engine.ExportConfig(ctx)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestExportCSVEntityReport(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	csvColumnList := ""
//	var flags int64 = 0
//	actual, err := g2engine.ExportCSVEntityReport(ctx, csvColumnList, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestExportJSONEntityReport(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	var flags int64 = 0
//	actual, err := g2engine.ExportJSONEntityReport(ctx, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestFetchNext(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	responseHandle := ""
//	actual, err := g2engine.FetchNext(ctx, responseHandle)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestFindInterestingEntitiesByEntityID(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	var entityID int64 = 1
//	var flags int64 = 0
//	actual, err := g2engine.FindInterestingEntitiesByEntityID(ctx, entityID, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestFindInterestingEntitiesByRecordID(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	dataSourceCode := "TEST"
//	recordID := "987654321"
//	var flags int64 = 0
//	actual, err := g2engine.FindInterestingEntitiesByRecordID(ctx, dataSourceCode, recordID, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestFindNetworkByEntityID(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	entityList := ""
//	maxDegree := 1
//	buildOutDegree := 2
//	maxEntities := 10
//	actual, err := g2engine.FindNetworkByEntityID(ctx, entityList, maxDegree, buildOutDegree, maxEntities)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestFindNetworkByEntityID_V2(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	entityList := ""
//	maxDegree := 1
//	buildOutDegree := 2
//	maxEntities := 10
//	var flags int64 = 0
//	actual, err := g2engine.FindNetworkByEntityID_V2(ctx, entityList, maxDegree, buildOutDegree, maxEntities, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestFindNetworkByRecordID(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	recordList := ""
//	maxDegree := 1
//	buildOutDegree := 2
//	maxEntities := 10
//	actual, err := g2engine.FindNetworkByRecordID(ctx, recordList, maxDegree, buildOutDegree, maxEntities)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestFindNetworkByRecordID_V2(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	recordList := ""
//	maxDegree := 1
//	buildOutDegree := 2
//	maxEntities := 10
//	var flags int64 = 0
//	actual, err := g2engine.FindNetworkByRecordID_V2(ctx, recordList, maxDegree, buildOutDegree, maxEntities, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestFindPathByEntityID(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	var entityID1 int64 = 1
//	var entityID2 int64 = 2
//	maxDegree := 1
//	actual, err := g2engine.FindPathByEntityID(ctx, entityID1, entityID2, maxDegree)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestFindPathByEntityID_V2(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	var entityID1 int64 = 1
//	var entityID2 int64 = 2
//	maxDegree := 1
//	var flags int64 = 0
//	actual, err := g2engine.FindPathByEntityID_V2(ctx, entityID1, entityID2, maxDegree, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestFindPathByRecordID(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	dataSourceCode1 := "TEST"
//	recordID1 := ""
//	dataSourceCode2 := "TEST"
//	recordID2 := ""
//	maxDegree := 1
//	actual, err := g2engine.FindPathByRecordID(ctx, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestFindPathByRecordID_V2(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	dataSourceCode1 := "TEST"
//	recordID1 := ""
//	dataSourceCode2 := "TEST"
//	recordID2 := ""
//	maxDegree := 1
//	var flags int64 = 0
//	actual, err := g2engine.FindPathByRecordID_V2(ctx, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestFindPathExcludingByEntityID(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	var entityID1 int64 = 1
//	var entityID2 int64 = 2
//	maxDegree := 1
//	excludedEntities := ""
//	actual, err := g2engine.FindPathExcludingByEntityID(ctx, entityID1, entityID2, maxDegree, excludedEntities)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestFindPathExcludingByEntityID_V2(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	var entityID1 int64 = 1
//	var entityID2 int64 = 2
//	maxDegree := 1
//	excludedEntities := ""
//	var flags int64 = 0
//	actual, err := g2engine.FindPathExcludingByEntityID_V2(ctx, entityID1, entityID2, maxDegree, excludedEntities, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestFindPathExcludingByRecordID(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	dataSourceCode1 := "TEST"
//	recordID1 := ""
//	dataSourceCode2 := "TEST"
//	recordID2 := ""
//	maxDegree := 1
//	excludedRecords := ""
//	actual, err := g2engine.FindPathExcludingByRecordID(ctx, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestFindPathExcludingByRecordID_V2(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	dataSourceCode1 := "TEST"
//	recordID1 := ""
//	dataSourceCode2 := "TEST"
//	recordID2 := ""
//	maxDegree := 1
//	excludedRecords := ""
//	var flags int64 = 0
//	actual, err := g2engine.FindPathExcludingByRecordID_V2(ctx, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestFindPathIncludingSourceByEntityID(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	var entityID1 int64 = 1
//	var entityID2 int64 = 2
//	maxDegree := 1
//	excludedEntities := ""
//	actual, err := g2engine.FindPathIncludingSourceByEntityID(ctx, entityID1, entityID2, maxDegree, excludedEntities)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestFindPathIncludingSourceByEntityID_V2(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	var entityID1 int64 = 1
//	var entityID2 int64 = 2
//	maxDegree := 1
//	excludedEntities := ""
//	requiredDsrcs := ""
//	var flags int64 = 0
//	actual, err := g2engine.FindPathIncludingSourceByEntityID_V2(ctx, entityID1, entityID2, maxDegree, excludedEntities, requiredDsrcs, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestFindPathIncludingSourceByRecordID(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	dataSourceCode1 := "TEST"
//	recordID1 := ""
//	dataSourceCode2 := "TEST"
//	recordID2 := ""
//	maxDegree := 1
//	excludedRecords := ""
//	requiredDsrcs := ""
//	actual, err := g2engine.FindPathIncludingSourceByRecordID(ctx, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, requiredDsrcs)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestFindPathIncludingSourceByRecordID_V2(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	dataSourceCode1 := "TEST"
//	recordID1 := ""
//	dataSourceCode2 := "TEST"
//	recordID2 := ""
//	maxDegree := 1
//	excludedRecords := ""
//	requiredDsrcs := ""
//	var flags int64 = 0
//	actual, err := g2engine.FindPathIncludingSourceByRecordID_V2(ctx, dataSourceCode1, recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords, requiredDsrcs, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestGetActiveConfigID(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	actual, err := g2engine.GetActiveConfigID(ctx)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}

//func TestGetEntityByEntityID(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	var entityID int64 = 1
//	actual, err := g2engine.GetEntityByEntityID(ctx, entityID)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestGetEntityByEntityID_V2(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	var entityID int64 = 1
//	var flags int64 = 0
//	actual, err := g2engine.GetEntityByEntityID_V2(ctx, entityID, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestGetEntityByRecordID(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	dataSourceCode := "TEST"
//	recordID := ""
//	actual, err := g2engine.GetEntityByRecordID(ctx, dataSourceCode, recordID)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestGetEntityByRecordID_V2(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	dataSourceCode := "TEST"
//	recordID := ""
//	var flags int64 = 0
//	actual, err := g2engine.GetEntityByRecordID_V2(ctx, dataSourceCode, recordID, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestGetLastException(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	actual, err := g2engine.GetLastException(ctx)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestGetLastExceptionCode(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	actual, err := g2engine.GetLastExceptionCode(ctx)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestGetRecord(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	dataSourceCode := "TEST"
//	recordID := ""
//	actual, err := g2engine.GetRecord(ctx, dataSourceCode, recordID)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestGetRecord_V2(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	dataSourceCode := "TEST"
//	recordID := ""
//	var flags int64 = 0
//	actual, err := g2engine.GetRecord_V2(ctx, dataSourceCode, recordID, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestGetRedoRecord(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	actual, err := g2engine.GetRedoRecord(ctx)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestGetRepositoryLastModifiedTime(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	actual, err := g2engine.GetRepositoryLastModifiedTime(ctx)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestGetVirtualEntityByRecordID(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	recordList := ""
//	actual, err := g2engine.GetVirtualEntityByRecordID(ctx, recordList)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestGetVirtualEntityByRecordID_V2(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	recordList := ""
//	var flags int64 = 0
//	actual, err := g2engine.GetVirtualEntityByRecordID_V2(ctx, recordList, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestHowEntityByEntityID(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	var entityID int64 = 1
//	actual, err := g2engine.HowEntityByEntityID(ctx, entityID)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestHowEntityByEntityID_V2(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	var entityID int64 = 1
//	var flags int64 = 0
//	actual, err := g2engine.HowEntityByEntityID_V2(ctx, entityID, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestInit(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	moduleName := "Test module name"
//	verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
//	iniParams, jsonErr := g2configuration.BuildSimpleSystemConfigurationJson("")
//	if jsonErr != nil {
//		logger.Fatalf("Cannot construct system configuration: %v", jsonErr)
//	}
//	err := g2engine.Init(ctx, moduleName, iniParams, verboseLogging)
//	testError(test, ctx, g2engine, err)
//}
//
//func TestInitWithConfigID(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	moduleName := "Test module name"
//	var initConfigID int64 = 1
//	verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
//	iniParams, jsonErr := g2configuration.BuildSimpleSystemConfigurationJson("")
//	if jsonErr != nil {
//		logger.Fatalf("Cannot construct system configuration: %v", jsonErr)
//	}
//	err := g2engine.InitWithConfigID(ctx, moduleName, iniParams, initConfigID, verboseLogging)
//	testError(test, ctx, g2engine, err)
//}
//
//func TestPrimeEngine(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	err := g2engine.PrimeEngine(ctx)
//	testError(test, ctx, g2engine, err)
//}
//
//func TestProcess(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	record := ""
//	err := g2engine.Process(ctx, record)
//	testError(test, ctx, g2engine, err)
//}
//
//func TestProcessRedoRecord(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	actual, err := g2engine.ProcessRedoRecord(ctx)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestProcessRedoRecordWithInfo(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	var flags int64 = 0
//	actualInfo, actual, err := g2engine.ProcessRedoRecordWithInfo(ctx, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual Info:", actualInfo)
//	test.Log("Actual:", actual)
//}
//
//func TestProcessWithInfo(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	record := ""
//	var flags int64 = 0
//	actual, err := g2engine.ProcessWithInfo(ctx, record, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestProcessWithResponse(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	record := ""
//	actual, err := g2engine.ProcessWithResponse(ctx, record)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestProcessWithResponseResize(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	record := ""
//	actual, err := g2engine.ProcessWithResponseResize(ctx, record)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestPurgeRepository(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	err := g2engine.PurgeRepository(ctx)
//	testError(test, ctx, g2engine, err)
//}
//
//func TestReevaluateEntity(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	var entityID int64 = 1
//	var flags int64 = 0
//	err := g2engine.ReevaluateEntity(ctx, entityID, flags)
//	testError(test, ctx, g2engine, err)
//}
//
//func TestReevaluateEntityWithInfo(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	var entityID int64 = 1
//	var flags int64 = 0
//	actual, err := g2engine.ReevaluateEntityWithInfo(ctx, entityID, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestReevaluateRecord(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	dataSourceCode := "TEST"
//	recordID := "987654321"
//	var flags int64 = 0
//	err := g2engine.ReevaluateRecord(ctx, dataSourceCode, recordID, flags)
//	testError(test, ctx, g2engine, err)
//}
//
//func TestReevaluateRecordWithInfo(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	dataSourceCode := "TEST"
//	recordID := "987654321"
//	var flags int64 = 0
//	actual, err := g2engine.ReevaluateRecordWithInfo(ctx, dataSourceCode, recordID, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestReinit(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	var initConfigID int64 = 0
//	err := g2engine.Reinit(ctx, initConfigID)
//	testError(test, ctx, g2engine, err)
//}
//
//func TestReplaceRecord(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	dataSourceCode := "TEST"
//	recordID := "987654321"
//	jsonData := ""
//	loadID := ""
//	err := g2engine.ReplaceRecord(ctx, dataSourceCode, recordID, jsonData, loadID)
//	testError(test, ctx, g2engine, err)
//}
//
//func TestReplaceRecordWithInfo(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	dataSourceCode := "TEST"
//	recordID := "987654321"
//	jsonData := ""
//	loadID := ""
//	var flags int64 = 0
//	actual, err := g2engine.ReplaceRecordWithInfo(ctx, dataSourceCode, recordID, jsonData, loadID, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestSearchByAttributes(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	jsonData := ""
//	actual, err := g2engine.SearchByAttributes(ctx, jsonData)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestSearchByAttributes_V2(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	jsonData := ""
//	var flags int64 = 0
//	actual, err := g2engine.SearchByAttributes_V2(ctx, jsonData, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestStats(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	actual, err := g2engine.Stats(ctx)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestWhyEntities(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	var entityID1 int64 = 1
//	var entityID2 int64 = 2
//	actual, err := g2engine.WhyEntities(ctx, entityID1, entityID2)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestWhyEntities_V2(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	var entityID1 int64 = 1
//	var entityID2 int64 = 2
//	var flags int64 = 0
//	actual, err := g2engine.WhyEntities_V2(ctx, entityID1, entityID2, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestWhyEntityByEntityID(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	var entityID int64 = 1
//	actual, err := g2engine.WhyEntityByEntityID(ctx, entityID)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestWhyEntityByEntityID_V2(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	var entityID int64 = 1
//	var flags int64 = 0
//	actual, err := g2engine.WhyEntityByEntityID_V2(ctx, entityID, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestWhyEntityByRecordID(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	dataSourceCode := "TEST"
//	recordID := "987654321"
//	actual, err := g2engine.WhyEntityByRecordID(ctx, dataSourceCode, recordID)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestWhyEntityByRecordID_V2(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	dataSourceCode := "TEST"
//	recordID := "987654321"
//	var flags int64 = 0
//	actual, err := g2engine.WhyEntityByRecordID_V2(ctx, dataSourceCode, recordID, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestWhyRecords(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	dataSourceCode1 := "TEST"
//	recordID1 := "987654321"
//	dataSourceCode2 := "TEST"
//	recordID2 := ""
//	actual, err := g2engine.WhyRecords(ctx, dataSourceCode1, recordID1, dataSourceCode2, recordID2)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
//
//func TestWhyRecords_V2(test *testing.T) {
//	ctx := context.TODO()
//	g2engine := getTestObject(ctx)
//	dataSourceCode1 := "TEST"
//	recordID1 := "987654321"
//	dataSourceCode2 := "TEST"
//	recordID2 := ""
//	var flags int64 = 0
//	actual, err := g2engine.WhyRecords_V2(ctx, dataSourceCode1, recordID1, dataSourceCode2, recordID2, flags)
//	testError(test, ctx, g2engine, err)
//	test.Log("Actual:", actual)
//}
