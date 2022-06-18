package main

import (
	"context"
	"fmt"

	"github.com/docktermj/xyzzygoapi/g2diagnostic"
	"github.com/docktermj/xyzzygoapi/g2helper"
)

// Values updated via "go install -ldflags" parameters.

var programName string = "unknown"
var buildVersion string = "0.0.0"
var buildIteration string = "0"

// ----------------------------------------------------------------------------
// Internal methods - names begin with lower case
// ----------------------------------------------------------------------------

func getG2diagnostic() (g2diagnostic.G2diagnostic, error) {
	var err error = nil
	g2diagnostic := g2diagnostic.G2diagnosticImpl{}
	ctx := context.TODO()

	moduleName := "Test module name"
	verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
	iniParams, jsonErr := g2helper.GetSimpleSystemConfigurationJson(ctx)
	if jsonErr != nil {
		return &g2diagnostic, jsonErr
	}

	err = g2diagnostic.Init(ctx, moduleName, iniParams, verboseLogging)
	return &g2diagnostic, err
}

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {
	ctx := context.TODO()
	g2diagnostic, _ := getG2diagnostic()
	secondsToRun := 1
	actual, err := g2diagnostic.CheckDBPerf(ctx, secondsToRun)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(actual)
}
