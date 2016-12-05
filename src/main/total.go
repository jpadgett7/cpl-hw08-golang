package main

import (
	"fmt"
)

type total struct {
	id       int
	distance float64
}

func (t total) String() string {
	return fmt.Sprintf("Traveler %d traveled %f miles", t.id, t.distance)
}
