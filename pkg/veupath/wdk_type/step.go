package wdk_type

import "time"

type Step struct {
	Id                      uint64             `json:"id"`
	DisplayName             string             `json:"displayName"`
	ShortDisplayName        string             `json:"shortDisplayName"`
	CustomName              string             `json:"customName"`
	BaseCustomName          string             `json:"baseCustomName"`
	Expanded                bool               `json:"expanded"`
	ExpandedName            string             `json:"expandedName"`
	IsFiltered              bool               `json:"isFiltered"`
	Description             string             `json:"description"`
	OwnerId                 uint64             `json:"ownerId"`
	StrategyId              *uint64            `json:"strategyId"`
	HasCompleteStepAnalyses bool               `json:"hasCompleteStepAnalyses"`
	RecordClassName         *string            `json:"recordClassName"`
	SearchName              string             `json:"searchName"`
	SearchConfig            SearchConfig       `json:"searchConfig"`
	Validation              ValidationBundle   `json:"validation"`
	CreatedTime             time.Time          `json:"createdTime"`
	LastRunTime             time.Time          `json:"lastRunTime"`
	DisplayPreferences      DisplayPreferences `json:"displayPreferences"`
	EstimatedSize           *uint              `json:"estimatedSize,omitempty"`
}
