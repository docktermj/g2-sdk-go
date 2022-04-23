package g2diagnostic

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * The unit tests in this file simulate command line invocation.
 */

func TestGetPhysicalCores(test *testing.T) {
	expected := 8

	// Create expected file.
	localG2diagnostic := G2diagnosticImpl{
		Args: map[string]interface{}{
			key.BIXAGENT_OPTION_CONFIGURATION: expected,
		},
	}

	actual, _ := localG2diagnostic.GetPhysicalCores()
	assert.NotEqual(test, "", actual, "No configuration filename found")
	if _, err := os.Stat(actual); os.IsNotExist(err) {
		assert.Falsef(test, true, "Could not find %s", actual)
	}
}
