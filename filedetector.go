package trojansourcedetector

import (
	"io"
)

// SingleFileDetector is a detector for malicious unicode code points in a single file.
type SingleFileDetector interface {
	// Detect reads all inputfrom a buffered reader and returns a list of errors found in that stream.
	Detect(filename string, input io.Reader) Errors
}
