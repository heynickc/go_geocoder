package main

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ReadDataFile(filename string) ([]*Address, error) {
	file, closer, err := OpenDataFile(filename)
	if closer != nil {
		defer closer()
	}
	if err != nil {
		return nil, err
	}
	return ReadData(file, suffixOf(filename))
}

func OpenDataFile(filename string) (io.ReadCloser, func(), error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	closer := func() { file.Close() }
	var reader io.ReadCloser = file
	return reader, closer, nil
}

func ReadData(reader io.Reader, suffix string) ([]*Address, error) {
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

func writeDataFile(filename string, addresses []*Address) error {
	file, closer, err := createDataFile(filename)
	if closer != nil {
		defer closer()
	}
	if err != nil {
		return err
	}
	return writeData(file, suffixOf(filename), addresses)
}

func createDataFile(filename string) (io.WriteCloser, func(), error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, nil, err
	}
	closer := func() { file.Close() }
	var writer io.WriteCloser = file
	var compressor *gzip.Writer
	if strings.HasSuffix(filename, ".gz") {
		compressor = gzip.NewWriter(file)
		closer = func() { compressor.Close(); file.Close() }
		writer = compressor
	}
	return writer, closer, nil
}

func writeData(writer io.Writer, suffix string,
	addresses []*Address) error {
	var marshaler AddressMarshaler
	switch suffix {
	// case ".gob":
	// 	marshaler = GobMarshaler{}
	// case ".inv":
	// 	marshaler = InvMarshaler{}
	case ".jsn", ".json":
		marshaler = JSONMarshaler{}
		// case ".txt":
		// 	marshaler = TxtMarshaler{}
		// case ".xml":
		// 	marshaler = XMLMarshaler{}
	}
	if marshaler != nil {
		return marshaler.MarshalAddresses(writer, addresses)
	}
	return errors.New("unrecognized output suffix")
}

func suffixOf(filename string) string {
	suffix := filepath.Ext(filename)
	if suffix == ".gz" {
		suffix = filepath.Ext(filename[:len(filename)-3])
	}
	return suffix
}
