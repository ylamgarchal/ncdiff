package ncd_test

import (
	"testing"

	"github.com/ylamgarchal/ncdiff/ncd"
)

func TestGetOneCompressedSliceLength(t *testing.T) {
	var testData = []byte("aaa: 123\nbbb: 456\nccc: 789\nddd: 123\neee: 456\nfff: 789\nggg: 123\nhhh: 456")
	compressedLength := ncd.GetOneCompressedSliceLength(testData)
	if compressedLength >= len(testData) {
		t.Errorf("raw data length: %d, compressed: %d", len(testData), compressedLength)
	}
}

func TestGetTwoCompressedSlicesLength(t *testing.T) {
	var testData1 = []byte("aaa: 123\nbbb: 456\nccc: 789\nddd: 123\neee: 456\nfff: 789\nggg: 123\nhhh: 456")
	var testData2 = []byte("aaa: iii\nbbb: jjj\nccc: kkk\nddd: lll\neee: mmm\nfff: nnn\nggg: 123\nhhh: 456")
	compressedLength := ncd.GetTwoCompressedSlicesLength(testData1, testData2)

	if compressedLength >= (len(testData1) + len(testData2)) {
		t.Errorf("raw data length: %d, compressed: %d", len(testData1)+len(testData2), compressedLength)
	}
}
