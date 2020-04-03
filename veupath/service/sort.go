package service

import (
	"fmt"
	"strings"
)

type SortDirection string

func (s *SortDirection) UnmarshalJSON(b []byte) error {
	switch strings.ToUpper(strings.Trim(string(b), `"`)) {
	case "ASC":
		*s = SortDirectionAscending
	case "DESC":
		*s = SortDirectionDescending
	default:
		//noinspection GoErrorStringFormat
		return fmt.Errorf("Unrecognized sort direction: %s", string(b))
	}
	return nil
}

const (
	SortDirectionAscending  = "ASC"
	SortDirectionDescending = "DESC"
)


type SortSpec struct {
	AttributeName string        `json:"attributeName"`
	Direction     SortDirection `json:"direction"`
}


type ClientSortSpec struct {
	ItemName  string        `json:"itemName"`
	Direction SortDirection `json:"direction"`
}

