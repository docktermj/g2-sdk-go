#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include "g2config.h"

typedef void* ConfigHandle;

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

void* G2config_create_helper_1() {
    ConfigHandle configHandle;
    int returnCode = G2Config_create(&configHandle);
    return configHandle;
}
