package main

import (
	"github.com/gorilla/http"
	"log"
	"os"
)

type Address struct {
	street, city, state, zip string
}

func (a *Address) SingleLine() string {
	return a.street + " " + a.city + "," + a.state + " " + a.zip
}

func (m *MDiMapGeocoder) Geocode(url string) {
	if _, err := http.Get(os.Stdout, url); err != nil {
		log.Fatalf("unable to fetch %q: %v", os.Args[1], err)
	}
}

func main() {

	var g = MDiMapGeocoder.New()

	g.Geocode(g.URL.String())

}
