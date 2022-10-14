package g2product

import (
	"context"
	"fmt"
	"testing"

	"github.com/docktermj/go-xyzzy-helpers/g2configuration"
	"github.com/docktermj/go-xyzzy-helpers/logger"
	"github.com/stretchr/testify/assert"

	truncator "github.com/aquilax/truncate"
)

var (
	g2product G2product
)

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context) G2product {

	if g2product == nil {
		g2product = &G2productImpl{}

		moduleName := "Test module name"
		verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
		iniParams, jsonErr := g2configuration.BuildSimpleSystemConfigurationJson("")
		if jsonErr != nil {
			logger.Fatalf("Cannot construct system configuration: %v", jsonErr)
		}

		initErr := g2product.Init(ctx, moduleName, iniParams, verboseLogging)
		if initErr != nil {
			logger.Fatalf("Cannot Init: %v", initErr)
		}
	}
	return g2product
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

func testError(test *testing.T, ctx context.Context, g2product G2product, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
		lastException, _ := g2product.GetLastException(ctx)
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

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

func TestClearLastException(test *testing.T) {
	ctx := context.TODO()
	g2product := getTestObject(ctx)
	g2product.ClearLastException(ctx)
}

func TestGetLastException(test *testing.T) {
	ctx := context.TODO()
	g2product := getTestObject(ctx)
	actual, err := g2product.GetLastException(ctx)
	if err == nil {
		printActual(test, actual)
	}
}

func TestGetLastExceptionCode(test *testing.T) {
	ctx := context.TODO()
	g2product := getTestObject(ctx)
	actual, err := g2product.GetLastExceptionCode(ctx)
	testError(test, ctx, g2product, err)
	printActual(test, actual)
}

func TestInit(test *testing.T) {
	ctx := context.TODO()
	g2product := &G2productImpl{}
	moduleName := "Test module name"
	verboseLogging := 0
	iniParams, jsonErr := g2configuration.BuildSimpleSystemConfigurationJson("")
	testError(test, ctx, g2product, jsonErr)
	err := g2product.Init(ctx, moduleName, iniParams, verboseLogging)
	testError(test, ctx, g2product, err)
}

func TestLicense(test *testing.T) {
	ctx := context.TODO()
	g2product := getTestObject(ctx)
	actual, err := g2product.License(ctx)
	testError(test, ctx, g2product, err)
	printActual(test, actual)
}

func TestValidateLicenseFile(test *testing.T) {
	ctx := context.TODO()
	g2product := getTestObject(ctx)
	licenseFilePath := ""
	actual, _ := g2product.ValidateLicenseFile(ctx, licenseFilePath)
	// testError(test, ctx, g2product, err)
	printActual(test, actual)
}

func TestValidateLicenseStringBase64(test *testing.T) {
	ctx := context.TODO()
	g2product := getTestObject(ctx)
	licenseString := ""
	actual, _ := g2product.ValidateLicenseStringBase64(ctx, licenseString)
	// testError(test, ctx, g2product, err)
	printActual(test, actual)
}

func TestVersion(test *testing.T) {
	ctx := context.TODO()
	g2product := getTestObject(ctx)
	actual, err := g2product.Version(ctx)
	testError(test, ctx, g2product, err)
	printActual(test, actual)
}

func TestDestroy(test *testing.T) {
	ctx := context.TODO()
	g2product := getTestObject(ctx)
	err := g2product.Destroy(ctx)
	testError(test, ctx, g2product, err)
}
