package vips

/*
#cgo pkg-config: vips
#include "vips.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

const (
	InterestingOff       int = 0
	InterestingNone      int = C.VIPS_INTERESTING_NONE
	InterestingCentre    int = C.VIPS_INTERESTING_CENTRE
	InterestingEntropy   int = C.VIPS_INTERESTING_ENTROPY
	InterestingAttention int = C.VIPS_INTERESTING_ATTENTION
	InterestingLow       int = C.VIPS_INTERESTING_LOW
	InterestingHigh      int = C.VIPS_INTERESTING_HIGH
	InterestingAll       int = C.VIPS_INTERESTING_ALL
	InterestingLast      int = C.VIPS_INTERESTING_LAST
)

func Start() {
	appName := C.CString("photoprism")
	defer C.free(unsafe.Pointer(appName))

	err := C.vips_init(appName)
	if err != 0 {
		panic("unable to start vips!")
	}
	C.vips_concurrency_set(1)
}

func Shutdown() {
	C.vips_shutdown()
}

func ThumbnailDefault(fileName, thumbnailFileName string, width, height, crop, q int) error {
	return Thumbnail(fileName, thumbnailFileName, width, height, crop, q, "", true)
}

func Thumbnail(fileName, thumbnailFileName string, width, height, crop, q int, exportProfile string, strip bool) error {
	cFileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cFileName))

	keep := "icc"
	if strip {
		keep = "none"
	}
	cOutName := C.CString(fmt.Sprintf("%s[Q=%d,optimize_coding,keep=%s,subsample-mode=auto]", thumbnailFileName, q, keep))
	defer C.free(unsafe.Pointer(cOutName))

	var cExportProfile *C.char = nil
	if exportProfile != "" {
		cExportProfile = C.CString(exportProfile)
		defer C.free(unsafe.Pointer(cExportProfile))
	}

	err := C.thumbnail(cFileName,
		cOutName,
		C.int(width),
		C.int(height),
		C.int(crop),
		cExportProfile)
	if err != 0 {
		return handleVipsError()
	}
	return nil
}

func handleVipsError() error {
	s := C.GoString(C.vips_error_buffer())
	C.vips_error_clear()

	return fmt.Errorf("%v\n", s)
}
