#include <stdlib.h>
#include <stdio.h>
#include "libg2diagnostic.h"
#include "g2diagnostic.h"

void* G2Diagnostic_resizeStringBuffer(void *ptr, size_t size) {
    //deallocate old buffer
    if (ptr != 0)
        free(ptr);
    //allocate new buffer
    void* buffer = malloc(size);
    return buffer;
}

char* G2Diagnostic_checkDBPerf_local(int secondsToRun) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2Diagnostic_resizeStringBuffer;
    int returnCode = G2Diagnostic_checkDBPerf(secondsToRun, &charBuff,
            &bufferSize, resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2Diagnostic_findEntitiesByFeatureIDs_local(const char *features) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2Diagnostic_resizeStringBuffer;
    int returnCode = G2Diagnostic_findEntitiesByFeatureIDs(features, &charBuff,
            &bufferSize, resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2Diagnostic_getDataSourceCounts_local() {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2Diagnostic_resizeStringBuffer;
    int returnCode = G2Diagnostic_getDataSourceCounts(&charBuff, &bufferSize,
            resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2Diagnostic_getDBInfo_local() {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2Diagnostic_resizeStringBuffer;
    int returnCode = G2Diagnostic_getDBInfo(&charBuff, &bufferSize,
            resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2Diagnostic_getEntityDetails_local(const long long entityID,
        const int includeInternalFeatures) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2Diagnostic_resizeStringBuffer;
    int returnCode = G2Diagnostic_getEntityDetails(entityID,
            includeInternalFeatures, &charBuff, &bufferSize, resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2Diagnostic_getEntityResume_local(const long long entityID) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2Diagnostic_resizeStringBuffer;
    int returnCode = G2Diagnostic_getEntityResume(entityID, &charBuff,
            &bufferSize, resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2Diagnostic_getEntitySizeBreakdown_local(const size_t minimumEntitySize,
        const int includeInternalFeatures) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2Diagnostic_resizeStringBuffer;
    int returnCode = G2Diagnostic_getEntitySizeBreakdown(minimumEntitySize,
            includeInternalFeatures, &charBuff, &bufferSize, resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2Diagnostic_getFeature_local(const long long libFeatID) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2Diagnostic_resizeStringBuffer;
    int returnCode = G2Diagnostic_getFeature(libFeatID, &charBuff, &bufferSize,
            resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2Diagnostic_getGenericFeatures_local(const char *featureType,
        const size_t maximumEstimatedCount) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2Diagnostic_resizeStringBuffer;
    int returnCode = G2Diagnostic_getGenericFeatures(featureType,
            maximumEstimatedCount, &charBuff, &bufferSize, resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2Diagnostic_getMappingStatistics_local(const int includeInternalFeatures) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2Diagnostic_resizeStringBuffer;
    int returnCode = G2Diagnostic_getMappingStatistics(includeInternalFeatures,
            &charBuff, &bufferSize, resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2Diagnostic_getRelationshipDetails_local(const long long relationshipID,
        const int includeInternalFeatures) {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2Diagnostic_resizeStringBuffer;
    int returnCode = G2Diagnostic_getRelationshipDetails(relationshipID,
            includeInternalFeatures, &charBuff, &bufferSize, resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}

char* G2Diagnostic_getResolutionStatistics_local() {
    size_t bufferSize = 1;
    char *charBuff = (char *) malloc(1);
    resize_buffer_type resizeFuncPointer = &G2Diagnostic_resizeStringBuffer;
    int returnCode = G2Diagnostic_getResolutionStatistics(&charBuff,
            &bufferSize, resizeFuncPointer);
    if (returnCode != 0) {
        return "";
    }
    return charBuff;
}
