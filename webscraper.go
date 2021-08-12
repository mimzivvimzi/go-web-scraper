package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type Product struct {
	// ID int `json:"id"`
	Title string `json:"title"`
	Price string `json:"price"`
	// Link string `json:"link"`
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
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		title := e.ChildText(".title")
		price := e.ChildText(".pull-right.price")

		// test := e.Request.Visit(e.Attr("href"))
		// fmt.Println(link)

		product := Product{
			Title: title,
			Price: price,
		}
	
		allProducts = append(allProducts, product)
	})




	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	link = e.Attr("href")
	// 	// if strings.Index(link, "/en/rent/view/") != -1 {
	// 	// 	fmt.Println(link)
	// 	// 	// c.Visit(link)
	// 	// }
	// 	fmt.Println(link)
	// })




	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://webscraper.io/test-sites/e-commerce/static/computers/laptops")

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "")
	enc.Encode(allProducts)
	// log.Println(c)
}