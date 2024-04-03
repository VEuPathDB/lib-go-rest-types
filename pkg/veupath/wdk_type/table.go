package wdk_type

type Table struct {
	IsInReport     bool             `json:"isInReport"`
	DisplayName    string           `json:"displayName"`
	Name           string           `json:"name"`
	Description    string           `json:"description"`
	Attributes     []Attribute      `json:"attributes"`
	IsDisplayable  bool             `json:"isDisplayable"`
	Properties     Properties       `json:"properties"`
	ClientSortSpec []ClientSortSpec `json:"clientSortSpec"`
}
