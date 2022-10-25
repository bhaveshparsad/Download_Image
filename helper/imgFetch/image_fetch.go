package imgFetch

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"Download_Image/helper/imgDownload"

	"golang.org/x/net/html"
)

func ImageFetch(url string) string {
	result := make([]string, 0)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	html_1, err := ioutil.ReadAll(resp.Body) 
	if err != nil {
		log.Fatal(err)
	}

	html_2 := string(html_1) 

	doc, err := html.Parse(strings.NewReader(html_2))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {
			for _, img := range n.Attr {
				if img.Key == "src" {
					result = append(result, img.Val)

				}
			}

		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	images := result
	message := imgDownload.ImageDownload(images)
	return message
}
