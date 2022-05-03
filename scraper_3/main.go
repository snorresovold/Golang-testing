package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	var wg sync.WaitGroup

	a := []string{"https://www.vg.no", "https://www.nrk.no/"}
	for i, s := range a {
		wg.Add(i)
		fmt.Println(i, s)
		go check_site(s)
		{
			defer wg.Done()
			worker(i)
		}
	}
	wg.Wait()
}

func check_site(x string) {
	c := colly.NewCollector()
	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})
	// before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.Visit(x)
}

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}
