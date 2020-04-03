package service

import "encoding/json"

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
	DefaultIdList   *string  `json:"defaultIdList,omitempty"`
	RecordClassName *string  `json:"recordClassName,omitempty"`
	Parsers         []Parser `json:"parsers,omitempty"`

	//
	// DateParamFormatter & DateRangeParamFormatter
	//
	MinDate *string `json:"minDate,omitempty"`
	MaxDate *string `json:"maxDate,omitempty"`

	//
	// EnumParamFormatter
	//
	DisplayType      *string `json:"displayType,omitempty"`
	MaxSelectedCount *int    `json:"maxSelectedCount,omitempty"`
	MinSelectedCount *int    `json:"minSelectedCount,omitempty"`
	//Vocabulary       [][3]string `json:"vocabulary"`

	//
	// FilterParamNewFormatter
	//
	FilterDataTypeDisplayName *string             `json:"filterDataTypeDisplayName,omitempty"`
	Ontology                  []Ontology          `json:"ontology,omitempty"`
	Values                    map[string][]string `json:"values,omitempty"`
	HideEmptyOntologyNodes    *bool               `json:"hideEmptyOntologyNodes,omitempty"`
	SortLeavesBeforeBranches  *bool               `json:"sortLeavesBeforeBranches,omitempty"`
	CountPredictsAnswerCount  *bool               `json:"countPredictsAnswerCount,omitempty"`
	// MinSelectedCount int `json:"minSelectedCount"`

	//
	// NumberParamFormatter & NumberRangeFormatter
	//
	MinValue  *float64 `json:"minValue,omitempty"`
	MaxValue  *float64 `json:"maxValue,omitempty"`
	Increment *float64 `json:"increment,omitempty"`

	//
	// StringParamFormatter
	//
	Length *int `json:"length,omitempty"`

	//
	// TreeBoxEnumParamFormatter
	//
	//Vocabulary       EnumParamTermNode `json:"vocabulary"`

	//
	// EnumParamFormatter & TreeBoxEnumParamFormatter
	//
	Vocabulary json.RawMessage `json:"vocabulary,omitempty"`
}
