package main

import (
	"net/http"
	"net/http/httputil"
	"fmt"
)

func main() {
	request, err := http.NewRequest(http.MethodGet, "https://www.imooc.com", nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like)")
	resp, err := http.Get("https://www.imooc.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}
