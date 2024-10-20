package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("https://academicguides.waldenu.edu/"),//change the site name here 
	)


	c.OnHTML("h2", func(e *colly.HTMLElement) {
		fmt.Println("Text found:", e.Text)
	})

	g
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	
	err := c.Visit("https://academicguides.waldenu.edu/writingcenter/grammar/articles")//and here too
	if err != nil {
		log.Fatal(err)
	}
}
