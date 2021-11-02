package trojansourcedetector_test

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/haveyoudebuggedit/trojansourcedetector"
)

func TestUnicodeEmptyFile(t *testing.T) {
	detector := trojansourcedetector.UnicodeDetector()
	var data []byte

	errs := detector.Detect("test.txt", bufio.NewReader(bytes.NewReader(data)))
	if len(errs.Get()) != 0 {
		t.Fatalf(
			"scanning an empty file did not return an empty error list (%d items)",
			len(errs.Get()),
		)
	}
}

func TestNoUnicode(t *testing.T) {
	detector := trojansourcedetector.UnicodeDetector()
	data := []byte("Hello world!")
	errs := detector.Detect("test.txt", bufio.NewReader(bytes.NewReader(data)))
	if len(errs.Get()) != 0 {
		t.Fatalf(
			"scanning a file without unicode characters did not return an empty error list (%d items)",
			len(errs.Get()),
		)
	}
}

func TestUnicodeInFirstLine(t *testing.T) {
	detector := trojansourcedetector.UnicodeDetector()
	data := []byte("Hello world! á")
	errs := detector.Detect("test.txt", bufio.NewReader(bytes.NewReader(data)))
	if len(errs.Get()) < 1 {
		t.Fatalf(
			"scanning a file with a unicode character didn't return an error",
		)
	}
	if errs.Get()[0].Code() != trojansourcedetector.ErrUnicode {
		t.Fatalf("wrong error code returned from unicode error (%s)", errs.Get()[0].Code())
	}
	if errs.Get()[0].Line() != 1 {
		t.Fatalf("wrong line number returned from unicode error (%d)", errs.Get()[0].Line())
	}
	if errs.Get()[0].Column() != 14 {
		t.Fatalf("wrong column number returned from unicode error (%d)", errs.Get()[0].Column())
	}
}

func TestUnicodeInSecondLine(t *testing.T) {
	detector := trojansourcedetector.UnicodeDetector()
	data := []byte("Hello world!\ná")
	errs := detector.Detect("test.txt", bufio.NewReader(bytes.NewReader(data)))
	if len(errs.Get()) < 1 {
		t.Fatalf(
			"scanning a file with a unicode character didn't return an error",
		)
	}
	if errs.Get()[0].Code() != trojansourcedetector.ErrUnicode {
		t.Fatalf("wrong error code returned from unicode error (%s)", errs.Get()[0].Code())
	}
	if errs.Get()[0].Line() != 2 {
		t.Fatalf("wrong line number returned from unicode error (%d)", errs.Get()[0].Line())
	}
	if errs.Get()[0].Column() != 1 {
		t.Fatalf("wrong column number returned from unicode error (%d)", errs.Get()[0].Column())
	}
}
