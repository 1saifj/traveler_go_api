package utils

import (
	"github.com/disintegration/imaging"
	"image"
)

func GetImageInStorage(path string) (image.Image, error) {
	return imaging.Open(path, imaging.AutoOrientation(true))
}

func ImageResizerBySize(img image.Image, size string) image.Image {
	var nImg image.Image
	switch size {
	case "small":
		nImg = imaging.Resize(img, 256, 256, imaging.Lanczos)
	case "medium":
		nImg = imaging.Resize(img, 512, 512, imaging.Lanczos)

	case "large":
		nImg = imaging.Resize(img, 1000, 1000, imaging.Lanczos)
	default:
		nImg = imaging.Resize(img, 100, 12, imaging.Lanczos)
	}
	return nImg
}
