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
	secondsToRun := 10
	actual, _ := g2diagnostic.CheckDBPerf(ctx, secondsToRun)
	test.Log("Database performance:", actual)
}

func TestClearLastException(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	g2diagnostic.ClearLastException(ctx)
}

func TestGetAvailableMemory(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	actual, _ := g2diagnostic.GetAvailableMemory(ctx)
	assert.Greater(test, actual, int64(0))
	test.Log("Available memory:", actual)
}

func TestGetDBInfo(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	actual, _ := g2diagnostic.GetDBInfo(ctx)
	test.Log("Database info:", actual)
}

func TestGetLastException(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	actual, _ := g2diagnostic.GetLastException(ctx)
	test.Log("Last exception:", actual)
}

func TestGetLogicalCores(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	actual, _ := g2diagnostic.GetLogicalCores(ctx)
	assert.Greater(test, actual, 0)
	test.Log(" Logical cores:", actual)
}

func TestGetPhysicalCores(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	actual, _ := g2diagnostic.GetPhysicalCores(ctx)
	assert.Greater(test, actual, 0)
	test.Log("Physical cores:", actual)
}

func TestGetTotalSystemMemory(test *testing.T) {
	g2diagnostic, _ := getTestObject()
	ctx := context.TODO()
	actual, _ := g2diagnostic.GetTotalSystemMemory(ctx)
	assert.Greater(test, actual, int64(0))
	test.Log("Total system memory:", actual)
}

func TestInit(test *testing.T) {
	g2diagnostic := G2diagnosticImpl{}
	ctx := context.TODO()

	moduleName := "Test module name"
	verboseLogging := 0
	iniParams := getConfigurationJson()

	err := g2diagnostic.Init(ctx, moduleName, iniParams, verboseLogging)
	if err != nil {
		test.Log("iniParams:", iniParams)
		lastException, _ := g2diagnostic.GetLastException(ctx)
		assert.FailNow(test, lastException)
	}
}
