#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include "libg2configmgr.h"
#include "g2configmgr.h"

void* G2configmgr_resizeStringBuffer(void* ptr, size_t size) {
    //deallocate old buffer
    if (ptr != 0)
        free(ptr);
    //allocate new buffer
    void* buffer = malloc(size);
    return buffer;
}

struct G2ConfigMgr_addConfig_result G2ConfigMgr_addConfig_helper(const char* configStr, const char* configComments) {
    long long configID;
    int returnCode = G2ConfigMgr_addConfig(configStr, configComments, &configID);
    struct G2ConfigMgr_addConfig_result result;
    result.configID = configID;
    result.returnCode = returnCode;
    return result;
}