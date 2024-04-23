#ifndef VIPS_THUMBNAIL_H

#define VIPS_THUMBNAIL_H

#include <stdlib.h>
#include <vips/vips.h>

int thumbnail(const char *filename, const char *outputname, int width, int height, int crop, const char* export_profile);

#endif