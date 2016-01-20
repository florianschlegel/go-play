package client

import (
	"fmt"
	"net/http"
)

// CallURL calls a url and prints interesting things to std out
func CallURL(url string) {
	resp, err := http.Get(url)
	fmt.Println("	calling:", url)
	fmt.Println("	result:", "that is the response Status", resp.Status, "this is the error", err)
}
