package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func makeRequest(url string) string {
	res, err := http.Get(url);

	if err != nil {
		log.Fatalln("Couldn't get request");
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalln("Couldn't read the body");
	}

	bodyText := string(body)
	return bodyText;
}

func main() {

	learn := makeRequest("https://learn.api.dotmh.dev");

	fmt.Println(learn);
}