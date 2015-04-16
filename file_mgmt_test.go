package main

import (
	"io/ioutil"
	"testing"
)

func TestOpenDataFile(t *testing.T) {

	file, _, _ := openDataFile("./my_address.json")

	byteResult, _ := ioutil.ReadAll(file)
	result := string(byteResult)

	if len(result) == 0 {
		t.Errorf("OpenDataFile didn't open the file %v", len(result))
	}
}
