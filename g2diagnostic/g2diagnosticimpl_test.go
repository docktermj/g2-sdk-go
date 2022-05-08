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
			SupportPath:  "/opt/senzing/data/3.0.0",
		},
		Sql: XyzzyConfigurationSql{
			Connection: "postgresql://postgres:postgres@127.0.0.1:5432:G2/",
		},
	}

	resultBytes, _ := json.Marshal(resultStruct)
	return string(resultBytes)
}

func getConfigurationJsonDefault() string {
	resultStruct := XyzzyConfiguration{
		Pipeline: XyzzyConfigurationPipeline{
			ConfigPath:   "/etc/opt/senzing",
			ResourcePath: "/opt/senzing/g2/resources",
			SupportPath:  "/opt/senzing/data",
		},
		Sql: XyzzyConfigurationSql{
			Connection: "xyzzy",
		},
	}

	resultBytes, _ := json.Marshal(resultStruct)
	return string(resultBytes)
}

/*
 * The unit tests in this file...
 */

func TestGetLastException(test *testing.T) {
	g2diagnostic := G2diagnosticImpl{}
	ctx := context.TODO()
	actual, _ := g2diagnostic.GetLastException(ctx)
	test.Log("Last exception:", actual)

}

func TestCheckDBPerf(test *testing.T) {
	g2diagnostic := G2diagnosticImpl{}
	ctx := context.TODO()
	secondsToRun := 5
	actual, _, _ := g2diagnostic.CheckDBPerf(ctx, secondsToRun)
	test.Log("Database performance:", actual)
}

func TestGetAvailableMemory(test *testing.T) {
	g2diagnostic := G2diagnosticImpl{}
	ctx := context.TODO()
	actual, _ := g2diagnostic.GetAvailableMemory(ctx)
	test.Log("Available memory:", actual)
}

func TestClearLastException(test *testing.T) {
	g2diagnostic := G2diagnosticImpl{}
	ctx := context.TODO()
	g2diagnostic.ClearLastException(ctx)
}

func TestGetLogicalCores(test *testing.T) {
	g2diagnostic := G2diagnosticImpl{}
	ctx := context.TODO()
	actual, _ := g2diagnostic.GetLogicalCores(ctx)
	test.Log(" Logical cores:", actual)
}

func TestGetPhysicalCores(test *testing.T) {
	g2diagnostic := G2diagnosticImpl{}
	ctx := context.TODO()
	actual, _ := g2diagnostic.GetPhysicalCores(ctx)
	test.Log("Physical cores:", actual)
}

func TestGetTotalSystemMemory(test *testing.T) {
	g2diagnostic := G2diagnosticImpl{}
	ctx := context.TODO()
	actual, _ := g2diagnostic.GetTotalSystemMemory(ctx)
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
