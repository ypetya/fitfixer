package main

import (
	"os"

	f "github.com/ypetya/fitfixer"
)

func main() {
	tf, sf := os.Args[1], os.Args[2]

	hre := f.HrEnhancer{}

	hre.Enhance(tf+".withHR.fit", tf, sf)
}
