package main

import (
	"reflect"
	"testing"
)

func TestGeocoder(t *testing.T) {

	gc := NewGeocoder()

	if reflect.TypeOf(gc).String() != "*main.Geocoder" {
		t.Errorf("Type of gc is %v", reflect.TypeOf(gc).String())
	}
}
