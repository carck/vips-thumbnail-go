package main

import (
	"github.com/carck/vips-thumbnail-go"
)

func main() {
	vips.Start()
	defer vips.Shutdown()

	vips.Thumbnail("./abc.jpg", "123.jpg", 224, 224, 0)
}
