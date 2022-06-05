package g2diagnostic

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * Internal methods.
 */

type XyzzyConfigurationPipeline struct {
	ConfigPath   string `json:"CONFIGPATH"`
	ResourcePath string `json:"RESOURCEPATH"`
	SupportPath  string `json:"SUPPORTPATH"`
}

type XyzzyConfigurationSql struct {
	Connection string `json:"CONNECTION"`
}

type XyzzyConfiguration struct {
	Pipeline XyzzyConfigurationPipeline `json:"PIPELINE"`
	Sql      XyzzyConfigurationSql      `json:"SQL"`
}

func getConfigurationJson() string {
	resultStruct := XyzzyConfiguration{
		Pipeline: XyzzyConfigurationPipeline{
			ConfigPath:   "/etc/opt/senzing",
			ResourcePath: "/opt/senzing/g2/resources",
			SupportPath:  "/opt/senzing/data",
		},
		Sql: XyzzyConfigurationSql{
			Connection: "postgresql://postgres:postgres@127.0.0.1:5432:G2/",
		},
	}

	resultBytes, _ := json.Marshal(resultStruct)
	return string(resultBytes)
}

func getTestObject() (G2diagnostic, error) {
	var err error = nil
	g2diagnostic := G2diagnosticImpl{}
	ctx := context.TODO()

	moduleName := "Test module name"
	verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
	iniParams := getConfigurationJson()

	err = g2diagnostic.Init(ctx, moduleName, iniParams, verboseLogging)
	return &g2diagnostic, err
}

func testError(test *testing.T, ctx context.Context, g2diagnostic G2diagnostic, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
		lastException, _ := g2diagnostic.GetLastException(ctx)
		assert.FailNow(test, lastException)
	}
}

/*
 * The unit tests in this file...
 */

// ----------------------------------------------------------------------------
// Initialization
// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
// Work in progress
// ----------------------------------------------------------------------------

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

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

func TestCheckDBPerf(test *testing.T) {
	teardownSuite := setupSuite(test)
	defer teardownSuite(test)

	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	secondsToRun := 1
	actual, err := g2diagnostic.CheckDBPerf(ctx, secondsToRun)
	testError(test, ctx, g2diagnostic, err)
	test.Log("Database performance:", actual)
}

func TestClearLastException(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	g2diagnostic.ClearLastException(ctx)
}

func TestEntityListBySize(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	aSize := 10

	aHandle, err := g2diagnostic.GetEntityListBySize(ctx, aSize)
	testError(test, ctx, g2diagnostic, err)

	anEntity, err := g2diagnostic.FetchNextEntityBySize(ctx, aHandle)
	testError(test, ctx, g2diagnostic, err)
	test.Log("Entity:", anEntity)

	err = g2diagnostic.CloseEntityListBySize(ctx, aHandle)
	testError(test, ctx, g2diagnostic, err)
}

func TestDestroy(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	err := g2diagnostic.Destroy(ctx)
	testError(test, ctx, g2diagnostic, err)
}

func TestFindEntitiesByFeatureIDs(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	features := "{\"ENTITY_ID\":1}"
	actual, err := g2diagnostic.FindEntitiesByFeatureIDs(ctx, features)
	testError(test, ctx, g2diagnostic, err)
	test.Log("Entities:", actual)
}

func TestGetAvailableMemory(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	actual, err := g2diagnostic.GetAvailableMemory(ctx)
	testError(test, ctx, g2diagnostic, err)
	assert.Greater(test, actual, int64(0))
	test.Log("Available memory:", actual)
}

func TestGetDBInfo(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	actual, err := g2diagnostic.GetDBInfo(ctx)
	testError(test, ctx, g2diagnostic, err)
	test.Log("Database info:", actual)
}

func TestGetLastException(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	actual, err := g2diagnostic.GetLastException(ctx)
	testError(test, ctx, g2diagnostic, err)
	test.Log("Last exception:", actual)
}

func TestGetLogicalCores(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	actual, err := g2diagnostic.GetLogicalCores(ctx)
	testError(test, ctx, g2diagnostic, err)
	assert.Greater(test, actual, 0)
	test.Log(" Logical cores:", actual)
}

func TestGetPhysicalCores(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	actual, err := g2diagnostic.GetPhysicalCores(ctx)
	testError(test, ctx, g2diagnostic, err)
	assert.Greater(test, actual, 0)
	test.Log("Physical cores:", actual)
}

func TestGetTotalSystemMemory(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	actual, err := g2diagnostic.GetTotalSystemMemory(ctx)
	testError(test, ctx, g2diagnostic, err)
	assert.Greater(test, actual, int64(0))
	test.Log("Total system memory:", actual)
}

func TestInit(test *testing.T) {
	g2diagnostic := &G2diagnosticImpl{}
	ctx := context.TODO()
	moduleName := "Test module name"
	verboseLogging := 0
	iniParams := getConfigurationJson()
	err := g2diagnostic.Init(ctx, moduleName, iniParams, verboseLogging)
	testError(test, ctx, g2diagnostic, err)
}
