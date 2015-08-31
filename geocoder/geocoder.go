package geocoder

import (
	"net/http"
	"net/url"
	"strings"
)

const (
	mdNoZipGeocoderURL   = "MD_CompositeLocator"
	mdWithZipGeocoderURL = "MD_CompositeLocatorWithZIPCodeCentroids"
)

type IGeocoder interface {
	Geocode(address string) string
}

type Geocoder struct {
	URL *url.URL
}

type InRecord struct {
	Address string
	Zip     string
}

func NewGeocoder(withZips bool) *Geocoder {

	u := new(url.URL)

	u.Scheme = "http"
	u.Host = "geodata.md.gov"

	if withZips {
		u.Path = "imap/rest/services/GeocodeServices/" + mdWithZipGeocoderURL + "/GeocodeServer/findAddressCandidates"
	}
	u.Path = "imap/rest/services/GeocodeServices/" + mdNoZipGeocoderURL + "/GeocodeServer/findAddressCandidates"

	v := url.Values{
		"Street":       []string{""},
		"City":         []string{""},
		"State":        []string{"Maryland"},
		"ZIP":          []string{""},
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

func (g *Geocoder) SetURLValues(address *InRecord) {

	oldQuery := g.URL.Query()

	oldQuery.Set("Street", strings.ToUpper(address.Address))
	oldQuery.Set("ZIP", address.Zip)

	g.URL.RawQuery = oldQuery.Encode()
}

func (g Geocoder) geocodeToCandidates() ([]string, error) {

	res, err := http.Get(g.URL.String())
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	unmarshaler := JSONMarshaler{}
	candidates, err := unmarshaler.UnmarshalAddresses(res.Body)

	bestMatch := candidates.GetBestMatchLocation()

	return bestMatch, nil
}
