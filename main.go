package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"

	"github.com/ylamgarchal/ncdiff/ncd"
)

func setupFlagCli() {

	var CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flag.Usage = func() {
		fmt.Fprintf(CommandLine.Output(), "Usage: %s file-x file-y\n", os.Args[0])
	}

	flag.Parse()

	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(1)
	}
}

// ComputeNcd compute the NCD given two slices of bytes
func ComputeNcd(b1 []byte, b2 []byte) float64 {

	fileXCompressedLength := ncd.GetOneCompressedSliceLength(b1)
	fileYCompressedLength := ncd.GetOneCompressedSliceLength(b2)
	fileXYCompressedLength := float64(ncd.GetTwoCompressedSlicesLength(b1, b2))

	minXY := math.Min(float64(fileXCompressedLength), float64(fileYCompressedLength))
	maxXY := math.Max(float64(fileXCompressedLength), float64(fileYCompressedLength))

	return ncd.Ncd(fileXYCompressedLength, minXY, maxXY)
}

func main() {

	setupFlagCli()

	fileXContent, err := ioutil.ReadFile(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}

	fileYContent, err := ioutil.ReadFile(flag.Args()[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%.2f\n", ComputeNcd(fileXContent, fileYContent)*100)
}
