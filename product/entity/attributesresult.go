package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type AttributesResult struct {
	Code      int            `json:"code"`
	Message   string         `json:"message"`
	Data      AttributesData `json:"data"`
	RequestId string         `json:"request_id"`
}

func (e *AttributesResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type AttributesData struct {
	Attributes Attributes `json:"attributes"`
}

func (e *AttributesData) String() string {
	return lib.ObjectToString(e)
}

type Attributes []Attribute

func (e *Attributes) String() string {
	return lib.ObjectToString(e)
}

// @json
type Attribute struct {
	Id                  string                `json:"id"`
	Name                string                `json:"name"`
	Type                string                `json:"type"`
	ValueDataFormat     string                `json:"value_data_format"`
	IsRequired          bool                  `json:"is_required"`
	IsMultipleSelection bool                  `json:"is_multiple_selection"`
	IsCustomizable      bool                  `json:"is_customizable"`
	Values              AttributeValues       `json:"values"`
	RequirementConditions RequirementConditions `json:"requirement_conditions"`
}

func (e *Attribute) String() string {
	return lib.ObjectToString(e)
}

type AttributeValues []AttributeValue

func (e *AttributeValues) String() string {
	return lib.ObjectToString(e)
}

// @json
type AttributeValue struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (e *AttributeValue) String() string {
	return lib.ObjectToString(e)
}
