package service

type Strategy struct {
	StrategyListItem

	StepTree      StepTree
	Steps         map[string]Step
	EstimatedSize uint `json:"estimatedSize"`
}
