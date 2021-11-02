package trojansourcedetector_test

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/haveyoudebuggedit/trojansourcedetector"
)

func TestNoContent(t *testing.T) {
	detector := trojansourcedetector.BIDIDetector()
	var data []byte

	errs := detector.Detect("test.txt", bufio.NewReader(bytes.NewReader(data)))
	if len(errs.Get()) != 0 {
		t.Fatalf("scanning an empty file did not return an empty error list (%d items)", len(errs.Get()))
	}
}

func TestNoBIDI(t *testing.T) {
	detector := trojansourcedetector.BIDIDetector()
	data := []byte("Hello world!")
	errs := detector.Detect("test.txt", bufio.NewReader(bytes.NewReader(data)))
	if len(errs.Get()) != 0 {
		t.Fatalf(
			"scanning a file without BIDI characters did not return an empty error list (%d items)",
			len(errs.Get()),
		)
	}
}

func TestBIDIInFirstLine(t *testing.T) {
	detector := trojansourcedetector.BIDIDetector()
	data := []byte("Hello world!\u202A")
	errs := detector.Detect("test.txt", bufio.NewReader(bytes.NewReader(data)))
	errList := errs.Get()
	if len(errList) < 1 {
		t.Fatalf(
			"No errors found while scanning content containing a BIDI character",
		)
	}
	if errList[0].File() != "test.txt" {
		t.Fatalf("Incorrect file ErrorCode reported (%s)", errList[0].File())
	}
	if errList[0].Line() != 1 {
		t.Fatalf("Incorrect line reported (%d)", errList[0].Line())
	}
	if errList[0].Column() != 13 {
		t.Fatalf("Incorrect column reported (%d)", errList[0].Column())
	}
}
func TestBIDIInSecondLine(t *testing.T) {
	detector := trojansourcedetector.BIDIDetector()
	data := []byte("Hello world!\n\u202A")
	errs := detector.Detect("test.txt", bufio.NewReader(bytes.NewReader(data)))
	errList := errs.Get()
	if len(errList) < 1 {
		t.Fatalf(
			"No errors found while scanning content containing a BIDI character",
		)
	}
	if errList[0].File() != "test.txt" {
		t.Fatalf("Incorrect file ErrorCode reported (%s)", errList[0].File())
	}
	if errList[0].Line() != 2 {
		t.Fatalf("Incorrect line reported (%d)", errList[0].Line())
	}
	if errList[0].Column() != 1 {
		t.Fatalf("Incorrect column reported (%d)", errList[0].Column())
	}
}
