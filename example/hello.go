package main

import (
	"github.com/carck/vips-thumbnail-go"
)

func main() {
	vips.Start()
	defer vips.Shutdown()

	err := vips.Thumbnail("./IMG_8638.heic", "123.jpg", 500, 500, -1, 90, "srgb", true)
	if err != nil {
		panic(err)
	}
}
