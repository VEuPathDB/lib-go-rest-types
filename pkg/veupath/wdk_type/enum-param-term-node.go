package wdk_type

type EnumParamTermNode struct {
	Data struct {
		Term    string `json:"term"`
		Display string `json:"display"`
	} `json:"data"`

	Children []EnumParamTermNode `json:"children"`
}
