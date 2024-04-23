package vips_test

import (
	_ "embed"
	"testing"

	"github.com/carck/vips-thumbnail-go"
)

func setupTest() func() {
	vips.Start()

	return func() {
		vips.Shutdown()
	}
}

func BenchmarkThumbnail(b *testing.B) {

	for i := 0; i < b.N; i++ {
		err := vips.Thumbnail("./testdata/rgba.jpg", "a.jpg[optimize_coding,keep=none,Q=90]", 224, 224, 0, "srgb")
		if err != nil {
			b.Fatal(err)
		}
	}
}
