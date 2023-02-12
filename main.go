package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/fogleman/gg"
)

// The data struct for the decoded data
// Notice that all fields must be exportable! ie start with uppercase !!!
type Config struct {
	Apikey  string
	Symbols []string
}

type Quote struct {
	Price          float32 `json:"c"`
	change         float32 `json:"d"`
	Percent_change float32 `json:"dp"`
	High           float32 `json:"h"`
	Low            float32 `json:"l"`
	open           float32 `json:"o"`
	previous_close float32 `json:"pc"`
	t              int
}

func main() {

	// Let's first read the `config.json` file
	jsonBlob, err := os.ReadFile("./config.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `config`
	var config Config
	err = json.Unmarshal(jsonBlob, &config)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// loop through symbols
	for _, symbol := range config.Symbols {

		log.Print("symbol is " + symbol)
		getQuote(symbol, config.Apikey)
	}
}

func getQuote(aSymbol string, apiKey string) {

	url := "https://finnhub.io/api/v1/quote?symbol=" + aSymbol + "&token=" + apiKey

	// factorize
	log.Print(url)

	req, _ := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal("Error during quote fetching: ", err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//	fmt.Println(res)
	fmt.Println(string(body))

	var quote Quote
	err = json.Unmarshal(body, &quote)
	if err != nil {
		log.Fatal("Unable to parse quote for "+aSymbol+": ", err)
	}

	log.Printf("Ticker %s value is %.2f (%.1f PCT)", aSymbol, quote.Price, quote.Percent_change)

	outputImage("$" + aSymbol + "=" + fmt.Sprint(quote.Price) + "(" + fmt.Sprintf("%.2f", quote.Percent_change) + "%)")

}

func outputImage(s string) {
	dc := gg.NewContext(800, 600)
	dc.LoadFontFace("fonts/impact.ttf", 80)

	// We declare a Rectangle with a given Width and Height, starting in the (0, 0) pixel.
	dc.DrawRectangle(0, 0, float64(dc.Width()), float64(dc.Height()))
	dc.SetHexColor("#FFFFFF") // We select it's colour.
	dc.Fill()                 // And fill the context with it.

	// text

	/*
		err := loadFont(dc, fontname)
		if err != nil {
			return err
		}
	*/
	c := "#777"
	dc.SetHexColor(c) // Set the text colour.
	dc.DrawString(s, 10, 100)

	dc.SavePNG("out.png")
}

// generate PNG 800x600 landscape

//
