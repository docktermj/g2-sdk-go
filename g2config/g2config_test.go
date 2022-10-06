package g2config

import (
	"context"
	"github.com/docktermj/go-xyzzy-helpers/g2configuration"
	"github.com/docktermj/go-xyzzy-helpers/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	g2config G2config
)

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context) G2config {

	if g2config == nil {
		g2config = &G2configImpl{}

		moduleName := "Test module name"
		verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
		iniParams, jsonErr := g2configuration.BuildSimpleSystemConfigurationJson("")
		if jsonErr != nil {
			logger.Fatalf("Cannot construct system configuration: %v", jsonErr)
		}

		initErr := g2config.Init(ctx, moduleName, iniParams, verboseLogging)
		if initErr != nil {
			logger.Fatalf("Cannot Init: %v", initErr)
		}
	}
	return g2config
}

func testError(test *testing.T, ctx context.Context, g2config G2config, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
		lastException, _ := g2config.GetLastException(ctx)
		assert.FailNow(test, lastException)
	}
}

// ----------------------------------------------------------------------------
// Test harness
// ----------------------------------------------------------------------------

func TestBuildSimpleSystemConfigurationJson(test *testing.T) {
	actual, err := g2configuration.BuildSimpleSystemConfigurationJson("")
	if err != nil {
		test.Log("Error:", err.Error())
		assert.FailNow(test, actual)
	}
}

func TestGetObject(test *testing.T) {
	ctx := context.TODO()
	getTestObject(ctx)
}

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

//func TestAddDataSource(test *testing.T) {
//	ctx := context.TODO()
//	g2config := getTestObject(ctx)
//	configHandle := "TEST"
//	inputJson := `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "111", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`
//	err := g2config.AddDataSource(ctx, configHandle, inputJson)
//	testError(test, ctx, g2config, err)
//}

func TestClearLastException(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx)
	err := g2config.ClearLastException(ctx)
	testError(test, ctx, g2config, err)
}

func TestDestroy(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx)
	err := g2config.Destroy(ctx)
	testError(test, ctx, g2config, err)
}
