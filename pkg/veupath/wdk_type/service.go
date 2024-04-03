package wdk_type

type Service struct {
	ChangePasswordUrl      string `json:"changePasswordUrl"`
	ReleaseDate            string `json:"releaseDate"`
	DisplayName            string `json:"displayName"`
	UserDatasetStoreStatus struct {
		IsAvailable bool `json:"isAvailable"`
	} `json:"userDatasetStoreStatus"`
	CategoriesOntologyName string `json:"categoriesOntologyName"`
	Description            string `json:"description"`
	StartupTime            uint64 `json:"startupTime"`
	ProjectId              string `json:"projectId"`
	BuildNumber            string `json:"buildNumber"`
	UserProfileProperties  []struct {
		IsRequired  bool   `json:"isRequired"`
		DisplayName string `json:"displayName"`
		Name        string `json:"name"`
		IsPublic    bool   `json:"isPublic"`
	} `json:"userProfileProperties"`
	Authentication interface{} `json:"authentication"`
}
