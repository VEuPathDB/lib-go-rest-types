package wdk_type

type User struct {
	Id      uint64 `json:"id"`
	IsGuest bool   `json:"isGuest"`

	// Private fields
	Email       *string          `json:"email,omitempty"`
	Properties  UserProperties   `json:"properties,omitempty"`
	Preferences *UserPreferences `json:"preferences,omitempty"`
}
