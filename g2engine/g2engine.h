#include <stdlib.h>
#include <stdio.h>
#include "libg2.h"

typedef void* ExportHandle;
typedef void*(*resize_buffer_type)(void *, size_t);

struct G2_addRecordWithInfoWithReturnedRecordID_result {
    char* recordID;
    char* withInfo;
    int returnCode;
};

struct G2_exportConfigAndConfigID_result {
    long long configID;
    char* config;
    int returnCode;
};

struct G2_getActiveConfigID_result {
    long long configID;
    int returnCode;
};

void* G2_resizeStringBuffer(void *ptr, size_t size);
char* G2_addRecordWithInfo_local(const char* dataSourceCode, const char* recordID, const char* jsonData, const char *loadID, const long long flags);
struct G2_addRecordWithInfoWithReturnedRecordID_result G2_addRecordWithInfoWithReturnedRecordID_local(const char* dataSourceCode, const char* jsonData, const char *loadID, const long long flags);
char* G2_checkRecord_local(const char* record, const char* recordQueryList);
char* G2_deleteRecordWithInfo_local(const char* dataSourceCode, const char* recordID, const char *loadID, const long long flags);
struct G2_exportConfigAndConfigID_result G2_exportConfigAndConfigID_local();
char* G2_exportConfig_local();
char* G2_findInterestingEntitiesByEntityID_local(long long entityID, long long flags);
char* G2_findInterestingEntitiesByRecordID_local(const char* dataSourceCode, const char* recordID, long long flags);
char* G2_findNetworkByEntityID_local(const char* entityList, const int maxDegree, const int buildOutDegree, const int maxEntities);
char* G2_findNetworkByEntityID_V2_local(const char* entityList, const int maxDegree, const int buildOutDegree, const int maxEntities, long long flags);
char* G2_findNetworkByRecordID_local(const char* recordList, const int maxDegree, const int buildOutDegree, const int maxEntities);
char* G2_findNetworkByRecordID_V2_local(const char* recordList, const int maxDegree, const int buildOutDegree, const int maxEntities, const long long flags);
char* G2_findPathByEntityID_local(const long long entityID1, const long long entityID2, const int maxDegree);
char* G2_findPathByEntityID_V2_local(const long long entityID1, const long long entityID2, const int maxDegree, const long long flags);
char* G2_findPathByRecordID_local(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree);
char* G2_findPathByRecordID_V2_local(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const long long flags);
char* G2_findPathExcludingByEntityID_local(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities);
char* G2_findPathExcludingByEntityID_V2_local(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities, const long long flags);
char* G2_findPathExcludingByRecordID_local(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords);
char* G2_findPathExcludingByRecordID_V2_local(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords, const long long flags);
char* G2_findPathIncludingSourceByEntityID_local(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities, const char* requiredDsrcs);
char* G2_findPathIncludingSourceByEntityID_V2_local(const long long entityID1, const long long entityID2, const int maxDegree, const char* excludedEntities, const char* requiredDsrcs, const long long flags);
char* G2_findPathIncludingSourceByRecordID_local(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords, const char* requiredDsrcs);
char* G2_findPathIncludingSourceByRecordID_V2_local(const char* dataSourceCode1, const char* recordID1, const char* dataSourceCode2, const char* recordID2, const int maxDegree, const char* excludedRecords, const char* requiredDsrcs, const long long flags);
struct G2_getActiveConfigID_result G2_getActiveConfigID_local();

char* G2_stats_local();
