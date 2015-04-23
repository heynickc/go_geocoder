package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Geocoder struct {
	URL *url.URL
}

func NewGeocoder() *Geocoder {

	u := new(url.URL)

	u.Scheme = "http"
	u.Host = "geodata.md.gov"
	u.Path = "imap/rest/services/GeocodeServices/MD_CompositeLocatorWithZIPCodeCentroids/GeocodeServer/findAddressCandidates"

	v := url.Values{
		"Street":       []string{"507 S Pinehurst Ave"},
		"City":         []string{""},
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

	return &Geocoder{u}
}

func (g *Geocoder) SetUrlValues(address *InRecord) {

	oldQuery := g.URL.Query()

	oldQuery.Set("Street", address.Address)
	oldQuery.Set("ZIP", address.Zip)

	g.URL.RawQuery = oldQuery.Encode()
}

func (g Geocoder) Geocode() ([]byte, error) {

	res, err := http.Get(g.URL.String())
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (g Geocoder) GeocodeToCandidates() ([]byte, error) {

	res, err := http.Get(g.URL.String())
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	unmarshaler := JSONMarshaler{}
	candidates, err := unmarshaler.UnmarshalAddresses(res.Body)

	fmt.Println(candidates.Candidates)

	return nil, nil
}
