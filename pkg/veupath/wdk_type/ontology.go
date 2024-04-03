package wdk_type

type Ontology struct {
	Term      string `json:"term"`
	Parent    string `json:"parent"`
	Display   string `json:"display"`
	Type      string `json:"type"`
	Units     string `json:"units"`
	Precision int64  `json:"precision"`
	IsRange   bool   `json:"isRange"`
}
