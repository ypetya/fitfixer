package main

import (
	"fmt"
	"os"

	f "github.com/ypetya/fitfixer"
)

func main() {
	switch {
	case len(os.Args) == 3 && os.Args[1] == "print":
		{
			p := f.PrintInfo{}
			p.PrintInfo(os.Args[2])
		}
	case len(os.Args) == 3 && os.Args[1] != "print":
		{
			tf, sf := os.Args[1], os.Args[2]
			hre := f.HrEnhancer{}
			hre.Enhance(tf+".withHR.fit", tf, sf)
		}
	default:
		fmt.Println("Usage:\nfitfixer <fit file to enhance> <fit file with Hr>\nfitfixer print <fit-file>")
	}
}
