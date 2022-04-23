package g2diagnostic

/*
Environment

The purpose of a Environment object is:

 1. Single source of "truth" for values submitted by user in this priority order:
    1. Commandline options
    1. OS environment variables
    1. configuration file
    1. default values
*/

/*
#include <stdlib.h>
#include <stdio.h>
#include "libg2diagnostic.h"
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -shared
*/
import "C"
import (
  "context"
)

// ----------------------------------------------------------------------------
// Structure
// ----------------------------------------------------------------------------

type G2diagnosticImpl struct {
  Args              map[string]interface{} `json:"args"`
  configuration     gonfig.Gonfig
  configurationFile string
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

func (g2diagnostic *G2diagnosticImpl) GetPhysicalCores(ctx context.Context) (int, error) {

    result := C.G2Diagnostic_getPhysicalCores()
    return int(result), nil

}
