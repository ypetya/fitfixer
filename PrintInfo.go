package fitfixer

import (
	"fmt"

	"github.com/tormoder/fit"
)

// Implements IPrintInfo
// examine fit-file data by printing on the console
type PrintInfo struct {
}

func (PrintInfo) PrintInfo(fitFile string) {
	ff := readFit(fitFile)

	// Inspect the dynamic Product field in the FileId message
	fmt.Println(ff.FileId.GetProduct())

	// Inspect the TimeCreated field in the FileId message
	fmt.Println(ff.FileId.TimeCreated)

	// Inspect the FIT file type
	fmt.Println(ff.Type())

	// Get the actual activity
	activity, err := ff.Activity()
	perr(err)

	fmt.Println("Sessions:")
	for _, session := range activity.Sessions {
		fmt.Println(session.Sport)
		fmt.Println(session.AvgHeartRate)
	}
}

func printTs(activity *fit.ActivityFile) {
	for _, record := range activity.Records {
		fmt.Println(record.Timestamp)
		break
	}
}
