package arn

import "math"

// EarthRadius is the radius of the earth in kilometers.
const EarthRadius = 6371

// Location ...
type Location struct {
	CountryName string  `json:"countryName"`
	CountryCode string  `json:"countryCode"`
	Latitude    float64 `json:"latitude" editable:"true"`
	Longitude   float64 `json:"longitude" editable:"true"`
	CityName    string  `json:"cityName"`
	RegionName  string  `json:"regionName"`
	TimeZone    string  `json:"timeZone"`
	ZipCode     string  `json:"zipCode"`
}

// IPInfoDBLocation ...
type IPInfoDBLocation struct {
	CountryName string `json:"countryName"`
	CountryCode string `json:"countryCode"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	CityName    string `json:"cityName"`
	RegionName  string `json:"regionName"`
	TimeZone    string `json:"timeZone"`
	ZipCode     string `json:"zipCode"`
}

// IsValid returns true if latitude and longitude are available.
func (p *Location) IsValid() bool {
	return p.Latitude != 0 && p.Longitude != 0
}

// Distance calculates the distance in kilometers to the second location.
// Original implementation: https://www.movable-type.co.uk/scripts/latlong.html
func (p *Location) Distance(p2 *Location) float64 {
	dLat := (p2.Latitude - p.Latitude) * (math.Pi / 180.0)
	dLon := (p2.Longitude - p.Longitude) * (math.Pi / 180.0)

	lat1 := p.Latitude * (math.Pi / 180.0)
	lat2 := p2.Latitude * (math.Pi / 180.0)

	a1 := math.Sin(dLat/2) * math.Sin(dLat/2)
	a2 := math.Sin(dLon/2) * math.Sin(dLon/2) * math.Cos(lat1) * math.Cos(lat2)

	a := a1 + a2
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return EarthRadius * c
}
