package g2helper

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func testError(test *testing.T, ctx context.Context, err error) {
	if err != nil {
		assert.FailNow(test, err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

func TestGetSimpleSystemConfigurationJson(test *testing.T) {

	ctx := context.TODO()
	actual, err := GetSimpleSystemConfigurationJson(ctx)
	testError(test, ctx, err)
	test.Log("Actual:", actual)
}
