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

func UnmarshalInRecordsAppendXY(reader *csv.Reader) (inRecords []*InRecord, err error) {

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
		inRecords = append(inRecords, ParseInRecordAppendXY(line))
	}

	return inRecords, nil
}

func ParseInRecord(line []string) (inRecord *InRecord) {

	inRecord = &InRecord{}

	inRecord.Address = strings.ToUpper(line[8])
	inRecord.Zip = strings.ToUpper(line[9])

	return inRecord
}

func ParseInRecordAppendXY(line []string) (inRecord *InRecord) {

	inRecord = &InRecord{}

	inRecord.Address = strings.ToUpper(line[8])
	inRecord.Zip = strings.ToUpper(line[9])

	return inRecord
}
