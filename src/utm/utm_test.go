package utm

import (
	"latlong"
	"math"
	"math/rand"
	"testing"
)

const (
	// Maximum difference between floating point values. UTM is
	// less precise than n-vector, so this value is more relaxed.
	closeEnough = 0.00001
)

// Generate 1,000,000 random lat/long coordinates, convert them to
// UTM, convert them back, and assert that we got something close
// enough to the original.
func TestRandPoints(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		// Generate a random, valid lat/long within the range that UTM can tolerate
		want := &latlong.Coordinate{
			Latitude:  -79 + rand.Float64()*162,
			Longitude: -180 + rand.Float64()*360,
		}

		// Convert to UTM
		coord, err := ToCoordinate(want)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		// Convert back to lat/long
		got, err := coord.ToLatLong()
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		// Make sure that the latitude is close enough to the original
		if d := math.Abs(want.Latitude - got.Latitude); d > closeEnough {
			t.Errorf("Difference in latitude (%f) outside of acceptable range (%f)", d, closeEnough)
			t.FailNow()
		}

		// Make sure that the longitude is close enough to the original
		if d := math.Abs(want.Longitude - got.Longitude); d > closeEnough {
			t.Errorf("Difference in longitude (%f) outside of acceptable range (%f)", d, closeEnough)
			t.FailNow()
		}

	}
}
