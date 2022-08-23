#include <stdlib.h>
#include <stdio.h>
#include "libg2diagnostic.h"

typedef void* EntityListBySizeHandle;
typedef void*(*resize_buffer_type)(void *, size_t);

void* G2Diagnostic_resizeStringBuffer(void *ptr, size_t size);
char* G2Diagnostic_checkDBPerf_local(int secondsToRun);
char* G2Diagnostic_findEntitiesByFeatureIDs_local(const char *features);
char* G2Diagnostic_getDataSourceCounts_local();
char* G2Diagnostic_getDBInfo_local();
char* G2Diagnostic_getEntityDetails_local(const long long entityID, const int includeInternalFeatures);
char* G2Diagnostic_getEntityResume_local(const long long entityID);
char* G2Diagnostic_getEntitySizeBreakdown_local(const size_t minimumEntitySize, const int includeInternalFeatures);
char* G2Diagnostic_getFeature_local(const long long libFeatID);
char* G2Diagnostic_getGenericFeatures_local(const char *featureType, const size_t maximumEstimatedCount);
char* G2Diagnostic_getMappingStatistics_local(const int includeInternalFeatures);
char* G2Diagnostic_getRelationshipDetails_local(const long long relationshipID, const int includeInternalFeatures);
char* G2Diagnostic_getResolutionStatistics_local();
