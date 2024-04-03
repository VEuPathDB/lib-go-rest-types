package wdk_type

type OrganismSearchRequest struct {
	SearchConfig SearchConfig `json:"searchConfig"`
	ReportConfig ReportConfig `json:"reportConfig"`
}
