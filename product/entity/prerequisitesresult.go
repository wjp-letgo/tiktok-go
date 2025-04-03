package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type PrerequisitesResult struct {
	Code      int               `json:"code"`
	Message   string            `json:"message"`
	Data      PrerequisitesData `json:"data"`
	RequestId string            `json:"request_id"`
}

func (e *PrerequisitesResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type PrerequisitesData struct {
	CheckResults CheckResults `json:"check_results"`
}

func (e *PrerequisitesData) String() string {
	return lib.ObjectToString(e)
}

// @json
type CheckResults []CheckResult

func (e *CheckResults) String() string {
	return lib.ObjectToString(e)
}

// @json
type CheckResult struct {
	CheckItem   string   `json:"check_item"`
	IsFailed    bool     `json:"is_failed"`
	FailReasons []string `json:"fail_reasons"`
}

func (e *CheckResult) String() string {
	return lib.ObjectToString(e)
}
