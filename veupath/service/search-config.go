package service

import "encoding/json"

type SearchConfig struct {
	Parameters       map[string]string `json:"parameters"`
	Filters          []FormatFilter    `json:"filters,omitempty"`
	ColumnFilters    ColumnFilters     `json:"columnFilters"`
	WdkWeight        int               `json:"wdkWeight"`
	LegacyFilterName *string           `json:"legacyFilterName,omitempty"`
}

type FormatFilter struct {
	Name     string          `json:"name"`
	Value    json.RawMessage `json:"value"`
	Disabled bool            `json:"disabled"`
}

type ColumnFilters map[string]json.RawMessage
