package g2diagnostic

import (
	"context"
	"fmt"
	"log"
	"testing"

	truncator "github.com/aquilax/truncate"
	"github.com/docktermj/go-xyzzy-helpers/g2configuration"
	"github.com/docktermj/go-xyzzy-helpers/logger"
	"github.com/stretchr/testify/assert"
)

var (
	g2diagnostic G2diagnostic
)

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context) G2diagnostic {
	if g2diagnostic == nil {
		g2diagnostic = &G2diagnosticImpl{}

		moduleName := "Test module name"
		verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
		iniParams, jsonErr := g2configuration.BuildSimpleSystemConfigurationJson("")
		if jsonErr != nil {
			logger.Fatalf("Cannot construct system configuration: %v", jsonErr)
		}

		initErr := g2diagnostic.Init(ctx, moduleName, iniParams, verboseLogging)
		if initErr != nil {
			logger.Fatalf("Cannot Init: %v", initErr)
		}
	}
	return g2diagnostic
}

func truncate(aString string) string {
	return truncator.Truncate(aString, 30, "...", truncator.PositionEnd)
}

func printResult(test *testing.T, title string, result interface{}) {
	if 1 == 0 {
		test.Logf("%s: %v", title, truncate(fmt.Sprintf("%v", result)))
	}
}

func printActual(test *testing.T, actual interface{}) {
	printResult(test, "Actual", actual)
}

func testError(test *testing.T, ctx context.Context, g2diagnostic G2diagnostic, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
		lastException, _ := g2diagnostic.GetLastException(ctx)
		assert.FailNow(test, lastException)
	}
}

// ----------------------------------------------------------------------------
// Test harness
// ----------------------------------------------------------------------------

// Reference: https://medium.com/nerd-for-tech/setup-and-teardown-unit-test-in-go-bd6fa1b785cd
func setupSuite(test testing.TB) func(test testing.TB) {
	test.Log("setup suite")

	// Return a function to teardown the test
	return func(test testing.TB) {
		test.Log("teardown suite")
	}
}

func TestBuildSimpleSystemConfigurationJson(test *testing.T) {
	actual, err := g2configuration.BuildSimpleSystemConfigurationJson("")
	if err != nil {
		test.Log("Error:", err.Error())
		assert.FailNow(test, actual)
	}
	printActual(test, actual)
}

func TestGetObject(test *testing.T) {
	ctx := context.TODO()
	getTestObject(ctx)
}

func TestLogger(test *testing.T) {
	// Configure the "log" standard library.

	log.SetFlags(log.Llongfile | log.Ldate | log.Lmicroseconds | log.LUTC)
	logger.SetLevel(logger.LevelInfo)

	// Test logger.

	logger.LogMessage(MessageIdFormat, 99, "Test message 1", "Variable1", "Variable2")
}

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

func TestCheckDBPerf(test *testing.T) {
	ctx := context.TODO()
	teardownSuite := setupSuite(test)
	defer teardownSuite(test)
	g2diagnostic := getTestObject(ctx)
	secondsToRun := 1
	actual, err := g2diagnostic.CheckDBPerf(ctx, secondsToRun)
	testError(test, ctx, g2diagnostic, err)
	printActual(test, actual)
}

func TestClearLastException(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx)
	g2diagnostic.ClearLastException(ctx)
}

func TestEntityListBySize(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx)
	aSize := 1000

	aHandle, err := g2diagnostic.GetEntityListBySize(ctx, aSize)
	testError(test, ctx, g2diagnostic, err)

	anEntity, err := g2diagnostic.FetchNextEntityBySize(ctx, aHandle)
	testError(test, ctx, g2diagnostic, err)
	printResult(test, "Entity", anEntity)

	err = g2diagnostic.CloseEntityListBySize(ctx, aHandle)
	testError(test, ctx, g2diagnostic, err)

}

func TestFindEntitiesByFeatureIDs(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx)
	features := "{\"ENTITY_ID\":1,\"LIB_FEAT_IDS\":[1,3,4]}"
	actual, err := g2diagnostic.FindEntitiesByFeatureIDs(ctx, features)
	testError(test, ctx, g2diagnostic, err)
	printResult(test, "len(Actual)", len(actual))

}

func TestGetAvailableMemory(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx)
	actual, err := g2diagnostic.GetAvailableMemory(ctx)
	testError(test, ctx, g2diagnostic, err)
	assert.Greater(test, actual, int64(0))
	printActual(test, actual)
}

func TestGetDataSourceCounts(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx)
	actual, err := g2diagnostic.GetDataSourceCounts(ctx)
	testError(test, ctx, g2diagnostic, err)
	printResult(test, "Data Source counts", actual)
}

func TestGetDBInfo(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx)
	actual, err := g2diagnostic.GetDBInfo(ctx)
	testError(test, ctx, g2diagnostic, err)
	printActual(test, actual)
}

func TestGetEntityDetails(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx)
	entityID := int64(1)
	includeInternalFeatures := 1
	actual, err := g2diagnostic.GetEntityDetails(ctx, entityID, includeInternalFeatures)
	testError(test, ctx, g2diagnostic, err)
	printActual(test, actual)
}

func TestGetEntityResume(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx)
	entityID := int64(1)
	actual, err := g2diagnostic.GetEntityResume(ctx, entityID)
	testError(test, ctx, g2diagnostic, err)
	printActual(test, actual)
}

func TestGetEntitySizeBreakdown(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx)
	minimumEntitySize := 1
	includeInternalFeatures := 1
	actual, err := g2diagnostic.GetEntitySizeBreakdown(ctx, minimumEntitySize, includeInternalFeatures)
	testError(test, ctx, g2diagnostic, err)
	printActual(test, actual)
}

func TestGetFeature(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx)
	libFeatID := int64(1)
	actual, err := g2diagnostic.GetFeature(ctx, libFeatID)
	testError(test, ctx, g2diagnostic, err)
	printActual(test, actual)
}

func TestGetGenericFeatures(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx)
	featureType := "PHONE"
	maximumEstimatedCount := 10
	actual, err := g2diagnostic.GetGenericFeatures(ctx, featureType, maximumEstimatedCount)
	testError(test, ctx, g2diagnostic, err)
	printActual(test, actual)
}

func TestGetLastException(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx)
	actual, err := g2diagnostic.GetLastException(ctx)
	if err != nil {
		test.Log("Error:", err.Error())
	} else {
		printActual(test, actual)
	}
}

func TestGetLastExceptionCode(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx)
	actual, err := g2diagnostic.GetLastExceptionCode(ctx)
	testError(test, ctx, g2diagnostic, err)
	printActual(test, actual)
}

func TestGetLogicalCores(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx)
	actual, err := g2diagnostic.GetLogicalCores(ctx)
	testError(test, ctx, g2diagnostic, err)
	assert.Greater(test, actual, 0)
	printActual(test, actual)
}

func TestGetMappingStatistics(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx)
	includeInternalFeatures := 1
	actual, err := g2diagnostic.GetMappingStatistics(ctx, includeInternalFeatures)
	testError(test, ctx, g2diagnostic, err)
	printActual(test, actual)
}

func TestGetPhysicalCores(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx)
	actual, err := g2diagnostic.GetPhysicalCores(ctx)
	testError(test, ctx, g2diagnostic, err)
	assert.Greater(test, actual, 0)
	printActual(test, actual)
}

func TestGetRelationshipDetails(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx)
	relationshipID := int64(1)
	includeInternalFeatures := 1
	actual, err := g2diagnostic.GetRelationshipDetails(ctx, relationshipID, includeInternalFeatures)
	testError(test, ctx, g2diagnostic, err)
	printActual(test, actual)
}

func TestGetResolutionStatistics(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx)
	actual, err := g2diagnostic.GetResolutionStatistics(ctx)
	testError(test, ctx, g2diagnostic, err)
	printActual(test, actual)
}

func TestGetTotalSystemMemory(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx)
	actual, err := g2diagnostic.GetTotalSystemMemory(ctx)
	testError(test, ctx, g2diagnostic, err)
	assert.Greater(test, actual, int64(0))
	printActual(test, actual)
}

func TestInit(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := &G2diagnosticImpl{}
	moduleName := "Test module name"
	verboseLogging := 0
	iniParams, jsonErr := g2configuration.BuildSimpleSystemConfigurationJson("")
	testError(test, ctx, g2diagnostic, jsonErr)
	err := g2diagnostic.Init(ctx, moduleName, iniParams, verboseLogging)
	testError(test, ctx, g2diagnostic, err)
}

func TestInitWithConfigID(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := &G2diagnosticImpl{}
	moduleName := "Test module name"
	initConfigID := int64(1)
	verboseLogging := 0
	iniParams, jsonErr := g2configuration.BuildSimpleSystemConfigurationJson("")
	testError(test, ctx, g2diagnostic, jsonErr)
	err := g2diagnostic.InitWithConfigID(ctx, moduleName, iniParams, initConfigID, verboseLogging)
	testError(test, ctx, g2diagnostic, err)
}

func TestReinit(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := &G2diagnosticImpl{}
	initConfigID := int64(1)
	err := g2diagnostic.Reinit(ctx, initConfigID)
	testError(test, ctx, g2diagnostic, err)
}

func TestDestroy(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx)
	err := g2diagnostic.Destroy(ctx)
	testError(test, ctx, g2diagnostic, err)
}
