package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	client := http.Client{
		// Timeout: time.Second,
	}

	req, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
