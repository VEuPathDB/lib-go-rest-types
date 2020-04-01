package recordtypes

type Filter struct {
	DisplayName string `json:"displayName"`
	IsViewOnly  bool   `json:"isViewOnly"`
	Name        string `json:"name"`
}
