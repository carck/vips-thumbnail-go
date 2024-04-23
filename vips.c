#include "vips.h"

int thumbnail(const char *filename, const char *outputname, int width, int height, int crop, int q)
{
    int ret;
    VipsImage *image;

    if (crop == -1)
    {
        ret = vips_thumbnail(filename, &image, width, "export-profile", "srgb", NULL);
    }
    else
    {
        ret = vips_thumbnail(filename, &image, width, "height", height, "crop", crop, "export-profile", "srgb",NULL);
    }

    if (ret)
    {
        return -1;
    }

    if (vips_image_write_to_file(image, outputname, "Q", q, NULL))
    {
        VIPS_UNREF(image);
        return (-1);
    }
    VIPS_UNREF(image);
    return 0;
}