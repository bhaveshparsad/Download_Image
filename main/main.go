package main

import(
	"fmt"
	"Download_Image/imgFetch"
)

func main() {
	url := "https://www.tftus.com"

	message := imgFetch.ImageFetch(url)

	fmt.Println(message)
}