package dms

import (
	"fmt"
)

// LatLonError is used for errors with lat/lon values
type LatLonError struct {
	err string
}

func (e *LatLonError) Error() string {
	return e.err
}

// DMS coordinates
type DMS struct {
	Degrees   uint8
	Minutes   uint8
	Seconds   float64
	Direction string
}

func (d *DMS) String() string {
	return fmt.Sprintf(`%d°%d'%f" %s`, d.Degrees, d.Minutes, d.Seconds, d.Direction)
}

// NewDMS converts Decimal Degreees to Degree, Minute, Seconds coordinates
func NewDMS(lat, lon float64) (*DMS, *DMS, error) {
	if lat < 0 || lon < 0 {
		return nil, nil, &LatLonError{"Latitude or longitude must be positive."}
	}
	if lat > 90 || lon > 180 {
		return nil, nil, &LatLonError{"Latitude must be less than 90 and longitude must be less than 180."}
	}

	var latDirection string
	var lonDirection string
	if lat > 0 {
		latDirection = "N"
	} else {
		latDirection = "S"
	}

	if lon > 0 {
		lonDirection = "E"
	} else {
		lonDirection = "W"
	}

	latitude := uint8(lat)
	latitudeMinutes := uint8((lat - float64(latitude)) * 60)
	latitudeSeconds := (lat - float64(latitude) - float64(latitudeMinutes)/60) * 3600

	longitude := uint8(lon)
	longitudeMinutes := uint8((lon - float64(longitude)) * 60)
	longitudeSeconds := (lon - float64(longitude) - float64(longitudeMinutes)/60) * 3600

	return &DMS{Degrees: latitude, Minutes: latitudeMinutes, Seconds: latitudeSeconds, Direction: latDirection},
		&DMS{Degrees: longitude, Minutes: longitudeMinutes, Seconds: longitudeSeconds, Direction: lonDirection},
		nil
}
