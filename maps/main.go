package main

import (
	"context"
	"log"
	"strconv"

	"github.com/kelvins/geocoder"
	"github.com/kr/pretty"

	"googlemaps.github.io/maps"
)

func check(err error) {
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
}
func main() {
	address := geocoder.Address{
		// Street:  "Central Park West",
		// Number:  115,
		City: "Тербуны",
		// State:   "New York",
		Country: "Russia",
	}
	location, err := geocoder.Geocoding(address)
	newLat := strconv.FormatFloat(location.Latitude, 'f', 6, 64)
	newLng := strconv.FormatFloat(location.Longitude, 'f', 6, 64)
	newLocation := newLat + "," + newLng

	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyBL_7YvhKHVVdThonbVenBEJ92emxL4TUo"))
	check(err)

	l, err := maps.ParseLatLng(newLocation)
	check(err)
	r := &maps.NearbySearchRequest{
		Location: &l,
		// Name:     "some field",
		Radius: 300000,
		Type:   "storage",
	}

	route, err := c.NearbySearch(context.Background(), r)
	check(err)
	result := route.Results

	for i := 0; i < len(result); i++ {

		if result[i].Scope == "APP" {
			pretty.Println(result[i].Name, "\t\t", result[i].Scope, "\t\t", result[i].PlaceID)
		}
	}
}
