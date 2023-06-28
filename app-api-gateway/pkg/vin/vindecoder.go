package vin

import (
	"errors"
	"regexp"
	"strings"
)

var vinRegex, _ = regexp.Compile(`(?i)([A-HJ-NPR-Z0-9]{8})([0-9X])([A-HJ-NPR-Z0-9]{8})`)

var (
	InvalidVinErr    = errors.New("vin provided does not conform to ISO 3779")
	InvalidRegionErr = errors.New("invalid region")
)

type Vin struct {
	Manufacturer string
	Year         int
	Country      string
	Region       string
	Serial       string
	WMI          string
	VDS          string
	VIS          string
}

func getRegion(letter byte) (string, bool) {
	switch {
	case letter >= 'A' && letter <= 'H':
		return "Africa", true
	case letter >= 'J' && letter <= 'R':
		return "Asia", true
	case letter >= 'S' && letter <= 'Z':
		return "Europe", true
	case letter >= '1' && letter <= '5':
		return "North America", true
	case letter == '6' || letter == '7':
		return "Oceania", true
	case letter == '8' || letter == '9':
		return "South America", true
	}

	return "", false
}

func Decode(vinNumber string) (*Vin, error) {
	var vin Vin

	// normalise the vin number
	vinNumber = strings.ToUpper(vinNumber)

	// check if it is valid VIN
	if !vinRegex.MatchString(vinNumber) {
		return nil, InvalidVinErr
	}

	// Validate checksum.

	// split the vin parts
	vin = Vin{
		WMI: vinNumber[:3],
		VDS: vinNumber[3:9],
		VIS: vinNumber[9:17],
	}

	// parse region
	region, ok := getRegion(vinNumber[0])
	if !ok {
		return nil, InvalidRegionErr
	}

	vin.Region = region

	// parse year
	year, ok := yearCodes[vinNumber[9]]
	if !ok {
		return nil, InvalidVinErr
	}

	vin.Year = year

	// parse country

}
