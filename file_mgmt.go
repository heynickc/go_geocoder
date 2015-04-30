package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func readAddressDataFile(filename string) (*Candidates, error) {
	file, closer, err := openAddressDataFile(filename)
	if closer != nil {
		defer closer()
	}
	if err != nil {
		return nil, err
	}
	return readAddressData(file, suffixOf(filename))
}

func openAddressDataFile(filename string) (io.ReadCloser, func(), error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	closer := func() { file.Close() }
	var reader io.ReadCloser = file
	return reader, closer, nil
}

func readAddressData(reader io.Reader, suffix string) (*Candidates, error) {
	var unmarshaler AddressUnmarshaler
	unmarshaler = JSONMarshaler{}
	if unmarshaler != nil {
		return unmarshaler.UnmarshalAddresses(reader)
	}
	return nil, fmt.Errorf("unrecognized input suffix: %s", suffix)
}

func createCsvFile(filename string) (io.WriteCloser, func(), error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, nil, err
	}
	closer := func() { file.Close() }
	var writer io.WriteCloser = file
	return writer, closer, nil
}

func suffixOf(filename string) string {
	suffix := filepath.Ext(filename)
	if suffix == ".gz" {
		suffix = filepath.Ext(filename[:len(filename)-3])
	}
	return suffix
}
