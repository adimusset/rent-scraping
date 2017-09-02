package main

import (
	"rent/airbnb"
)

type Crawler interface {
	Crawl() ([]Object, error)
}

type AirbnbCrawler struct {
}

func NewAirbnbCrawler() *AirbnbCrawler {
	return &AirbnbCrawler{}
}

func (c *AirbnbCrawler) Crawl(numberOfGuests int) ([]Object, error) {
	response, err := airbnb.GetAirbnbResponse(numberOfGuests)
	return responseToObjects(response), err
}

func responseToObjects(r *airbnb.AirbnbResponse) []Object {
	out := []Object{}
	for _, t := range r.ExploreTabs {
		for _, s := range t.Sections {
			for _, l := range s.Listings {
				o := make(map[string]interface{})
				o["price"] = l.Price.Rate.Amount
				o["id"] = l.Info.ID
				o["capacity"] = l.Info.Capacity
				o["room_type"] = l.Info.RoomType
				o["beds"] = l.Info.Beds
				o["latitude"] = l.Info.Latitude
				o["longitude"] = l.Info.Longitude
				obj := Object(o)
				out = append(out, obj)
			}
		}
	}
	return out
}
