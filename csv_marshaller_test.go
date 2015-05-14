package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"testing"

	"github.com/mitchellh/ioprogress"
)

type InAddress struct {
	Address string
	Zip     string
}

func TestOpenCSVFile(t *testing.T) {
	file, err := os.Open("./sso_db_raw_sample.csv")
	ok(t, err)

	defer file.Close()

	reader := csv.NewReader(file)
	ok(t, err)

	data, err := reader.ReadAll()
	ok(t, err)

	equals(t, 20, len(data))
}

func TestUnmarshalInRecords(t *testing.T) {
	file, err := os.Open("./sso_db_raw_sample.csv")
	ok(t, err)

	defer file.Close()

	fileStat, err := file.Stat()
	ok(t, err)

	fmt.Println(fileStat.Size())

	progressR := &ioprogress.Reader{
		Reader: file,
		Size:   fileStat.Size(),
	}

	reader := csv.NewReader(progressR)

	ok(t, err)

	data, err := unmarshalInRecords(reader)
	ok(t, err)

	equals(t, 20, len(data))
}

func TestCsvWriter(t *testing.T) {
	t.Skip("Just to see how to do this appropriately")

	testData := [][]string{[]string{"test"}}
	writer, err := os.Create("./test.csv")
	ok(t, err)
	csvWriter := csv.NewWriter(writer)
	csvWriter.WriteAll(testData)
	writer.Close()
}
