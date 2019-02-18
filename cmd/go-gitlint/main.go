package main

import "github.com/llorllale/go-gitlint/internal/commits"

// @todo #4 The path should be passed in as a command line
//  flag instead of being hard coded. All other configuration
//  options should be passed in through CLI as well.
func main() {
	commits.Printed(
		commits.In("."),
	)()
}
