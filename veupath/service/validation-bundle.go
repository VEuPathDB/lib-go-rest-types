package service

type ValidationBundle struct {
	Level   string                  `json:"level"`
	IsValid bool                    `json:"isValid"`
	Errors  *ValidationBundleErrors `json:"errors,omitempty"`
}

type ValidationBundleErrors struct {
	General []string            `json:"general"`
	ByKey   map[string][]string `json:"byKey"`
}
