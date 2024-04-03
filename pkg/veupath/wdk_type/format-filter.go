package wdk_type

import "encoding/json"

type FormatFilter struct {
	Name     string          `json:"name"`
	Value    json.RawMessage `json:"value"`
	Disabled bool            `json:"disabled"`
}
