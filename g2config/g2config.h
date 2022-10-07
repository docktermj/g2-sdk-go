#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include "libg2config.h"

//typedef void* ConfigHandle;
typedef void*(*resize_buffer_type)(void *, size_t);

struct G2Config_addDataSource_result {
    char* response;
    int returnCode;
};

struct G2Config_listDataSources_result {
    char* response;
    int returnCode;
};


struct G2Config_addDataSource_result G2Config_addDataSource_helper(uintptr_t configHandle, const char *inputJson);
int G2config_close_helper(uintptr_t configHandle);
void* G2config_create_helper();
struct G2Config_listDataSources_result G2Config_listDataSources_helper(uintptr_t configHandle);


