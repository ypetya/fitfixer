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
