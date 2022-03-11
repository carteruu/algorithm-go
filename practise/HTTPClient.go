package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HTTPClient
func main() {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr, Timeout: 1 * time.Second}
	resp, err := client.Get("http://127.0.0.1:8080/")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	fmt.Println(ioutil.ReadAll(resp.Body))
}
