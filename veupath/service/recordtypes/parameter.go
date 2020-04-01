package recordtypes

type Parameter struct {
	Name                string   `json:"name"`
	DisplayName         string   `json:"displayName"`
	Help                string   `json:"help"`
	Type                string   `json:"type"`
	IsVisible           bool     `json:"isVisible"`
	Group               string   `json:"group"`
	IsReadOnly          bool     `json:"isReadOnly"`
	AllowEmptyValue     bool     `json:"allowEmptyValue"`
	VisibleHelp         string   `json:"visibleHelp"`
	DependentParams     []string `json:"dependentParams"`
	InitialDisplayValue string   `json:"initialDisplayValue"`

	//
	// DatasetParamFormatter
	//
	DefaultIdList   *string  `json:"defaultIdList"`
	RecordClassName *string  `json:"recordClassName"`
	Parsers         []Parser `json:"parsers"`

	//
	// DateParamFormatter & DateRangeParamFormatter
	//
	MinDate *string `json:"minDate"`
	MaxDate *string `json:"maxDate"`

	//
	// EnumParamFormatter
	//
	DisplayType      *string `json:"displayType"`
	MaxSelectedCount *int    `json:"maxSelectedCount"`
	MinSelectedCount *int    `json:"minSelectedCount"`
	//Vocabulary       [][3]string `json:"vocabulary"`

	//
	// FilterParamNewFormatter
	//
	FilterDataTypeDisplayName *string             `json:"filterDataTypeDisplayName"`
	Ontology                  []Ontology          `json:"ontology"`
	Values                    map[string][]string `json:"values"`
	HideEmptyOntologyNodes    *bool               `json:"hideEmptyOntologyNodes"`
	SortLeavesBeforeBranches  *bool               `json:"sortLeavesBeforeBranches"`
	CountPredictsAnswerCount  *bool               `json:"countPredictsAnswerCount"`
	// MinSelectedCount int `json:"minSelectedCount"`

	//
	// NumberParamFormatter & NumberRangeFormatter
	//
	MinValue  *float64 `json:"minValue"`
	MaxValue  *float64 `json:"maxValue"`
	Increment *float64 `json:"increment"`

	//
	// StringParamFormatter
	//
	Length *int `json:"length"`

	//
	// TreeBoxEnumParamFormatter
	//
	//Vocabulary       EnumParamTermNode `json:"vocabulary"`


	//
	// EnumParamFormatter & TreeBoxEnumParamFormatter
	//
	Vocabulary []byte `json:"vocabulary"`

}

