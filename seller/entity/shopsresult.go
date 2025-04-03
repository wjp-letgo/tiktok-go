package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type ShopsResult struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Data      ShopsData `json:"data"`
	RequestId string    `json:"request_id"`
}

func (e *ShopsResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type ShopsData struct {
	Shops Shops `json:"shops"`
}

func (e *ShopsData) String() string {
	return lib.ObjectToString(e)
}

type Shops []Shop

func (e *Shops) String() string {
	return lib.ObjectToString(e)
}

// @json
type Shop struct {
	Id         string `json:"id"`
	Region     string `json:"region"`
}

func (e *Shop) String() string {
	return lib.ObjectToString(e)
}
