package recordtypes

type OrganismSearchRequest struct {
	SearchConfig SearchConfig `json:"searchConfig"`
	ReportConfig ReportConfig `json:"reportConfig"`
}

type SearchConfig struct {
	Parameters map[string]string `json:"parameters"`
}

type ReportConfig struct {
	Attributes []string `json:"attributes"`
}

func NewOrganismSearchRequest() *OrganismSearchRequest {
	return &OrganismSearchRequest{
		SearchConfig: SearchConfig{Parameters: make(map[string]string)},
		ReportConfig: ReportConfig{Attributes: []string{"organism"}},
	}
}