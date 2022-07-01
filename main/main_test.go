package main

import (
	"testing"

	"Download_Image/imgFetch"
)


func TestGetImage(t *testing.T){
	url := "https://www.tftus.com"
	expectedOutput := "Successful"
	output := imgFetch.ImageFetch(url)

	if output != expectedOutput{
		t.Errorf("got %q, wanted %q", output, expectedOutput)
	}

}

