package main

import (
	"encoding/json"
	"io"
	"sort"
)

type SpatialReference struct {
	SpatialReference Wkid
}

type Wkid struct {
	Wkid       int
	LatestWkid int
}

type Candidates struct {
	Candidates []*Address
}

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
	MarshalAddresses(writer io.Writer, candidates Candidates) error
}

type AddressUnmarshaler interface {
	UnmarshalAddresses(reader io.Reader) (*Candidates, error)
}

type JSONMarshaler struct{}

func (JSONMarshaler) MarshalAddresses(writer io.Writer,
	candidates Candidates) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(candidates)
}

func (JSONMarshaler) UnmarshalAddresses(reader io.Reader) (*Candidates, error) {
	decoder := json.NewDecoder(reader)
	var candidates *Candidates
	if err := decoder.Decode(&candidates); err != nil {
		return nil, err
	}
	return candidates, nil
}

type ByScore []*Address

func (c ByScore) Len() int           { return len(c) }
func (c ByScore) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByScore) Less(i, j int) bool { return c[i].Score < c[j].Score }

func (c *Candidates) SortCandidates(inRec *InRecord) {
	sort.Sort(sort.Reverse(ByScore(c.Candidates)))

}

func (c Candidates) GetScores() []float32 {
	scores := []float32{}
	for i := 0; i < len(c.Candidates); i++ {
		scores = append(scores, c.Candidates[i].Score)
	}
	return scores
}
