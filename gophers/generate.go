package main

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/qeesung/image2ascii/convert"
)

func main() {

	// Display usage/help message
	if len(os.Args) == 1 || (len(os.Args) == 2 && os.Args[1] == "-h") || (len(os.Args) == 2 && os.Args[1] == "--help") {
		usage := "Generate a Gopher image to ASCII text.\n\nUsage:\n   generate $image\n\nExample:\n   generate path/image.png"

		fmt.Println(usage)
		return
	} else {

		imageFileName := os.Args[1]

		// Create convert options
		convertOptions := convert.DefaultOptions
		convertOptions.FixedWidth = 70
		convertOptions.FixedHeight = 40
		convertOptions.Reversed = true

		// Create the image converter
		converter := convert.NewImageConverter()
		fmt.Print(converter.ImageFile2ASCIIString(imageFileName, &convertOptions))
	}

}
