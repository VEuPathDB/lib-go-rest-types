package wdk_type

type Group struct {
	DisplayType string   `json:"displayType"`
	DisplayName string   `json:"displayName"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	IsVisible   bool     `json:"isVisible"`
	Parameters  []string `json:"parameters"`
}
