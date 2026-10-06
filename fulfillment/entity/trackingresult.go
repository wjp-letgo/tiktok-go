package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type TrackingResult struct {
	Code      int          `json:"code"`
	Message   string       `json:"message"`
	Data      TrackingData `json:"data"`
	RequestId string       `json:"request_id"`
}

func (e *TrackingResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type TrackingData struct {
	Tracking Trackings `json:"tracking"`
}

func (e *TrackingData) String() string {
	return lib.ObjectToString(e)
}

type Trackings []Tracking

func (e *Trackings) String() string {
	return lib.ObjectToString(e)
}

// @json
type Tracking struct {
	Description      string `json:"description"`
	UpdateTimeMillis int64  `json:"update_time_millis"`
}

func (e *Tracking) String() string {
	return lib.ObjectToString(e)
}
