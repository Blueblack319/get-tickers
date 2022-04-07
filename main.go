package main

import (
	"fmt"
	"net/http"

	"github.com/Blueblack319/go-crawler/tools"
	"github.com/PuerkitoBio/goquery"
)

// Variables
const URL string = "https://etfdb.com/etfs/asset-class/bond/"

//MAIN
// 1. Get etf 
// 2. Get page O
// 3. loop through paginations
func main(){
	doc, err := getDocument(URL)
	tools.CheckError(err)
	
	tickerList := doc.Find(`table[data-hash="etfs"] tbody tr`)
	tickerList.Each(func(idx int, s *goquery.Selection){
		ticker := s.Find(`td[data-th="Symbol"] a`).Text()
		fmt.Println(ticker)
	})
}

// Functions

func getDocument(baseURL string) (*goquery.Document, error){
	res, err := http.Get(baseURL)
	tools.CheckError(err)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	return doc, err
}
