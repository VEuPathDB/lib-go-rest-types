package service

type User struct {
	Id      uint64 `json:"id"`
	IsGuest bool   `json:"isGuest"`

	// Private fields
	Email       *string          `json:"email,omitempty"`
	Properties  UserProperties   `json:"properties,omitempty"`
	Preferences *UserPreferences `json:"preferences,omitempty"`
}

type UserPreferences struct {
	Global  GlobalUserPreferences  `json:"global"`
	Project ProjectUserPreferences `json:"project"`
}

type (
	UserProperties         map[string]string
	GlobalUserPreferences  map[string]string
	ProjectUserPreferences map[string]string
)
