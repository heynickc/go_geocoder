package main

import (
	"log"
	"net/url"
	"testing"
)

type TestPairURL struct {
	In  string
	Out map[string]string
}

var TestPairsURL = []TestPairURL{TestPairURL{"http://geodata.md.gov/imap/rest/services/GeocodeServices/MD_CompositeLocatorWithZIPCodeCentroids/GeocodeServer/findAddressCandidates?Street=507+S+Pinehurst+Ave&City=Salisbury&State=Maryland&ZIP=21801&SingleLine=&outFields=&maxLocations=United+States&outSR=4326&searchExtent=&f=pjson",
	map[string]string{
		"Scheme":   "http",
		"Host":     "geodata.md.gov",
		"Path":     "/imap/rest/services/GeocodeServices/MD_CompositeLocatorWithZIPCodeCentroids/GeocodeServer/findAddressCandidates",
		"RawQuery": "Street=507+S+Pinehurst+Ave&City=Salisbury&State=Maryland&ZIP=21801&SingleLine=&outFields=&maxLocations=United+States&outSR=4326&searchExtent=&f=pjson"},
}}

func TestNewUrl(t *testing.T) {

	result, err := url.Parse(TestPairsURL[0].In)
	if err != nil {
		log.Fatal(err)
	}

	if result.Scheme != TestPairsURL[0].Out["Scheme"] {
		t.Errorf("Expected %v, but got %v", result.Scheme, TestPairsURL[0].Out["Scheme"])
	}
	if result.Host != TestPairsURL[0].Out["Host"] {
		t.Errorf("Expected %v, but got %v", result.Host, TestPairsURL[0].Out["Host"])
	}
	if result.Path != TestPairsURL[0].Out["Path"] {
		t.Errorf("Expected %v, but got %v", result.Path, TestPairsURL[0].Out["Path"])
	}
	if result.RawQuery != TestPairsURL[0].Out["RawQuery"] {
		t.Errorf("Expected %v, but got %v", result.RawQuery, TestPairsURL[0].Out["RawQuery"])
	}
}

func TestUrlBuilder(t *testing.T) {

	u := new(url.URL)

	u.Scheme = "http"
	u.Host = "geodata.md.gov"
	u.Path = "imap/rest/services/GeocodeServices/MD_CompositeLocatorWithZIPCodeCentroids/GeocodeServer/findAddressCandidates"

	v := url.Values{
		"Street":       []string{"507 S Pinehurst Ave"},
		"City":         []string{"Salisbury"},
		"State":        []string{"Maryland"},
		"ZIP":          []string{"21801"},
		"SingleLine":   []string{""},
		"outFields":    []string{""},
		"maxLocations": []string{"United States"},
		"outSR":        []string{"4326"},
		"searchExtent": []string{""},
		"f":            []string{"json"},
	}
	u.RawQuery = v.Encode()

	expected := "http://geodata.md.gov/imap/rest/services/GeocodeServices/MD_CompositeLocatorWithZIPCodeCentroids/GeocodeServer/findAddressCandidates?City=Salisbury&SingleLine=&State=Maryland&Street=507+S+Pinehurst+Ave&ZIP=21801&f=json&maxLocations=United+States&outFields=&outSR=4326&searchExtent="

	if u.String() != expected {
		t.Errorf("Expected %v, but got %v", u.String(), expected)
	}
}

func TestUrlBuilderSingleLine(t *testing.T) {

	u := new(url.URL)

	u.Scheme = "http"
	u.Host = "geodata.md.gov"
	u.Path = "imap/rest/services/GeocodeServices/MD_CompositeLocatorWithZIPCodeCentroids/GeocodeServer/findAddressCandidates"

	v := url.Values{
		"SingleLine":   []string{"507 S Pinehurst Ave Salisbury, MD 21801"},
		"outFields":    []string{""},
		"maxLocations": []string{"United States"},
		"outSR":        []string{"4326"},
		"searchExtent": []string{""},
		"f":            []string{"json"},
	}
	u.RawQuery = v.Encode()

	expected := "http://geodata.md.gov/imap/rest/services/GeocodeServices/MD_CompositeLocatorWithZIPCodeCentroids/GeocodeServer/findAddressCandidates?SingleLine=507+S+Pinehurst+Ave+Salisbury%2C+MD+21801&f=json&maxLocations=United+States&outFields=&outSR=4326&searchExtent="

	if u.String() != expected {
		t.Errorf("Expected %v, but got %v", u.String(), expected)
	}
}
