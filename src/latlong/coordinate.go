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
