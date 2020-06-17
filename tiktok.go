package main

import (
	"encoding/json"
	"github.com/gocolly/colly"
	"log"
)
type Data struct {
	ContentUrl string `json:"contentUrl"`
}

func newScraper() *colly.Collector {
	c := colly.NewCollector()
	c.UserAgent="Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:28.0) Gecko/20100101 Firefox/28.0"
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		r.Headers.Set("Accept-Encoding","gzip, deflate")
		r.Headers.Set("Accept-Language","en-US,en;q=0.9")
	})
	return c
}

func getVideoLink(copiedLink string,scraper *colly.Collector)(Data,error){
	data := Data{}
	scraper.OnHTML("script[id=videoObject]", func(e *colly.HTMLElement) {
		err:=json.Unmarshal([]byte(e.Text),&data)
		checkErr(err)
	})
	err:=scraper.Visit(copiedLink)
	checkErr(err)
	return data,nil
}

func checkErr(err error) {
	if err!=nil{
		log.Fatal(err)
	}
}