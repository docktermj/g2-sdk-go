#include <stdlib.h>
#include <stdio.h>
#include "libg2config.h"

typedef void* ExportHandle;
typedef void*(*resize_buffer_type)(void *, size_t);


long long G2config_create_helper();
