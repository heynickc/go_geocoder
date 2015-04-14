package main

import (
	"io/ioutil"
	"log"
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

	return &Geocoder{u}
}

func (g *Geocoder) Geocode() []byte {

	res, err := http.Get(g.URL.String())
	if err != nil {
		log.Fatalf("Unable to get %q: %v", g.URL.String(), err)
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		log.Fatalf("Error parsing response body: %v", err)
	}

	return data
}

func main() {

}
