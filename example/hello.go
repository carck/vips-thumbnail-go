package main

import (
	"github.com/carck/vips-thumbnail-go"
)

func main() {
	vips.Start()
	defer vips.Shutdown()

	err := vips.Thumbnail("./abc.jpg", "123.jpg[Q=90,optimize_coding,keep=none,subsample-mode=on]", 224, 224, 0, "srgb")
	if err != nil {
		panic(err)
	}
}
