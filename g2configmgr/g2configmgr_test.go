package g2configmgr

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/docktermj/g2-sdk-go/g2config"
	"github.com/docktermj/go-xyzzy-helpers/g2configuration"
	"github.com/docktermj/go-xyzzy-helpers/logger"
	"github.com/senzing/go-logging/messagelogger"
	"github.com/stretchr/testify/assert"

	truncator "github.com/aquilax/truncate"
)

var (
	g2configmgr     G2configmgr
	g2configX       g2config.G2config
	loggerSingleton messagelogger.MessageLoggerInterface
)

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context) G2configmgr {

	if g2configmgr == nil {
		g2configmgr = &G2configmgrImpl{}
		logger := getLogger(ctx)

		moduleName := "Test module name"
		verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
		iniParams, jsonErr := g2configuration.BuildSimpleSystemConfigurationJson("")
		if jsonErr != nil {
			logger.Log(1001, "Cannot construct system configuration: %v", jsonErr)
		}

		initErr := g2configmgr.Init(ctx, moduleName, iniParams, verboseLogging)
		if initErr != nil {
			logger.Log(1002, "Cannot Init: %v", initErr)
		}
	}
	return g2configmgr
}

func getG2Config(ctx context.Context) g2config.G2config {

	if g2configX == nil {
		g2configX = &g2config.G2configImpl{}
		logger := getLogger(ctx)

		moduleName := "Test module name"
		verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
		iniParams, jsonErr := g2configuration.BuildSimpleSystemConfigurationJson("")
		if jsonErr != nil {
			logger.Log(1001, "Cannot construct system configuration: %v", jsonErr)
		}

		initErr := g2configX.Init(ctx, moduleName, iniParams, verboseLogging)
		if initErr != nil {
			logger.Log(1002, "Cannot Init: %v", initErr)
		}
	}
	return g2configX
}

func getLogger(ctx context.Context) messagelogger.MessageLoggerInterface {
	if loggerSingleton == nil {
		log.SetFlags(log.LstdFlags)
		loggerSingleton, _ = messagelogger.New()
	}
	return loggerSingleton
}

func truncate(aString string) string {
	return truncator.Truncate(aString, 50, "...", truncator.PositionEnd)
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
	ctx := context.TODO()
	logger := getLogger(ctx)
	logger.Log(1003, "Test message 1", "Variable1", "Variable2")
}

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

func TestAddConfig(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx)
	now := time.Now()

	// Create an in-memory configuration.

	g2config := getG2Config(ctx)
	configHandle, err1 := g2config.Create(ctx)
	if err1 != nil {
		test.Log("Error:", err1.Error())
		assert.FailNow(test, "g2config.Create()")
	}

	// Modify the in-memory configuration so it is different from the created configuration.
	// If not, on Save Senzing will detect that it is the same and no Save occurs.

	inputJson := `{"DSRC_CODE": "GO_TEST_` + strconv.FormatInt(now.Unix(), 10) + `"}`
	_, err2 := g2config.AddDataSource(ctx, configHandle, inputJson)
	if err2 != nil {
		test.Log("Error:", err2.Error())
		assert.FailNow(test, "g2config.AddDataSource()")
	}

	// Create a JSON string from the in-memory version of the configuration.

	configStr, err3 := g2config.Save(ctx, configHandle)
	if err3 != nil {
		test.Log("Error:", err2.Error())
		assert.FailNow(test, configStr)
	}

	// Perform the test.

	configComments := fmt.Sprintf("g2configmgr_test at %s", now.UTC())
	actual, err := g2configmgr.AddConfig(ctx, configStr, configComments)
	testError(test, ctx, g2configmgr, err)
	printActual(test, actual)
}

func TestClearLastException(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx)
	err := g2configmgr.ClearLastException(ctx)
	testError(test, ctx, g2configmgr, err)
}

func TestGetConfig(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx)

	// Get a ConfigID.

	configID, err1 := g2configmgr.GetDefaultConfigID(ctx)
	if err1 != nil {
		test.Log("Error:", err1.Error())
		assert.FailNow(test, "g2configmgr.GetDefaultConfigID()")
	}

	actual, err := g2configmgr.GetConfig(ctx, configID)
	testError(test, ctx, g2configmgr, err)
	printActual(test, actual)
}

func TestGetConfigList(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx)
	actual, err := g2configmgr.GetConfigList(ctx)
	testError(test, ctx, g2configmgr, err)
	printActual(test, actual)
}

func TestGetDefaultConfigID(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx)
	actual, err := g2configmgr.GetDefaultConfigID(ctx)
	testError(test, ctx, g2configmgr, err)
	printActual(test, actual)
}

func TestGetLastException(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx)
	actual, err := g2configmgr.GetLastException(ctx)
	if err != nil {
		test.Log("Error:", err.Error())
	} else {
		printActual(test, actual)
	}
}

func TestGetLastExceptionCode(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx)
	actual, err := g2configmgr.GetLastExceptionCode(ctx)
	testError(test, ctx, g2configmgr, err)
	printActual(test, actual)
}

func TestInit(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx)
	moduleName := "Test module name"
	verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
	iniParams, jsonErr := g2configuration.BuildSimpleSystemConfigurationJson("")
	if jsonErr != nil {
		logger.Fatalf("Cannot construct system configuration: %v", jsonErr)
	}
	err := g2configmgr.Init(ctx, moduleName, iniParams, verboseLogging)
	testError(test, ctx, g2configmgr, err)
}

func TestReplaceDefaultConfigID(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx)

	oldConfigID, err1 := g2configmgr.GetDefaultConfigID(ctx)
	if err1 != nil {
		test.Log("Error:", err1.Error())
		assert.FailNow(test, "g2configmgr.GetDefaultConfigID()")
	}

	// FIXME: This is kind of a cheeter.

	newConfigID, err2 := g2configmgr.GetDefaultConfigID(ctx)
	if err2 != nil {
		test.Log("Error:", err2.Error())
		assert.FailNow(test, "g2configmgr.GetDefaultConfigID()-2")
	}

	err := g2configmgr.ReplaceDefaultConfigID(ctx, oldConfigID, newConfigID)
	testError(test, ctx, g2configmgr, err)
}

func TestSetDefaultConfigID(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx)

	configID, err1 := g2configmgr.GetDefaultConfigID(ctx)
	if err1 != nil {
		test.Log("Error:", err1.Error())
		assert.FailNow(test, "g2configmgr.GetDefaultConfigID()")
	}
	err := g2configmgr.SetDefaultConfigID(ctx, configID)
	testError(test, ctx, g2configmgr, err)
}

func TestDestroy(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx)
	err := g2configmgr.Destroy(ctx)
	testError(test, ctx, g2configmgr, err)
}
