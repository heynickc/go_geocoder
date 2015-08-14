package main

import "sort"

type GeocodeResp struct {
	SpatialReference struct {
		Wkid       int
		LatestWkId int
	}
	Candidates []*Candidate
}

type Candidate struct {
	Address  string
	Location struct {
		X float32
		Y float32
	}
	Score float32
}

type ByCandidateScore []*Candidate

func (c ByCandidateScore) Len() int           { return len(c) }
func (c ByCandidateScore) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByCandidateScore) Less(i, j int) bool { return c[i].Score < c[j].Score }

func (g GeocodeResp) GetBestCandidate() *Candidate {
	sort.Sort(ByCandidateScore(g.Candidates))
	if len(g.Candidates) > 0 {
		return g.Candidates[len(g.Candidates)-1]
	} else {
		return &Candidate{}
	}
}
