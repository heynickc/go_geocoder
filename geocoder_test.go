package main

import (
	// "fmt"
	// "net/url"
	"reflect"
	"testing"
)

func TestGeocoder(t *testing.T) {

	gc := NewGeocoder()

	if reflect.TypeOf(gc).String() != "*main.Geocoder" {
		t.Errorf("Type of gc is %v", reflect.TypeOf(gc).String())
	}
}

func TestMakeUrlValues(t *testing.T) {
	t.Skip()

	inRec := &InRecord{"507 N PINEHURST AVE", "21801"}
	gc := NewGeocoder()

	gc.SetUrlValues(inRec)

	equals(t, "http://geodata.md.gov/imap/rest/services/GeocodeServices/MD_CompositeLocatorWithZIPCodeCentroids/GeocodeServer/findAddressCandidates?City=&SingleLine=&State=Maryland&Street=507+N+PINEHURST+AVE&ZIP=21801&f=json&maxLocations=United+States&outFields=&outSR=4326&searchExtent=", gc.URL.String())
}
