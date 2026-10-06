package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type CreateBrandRequest struct {
	Name string `json:"name,omitempty"`
}

func (e *CreateBrandRequest) String() string {
	return lib.ObjectToString(e)
}
