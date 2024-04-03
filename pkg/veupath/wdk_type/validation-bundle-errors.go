package wdk_type

type ValidationBundleErrors struct {
	General []string            `json:"general"`
	ByKey   map[string][]string `json:"byKey"`
}
