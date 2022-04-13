package etf

import (
	"fmt"

	"github.com/Blueblack319/go-crawler/tools"
	"github.com/PuerkitoBio/goquery"
)

const URL string = "https://etfdb.com/etfs/asset-class/"

func getTickersEtfdb(class string, ch chan<- []string){
	doc, err := getDocument(URL + class)
	tools.CheckError(err)

	sub := []string{class}

	tickerList := doc.Find(`table[data-hash="etfs"] tbody tr`)
	tickerList.Each(func(idx int, s *goquery.Selection){
		sub = append(sub, s.Find(`td[data-th="Symbol"] a`).Text())
	})

	ch <- sub
}

func getTickersByClass(classes []string) (map[string][]string, error){
	tickers := make(map[string][]string)
	ch := make(chan []string)

	for _, s := range classes{
		go getTickersEtfdb(s, ch)
	}
	OutLoop:
		for i := 0; i < len(classes); i++{
			extracted := <-ch
			for _, x := range classes{
				if extracted[0] == x{
					tickers[x] = extracted[1:]
					continue OutLoop
				}
			}
		}
	
	return tickers, nil
}

func GetFromEtfdb(){
	var classes = []string{"equity", "bond", "real-estate", "currency", "alternatives"}
	tickers, err := getTickersByClass(classes)
	tools.CheckError(err)

	for i, t := range tickers{
		fmt.Println(i, t)
	}
}