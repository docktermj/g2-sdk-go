package g2diagnostic

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * The unit tests in this file simulate command line invocation.
 */

func TestGetPhysicalCores(test *testing.T) {
	expected := 8

	ctx := context.TODO()

	// Create expected file.
	localG2diagnostic := G2diagnosticImpl{
		Args: map[string]interface{}{},
	}

	actual, _ := localG2diagnostic.GetPhysicalCores(ctx)

	fmt.Println(">>> ", actual)
	assert.Equal(test, expected, actual, "Not expected number of cores")

}
