package geocoder

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"
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
	Location struct {
		X float32
		Y float32
	}
	Score float32
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
	candidates.SortCandidates()
	return candidates, nil
}

type ByScore []*Address

func (c ByScore) Len() int           { return len(c) }
func (c ByScore) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByScore) Less(i, j int) bool { return c[i].Score < c[j].Score }

func (c *Candidates) SortCandidates() {
	sort.Sort(sort.Reverse(ByScore(c.Candidates)))
}

func (c *Candidates) GetBestMatchLocation() []string {
	var bestMatch []string
	c.SortCandidates()
	if len(c.Candidates) > 0 && c.Candidates[0] != nil {
		xVal := fmt.Sprintf("%.6f", c.Candidates[0].Location.X)
		yVal := fmt.Sprintf("%.6f", c.Candidates[0].Location.Y)
		matchedAddr := strings.ToUpper(c.Candidates[0].Address)
		score := fmt.Sprintf("%.6f", c.Candidates[0].Score)

		bestMatch = []string{xVal, yVal, matchedAddr, score}
	} else {
		bestMatch = []string{}
	}
	return bestMatch
}
