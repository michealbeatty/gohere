package here

// Base is the base for the JSON response from the Here API
type Base struct {
	Response *Response
}

//Response contains the meat of the JSON response
type Response struct {
	View []struct {
		Result []struct {
			Location struct {
				NavigationPosition []struct {
					Latitude  float64
					Longitude float64
				}
				Address struct {
					Label string
				}
			}
		}
	}
}

//type View struct {
// 	Result *Result
//}

// type Result struct {
// 	Location *Location
// }

// type Location struct {
// 	LocationType       string
// 	NavigationPosition []struct {
// 		Latitude  float64
// 		Longitude float64
// 	}
// }

// type Address struct {
// 	Country     string
// 	State       string
// 	County      string
// 	City        string
// 	Street      string
// 	HouseNumber string
// 	PostalCode  string
// }

// type NavigationPosition struct {
// 	Latitude  float64
// 	Longitude float64
// }

//Geolocation is a struct containing the parts of the JSON response
// that we care about most.
type Geolocation struct {
	AddressLabel string
	Latitude     float64
	Longitude    float64
}
