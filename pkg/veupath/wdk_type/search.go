package wdk_type

type Search struct {
	UrlSegment                           string      `json:"urlSegment"`
	FullName                             string      `json:"fullName"`
	DisplayName                          string      `json:"displayName"`
	ShortDisplayName                     string      `json:"shortDisplayName"`
	Description                          string      `json:"description"`
	IconName                             string      `json:"iconName"`
	Summary                              string      `json:"summary"`
	Help                                 *string     `json:"help,omitempty"`
	NewBuild                             *string     `json:"newBuild,omitempty"`
	ReviseBuild                          *string     `json:"reviseBuild,omitempty"`
	OutputRecordClassName                string      `json:"outputRecordClassName"`
	Filters                              []Filter    `json:"filters"`
	DefaultAttributes                    []string    `json:"defaultAttributes"`
	DefaultSorting                       []SortSpec  `json:"defaultSorting"`
	DynamicAttributes                    []Attribute `json:"dynamicAttributes"`
	DefaultSummaryView                   string      `json:"defaultSummaryView"`
	SummaryViewPlugins                   []Plugin    `json:"summaryViewPlugins"`
	IsAnalyzable                         bool        `json:"isAnalyzable"`
	AllowedPrimaryInputRecordClassNames  []string    `json:"allowedPrimaryInputRecordClassNames"`
	AllowedSecondaryInputRecordClassName []string    `json:"allowedSecondaryInputRecordClassName"`
	QueryName                            string      `json:"queryName"`
	NoSummaryOnSingleRecord              bool        `json:"noSummaryOnSingleRecord"`
	Properties                           Properties  `json:"properties"`
	Groups                               []Group     `json:"groups"`
	ParamNames                           []string    `json:"paramNames"`
	Parameters                           []Parameter `json:"parameters"`
}
