package wdk_type

type RecordType struct {
	NativeDisplayName            string      `json:"nativeDisplayName"`
	Searches                     []Search    `json:"searches"`
	UrlSegment                   string      `json:"urlSegment"`
	UseBasket                    bool        `json:"useBasket"`
	NativeShortDisplayName       string      `json:"nativeShortDisplayName"`
	Formats                      []Format    `json:"formats"`
	HasAllRecordsQuery           bool        `json:"hasAllRecordsQuery"`
	DisplayName                  string      `json:"displayName"`
	NativeShortDisplayNamePlural string      `json:"nativeShortDisplayNamePlural"`
	FullName                     string      `json:"fullName"`
	ShortDisplayName             string      `json:"shortDisplayName"`
	ShortDisplayNamePlural       string      `json:"shortDisplayNamePlural"`
	Tables                       []Table     `json:"tables"`
	NativeDisplayNamePlural      string      `json:"nativeDisplayNamePlural"`
	DisplayNamePlural            string      `json:"displayNamePlural"`
	RecordIdAttributeName        string      `json:"recordIdAttributeName"`
	Attributes                   []Attribute `json:"attributes"`
	PrimaryKeyColumnRefs         []string    `json:"primaryKeyColumnRefs"`
	Properties                   Properties  `json:"properties"`
}
