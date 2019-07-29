package ncd

import (
	"bytes"
	"compress/zlib"
	"log"
)

// GetOneCompressedSliceLength compress the given slice and
// returns the length of the result
func GetOneCompressedSliceLength(b []byte) int {
	var buf bytes.Buffer
	w, _ := zlib.NewWriterLevel(&buf, 9)
	defer w.Close()

	_, err := w.Write(b)
	if err != nil {
		log.Fatal(err)
	}

	w.Flush()
	return buf.Len()
}

// GetTwoCompressedSlicesLength compress the two given slices
// and returns the length of the result
func GetTwoCompressedSlicesLength(b1 []byte, b2 []byte) int {

	var buf bytes.Buffer
	w := zlib.NewWriter(&buf)
	defer w.Close()

	_, err := w.Write(b1)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(b2)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	w.Flush()
	return buf.Len()
}

// Ncd compute the normalized compression distance
func Ncd(zXY, minXY, maxXY float64) float64 {
	return (zXY - minXY) / maxXY
}
