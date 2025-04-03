package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type PermissionsResult struct {
	Code      int             `json:"code"`
	Message   string          `json:"message"`
	Data      PermissionsData `json:"data"`
	RequestId string          `json:"request_id"`
}

func (e *PermissionsResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type PermissionsData struct {
	Permissions []string `json:"permissions"`
}

func (e *PermissionsData) String() string {
	return lib.ObjectToString(e)
}
