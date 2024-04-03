package wdk_type

import (
	"fmt"
	"strings"
)

const (
	SortDirectionAscending  SortDirection = "ASC"
	SortDirectionDescending SortDirection = "DESC"
)

type SortDirection string

func (s *SortDirection) UnmarshalJSON(b []byte) error {
	switch strings.ToUpper(strings.Trim(string(b), `"`)) {
	case "ASC":
		*s = SortDirectionAscending
	case "DESC":
		*s = SortDirectionDescending
	default:
		return fmt.Errorf("unrecognized sort direction: %s", string(b))
	}

	return nil
}
