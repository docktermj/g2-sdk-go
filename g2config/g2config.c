#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include "g2config.h"

typedef void* ConfigHandle;

uintptr_t G2config_create_helper() {
    ConfigHandle configHandle;
    int returnCode = G2Config_create(&configHandle);
    return (uintptr_t)configHandle;
}

void* G2config_create_helper_1() {
    ConfigHandle configHandle;
    int returnCode = G2Config_create(&configHandle);
    return configHandle;
}
