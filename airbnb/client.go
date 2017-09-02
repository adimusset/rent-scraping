package airbnb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const u = "https://www.airbnb.fr/api/v2/explore_tabs?"

// const airbnbUrl = "https://www.airbnb.fr/api/v2/explore_tabs?" +
// 	"version=1.2.7&_format=for_explore_search_web&items_per_grid=18&experiences_per_grid=20" +
// 	"&guidebooks_per_grid=20&fetch_filters=true&is_guided_search=false&supports_for_you_v3=true" +
// 	"&screen_size=medium&timezone_offset=120&auto_ib=false&guest_from_sem_traffic=false" +
// 	"&is_luxury_request=false&metadata_only=false&is_standard_search=true&tab_id=home_tab" +
// 	"&location=Paris&allow_override%5B%5D=&ne_lat=48.9052248670166&ne_lng=2.477368999566579" +
// 	"&sw_lat=48.79506752946523&sw_lng=2.2006508599181416&zoom=12&search_by_map=true" +
// 	"&federated_search_session_id=c71c28e3-b0f6-4979-9a26-f8e950f7f564&_intents=p1" +
// 	"&key=d306zoyjsyarp7ifhu67rjxn52tv0t20&currency=EUR&locale=fr"

func GetAirbnbResponse(numberOfGuests int) (*AirbnbResponse, error) {
	v := url.Values{}
	v.Set("version", "1.2.7") // api version ?
	v.Set("_format", "for_explore_search_web")
	v.Set("items_per_grid", "10000")    // max returned items
	v.Set("experiences_per_grid", "20") // ignored
	v.Set("guidebooks_per_grid", "20")  // ignored
	v.Set("fetch_filters", "true")
	v.Set("is_guided", "false")
	v.Set("supports_for_you_v3", "true")
	v.Set("screen_size", "medium") // matters ?
	v.Set("timezone_offset", "120")
	v.Set("auto_ib", "false")
	v.Set("guest_from_sem_traffic", "false")
	v.Set("is_luxury_request", "false")
	v.Set("metadata_only", "false")
	v.Set("is_standard_search", "true")
	v.Set("tab_id", "home_tab")
	v.Set("location", "Paris")
	v.Set("allow_override[]", "")
	// define the part of the map we are looking into
	v.Set("ne_lat", "48.9052248670166")
	v.Set("ne_lng", "2.477368999566579")
	v.Set("sw_lat", "48.79506752946523")
	v.Set("sw_lng", "2.2006508599181416")
	v.Set("search_by_map", "true")

	v.Set("federated_search_session_id", "c71c28e3-b0f6-4979-9a26-f8e950f7f564")
	v.Set("_intents", "p1")
	v.Set("key", "d306zoyjsyarp7ifhu67rjxn52tv0t20	") // api key
	v.Set("currency", "EUR")
	v.Set("locale", "fr")

	v.Set("adults", fmt.Sprintf("%d", numberOfGuests))
	v.Set("guests", fmt.Sprintf("%d", numberOfGuests))

	airbnbUrl := u + v.Encode()
	resp, err := http.Get(airbnbUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response AirbnbResponse
	err = json.Unmarshal(body, &response)
	return &response, err
}
