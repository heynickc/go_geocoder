package main

import (
	geocoder "github.com/heynickc/go_mdimapgeocoder"
	"io/ioutil"
	"testing"
)

func TestOpenDataFile(t *testing.T) {

	file, _, _ := geocoder.OpenDataFile("./my_address.json")

	byteResult, _ := ioutil.ReadAll(file)
	result := string(byteResult)

	if len(result) == 0 {
		t.Errorf("OpenDataFile didn't open the file %v", len(result))
	}
}

func TestReadDataFile(t *testing.T) {

	file, _, _ := geocoder.OpenDataFile("./my_address.json")

	marshaler := geocoder.JSONMarshaler{}

	byteResult, _ := marshaler.UnmarshalAddresses(file)

	t.Error(byteResult)

}
