package main

import (
	"encoding/csv"
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
	defer file.Close() // closes the file

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		colly.AllowedDomains("vg.no"),
	)

	c.OnHTML("article article-extract article-width-full has-article-image has-full-width-image newsroom-vg skin-default", func(e *colly.HTMLElement) {

		writer.Write([]string{
			e.ChildText("h3"),
			e.ChildText("a"),
		})
	})

}
