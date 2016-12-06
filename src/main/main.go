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
	// Try to unmarshal a latlong
	// fmt.Println([]byte(s))
	c1 := new(latlong.Coordinate)
	if e := json.Unmarshal([]byte(s), c1); e == nil {
		l = c1
		err = nil
		return
	}

	// Try to unmarshal an nvector
	c2 := new(nvector.Coordinate)
	if e := json.Unmarshal([]byte(s), c2); e == nil {
		l = c2
		err = nil
		return
	}

	// Try to unmarshal a utm
	c3 := new(utm.Coordinate)
	if e := json.Unmarshal([]byte(s), c3); e == nil {
		l = c3
		err = nil
		return
	}

	// Unmarshaling unsuccesful
	l = nil
	msg := "Cannot unmarshal coordinate: " + s
	err = errors.New(msg)
	return
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
	// Open the file
	file, err := os.Open(fname)
	if err != nil {
		// Error opening the file, presumably does not exist
		fmt.Println("open %s: no such file or directory", fname)
		os.Exit(1)
	}
	defer file.Close()

	// Setting up some variables for loading trips from the file
	// currentID holds the ID of the trip we are collecting coordinates for
	// tmpID holds the ID found in the file
	// tmpJSON holds the raw coordinate found in the file
	// tmpCoords holds the unmarshaled coordinates to be sent thru trips
	currentID := 0
	var tmpID int
	var tmpJSON string
	var myCoord latlong.LatLonger
	tmpCoords := make([]latlong.LatLonger, 1)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		_, err := fmt.Sscanf(scanner.Text(), "%d\t%s", &tmpID, &tmpJSON)
		if err == nil {
			if tmpID != currentID {
				// Done collecting coordinates for the current trip
				// Send what we have thru channel, and reset our variables
				trips <- trip{currentID, tmpCoords}
				currentID = tmpID
				tmpCoords = nil
			}
			myCoord, err = unmarshalLatLonger(tmpJSON)
			if err == nil {
				tmpCoords = append(tmpCoords, myCoord)
			} else {
						fmt.Println(err)
						os.Exit(1)
			}
		} else {
			// Sscanf failed. Shouldn't happen, hopefully
			log.Fatal("fmt.Sscanf() is broken. Pls fix.")
		}
	}
	// One last trip sent thru channel before closing it
	trips <- trip{currentID, tmpCoords}
	close(trips)
	return
}

// computeDistances continually receives trips over a channel and
// computes the total travel distance for each trip, sending the
// totalled results over a channel.
//
// After the distance of the last trip has been calculated and sent
// over the output channel (totals), computeDistances closes the
// channel to indicate that there will be no more results.
func computeDistances(trips chan trip, totals chan total) {
	var currentDist float64 = 0
	var pPrev, pNext latlong.LatLonger
	for trip := range trips {
		for _, point := range trip.trajectory {
			pNext = point
			if pPrev == nil {
				// Can't find the distance with just one coordinate!
				pPrev = point
				continue
			}
			currentDist = currentDist + latlong.Distance(pPrev, pNext)
			pPrev = pNext
		}
		totals <- total{trip.id, currentDist}
		pPrev = nil
		currentDist = 0
	}
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
	if fname != "" {
		go loadTrips(fname, trips)
		go computeDistances(trips, totals)
		for tot := range totals {
			fmt.Println(tot)
		}
	} else {
		log.Println("Need a file to process!")
	}

	return
}
