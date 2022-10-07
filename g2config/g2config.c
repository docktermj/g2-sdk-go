#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include "g2config.h"

int G2config_close_helper(uintptr_t configHandle) {
    printf(">>>> Close >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\n");
    printf(" configHandle: %lui\n", configHandle);
    printf("&configHandle: %p\n", &configHandle);
    fflush(stdout);
    printf("<<<< Close <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<\n");
    int returnCode = G2Config_close((void*)configHandle);
    return returnCode;
}

void* G2config_create_helper() {
    ConfigHandle configHandle;
    printf(">>>> Create >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\n");
    fflush(stdout);
    int returnCode = G2Config_create(&configHandle);
    printf("Return  code: %i\n", returnCode);
    printf("configHandle: %p\n", configHandle);
    printf("<<<< Create <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<\n");
    fflush(stdout);
    return configHandle;
}
