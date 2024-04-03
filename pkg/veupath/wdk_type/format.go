package wdk_type

type Format struct {
	IsInReport  bool     `json:"isInReport"`
	DisplayName string   `json:"displayName"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Scopes      []string `json:"scopes"`
	Types       string   `json:"types"`
}
