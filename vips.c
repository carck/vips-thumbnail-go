#include "vips.h"

int thumbnail(const char *filename, const char *outputname, int width, int height, int crop, const char* export_profile)
{
    int ret;
    VipsImage *image;

    if (crop == -1)
    {
        ret = vips_thumbnail(filename, &image, width, "export-profile", export_profile, NULL);
    }
    else
    {
        ret = vips_thumbnail(filename, &image, width, "height", height, "crop", crop, "export-profile", export_profile, NULL);
    }

    if (ret)
    {
        return -1;
    }

    if (vips_image_write_to_file(image, outputname, NULL))
    {
        VIPS_UNREF(image);
        return (-1);
    }
    VIPS_UNREF(image);
    return 0;
}