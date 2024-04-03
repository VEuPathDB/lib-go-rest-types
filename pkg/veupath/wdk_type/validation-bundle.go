package wdk_type

type ValidationBundle struct {
	Level   string                  `json:"level"`
	IsValid bool                    `json:"isValid"`
	Errors  *ValidationBundleErrors `json:"errors,omitempty"`
}
