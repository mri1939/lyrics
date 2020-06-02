package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	if len(os.Args) != 3 {
		fmt.Printf("usage: %s <title> <artist>\n", os.Args[0])
		os.Exit(1)
	}
	title := os.Args[1]
	artist := os.Args[2]
	req, err := http.NewRequest("GET", "https://makeitpersonal.co/lyrics", nil)
	if err != nil {
		log.Fatal("Failed to create request")
	}
	q := req.URL.Query()
	q.Add("artist", artist)
	q.Add("title", title)
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Failed while doing the request", err)
	}
	io.Copy(os.Stdout, resp.Body)
	fmt.Println()
}
