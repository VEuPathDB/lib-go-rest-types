package wdk_type

type SearchConfig struct {
	Parameters       map[string]string `json:"parameters"`
	Filters          []FormatFilter    `json:"filters,omitempty"`
	ColumnFilters    ColumnFilters     `json:"columnFilters"`
	WdkWeight        int               `json:"wdkWeight"`
	LegacyFilterName *string           `json:"legacyFilterName,omitempty"`
}
