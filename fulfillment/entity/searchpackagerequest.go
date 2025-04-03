package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type SearchPackageRequest struct {
	CreateTimeGe  int    `json:"create_time_ge,omitempty"`
	CreateTimeLt  int    `json:"create_time_lt,omitempty"`
	UpdateTimeGe  int    `json:"update_time_ge,omitempty"`
	UpdateTimeLt  int    `json:"update_time_lt,omitempty"`
	PackageStatus string `json:"package_status,omitempty"`
}

func (e *SearchPackageRequest) String() string {
	return lib.ObjectToString(e)
}
