package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type PackageShippingDocumentResult struct {
	Code      int                         `json:"code"`
	Message   string                      `json:"message"`
	Data      PackageShippingDocumentData `json:"data"`
	RequestId string                      `json:"request_id"`
}

func (e *PackageShippingDocumentResult) String() string {
	return lib.ObjectToString(e)
}
// @json
type PackageShippingDocumentData struct {
	DocUrl         string `json:"doc_url"`
	TrackingNumber string `json:"tracking_number"`
}

func (e *PackageShippingDocumentData) String() string {
	return lib.ObjectToString(e)
}
