package main

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type Address struct {
	Address  string
	Location Location
	Score    float32
}

type Location struct {
	X float32
	Y float32
}

type AddressMarshaler interface {
	MarshalAddresses(writer io.Writer, addresses []*Address) error
}

type AddressUnmarshaler interface {
	UnmarshalAddresses(reader io.Reader) ([]*Address, error)
}

type JSONMarshaler struct{}

func (JSONMarshaler) MarshalAddresses(writer io.Writer,
	addresses []*Address) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(addresses)
}

func (JSONMarshaler) UnmarshalAddresses(reader io.Reader) ([]*Address,
	error) {
	decoder := json.NewDecoder(reader)
	var addresses []*Address
	if err := decoder.Decode(&addresses); err != nil {
		return nil, err
	}
	return addresses, nil
}

type Geocoder struct {
	URL *url.URL
}

func NewGeocoder() *Geocoder {

	u := new(url.URL)

	u.Scheme = "http"
	u.Host = "geodata.md.gov"
	u.Path = "imap/rest/services/GeocodeServices/MD_CompositeLocatorWithZIPCodeCentroids/GeocodeServer/findAddressCandidates"

	v := url.Values{
		"Street":       []string{"507 S Pinehurst Ave"},
		"City":         []string{"Salisbury"},
		"State":        []string{"Maryland"},
		"ZIP":          []string{"21801"},
		"SingleLine":   []string{""},
		"outFields":    []string{""},
		"maxLocations": []string{"United States"},
		"outSR":        []string{"4326"},
		"searchExtent": []string{""},
		"f":            []string{"json"},
	}
	u.RawQuery = v.Encode()

	return &Geocoder{u}
}

func (g *Geocoder) Geocode() []byte {

	res, err := http.Get(g.URL.String())
	if err != nil {
		log.Fatalf("Unable to get %q: %v", g.URL.String(), err)
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		log.Fatalf("Error parsing response body: %v", err)
	}

	return data
}

func main() {
	addresses, err := readDataFile("./my_address.json")
	if err != nil {
		log.Fatalln("Failed to read:", err)
	}

	fmt.Println(addresses)
}

func readDataFile(filename string) ([]*Address, error) {
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
	var decompressor *gzip.Reader
	if strings.HasSuffix(filename, ".gz") {
		if decompressor, err = gzip.NewReader(file); err != nil {
			return file, closer, err
		}
		closer = func() { decompressor.Close(); file.Close() }
		reader = decompressor
	}
	return reader, closer, nil
}

func readData(reader io.Reader, suffix string) ([]*Address, error) {
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
