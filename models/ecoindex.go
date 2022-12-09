package models

type Ecoindex struct {
	Id       string
	Grade    string
	Score    float32
	Date     string
	Requests int
	Size     float32
	Nodes    int
	Url      string
}

type EcoindexSearchResults struct {
	Count        int
	LatestResult Ecoindex
	OlderResults []Ecoindex
	HostResults  []Ecoindex
}
