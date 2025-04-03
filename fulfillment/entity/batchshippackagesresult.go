package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type BatchShipPackagesResult struct {
	Code      int                   `json:"code"`
	Message   string                `json:"message"`
	Data      BatchShipPackagesData `json:"data"`
	RequestId string                `json:"request_id"`
}

func (e *BatchShipPackagesResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type BatchShipPackagesData struct {
	Errors BatchShipErrors `json:"errors"`
}

func (e *BatchShipPackagesData) String() string {
	return lib.ObjectToString(e)
}

type BatchShipErrors []BatchShipError

func (e *BatchShipErrors) String() string {
	return lib.ObjectToString(e)
}

// @json
type BatchShipError struct {
	Code            int             `json:"code"`
	Message         string          `json:"message"`
	BatchShipDetail BatchShipDetail `json:"detail"`
}

func (e *BatchShipError) String() string {
	return lib.ObjectToString(e)
}

// @json
type BatchShipDetail struct {
	PackageId string `json:"package_id"`
}

func (e *BatchShipDetail) String() string {
	return lib.ObjectToString(e)
}
