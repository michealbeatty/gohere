package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/michealbeatty/gohere/here"
)

// APP_ID and APP_CODE are the Here credentials
const appID = ""
const appCode = ""

func main() {
	var lat float64
	var lon float64

	var searchtext = flag.String("geo", "", "An address or landmark to geocode")
	var rev = flag.String("rev", "", "A coordinate to reverse geocode")
	flag.Parse()

	// TODO: Change here.go to only return the desired information. This parsing of the result should
	//be moved into the proper function.
	if *searchtext != "" {
		searchTerm := *searchtext
		location, err := here.GeocodeFreeform(appID, appCode, searchTerm)
		if err != nil {
			fmt.Println(err)
		}
		if len(location.Response.View) == 0 {
			fmt.Printf("There was a problem with your search for %s\n", searchTerm)
			fmt.Println("Please try more specific terms")
			os.Exit(1)
		} else if len(location.Response.View[0].Result) >= 1 {
			for _, item := range location.Response.View[0].Result {
				lat = item.Location.NavigationPosition[0].Latitude
				lon = item.Location.NavigationPosition[0].Longitude
				loc := here.Geolocation{
					AddressLabel: item.Location.Address.Label,
					Latitude:     lat,
					Longitude:    lon,
				}

				fmt.Printf("%s\t%f, %f\n", loc.AddressLabel, loc.Latitude, loc.Longitude)
				// TODO: Use templates to better present the information
			}
		}
	}

	if *rev != "" {
		coordinates := strings.Split(*rev, ",")
		location, err := here.ReverseGeocode(appID, appCode, coordinates[0], coordinates[1], 250)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(location.Response.View[0].Result[0].Location.Address.Label)
	}

}
