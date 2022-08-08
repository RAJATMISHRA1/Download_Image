package main

import (
	"Download_Image/imgFetch"
	"fmt"
)

func main() {
	url := "https://www.tftus.com/"

	message := imgFetch.ImageFetch(url)

	fmt.Println(message)
}
