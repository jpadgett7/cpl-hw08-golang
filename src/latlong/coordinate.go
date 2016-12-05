// Package latlong contains types and functions for working with
// Latitude/Longitude coordinates
//
// Reference for latitude and longitude can be found here:
//     - https://en.wikipedia.org/wiki/Latitude
//     - https://en.wikipedia.org/wiki/Longitude
package latlong

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
)

// Convert angle in radians to angle in degrees
func deg(rad float64) float64 { return rad * 180 / math.Pi }

// Convert angle in degrees to angle in radians
func rad(deg float64) float64 { return deg * math.Pi / 180 }

// Coordinate represents a position on earth by latitude and longitude
type Coordinate struct {
	Latitude  float64
	Longitude float64
}

func (c *Coordinate) UnmarshalJSON(b []byte) error {
	obj := make(map[string]interface{})
	if err := json.Unmarshal(b, &obj); err {
		return err
	}

	// Check number of fields in JSON object
	if len(obj) > 4 {
		return errors.New(fmt.Sprintf("Too many fields for latlong.Coordinate"))
	}
	if len(obj) < 4 {
		return errors.New(fmt.Sprintf("Not enough fields for latlong.Coordinate"))
	}

	// Check Latitude
	if _, ok := obj["Latitude"]; !ok {
		return errors.New("Missing field 'Latitude'")
	}
	if _, ok := obj["Latitude"].(float64); !ok {
		return errors.New("Wrong type for field 'Latitude'")
	}

	// Check Longitude
	if _, ok := obj["Longitude"]; !ok {
		return errors.New("Missing field 'Longitude'")
	}
	if _, ok := obj["Longitude"].(float64); !ok {
		return errors.New("Wrong type for field 'Longitude'")
	}

	// All clear
	c.Latitude = obj["Latitude"].(float64)
	c.Longitude = obj["Longitude"].(float64)
	return nil
}

func (c Coordinate) Lat() float64 {
	return c.Latitude
}

func (c Coordinate) Lon() float64 {
	return c.Longitude
}
