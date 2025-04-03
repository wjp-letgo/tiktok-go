package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type OrdersSplitAttributesResult struct {
	Code      int                       `json:"code"`
	Message   string                    `json:"message"`
	Data      OrdersSplitAttributesData `json:"data"`
	RequestId string                    `json:"request_id"`
}

func (e *OrdersSplitAttributesResult) String() string {
	return lib.ObjectToString(e)
}

type OrdersSplitAttributesData struct {
	SplitAttributes SplitAttributes `json:"split_attributes"`
}

func (e *OrdersSplitAttributesData) String() string {
	return lib.ObjectToString(e)
}

type SplitAttributes []SplitAttribute

func (e *SplitAttributes) String() string {
	return lib.ObjectToString(e)
}

// @json
type SplitAttribute struct {
	OrderId  string `json:"order_id"`
	CanSplit bool   `json:"can_split"`
	Reason   string `json:"reason"`
}

func (e *SplitAttribute) String() string {
	return lib.ObjectToString(e)
}
