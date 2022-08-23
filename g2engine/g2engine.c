#include <stdlib.h>
#include <stdio.h>
#include "libg2.h"
#include "g2engine.h"

void* G2_resizeStringBuffer(void *ptr, size_t size) {
    //deallocate old buffer
    if (ptr != 0)
        free(ptr);
    //allocate new buffer
    void* buffer = malloc(size);
    return buffer;
}

char* G2_addRecordWithInfo_local(const char* dataSourceCode,
        const char* recordID, const char* jsonData, const char *loadID,
        const long long flags) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    int returnCode = G2_addRecordWithInfo(dataSourceCode, recordID, jsonData,
            loadID, flags, &charBuff, &bufferSize, resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

struct G2_addRecordWithInfoWithReturnedRecordID_result G2_addRecordWithInfoWithReturnedRecordID_local(
        const char* dataSourceCode, const char* jsonData, const char *loadID,
        const long long flags) {
    size_t bufferSize = 1;
    size_t recordIDBufSize = 256;
    char recordIDBuf[recordIDBufSize];
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    int returnCode = G2_addRecordWithInfoWithReturnedRecordID(dataSourceCode,
            jsonData, loadID, flags, recordIDBuf, recordIDBufSize, &charBuff,
            &bufferSize, resizeFuncPointer);
    struct G2_addRecordWithInfoWithReturnedRecordID_result result;
    result.recordID = recordIDBuf;
    result.withInfo = charBuff;
    result.returnCode = returnCode;
    return result;
}

char* G2_checkRecord_local(const char* record, const char* recordQueryList) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    int returnCode = G2_checkRecord(record, recordQueryList, &charBuff,
            &bufferSize, resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2_deleteRecordWithInfo_local(const char* dataSourceCode,
        const char* recordID, const char *loadID, const long long flags) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    int returnCode = G2_deleteRecordWithInfo(dataSourceCode, recordID, loadID,
            flags, &charBuff, &bufferSize, resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

struct G2_exportConfigAndConfigID_result G2_exportConfigAndConfigID_local() {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    long long configID;
    int returnCode = G2_exportConfigAndConfigID(&charBuff, &bufferSize,
            resizeFuncPointer, &configID);
    struct G2_exportConfigAndConfigID_result result;
    result.configID = configID;
    result.config = charBuff;
    result.returnCode = returnCode;
    return result;
}

char* G2_exportConfig_local() {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    int returnCode = G2_exportConfig(&charBuff, &bufferSize, resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2_findInterestingEntitiesByEntityID_local(long long entityID,
        long long flags) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    int returnCode = G2_findInterestingEntitiesByEntityID(entityID, flags,
            &charBuff, &bufferSize, resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2_findInterestingEntitiesByRecordID_local(const char* dataSourceCode,
        const char* recordID, long long flags) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    int returnCode = G2_findInterestingEntitiesByRecordID(dataSourceCode,
            recordID, flags, &charBuff, &bufferSize, resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2_findNetworkByEntityID_local(const char* entityList,
        const int maxDegree, const int buildOutDegree, const int maxEntities) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    int returnCode = G2_findNetworkByEntityID(entityList, maxDegree,
            buildOutDegree, maxEntities, &charBuff, &bufferSize,
            resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2_findNetworkByEntityID_V2_local(const char* entityList,
        const int maxDegree, const int buildOutDegree, const int maxEntities,
        long long flags) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    int returnCode = G2_findNetworkByEntityID_V2(entityList, maxDegree,
            buildOutDegree, maxEntities, flags, &charBuff, &bufferSize,
            resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2_findNetworkByRecordID_local(const char* recordList,
        const int maxDegree, const int buildOutDegree, const int maxEntities) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    int returnCode = G2_findNetworkByRecordID(recordList, maxDegree,
            buildOutDegree, maxEntities, &charBuff, &bufferSize,
            resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2_findNetworkByRecordID_V2_local(const char* recordList,
        const int maxDegree, const int buildOutDegree, const int maxEntities,
        const long long flags) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    int returnCode = G2_findNetworkByRecordID_V2(recordList, maxDegree,
            buildOutDegree, maxEntities, flags, &charBuff, &bufferSize,
            resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2_findPathByEntityID_local(const long long entityID1,
        const long long entityID2, const int maxDegree) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    int returnCode = G2_findPathByEntityID(entityID1, entityID2, maxDegree,
            &charBuff, &bufferSize, resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2_findPathByEntityID_V2_local(const long long entityID1,
        const long long entityID2, const int maxDegree, const long long flags) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    int returnCode = G2_findPathByEntityID_V2(entityID1, entityID2, maxDegree,
            flags, &charBuff, &bufferSize, resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2_findPathByRecordID_local(const char* dataSourceCode1,
        const char* recordID1, const char* dataSourceCode2,
        const char* recordID2, const int maxDegree) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    int returnCode = G2_findPathByRecordID(dataSourceCode1, recordID1,
            dataSourceCode2, recordID2, maxDegree, &charBuff, &bufferSize,
            resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2_findPathByRecordID_V2_local(const char* dataSourceCode1,
        const char* recordID1, const char* dataSourceCode2,
        const char* recordID2, const int maxDegree, const long long flags) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    int returnCode = G2_findPathByRecordID_V2(dataSourceCode1, recordID1,
            dataSourceCode2, recordID2, maxDegree, flags, &charBuff,
            &bufferSize, resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2_findPathExcludingByEntityID_local(const long long entityID1,
        const long long entityID2, const int maxDegree,
        const char* excludedEntities) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    int returnCode = G2_findPathExcludingByEntityID(entityID1, entityID2,
            maxDegree, excludedEntities, &charBuff, &bufferSize,
            resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2_findPathExcludingByEntityID_V2_local(const long long entityID1,
        const long long entityID2, const int maxDegree,
        const char* excludedEntities, const long long flags) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    int returnCode = G2_findPathExcludingByEntityID_V2(entityID1, entityID2,
            maxDegree, excludedEntities, flags, &charBuff, &bufferSize,
            resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2_findPathExcludingByRecordID_local(const char* dataSourceCode1,
        const char* recordID1, const char* dataSourceCode2,
        const char* recordID2, const int maxDegree, const char* excludedRecords) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    int returnCode = G2_findPathExcludingByRecordID(dataSourceCode1, recordID1,
            dataSourceCode2, recordID2, maxDegree, excludedRecords, &charBuff,
            &bufferSize, resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2_findPathExcludingByRecordID_V2_local(const char* dataSourceCode1,
        const char* recordID1, const char* dataSourceCode2,
        const char* recordID2, const int maxDegree, const char* excludedRecords,
        const long long flags) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    int returnCode = G2_findPathExcludingByRecordID_V2(dataSourceCode1,
            recordID1, dataSourceCode2, recordID2, maxDegree, excludedRecords,
            flags, &charBuff, &bufferSize, resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2_findPathIncludingSourceByEntityID_local(const long long entityID1,
        const long long entityID2, const int maxDegree,
        const char* excludedEntities, const char* requiredDsrcs) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    int returnCode = G2_findPathIncludingSourceByEntityID(entityID1, entityID2,
            maxDegree, excludedEntities, requiredDsrcs, &charBuff, &bufferSize,
            resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2_findPathIncludingSourceByEntityID_V2_local(const long long entityID1,
        const long long entityID2, const int maxDegree,
        const char* excludedEntities, const char* requiredDsrcs,
        const long long flags) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    int returnCode = G2_findPathIncludingSourceByEntityID_V2(entityID1,
            entityID2, maxDegree, excludedEntities, requiredDsrcs, flags,
            &charBuff, &bufferSize, resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2_stats_local() {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2_resizeStringBuffer;
    int returnCode = G2_stats(&charBuff, &bufferSize, resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}