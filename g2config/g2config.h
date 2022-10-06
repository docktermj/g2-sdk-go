#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include "libg2config.h"

typedef void* ExportHandle;
typedef void*(*resize_buffer_type)(void *, size_t);


uintptr_t G2config_create_helper();
