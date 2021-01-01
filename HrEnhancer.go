package fitfixer

import (
	"bytes"
	"encoding/binary"

	"github.com/tormoder/fit"
)

// Implements IEnhancer
// Behavior is to produce a new fit file from 2 existing fit files
// enhancing the first file by taking the HR from the second file
type HrEnhancer struct {
}

func (h HrEnhancer) Enhance(target string, toEnhance string, with string) {
	tfit, sfit := readFit(toEnhance), readFit(with)

	h.enhance(tfit, sfit)

	outBuf := &bytes.Buffer{}
	err := fit.Encode(outBuf, tfit, binary.BigEndian)
	perr(err)

	writeFile(target, outBuf)
}

func (HrEnhancer) enhance(target *fit.File, source *fit.File) {

	// For AvgHeartRate calculation
	var sumHR int
	var samplesCount int

	// seq 1.
	a1, err := target.Activity()
	perr(err)
	if len(a1.Sessions) != 1 {
		panic("Target activity has multiple sessions!")
	}

	// seq 2.
	a2, err := source.Activity()
	perr(err)

	// limits of seq 2.
	i2, maxI2 := 0, len(a2.Records)-1

	// keep last known value
	var hr uint8

seq1:
	for i1, r1 := range a1.Records {
		t1 := r1.Timestamp.Unix()
		r2 := a2.Records[i2]
		for i2 < maxI2 {
			t2 := r2.Timestamp.Unix()
			switch {
			case t2 > t1 || t2 == t1:
				{
					// seq 2. is late => take the first known value
					hr = a2.Records[i2].HeartRate
					a1.Records[i1].HeartRate = hr
					sumHR += int(hr)
					samplesCount += 1
					continue seq1
				}
			case t2 < t1:
				{
					// seq 2. ran out, has newever => increment
					i2 += 1
					r2 = a2.Records[i2]
				}
			}
		}

		if i2 >= maxI2 {
			// seq 2. finished early => take the last known value
			a1.Records[i1].HeartRate = hr
			sumHR += int(hr)
			samplesCount += 1
		}
	}

	//
	var avgHr uint8
	avgHr = uint8(sumHR / samplesCount)
	a1.Sessions[0].AvgHeartRate = avgHr
}
