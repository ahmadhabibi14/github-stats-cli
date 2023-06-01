package app

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
)

func profile_image(url string) (image.Image, error) {
	fetchImage, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var images []image.Image
	img, _, err := image.Decode(fetchImage.Body)
	if err != nil {
		return nil, err
	}
	images = append(images, img)

	return images[0], nil
}
