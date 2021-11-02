package trojansourcedetector

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strings"
)

// UnicodeDetector returns a detector that detects any non-ASCII Unicode character.
func UnicodeDetector() SingleFileDetector {
	return &unicodeDetector{}
}

type unicodeDetector struct {
}

func (u unicodeDetector) Detect(filename string, i io.Reader) Errors {
	input := bufio.NewReader(i)
	line := uint(0)
	column := uint(0)
	e := NewErrors()
	for {
		nextByte, err := input.ReadByte()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				e.Add(
					ErrIOFile,
					err.Error(),
					filename,
					line+1,
					column+1,
				)
			}
			return e
		}
		if nextByte > 0x7F {
			e.Add(
				ErrUnicode,
				fmt.Sprintf("character code: %s", strings.ToUpper(hex.EncodeToString([]byte{nextByte}))),
				filename,
				line+1,
				column+1,
			)
		}
		if nextByte == '\n' {
			line++
			column = 0
		} else {
			column += 1
		}
	}
}
