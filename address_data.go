package main

type Address struct {
	Address  string
	Location Location
	Score    float32
}

type Location struct {
	X float32
	Y float32
}

type AddressMarshaler interface {
	MarshalAddresses(writer io.Writer, invoices []*Address) error
}

type AddressUnmarshaler interface {
	UnmarshalAddresses(reader io.Reader) ([]*Address, error)
}
