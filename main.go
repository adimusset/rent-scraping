package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	// pour 1 type de donnée
	// 1 scheduler
	// n crawlers
	// p writers
	duration := flag.Int("duration", 10, "duration between each fetching and writing")
	var guests int
	flag.IntVar(&guests, "guests", 1, "number of guests")
	ticker := time.NewTicker(time.Duration(*duration) * time.Second)
	crawler := NewAirbnbCrawler()
	writer := NewLocalWriter("test_data", "/tmp")
	flag.Parse()

	for range ticker.C {
		fmt.Println("fetching airbnb data")
		objs, err := crawler.Crawl(guests)
		if err != nil {
			fmt.Println("Error while fetching airbnb data")
			return
		}
		fmt.Println("writing airbnb data")
		err = writer.Write(objs, guests)
		if err != nil {
			fmt.Println("Error while writing airbnb data")
			return
		}
		fmt.Println("Done")
		fmt.Println(time.Now().Format(time.RFC3339))
	}
}
