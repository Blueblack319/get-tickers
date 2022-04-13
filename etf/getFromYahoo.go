package etf

import (
	"fmt"

	"github.com/Blueblack319/go-crawler/tools"
	"github.com/PuerkitoBio/goquery"
)

const YAHOO_URL string = "https://finance.yahoo.com/world-indices"

func GetFromYahoo(){
	symbols_list := []string{}
	doc, err := getDocument(YAHOO_URL)
	tools.CheckError(err)

	symbols := doc.Find("div#list-res-table tr")
	symbols.Each(func(idx int, s *goquery.Selection){
		symbols_list = append(symbols_list, s.Find("td a").Text())
	})
	fmt.Print(symbols_list)
}