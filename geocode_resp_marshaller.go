package main

import (
	"encoding/json"
	"io"
)

type GeocodeRespMarshaler interface {
	MarshalAddresses(writer io.Writer, geocodeResp GeocodeResp) error
}

type GeocodeRespUnmarshaler interface {
	UnmarshalAddresses(reader io.Reader) (*GeocodeResp, error)
}

type GeocodeRespJSONMarshaler struct{}

func (GeocodeRespJSONMarshaler) MarshalAddresses(writer io.Writer, gr GeocodeResp) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(gr)
}

func (GeocodeRespJSONMarshaler) UnmarshalAddresses(reader io.Reader) (*GeocodeResp, error) {
	decoder := json.NewDecoder(reader)
	var gr *GeocodeResp
	if err := decoder.Decode(&gr); err != nil {
		return nil, err
	}
	return gr, nil
}
