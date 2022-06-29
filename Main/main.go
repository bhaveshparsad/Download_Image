package main

import(
	"fmt"
	"Download_Image/Image_Fetch"
)

func main() {
	url := "https://www.tftus.com"

	message := Image_Fetch.ImageFetch(url)

	fmt.Println(message)
}