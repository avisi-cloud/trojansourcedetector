package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data := map[string]string{
		"testdata/unicode.txt": "Hello world with an accent: \u00E1",
		"testdata/bidi.txt":    "Hello world with a BIDI control character: \u202A",
	}

	for file, contents := range data {
		if err := ioutil.WriteFile(
			file,
			[]byte(contents),
			0,
		); err != nil {
			panic(fmt.Errorf("failed to write %s (%w)", file, err))
		}
	}
}
