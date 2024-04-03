package wdk_type

type SortSpec struct {
	AttributeName string        `json:"attributeName"`
	Direction     SortDirection `json:"direction"`
}
