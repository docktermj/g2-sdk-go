#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include "libg2configmgr.h"

typedef void* (*resize_buffer_type)(void*, size_t);

struct G2ConfigMgr_addConfig_result {
    long long configID;
    int returnCode;
};

struct G2ConfigMgr_addConfig_result G2ConfigMgr_addConfig_helper(const char* configStr, const char* configComments);
