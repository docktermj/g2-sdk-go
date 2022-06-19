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

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

func TestNoop(test *testing.T) {
	ctx := context.TODO()
	g2engine, err := getTestObject(ctx)
	testError(test, ctx, g2engine, err)
}
