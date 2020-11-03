package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/tormoder/fit"
	"io/ioutil"
	"os"
)

func main() {
	tf, sf := os.Args[1], os.Args[2]
	tfit, sfit := readFit(tf), readFit(sf)

	enhance(tfit, sfit)

	outBuf := &bytes.Buffer{}
	err := fit.Encode(outBuf, tfit, binary.BigEndian)
	perr(err)

	writeFile(tf+".withHR.fit", outBuf)

}

func printInfo(tfit *fit.File, sfit *fit.File) {
	fmt.Println("fitfixer v0.0.1")
	fmt.Println("target file:")
	info(tfit)
	fmt.Println("source file:")
	info(sfit)
	fmt.Println("-------")
}

func info(ff *fit.File) {

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
	}
}

func printTs(activity *fit.ActivityFile) {
	for _, record := range activity.Records {
		fmt.Println(record.Timestamp)
		break
	}
}

func enhance(target *fit.File, source *fit.File) {

	a1, err := target.Activity()
	perr(err)

	a2, err := source.Activity()
	perr(err)

	i2, maxI2 := 0, len(a2.Records)-1

	for i1, r1 := range a1.Records {
	inner:
		for i2 < maxI2 {
			r2 := a2.Records[i2]
			t1, t2 := r1.Timestamp.Unix(), r2.Timestamp.Unix()
			switch {
			case t2 > t1 || t2 == t1:
				{
					hr := a2.Records[i2].HeartRate
					a1.Records[i1].HeartRate = hr
					break inner
				}
			case t2 < t1:
				{
					i2 += 1
				}
			}
		}
	}
}

func readFit(target string) *fit.File {
	td := readFile(target)

	tfit, err := fit.Decode(bytes.NewReader(td))
	perr(err)
	return tfit
}

func writeFile(target string, data *bytes.Buffer) {

	f, err := os.Create(target)
	perr(err)
	defer f.Close()
	f.Sync()
	wr := bufio.NewWriter(f)

	data.WriteTo(wr)
	perr(err)

	wr.Flush()
}

func readFile(target string) []byte {
	tData, err := ioutil.ReadFile(target)
	perr(err)
	return tData
}

func perr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
