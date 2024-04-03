package wdk_type

type UserPreferences struct {
	Global  GlobalUserPreferences  `json:"global"`
	Project ProjectUserPreferences `json:"project"`
}
