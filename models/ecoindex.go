package models

type Ecoindex struct {
	Id       string  `json:"id"`
	Grade    string  `json:"grade"`
	Score    float32 `json:"score"`
	Date     string  `json:"date"`
	Requests int     `json:"requests"`
	Size     float32 `json:"size"`
	Nodes    int     `json:"nodes"`
	Url      string  `json:"url"`
}

type EcoindexSearchResults struct {
	Count        int        `json:"count"`
	LatestResult Ecoindex   `json:"latest-result"`
	OlderResults []Ecoindex `json:"older-results"`
	HostResults  []Ecoindex `json:"host-results"`
}
