#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include "libg2product.h"
#include "g2product.h"

void* G2Product_resizeStringBuffer(void* ptr, size_t size) {
    //deallocate old buffer
    if (ptr != 0)
        free(ptr);
    //allocate new buffer
    void* buffer = malloc(size);
    return buffer;
}