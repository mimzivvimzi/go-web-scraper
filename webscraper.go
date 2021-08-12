package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type Listing struct {
	// ID int `json:"id"`
	Description string `json:"description"`
	Location string `json:"location"`
	// Link string `json:"link"`
}

func main() {
	allListings := make([]Listing, 0)

	c := colly.NewCollector()

	c.OnHTML(".listing-item.listing-title", func(e *colly.HTMLElement) {
		// fmt.Println(e.ChildText("span"))
		listingDesc := e.ChildText(".text-semi-strong")
		listingLoc := e.ChildText("span")
		// link := e.Attr("href")
		

		listing := Listing{
			Description: listingDesc,
			Location: listingLoc,
			// Link: link,
		}

		allListings = append(allListings, listing)
		// test := e.Request.Visit(e.Attr("href"))
		// fmt.Println(test)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Index(link, "/en/rent/view/") != -1 {
			fmt.Println(link)
			// c.Visit(link)
		}
		// fmt.Println(link)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://apartments.gaijinpot.com/en/rent/listing?page=3")

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "")
	enc.Encode(allListings)
	// log.Println(c)
}