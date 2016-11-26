# Go Travel : A tool for computing travel distances for eventual travel reimbursement

This tool is designed to aid in the calculation of travel reimbursement for Team Rocket's fleet of flying Pokemon.

Documentation for its behavior can be found publicly here: http://cpl.mwisely.xyz
To verify that the code works properly, compare its output against provided samples.

**Note: Follow the design specifications EXACTLY.**
Not doing so will hurt your grade.

**Note: the following commands will work on campus machines.**
**If you use your own machine or editors, you are on your own.**

## Setup Go 1.7.3

~~~shell
$ bash setup.sh
$ GOROOT="$(pwd)/go" ./go/bin/go version
go version go1.7.3 linux/amd64
~~~

## Check and Correct Style

~~~shell
# List the packages you wish to fix
$ GOROOT="$(pwd)/go" GOPATH="$(pwd)" ./go/bin/go fmt main latlong nvector utm
~~~~

This will run the `go fmt` tool to properly format your Go code.

**Warning** `go fmt` will modify your `.go` file(s).
You should close your text editors before you run it.
Not all editors are smart enough to handle the "file changed beneath its feet" situation.

## Building and Running the Program

### With `go install` (the preferred method)

~~~shell
# Build the program and drop the executable in bin/
$ GOROOT="$(pwd)/go" GOPATH="$(pwd)" ./go/bin/go install main

# Run the program and check its usage
$ ./bin/main -help
Usage:  ./bin/main <filename>
  -debug
        enable debug output

# Run the program and give it a file to process
$ ./bin/main test.dat
Traveler 0 traveled 190.12 miles
Traveler 1 traveled 191.26 miles
Traveler 2 traveled 184.31 miles
Traveler 3 traveled 163.49 miles
...

~~~


### With `go build`

~~~shell
# Build the program and drop the executable in the current directory
$ GOROOT="$(pwd)/go" GOPATH="$(pwd)" ./go/bin/go build main

# Run the program and check its usage
$ ./main -help
Usage:  ./bin/main <filename>
  -debug
        enable debug output

# Run the program and give it a file to process
$ ./main test.dat
Traveler 0 traveled 190.12 miles
Traveler 1 traveled 191.26 miles
Traveler 2 traveled 184.31 miles
Traveler 3 traveled 163.49 miles
...

~~~

**DO NOT** commit compiled files to your git repository.

**DO NOT** add the go compiler (or its `.tar.gz`) to your git repository.
