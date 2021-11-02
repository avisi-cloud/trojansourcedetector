package trojansourcedetector_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/haveyoudebuggedit/trojansourcedetector"
)

//go:generate go run cmd/testdata/main.go

func TestE2E(t *testing.T) {
	assertFileExists(t, "testdata/bidi.txt")
	assertFileExists(t, "testdata/unicode.txt")

	detector := trojansourcedetector.New(&trojansourcedetector.Config{
		Directory:     "testdata",
		DetectUnicode: true,
		DetectBIDI:    true,
		Parallelism:   10,
	})

	errs := detector.Run()

	assertHasError(t, errs, trojansourcedetector.ErrBIDI, "bidi.txt", 1, 44)
	assertHasError(t, errs, trojansourcedetector.ErrUnicode, "unicode.txt", 1, 29)
	assertHasError(t, errs, trojansourcedetector.ErrUnicode, "unicode.txt", 1, 30)
}

func assertFileExists(t *testing.T, file string) {
	if _, err := os.Stat(file); err != nil {
		t.Fatalf("file does not exist: %s did you run go generate? (%v)", file, err)
	}
}

func assertHasError(
	t *testing.T,
	errs trojansourcedetector.Errors,
	code trojansourcedetector.ErrorCode,
	file string,
	line uint,
	column uint,
) {
	for _, err := range errs.Get() {
		if err.Code() == code && filepath.ToSlash(err.File()) == file && err.Line() == line && err.Column() == column {
			return
		}
	}
	t.Fatalf("Did not find expected '%s' error in %s line %d column %d", code, file, line, column)
}
