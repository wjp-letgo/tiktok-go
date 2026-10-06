package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type CategoryRulesResult struct {
	Code      int               `json:"code"`
	Message   string            `json:"message"`
	Data      CategoryRulesData `json:"data"`
	RequestId string            `json:"request_id"`
}

func (e *CategoryRulesResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type CategoryRulesData struct {
	ProductCertifications ProductCertifications `json:"product_certifications"`
	SizeChart             SizeChartRule         `json:"size_chart"`
	Cod                   CodRule               `json:"cod"`
	PackageDimension      PackageDimensionRule  `json:"package_dimension"`
	Epr                   EprRule               `json:"epr"`
}

func (e *CategoryRulesData) String() string {
	return lib.ObjectToString(e)
}

type ProductCertifications []ProductCertification

func (e *ProductCertifications) String() string {
	return lib.ObjectToString(e)
}

// @json
type ProductCertification struct {
	Id                    string                 `json:"id"`
	Name                  string                 `json:"name"`
	SampleImageUrl        string                 `json:"sample_image_url"`
	IsRequired            bool                   `json:"is_required"`
	RequirementConditions RequirementConditions  `json:"requirement_conditions"`
}

func (e *ProductCertification) String() string {
	return lib.ObjectToString(e)
}

type RequirementConditions []RequirementCondition

func (e *RequirementConditions) String() string {
	return lib.ObjectToString(e)
}

// @json
type RequirementCondition struct {
	ConditionType    string `json:"condition_type"`
	AttributeId      string `json:"attribute_id"`
	AttributeValueId string `json:"attribute_value_id"`
}

func (e *RequirementCondition) String() string {
	return lib.ObjectToString(e)
}

// @json
type SizeChartRule struct {
	IsSupported bool `json:"is_supported"`
	IsRequired  bool `json:"is_required"`
}

func (e *SizeChartRule) String() string {
	return lib.ObjectToString(e)
}

// @json
type CodRule struct {
	IsSupported bool `json:"is_supported"`
}

func (e *CodRule) String() string {
	return lib.ObjectToString(e)
}

// @json
type PackageDimensionRule struct {
	IsRequired bool `json:"is_required"`
}

func (e *PackageDimensionRule) String() string {
	return lib.ObjectToString(e)
}

// @json
type EprRule struct {
	IsRequired bool `json:"is_required"`
}

func (e *EprRule) String() string {
	return lib.ObjectToString(e)
}
