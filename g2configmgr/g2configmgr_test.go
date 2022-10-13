package g2configmgr

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/docktermj/go-xyzzy-helpers/g2configuration"
	"github.com/docktermj/go-xyzzy-helpers/logger"
	"github.com/stretchr/testify/assert"

	truncator "github.com/aquilax/truncate"
)

var (
	g2configmgr G2configmgr
)

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context) G2configmgr {

	if g2configmgr == nil {
		g2configmgr = &G2configmgrImpl{}

		moduleName := "Test module name"
		verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
		iniParams, jsonErr := g2configuration.BuildSimpleSystemConfigurationJson("")
		if jsonErr != nil {
			logger.Fatalf("Cannot construct system configuration: %v", jsonErr)
		}

		initErr := g2configmgr.Init(ctx, moduleName, iniParams, verboseLogging)
		if initErr != nil {
			logger.Fatalf("Cannot Init: %v", initErr)
		}
	}
	return g2configmgr
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

func testError(test *testing.T, ctx context.Context, g2configmgr G2configmgr, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
		lastException, _ := g2configmgr.GetLastException(ctx)
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

func TestClearLastException(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx)
	err := g2config.ClearLastException(ctx)
	testError(test, ctx, g2config, err)
}

func TestGetLastException(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx)
	actual, err := g2config.GetLastException(ctx)
	if err != nil {
		test.Log("Error:", err.Error())
	} else {
		printActual(test, actual)
	}
}

func TestGetLastExceptionCode(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx)
	actual, err := g2config.GetLastExceptionCode(ctx)
	testError(test, ctx, g2config, err)
	printActual(test, actual)
}

func TestInit(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx)
	moduleName := "Test module name"
	verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
	iniParams, jsonErr := g2configuration.BuildSimpleSystemConfigurationJson("")
	if jsonErr != nil {
		logger.Fatalf("Cannot construct system configuration: %v", jsonErr)
	}
	err := g2config.Init(ctx, moduleName, iniParams, verboseLogging)
	testError(test, ctx, g2config, err)
}

func TestDestroy(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx)
	err := g2config.Destroy(ctx)
	testError(test, ctx, g2config, err)
}
