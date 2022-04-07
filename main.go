package main

import (
	"fmt"
	"net/http"

	"github.com/Blueblack319/go-crawler/tools"
	"github.com/PuerkitoBio/goquery"
)

// Variables
const URL string = "https://etfdb.com/etfs/asset-class/"


//MAIN
// Get page O
// Get etf O 
// Get tickers O 
// loop through paginations
func main(){
	var classes = []string{"equity", "bond", "real-estate", "currency", "alternatives"}
	tickers, err := getTickersByClass(classes)
	tools.CheckError(err)

	fmt.Print(tickers)
}

// Functions
func getDocument(url string) (*goquery.Document, error){
	res, err := http.Get(url)
	tools.CheckError(err)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	return doc, err
}

func getTickers(class string, ch chan<- []string){
	doc, err := getDocument(URL+class)
	tools.CheckError(err)

	sub := []string{}

	tickerList := doc.Find(`table[data-hash="etfs"] tbody tr`)
	tickerList.Each(func(idx int, s *goquery.Selection){
		sub = append(sub, s.Find(`td[data-th="Symbol"] a`).Text())
	})

	ch <- sub
}

func getTickersByClass(classes []string) ([]string, error){
	tickers := []string{}
	ch := make(chan []string)

	for _, s := range classes{
		go getTickers(s, ch)
	}
	
	for i := 0; i < len(classes); i++{
		extracted := <-ch
		tickers = append(tickers, extracted...)
	}
	return tickers, nil
}