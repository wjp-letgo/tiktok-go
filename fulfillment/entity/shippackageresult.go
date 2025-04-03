package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type ShipPackageResult struct {
	Code      int             `json:"code"`
	Message   string          `json:"message"`
	Data      ShipPackageData `json:"data"`
	RequestId string          `json:"request_id"`
}

func (e *ShipPackageResult) String() string {
	return lib.ObjectToString(e)
}

type ShipPackageData struct{
}
func (e *ShipPackageData) String() string {
	return lib.ObjectToString(e)
}