package main

import (
	"net/url"
)

type MDiMapGeocoder struct {
	*url.URL
}

func New() *MDiMapGeocoder {

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

	return &MDiMapGeocoder{u}
}
