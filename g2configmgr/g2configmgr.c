#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include "libg2configmgr.h"
#include "g2configmgr.h"

void* G2config_resizeStringBuffer(void* ptr, size_t size) {
    //deallocate old buffer
    if (ptr != 0)
        free(ptr);
    //allocate new buffer
    void* buffer = malloc(size);
    return buffer;
}