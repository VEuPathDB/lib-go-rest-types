package wdk_type

type Attribute struct {
	Align          string     `json:"align"`
	DisplayName    string     `json:"displayName"`
	Formats        []Format   `json:"formats"`
	Help           string     `json:"help"`
	IsDisplayable  bool       `json:"isDisplayable"`
	IsInReport     bool       `json:"isInReport"`
	IsRemovable    bool       `json:"isRemovable"`
	IsSortable     bool       `json:"isSortable"`
	Name           string     `json:"name"`
	Properties     Properties `json:"properties"`
	TruncateTo     uint       `json:"truncateTo"`
	Type           string     `json:"type"`
	Tools          Tools      `json:"tools"`
	ColumnDataType string     `json:"columnDataType"`
}
