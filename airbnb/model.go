package airbnb

type AirbnbResponse struct {
	ExploreTabs []Tab `json:"explore_tabs"`
}

type Tab struct {
	Sections []Section `json:"sections"`
}

type Section struct {
	Listings []Listing `json:"listings"`
}

type Listing struct {
	Info  Info  `json:"listing"`
	Price Price `json:"pricing_quote"`
}

type Info struct {
	ID        int     `json:"id"`
	Beds      int     `json:"beds"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
	Capacity  int     `json:"person_capacity"`
	RoomType  string  `json:"room_type"`
}

type Price struct {
	Rate Rate `json:"rate"`
}

type Rate struct {
	Amount int `json:"amount"`
}
