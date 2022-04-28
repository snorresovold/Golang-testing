package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {

	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, err :%q", err)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector()

	// Visiting all links
	c.OnHTML("a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
		writer.Write([]string{
			e.Attr("href"),
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://go-colly.org/")

}
