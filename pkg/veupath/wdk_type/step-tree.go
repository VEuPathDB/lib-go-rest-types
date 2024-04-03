package wdk_type

type StepTree struct {
	StepId uint64 `json:"stepId"`

	PrimaryInput   *StepTree `json:"primaryInput"`
	SecondaryInput *StepTree `json:"secondaryInput"`
}
