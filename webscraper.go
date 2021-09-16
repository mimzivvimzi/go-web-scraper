package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Product struct {
	Title string `json:"title"`
	Price string `json:"price"`
}

func main() {
	allProducts := make([]Product, 0)

	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if !strings.HasPrefix(link, "/product/") {
			return
		}
		e.Request.Visit(link)
	})
	
	c.OnHTML(".col-sm-4.col-lg-4.col-md-4", func(e *colly.HTMLElement) {

		title := e.ChildText(".title")
		price := e.ChildText(".pull-right.price")

		product := Product{
			Title: title,
			Price: price,
		}
	
		allProducts = append(allProducts, product)
	})

	for i := 2; i <= 20; i++ {
		fmt.Println("Scraping page: ", i)

		c.Visit("https://webscraper.io/test-sites/e-commerce/static/computers/laptops?page=" + strconv.Itoa(i))
	}

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://webscraper.io/test-sites/e-commerce/static/computers/laptops")

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "")
	enc.Encode(allProducts)
	// log.Println(c)
}

