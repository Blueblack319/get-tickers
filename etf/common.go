package etf

import (
	"net/http"

	"github.com/Blueblack319/go-crawler/tools"
	"github.com/PuerkitoBio/goquery"
)

func getDocument(url string) (*goquery.Document, error){
	res, err := http.Get(url)
	tools.CheckError(err)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	return doc, err
}