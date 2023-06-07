package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type GBP struct {
	GBP float64 `json:"GBP"`
}

type CurrencyData struct {
	Data GBP `json:"data"`
}


func main() {
	key := os.Getenv("FREE_CURRENCY_API_KEY")
	urlTemplate := "https://api.freecurrencyapi.com/v1/latest?apikey=%s&currencies=GBP&base_currency=JPY"

	if (key == "") { 
		log.Fatalln("Please set you Free Currency API Key")
	}

	url := fmt.Sprintf(urlTemplate, key)

	res, err := http.Get(url);

	if err != nil {
		log.Fatalln("Couldn't get request");
	}

	defer res.Body.Close()

	var jsonResponse CurrencyData

	if err := json.NewDecoder(res.Body).Decode(&jsonResponse); err != nil {
		log.Println(err)
		log.Fatalln("Can not convert to JSON")
	}

	fmt.Println(jsonResponse.Data.GBP)
}