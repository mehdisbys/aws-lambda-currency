package main

import (
	"encoding/json"
	"github.com/apex/go-apex"
	"os"
	"fmt"
	"io/ioutil"
	"net/http"
)
import _ "github.com/joho/godotenv/autoload"

type message struct {
	Hello string `json:"hello"`
}

type CurrencyAPIResponse struct {
	Success   bool   `json:"success"`
	Timestamp int    `json:"timestamp"`
	Source    string `json:"source"`
	Quotes struct {
		GBP     float64 `json:"USDGBP"`
		EUR     float64 `json:"USDEUR"`
		Bitcoin float64 `json:"USDBTC"`
	} `json:"quotes"`
}

func GetQuotes() (CurrencyAPIResponse) {

	response, err := http.Get("http://apilayer.net/api/live?access_key=" + os.Getenv("CURRENCY_API_KEY") + "&format=1")

	var currencies CurrencyAPIResponse

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		quotes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		if err := json.Unmarshal(quotes, &currencies); err != nil {
			panic(err)
		}

		if currencies.Success != true {
			panic("Could not retrieve response from API.")
		}
	}

	return currencies;
}

func SayHello() map[string]string {

	return map[string]string{"hello": "go"}
}

func main() {

//	fmt.Printf("%+v\n", GetQuotes());

	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		return GetQuotes(), nil
	})

}
