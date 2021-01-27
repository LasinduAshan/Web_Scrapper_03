package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	c.OnHTML("h1", func(element *colly.HTMLElement) {
		fmt.Println("Vehicle: ", element.Text)
	})
	c.OnHTML(".amount--3NTpl", func(element *colly.HTMLElement) {
		fmt.Println("Price: ", element.Text)
	})
	c.OnHTML(".two-columns--19Hyo", func(element *colly.HTMLElement) {
		fmt.Println(element.Text)
	})
	desc := false
	c.OnHTML(".description--1nRbz > p", func(element *colly.HTMLElement) {
		if !desc {
			fmt.Println("Description: ", element.Text)
		} else {
			fmt.Println(element.Text)
		}
		desc = true
	})

	_ = c.Visit("https://ikman.lk/en/ad/suzuki-wagon-r-stingray-2015-for-sale-colombo-1416")
}
