package xyzzyapi

/*
#include <stdlib.h>
#include <stdio.h>
#include "libg2diagnostic.h"
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -shared
*/
import "C"
import (
  "fmt"
)

// Values updated via "go install -ldflags" parameters.

var moduleName string = "github.com/docktermj/xyzzyapi"
var buildVersion string = "0.0.1"
var buildIteration string = "0"

// ----------------------------------------------------------------------------
// libg2diagnostic.h
// ----------------------------------------------------------------------------

func G2Diagnostic_getPhysicalCores() int {
    result := C.G2Diagnostic_getPhysicalCores()
    return int(result)
}
