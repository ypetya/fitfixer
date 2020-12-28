package fitfixer

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tormoder/fit"
)

func perr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
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

func readFit(target string) *fit.File {
	td := readFile(target)

	tfit, err := fit.Decode(bytes.NewReader(td))
	perr(err)
	return tfit
}
