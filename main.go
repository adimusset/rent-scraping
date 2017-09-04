package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/DataDog/datadog-go/statsd"
)

func main() {
	var duration int
	flag.IntVar(&duration, "duration", 120, "duration between each write")
	var guests int
	flag.IntVar(&guests, "guests", 1, "number of guests")
	var path string
	flag.StringVar(&path, "path", "/tmp", "path where the data is saved")
	crawler := NewAirbnbCrawler()
	writer := NewLocalWriter("test_data", path)
	flag.Parse()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	ddClient, err := statsd.New("127.0.0.1:8125")
	if err != nil {
		log.Fatal(err)
	}

	logger.Printf("rent-scraping starting with configuration: guests %d duration %d path %s",
		guests, duration, path)
	ddClient.Count("app.started", 1, []string{"app:rent"}, 1)

	for {
		start := time.Now()
		logger.Printf("Starting to fetch data")
		objs, err := crawler.Crawl(guests)
		if err != nil {
			logger.Printf("Error while fetching data: %s", err.Error())
			return
		}
		end := time.Now()
		ddClient.Histogram("fetching_time_ms", end.Sub(start).Seconds()*1000, []string{"app:rent"}, 1)
		logger.Printf("Fetching completed in %.2f seconds", end.Sub(start).Seconds())
		logger.Printf("Writing %d airbnb objects", len(objs))
		err = writer.Write(objs, guests)
		if err != nil {
			logger.Printf("Error while writing data: %s", err.Error())
			return
		}
		total := time.Now().Sub(start)
		ddClient.Histogram("writing_time_ms", time.Now().Sub(end).Seconds()*1000, []string{"app:rent"}, 1)
		logger.Printf("Writing completed in %.2f seconds - total time %.2f seconds",
			time.Now().Sub(end).Seconds(), total.Seconds())

		if total < time.Duration(duration)*time.Second {
			time.Sleep(time.Duration(duration)*time.Second - total)
		}
	}
}
