#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include "libg2config.h"

typedef void* ConfigHandle;
typedef void*(*resize_buffer_type)(void *, size_t);


int G2config_close_helper(ConfigHandle configHandle);
void* G2config_create_helper();

