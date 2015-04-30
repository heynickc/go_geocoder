package main

import (
	"encoding/csv"
	"fmt"
	"github.com/mitchellh/ioprogress"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	mdNoZipGeocoderUrl   = "MD_CompositeLocator"
	mdWithZipGeocoderUrl = "MD_CompositeLocatorWithZIPCodeCentroids"
)

type Geocoder struct {
	URL *url.URL
}

func NewGeocoder(withZips bool) *Geocoder {

	u := new(url.URL)

	u.Scheme = "http"
	u.Host = "geodata.md.gov"

	if withZips {
		u.Path = "imap/rest/services/GeocodeServices/" + mdWithZipGeocoderUrl + "/GeocodeServer/findAddressCandidates"
	}
	u.Path = "imap/rest/services/GeocodeServices/" + mdNoZipGeocoderUrl + "/GeocodeServer/findAddressCandidates"

	v := url.Values{
		"Street":       []string{""},
		"City":         []string{""},
		"State":        []string{"Maryland"},
		"ZIP":          []string{""},
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

func GeocodeFile(inFileName, outFileName string) error {

	file, err := os.Open(inFileName)
	defer file.Close()
	if err != nil {
		return err
	}

	fileStat, err := file.Stat()
	if err != nil {
		return err
	}

	progressR := &ioprogress.Reader{
		Reader:       file,
		Size:         fileStat.Size(),
		DrawFunc:     drawTerminalBar(os.Stdout),
		DrawInterval: time.Microsecond,
	}

	reader := csv.NewReader(progressR)
	err = unmarshalAndGeocodeInRecords(reader, outFileName)
	if err != nil {
		return err
	}
	return nil
}

func drawTerminalBar(w io.Writer) ioprogress.DrawFunc {
	bar := ioprogress.DrawTextFormatBar(20)
	return ioprogress.DrawTerminalf(w, func(progress, total int64) string {
		return fmt.Sprintf(
			"%s %s",
			bar(progress, total),
			ioprogress.DrawTextFormatBytes(progress, total))

	})
}

func (g *Geocoder) setUrlValues(address *InRecord) {

	oldQuery := g.URL.Query()

	oldQuery.Set("Street", strings.ToUpper(address.Address))
	oldQuery.Set("ZIP", address.Zip)

	g.URL.RawQuery = oldQuery.Encode()
}

func (g Geocoder) Geocode() ([]byte, error) {

	res, err := http.Get(g.URL.String())
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (g Geocoder) geocodeToCandidates() ([]string, error) {

	res, err := http.Get(g.URL.String())
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	unmarshaler := JSONMarshaler{}
	candidates, err := unmarshaler.UnmarshalAddresses(res.Body)

	bestMatch := candidates.GetBestMatchLocation()

	return bestMatch, nil
}

func unmarshalAndGeocodeInRecords(reader *csv.Reader, outFileName string) error {

	var outRecords [][]string

	eof := false
	for lino := 1; !eof; lino++ {
		line, err := reader.Read()
		if lino == 1 {
			line = append(line, []string{"X", "Y", "AddressMatch", "MatchScore"}...)
			outRecords = append(outRecords, line)
			continue
		}
		if err == io.EOF {
			err = nil
			eof = true
			continue
		} else if err != nil {
			return err
		}

		parsedLine, err := parseAndGeocodeInRecord(line)
		outRecords = append(outRecords, parsedLine)
	}

	return outputNewRecords(outRecords, outFileName)
}

func outputNewRecords(newRecords [][]string, outFileName string) error {
	writer, closer, err := createCsvFile(outFileName)
	if closer != nil {
		defer closer()
	}
	if err != nil {
		return err
	}

	csvWriter := csv.NewWriter(writer)
	return csvWriter.WriteAll(newRecords)
}

func parseAndGeocodeInRecord(line []string) ([]string, error) {

	inRecord := &InRecord{}

	inRecord.Address = strings.ToUpper(line[8])
	inRecord.Zip = strings.ToUpper(line[9])

	gc := NewGeocoder(false)
	gc.setUrlValues(inRecord)

	xyVals, err := gc.geocodeToCandidates()

	if err != nil {
		return nil, err
	}

	line = append(line, xyVals...)
	return line, nil
}
