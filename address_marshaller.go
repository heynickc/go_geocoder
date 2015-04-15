package main

import (
	"encoding/json"
	"io"
)

type Address struct {
	address  string
	location Location
	score    float32
}

type Location struct {
	x float32
	y float32
}

type AddressMarshaler interface {
	MarshalAddresses(writer io.Writer, addresses []*Address) error
}

type AddressUnmarshaler interface {
	UnmarshalAddresses(reader io.Reader) ([]*Address, error)
}

type JSONMarshaler struct{}

func (JSONMarshaler) MarshalAddresses(writer io.Writer,
	addresses []*Address) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(addresses)
}

func (JSONMarshaler) UnmarshalAddresses(reader io.Reader) ([]*Address,
	error) {
	decoder := json.NewDecoder(reader)
	var addresses []*Address
	if err := decoder.Decode(&addresses); err != nil {
		return nil, err
	}
	return addresses, nil
}
