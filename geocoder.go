package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/heynickc/go_geocoder/geocoder"
	"github.com/mitchellh/ioprogress"
)

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

	gc := geocoder.NewGeocoder(false)
	gc.SetURLValues(inRecord)

	xyVals, err := gc.geocodeToCandidates()

	if err != nil {
		return nil, err
	}

	line = append(line, xyVals...)
	return line, nil
}
