package g2diagnostic

import (
	"context"
	"testing"
	//	"github.com/stretchr/testify/assert"
)

/*
 * The unit tests in this file...
 */

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
