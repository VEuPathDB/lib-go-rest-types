package wdk_type

type StrategyListItem struct {
	StrategyId                uint64  `json:"strategyId"`
	Description               string  `json:"description"`
	Name                      string  `json:"name"`
	Author                    string  `json:"author"`
	RootStepId                uint64  `json:"rootStepId"`
	RecordClassName           *string `json:"recordClassName"`
	Signature                 string  `json:"signature"`
	CreatedTime               string  `json:"createdTime"`
	LastViewed                string  `json:"lastViewed"`
	LastModified              string  `json:"lastModified"`
	ReleaseVersion            string  `json:"releaseVersion"`
	IsPublic                  bool    `json:"isPublic"`
	IsSaved                   bool    `json:"isSaved"`
	IsValid                   bool    `json:"isValid"`
	IsDeleted                 bool    `json:"isDeleted"`
	Organization              string  `json:"organization"`
	EstimatedSize             *uint   `json:"estimatedSize"`
	NameOfFirstStep           string  `json:"nameOfFirstStep"`
	LeafAndTransformStepCount uint64  `json:"leafAndTransformStepCount"`
}
