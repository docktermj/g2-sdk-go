#include <stdlib.h>
#include <stdio.h>
#include "g2config.h"


long long G2config_create_helper() {
    long long configHandle;
    int returnCode = G2Config_create(&configHandle);
    return configHandle;
}
