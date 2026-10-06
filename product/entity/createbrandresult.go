package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type CreateBrandResult struct {
	Code      int             `json:"code"`
	Message   string          `json:"message"`
	Data      CreateBrandData `json:"data"`
	RequestId string          `json:"request_id"`
}

func (e *CreateBrandResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type CreateBrandData struct {
	Id string `json:"id"`
}

func (e *CreateBrandData) String() string {
	return lib.ObjectToString(e)
}
