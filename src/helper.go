package jengo_src

import (
	"bytes"
	"errors"
	"fmt"
	"os"
)

func GetFileBytes(fileName string) []byte {
	f, err := os.Open(fileName)

	if err != nil {
		Info.Printf("Could not open file %s", fileName)
		os.Exit(1)
	}

	// Package bytes implements functions for the manipulation of byte slices. It is analogous to the facilities of the strings package.
	// Read in bytes in a single buffer, and convert them into a string when we are done.
	//
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(f)
	if err != nil {
		ErrorLog(errors.New(fmt.Sprintf("Could not read file %s", fileName)))
	}

	return buf.Bytes()
}