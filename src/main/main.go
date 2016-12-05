package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"latlong"
	"log"
	"nvector"
	"os"
	"utm"
)

var (
	// True if we want to see debug output, otherwise false.
	// Set by the user with the -debug flag
	debug bool
)

// parseCLIArgs parses options from the command line.
//
// Returns the name of the user-provided data file
func parseCLIArgs() string {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:  %s <filename>\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.BoolVar(&debug, "debug", false, "enable debug output")

	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "Need a file to process!\n\n")
		flag.Usage()
		os.Exit(1)
	}

	return flag.Arg(0)
}

// unmarshalLatLonger attempts to unmarshal a JSON encoded
// latlong.LatLonger coordinate.
//
// The coordinate may be a JSON encoded latlong.Coordinate,
// nvector.Coordinate, or utm.Coordinate.
//
// For each of the above coordinate types, unmarshalLatLonger attempts
// to unmarshal the string. It starts with latlong.Coordinate. If it
// successfully unmarshals the string as a latlong.Coordinate, it
// returns it along with a nil error. If it fails, it tries to
// unmarshal it as a nvector.Coordinate. unmarshalLatLonger tries each
// type until one succeeds. If it fails to unmarshal the string to
// **any** of the above coordinate types, it returns a non-nil error.
//
// If unmarshaling is successful, the coordinate is returned as a latlong.LatLonger.
func unmarshalLatLonger(s string) (l latlong.LatLonger, err error) {
	if e := l.UnmarshalJSON([]byte(s)); e == nil {
		err = nil
		return
	} else {
		l, err = nvector.ToCoordinate(l)
		if err == nil {
			if e := l.UnmarshalJSON([]byte(s)); e == nil {
				err = nil
				return
			} else {
				l, err = utm.ToCoordinate(l)
				
			}
		} else {
			return
		}
	}

	return nil, nil
}

// loadTrips loads trip information line-by-line from a file and sends
// results over a channel.
//
// Refer to online documentatin for format expectations
//
// Attempts to open a file and read its contents line-by-line. As it
// reads through the file, loadTrips groups coordinates by traveler
// ID. Records are aggregated for each traveler ID. Once all
// coordinates for a traveler have been seen, the trip information is
// sent over the trips channel.
//
// When loadTrips finishes processing all of the lines in the file and
// sends the final trip over the output channel, it closes the output
// channel to signal that nothing is left.
func loadTrips(fname string, trips chan trip) {
	close(trips)
}

// computeDistances continually receives trips over a channel and
// computes the total travel distance for each trip, sending the
// totalled results over a channel.
//
// After the distance of the last trip has been calculated and sent
// over the output channel (totals), computeDistances closes the
// channel to indicate that there will be no more results.
func computeDistances(trips chan trip, totals chan total) {
	close(totals)
}

func main() {
	fname := parseCLIArgs()
	trips := make(chan trip)
	totals := make(chan total)

	log.SetFlags(0) // Dial back the log output
	if debug {
		log.Printf("Starting program %s", os.Args[0])
	}
	if fname != nil {
		go loadTrips(fname, trips)
		go computeDistances(trips, totals)
	} else {
		log.Println("Need a file to process!")
	}

	return
}
