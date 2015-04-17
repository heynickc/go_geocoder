package main

import (
	"io/ioutil"
	"testing"
)

func TestOpenDataFile(t *testing.T) {

	file, _, err := openAddressDataFile("./my_address.json")

	byteResult, _ := ioutil.ReadAll(file)
	result := string(byteResult)

	ok(t, err)
	equals(t, 1574, len(result))
}
