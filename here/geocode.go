package here

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const gen = "9"

// These are the test environment URLs
const baseGeocodeURL = "https://geocoder.cit.api.here.com/6.2/geocode.json?"
const baseReverseGeocodeURL = "https://reverse.geocoder.cit.api.here.com/6.2/reversegeocode.json"

// These are the production URLs
// const baseGeocodeURL = "https://geocoder.api.here.com"
// const baseReverseGeocodeURL = "https://reverse.geocoder.api.here.com"

//GeocodeFreeform takes an address as a string, geocodes it using the Here geocoder API then returns the result as a string.
func GeocodeFreeform(id string, code string, searchtext string) (*Base, error) {

	u, err := url.Parse(baseGeocodeURL)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("app_code", code)
	q.Set("app_id", id)
	q.Set("gen", gen)
	q.Set("searchtext", searchtext)
	u.RawQuery = q.Encode()

	res, err := http.Get(u.String())
	if err != nil {
		fmt.Println("Hi I'm Get error", err)
		return nil, err
	}

	var result Base
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		res.Body.Close()
		fmt.Println("Error: ", err)
		return nil, err
	}

	res.Body.Close()
	// TODO: Handle parsing of JSON and only return the coordinates.
	return &result, err
}

//ReverseGeocode takes a latitude and a longitude as float64s and a radius as an int and returns an address as a string.
func ReverseGeocode(id string, code string, lat string, lon string, radius int) (*Base, error) {
	prox := []string{lat, lon, strconv.Itoa(radius)}
	u, err := url.Parse(baseReverseGeocodeURL)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("prox", strings.Join(prox, ","))
	q.Set("mode", "retrieveAddresses")
	u.RawQuery = q.Encode()

	res, err := http.Get(u.String())
	if err != nil {
		fmt.Println("Hi I'm Get error", err)
		return nil, err
	}

	var result Base
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		res.Body.Close()
		fmt.Println("Error: ", err)
		return nil, err
	}

	res.Body.Close()
	// TODO: Handle parsing of JSON and only return the address.
	return &result, err
}
