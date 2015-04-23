package main

import (
	"encoding/csv"
	// "fmt"
	"io"
	"strings"
)

type Aliases struct {
	Street []string
	City   []string
	State  []string
	ZIP    []string
}

type InRecord struct {
	Address string
	Zip     string
}

func UnmarshalInRecords(reader *csv.Reader) (inRecords []*InRecord, err error) {

	eof := false
	for lino := 2; !eof; lino++ {
		line, err := reader.Read()
		if err == io.EOF {
			err = nil
			eof = true
			return inRecords, nil
		} else if err != nil {
			return nil, err
		}
		inRecords = append(inRecords, ParseInRecord(line))
	}

	return inRecords, nil
}

func UnmarshalAndGeocodeInRecords(reader *csv.Reader) (outRecords [][]string, err error) {

	eof := false
	for lino := 2; !eof; lino++ {
		line, err := reader.Read()
		if err == io.EOF {
			err = nil
			eof = true
			return outRecords, nil
		} else if err != nil {
			return nil, err
		}
		outRecords = append(outRecords, ParseAndGeocodeInRecord(line))
	}

	return outRecords, nil
}

func ParseInRecord(line []string) (inRecord *InRecord) {

	inRecord = &InRecord{}

	inRecord.Address = strings.ToUpper(line[8])
	inRecord.Zip = strings.ToUpper(line[9])

	return inRecord
}

func ParseAndGeocodeInRecord(line []string) []string {

	inRecord := &InRecord{}

	inRecord.Address = strings.ToUpper(line[8])
	inRecord.Zip = strings.ToUpper(line[9])

	gc := NewGeocoder()
	gc.SetUrlValues(inRecord)
	gCode := gc.Geocode()

	line = append(line, string(gCode))
	return line
}
