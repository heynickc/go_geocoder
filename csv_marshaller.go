package main

import (
	"encoding/csv"
	"io"
	"strings"
)

type Aliases struct {
	Street []string
	City   []string
	State  []string
	ZIP    []string
}

func unmarshalInRecords(reader *csv.Reader) (inRecords []*InRecord, err error) {

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
		inRecords = append(inRecords, parseInRecord(line))
	}

	return inRecords, nil
}

func parseInRecord(line []string) (inRecord *InRecord) {

	inRecord = &InRecord{}

	inRecord.Address = strings.ToUpper(line[8])
	inRecord.Zip = strings.ToUpper(line[9])

	return inRecord
}
