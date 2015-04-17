package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func readDataFile(filename string) (*Candidates, error) {
	file, closer, err := openDataFile(filename)
	if closer != nil {
		defer closer()
	}
	if err != nil {
		return nil, err
	}
	return readData(file, suffixOf(filename))
}

func openDataFile(filename string) (io.ReadCloser, func(), error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	closer := func() { file.Close() }
	var reader io.ReadCloser = file
	return reader, closer, nil
}

func readData(reader io.Reader, suffix string) (*Candidates, error) {
	var unmarshaler AddressUnmarshaler
	switch suffix {
	// case ".gob":
	// 	unmarshaler = GobMarshaler{}
	// case ".inv":
	// 	unmarshaler = InvMarshaler{}
	case ".jsn", ".json":
		unmarshaler = JSONMarshaler{}
		// case ".txt":
		// 	unmarshaler = TxtMarshaler{}
		// case ".xml":
		// 	unmarshaler = XMLMarshaler{}
	}
	if unmarshaler != nil {
		return unmarshaler.UnmarshalAddresses(reader)
	}
	return nil, fmt.Errorf("unrecognized input suffix: %s", suffix)
}

func suffixOf(filename string) string {
	suffix := filepath.Ext(filename)
	if suffix == ".gz" {
		suffix = filepath.Ext(filename[:len(filename)-3])
	}
	return suffix
}
