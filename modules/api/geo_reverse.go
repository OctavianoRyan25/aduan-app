package api

import (
	"fmt"
	"os"

	"github.com/kelvins/geocoder"
)

func ReverseGeocode(lat, long float64) (string, error) {
	geocoder.ApiKey = os.Getenv("GEOAPIKEY")

	location := geocoder.Location{
		Latitude:  lat,
		Longitude: long,
	}

	addresses, err := geocoder.GeocodingReverse(location)

	if err != nil {
		fmt.Println("Could not get the addresses: ", err)
		return err.Error(), err
	}

	// Usually, the first address returned from the API
	// is more detailed, so let's work with it
	address := addresses[0]

	// // Print the address formatted by the geocoder package
	// fmt.Println(address.FormatAddress())
	// // Print the formatted address from the API
	// fmt.Println(address.FormattedAddress)
	// // Print the type of the address
	// fmt.Println(address.Types)

	return address.FormattedAddress, nil
}
