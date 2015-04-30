package main

import (
	"encoding/csv"
	"fmt"
	"github.com/mitchellh/ioprogress"
	"os"
	"testing"
)

type InAddress struct {
	Address string
	Zip     string
}

func TestOpenCSVFile(t *testing.T) {
	file, err := os.Open("./sso_db_raw.csv")
	ok(t, err)

	defer file.Close()

	reader := csv.NewReader(file)
	ok(t, err)

	data, err := reader.ReadAll()
	ok(t, err)

	equals(t, 3758, len(data))
}

func TestUnmarshalInRecords(t *testing.T) {
	file, err := os.Open("./sso_db_raw.csv")
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

	equals(t, 3758, len(data))
}

func TestGeocodeInRecords(t *testing.T) {
	t.Skip("Just to see how to do this appropriately")

	file, err := os.Open("./sso_db_raw.csv")
	ok(t, err)

	defer file.Close()

	reader := csv.NewReader(file)
	ok(t, err)

	data, err := unmarshalInRecords(reader)
	ok(t, err)

	gc := NewGeocoder(false)
	for i := 0; i < 5; i++ {
		gc.setUrlValues(data[i])

		fmt.Println(data[i].Address)
		parsedData, err := gc.Geocode()
		ok(t, err)

		fmt.Println(string(parsedData) + "\n")
	}
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
