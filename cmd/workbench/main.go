package main

import (
	"fmt"
	"os"

	"github.com/h2non/bimg"
)

func main() {
	buffer, err := bimg.Read("./tmp/files/ce8815c1-cb05-448d-94c8-f7dd5a56461b_1698265722_0.jpg")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	newImage, err := bimg.NewImage(buffer).Resize(200, 200)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	size, err := bimg.NewImage(newImage).Size()
	if size.Width == 200 && size.Height == 200 {
		fmt.Println("The image size is valid")
	}

	bimg.Write("./tmp/files/new.jpg", newImage)
}
