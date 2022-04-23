package xyzzygoapi

/*
#include <stdlib.h>
#include <stdio.h>
#include "libg2diagnostic.h"
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -shared
*/
import "C"
import ()

// Values updated via "go install -ldflags" parameters.

var moduleName string = "github.com/docktermj/xyzzyapi"
var buildVersion string = "0.0.1"
var buildIteration string = "0"
