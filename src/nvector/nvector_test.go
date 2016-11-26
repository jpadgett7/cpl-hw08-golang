package nvector

import (
	"latlong"
	"math"
	"math/rand"
	"testing"
)

const (
	closeEnough = 0.00000001 // Maximum difference between floating point values
)

// Generate 1,000,000 random lat/long coordinates, convert them to
// n-vector, convert them back, and assert that we got something close
// enough to the original.
func TestRandPoints(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		// Random, valid lat/long
		want := &latlong.Coordinate{
			Latitude:  -90 + rand.Float64()*180,
			Longitude: -180 + rand.Float64()*360,
		}

		// Convert it to an n-vector
		coord := ToCoordinate(want)

		// Convert it back
		got := coord.ToLatLong()

		// Make sure it matches the original latitude
		if d := math.Abs(want.Latitude - got.Latitude); d > closeEnough {
			t.Errorf("Difference in latitude (%f) outside of acceptable range (%f)", d, closeEnough)
			t.FailNow()
		}

		// Make sure it matches the original longitude
		if d := math.Abs(want.Longitude - got.Longitude); d > closeEnough {
			t.Errorf("Difference in longitude (%f) outside of acceptable range (%f)", d, closeEnough)
			t.FailNow()
		}

	}
}
