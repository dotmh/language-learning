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

type Coins struct {
	C500 int
	C100 int
	C50 int
	C10 int
	C5 int
	C1 int
}

type Notes struct {
	N10000 int
	N5000 int
	N1000 int
}

func main() {
	key := os.Getenv("FREE_CURRENCY_API_KEY")
	urlTemplate := "https://api.freecurrencyapi.com/v1/latest?apikey=%s&currencies=GBP&base_currency=JPY"
	coins := Coins {
		C500: 500,
		C100: 100,
		C50: 50,
		C10: 10,
		C5: 5,
		C1: 1,
	}

	notes := Notes {
		N10000: 10000,
		N5000: 5000,
		N1000: 1000,
	}

	wallet := [9]int {
		notes.N10000 * 5,
		notes.N5000 * 1,
		notes.N1000 * 5,
		coins.C500 * 1,
		coins.C100 * 4,
		coins.C50 * 5,
		coins.C10 * 8,
		coins.C5 * 7,
		coins.C1 * 14,
	}

	if (key == "") { 
		log.Fatalln("Please set you Free Currency API Key")
	}

	yen := 0

	for _, num := range wallet {
		yen += num
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

	gbp := float64(yen) * jsonResponse.Data.GBP;

	fmt.Printf("You have ¥%d , which is £%.2f \n", yen , gbp)

}