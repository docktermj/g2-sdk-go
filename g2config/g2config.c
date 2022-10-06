#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include "g2config.h"

int G2config_close_helper(void* configHandle) {
    int returnCode = G2Config_close(&configHandle);
    return returnCode;
}

void* G2config_create_helper() {
    ConfigHandle configHandle;
    printf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\n");
    fflush(stdout);
    int returnCode = G2Config_create(&configHandle);
    printf("%i\n", returnCode);
    printf("%p\n", configHandle);
    printf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\n");
    fflush(stdout);
    return configHandle;
}
