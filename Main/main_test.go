package main

import (
	"testing"

	"Download_Image/Image_Fetch"
)


func TestGetImage(t *testing.T){
	url := "https://www.tftus.com"
	expectedOutput := "Successful"
	output := Image_Fetch.ImageFetch(url)

	if output != expectedOutput{
		t.Errorf("got %q, wanted %q", output, expectedOutput)
	}

}

