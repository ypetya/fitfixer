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
	fmt.Println("Product:", ff.FileId.GetProduct())

	// Inspect the TimeCreated field in the FileId message
	fmt.Println("Time created:", ff.FileId.TimeCreated)

	// Inspect the FIT file type
	fmt.Println("File type:", ff.Type())

	// Get the actual activity
	activity, err := ff.Activity()
	perr(err)

	fmt.Println("Sessions:")
	for _, session := range activity.Sessions {
		fmt.Println(" Sport:", session.Sport)
		fmt.Println(" AvgHeartRate:", session.AvgHeartRate)
	}
}

func printTs(activity *fit.ActivityFile) {
	for _, record := range activity.Records {
		fmt.Println(record.Timestamp)
		break
	}
}
