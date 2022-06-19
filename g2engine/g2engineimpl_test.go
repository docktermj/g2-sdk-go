package g2engine

import (
	"context"
	"testing"

	"github.com/docktermj/xyzzygoapi/g2helper"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context) (G2engine, error) {
	var err error = nil
	g2engine := G2engineImpl{}

	moduleName := "Test module name"
	verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
	iniParams, jsonErr := g2helper.BuildSimpleSystemConfigurationJson()
	if jsonErr != nil {
		return &g2engine, jsonErr
	}

	err = g2engine.Init(ctx, moduleName, iniParams, verboseLogging)
	return &g2engine, err
}

func testError(test *testing.T, ctx context.Context, g2engine G2engine, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
		lastException, _ := g2engine.GetLastException(ctx)
		assert.FailNow(test, lastException)
	}
}

func TestGetObject(test *testing.T) {
	ctx := context.TODO()
	g2engine, err := getTestObject(ctx)
	testError(test, ctx, g2engine, err)
}

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

func TestAddRecord(test *testing.T) {
	ctx := context.TODO()
	g2engine, _ := getTestObject(ctx)

	dataSourceCode := "TEST"
	recordID := "987654321"
	jsonData := `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "987654321", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	loadID := "TEST"

	err := g2engine.AddRecord(ctx, dataSourceCode, recordID, jsonData, loadID)
	testError(test, ctx, g2engine, err)
}

func TestAddRecordWithInfo(test *testing.T) {
	ctx := context.TODO()
	g2engine, _ := getTestObject(ctx)

	dataSourceCode := "TEST"
	recordID := "987654321"
	jsonData := `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "987654321", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
	loadID := "TEST"
	var flags int64 = 0

	actual, err := g2engine.AddRecordWithInfo(ctx, dataSourceCode, recordID, jsonData, loadID, flags)
	testError(test, ctx, g2engine, err)
	test.Log("Actual:", actual)
}

func TestDeleteRecord(test *testing.T) {
	ctx := context.TODO()
	g2engine, _ := getTestObject(ctx)

	dataSourceCode := "TEST"
	recordID := "987654321"
	loadID := "TEST"

	err := g2engine.DeleteRecord(ctx, dataSourceCode, recordID, loadID)
	testError(test, ctx, g2engine, err)
}

func TestDeleteRecordWithInfo(test *testing.T) {
	ctx := context.TODO()
	g2engine, _ := getTestObject(ctx)

	dataSourceCode := "TEST"
	recordID := "987654321"
	loadID := "TEST"
	var flags int64 = 0

	actual, err := g2engine.DeleteRecordWithInfo(ctx, dataSourceCode, recordID, loadID, flags)
	testError(test, ctx, g2engine, err)
	test.Log("Actual:", actual)
}

func TestDestroy(test *testing.T) {
	ctx := context.TODO()
	g2engine, _ := getTestObject(ctx)
	err := g2engine.Destroy(ctx)
	testError(test, ctx, g2engine, err)
}

func TestStats(test *testing.T) {
	ctx := context.TODO()
	g2engine, _ := getTestObject(ctx)
	actual, err := g2engine.Stats(ctx)
	testError(test, ctx, g2engine, err)
	test.Log("Actual:", actual)
}
