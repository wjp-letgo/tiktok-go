package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type CreateProductRequest struct {
	SaveMode           string            `json:"save_mode,omitempty"`
	Title              string            `json:"title,omitempty"`
	Description        string            `json:"description,omitempty"`
	CategoryId         string            `json:"category_id,omitempty"`
	BrandId            string            `json:"brand_id,omitempty"`
	MainImages         ProductImages     `json:"main_images,omitempty"`
	Skus               ProductSkus       `json:"skus,omitempty"`
	PackageWeight      PackageWeight     `json:"package_weight,omitempty"`
	PackageDimensions  PackageDimensions `json:"package_dimensions,omitempty"`
	IsCodAllowed       bool              `json:"is_cod_allowed,omitempty"`
	CategoryVersion    string            `json:"category_version,omitempty"`
	ListingPlatforms   []string          `json:"listing_platforms,omitempty"`
	ProductAttributes  ProductAttrList   `json:"product_attributes,omitempty"`
	SizeChart          SizeChart         `json:"size_chart,omitempty"`
}

func (e *CreateProductRequest) String() string {
	return lib.ObjectToString(e)
}

type ProductImages []ProductImage

func (e *ProductImages) String() string {
	return lib.ObjectToString(e)
}

// @json
type ProductImage struct {
	Uri string `json:"uri,omitempty"`
	Url string `json:"url,omitempty"`
}

func (e *ProductImage) String() string {
	return lib.ObjectToString(e)
}

type ProductSkus []ProductSku

func (e *ProductSkus) String() string {
	return lib.ObjectToString(e)
}

// @json
type ProductSku struct {
	Id               string             `json:"id,omitempty"`
	SellerSku        string             `json:"seller_sku,omitempty"`
	Price            ProductPrice       `json:"price,omitempty"`
	Inventory        SkuInventories     `json:"inventory,omitempty"`
	SalesAttributes  SalesAttributes    `json:"sales_attributes,omitempty"`
	IdentifierCode   IdentifierCode     `json:"identifier_code,omitempty"`
}

func (e *ProductSku) String() string {
	return lib.ObjectToString(e)
}

type SalesAttributes []SalesAttribute

func (e *SalesAttributes) String() string {
	return lib.ObjectToString(e)
}

// @json
type SalesAttribute struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	ValueId   string `json:"value_id,omitempty"`
	ValueName string `json:"value_name,omitempty"`
	SkuImg    ProductImage `json:"sku_img,omitempty"`
}

func (e *SalesAttribute) String() string {
	return lib.ObjectToString(e)
}

// @json
type IdentifierCode struct {
	Code string `json:"code,omitempty"`
	Type string `json:"type,omitempty"`
}

func (e *IdentifierCode) String() string {
	return lib.ObjectToString(e)
}

// @json
type PackageWeight struct {
	Value string `json:"value,omitempty"`
	Unit  string `json:"unit,omitempty"`
}

func (e *PackageWeight) String() string {
	return lib.ObjectToString(e)
}

// @json
type PackageDimensions struct {
	Length string `json:"length,omitempty"`
	Width  string `json:"width,omitempty"`
	Height string `json:"height,omitempty"`
	Unit   string `json:"unit,omitempty"`
}

func (e *PackageDimensions) String() string {
	return lib.ObjectToString(e)
}

type ProductAttrList []ProductAttr

func (e *ProductAttrList) String() string {
	return lib.ObjectToString(e)
}

// @json
type ProductAttr struct {
	Id     string            `json:"id,omitempty"`
	Values ProductAttrValues `json:"values,omitempty"`
}

func (e *ProductAttr) String() string {
	return lib.ObjectToString(e)
}

type ProductAttrValues []ProductAttrValue

func (e *ProductAttrValues) String() string {
	return lib.ObjectToString(e)
}

// @json
type ProductAttrValue struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (e *ProductAttrValue) String() string {
	return lib.ObjectToString(e)
}

// @json
type SizeChart struct {
	Image ProductImage `json:"image,omitempty"`
	Template ProductImage `json:"template,omitempty"`
}

func (e *SizeChart) String() string {
	return lib.ObjectToString(e)
}

// @json
type CreateProductResult struct {
	Code      int               `json:"code"`
	Message   string            `json:"message"`
	Data      CreateProductData `json:"data"`
	RequestId string            `json:"request_id"`
}

func (e *CreateProductResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type CreateProductData struct {
	ProductId string      `json:"product_id"`
	Skus      ProductSkus `json:"skus"`
}

func (e *CreateProductData) String() string {
	return lib.ObjectToString(e)
}

// @json
type ProductDetailResult struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      ProductInfo `json:"data"`
	RequestId string      `json:"request_id"`
}

func (e *ProductDetailResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type ProductInfo struct {
	Id                string            `json:"id"`
	Title             string            `json:"title"`
	Status            string            `json:"status"`
	Description       string            `json:"description"`
	CategoryChains    RecommendCategories `json:"category_chains"`
	Brand             Brand             `json:"brand"`
	MainImages        ProductImages     `json:"main_images"`
	Skus              ProductSkus       `json:"skus"`
	PackageWeight     PackageWeight     `json:"package_weight"`
	PackageDimensions PackageDimensions `json:"package_dimensions"`
	CreateTime        int               `json:"create_time"`
	UpdateTime        int               `json:"update_time"`
	IsCodAllowed      bool              `json:"is_cod_allowed"`
}

func (e *ProductInfo) String() string {
	return lib.ObjectToString(e)
}

// @json
type UploadImageResult struct {
	Code      int             `json:"code"`
	Message   string          `json:"message"`
	Data      UploadImageData `json:"data"`
	RequestId string          `json:"request_id"`
}

func (e *UploadImageResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type UploadImageData struct {
	Uri    string `json:"uri"`
	Url    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
	UseCase string `json:"use_case"`
}

func (e *UploadImageData) String() string {
	return lib.ObjectToString(e)
}

// @json
type UploadFileResult struct {
	Code      int            `json:"code"`
	Message   string         `json:"message"`
	Data      UploadFileData `json:"data"`
	RequestId string         `json:"request_id"`
}

func (e *UploadFileResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type UploadFileData struct {
	Id   string `json:"id"`
	Url  string `json:"url"`
	Name string `json:"name"`
}

func (e *UploadFileData) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchSizeChartsRequest struct {
	Ids []string `json:"ids,omitempty"`
}

func (e *SearchSizeChartsRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchSizeChartsResult struct {
	Code      int                  `json:"code"`
	Message   string               `json:"message"`
	Data      SearchSizeChartsData `json:"data"`
	RequestId string               `json:"request_id"`
}

func (e *SearchSizeChartsResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchSizeChartsData struct {
	NextPageToken string     `json:"next_page_token"`
	SizeCharts    SizeCharts `json:"size_charts"`
}

func (e *SearchSizeChartsData) String() string {
	return lib.ObjectToString(e)
}

type SizeCharts []SizeChartItem

func (e *SizeCharts) String() string {
	return lib.ObjectToString(e)
}

// @json
type SizeChartItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (e *SizeChartItem) String() string {
	return lib.ObjectToString(e)
}

// @json
type OptimizeImagesRequest struct {
	Images ProductImages `json:"images,omitempty"`
}

func (e *OptimizeImagesRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type OptimizeImagesResult struct {
	Code      int                `json:"code"`
	Message   string             `json:"message"`
	Data      OptimizeImagesData `json:"data"`
	RequestId string             `json:"request_id"`
}

func (e *OptimizeImagesResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type OptimizeImagesData struct {
	Images ProductImages `json:"images"`
}

func (e *OptimizeImagesData) String() string {
	return lib.ObjectToString(e)
}

// @json
type ComplianceSearchRequest struct {
	Keyword string `json:"keyword,omitempty"`
}

func (e *ComplianceSearchRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type ManufacturersResult struct {
	Code      int               `json:"code"`
	Message   string            `json:"message"`
	Data      ManufacturersData `json:"data"`
	RequestId string            `json:"request_id"`
}

func (e *ManufacturersResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type ManufacturersData struct {
	NextPageToken  string        `json:"next_page_token"`
	Manufacturers  Manufacturers `json:"manufacturers"`
}

func (e *ManufacturersData) String() string {
	return lib.ObjectToString(e)
}

type Manufacturers []Manufacturer

func (e *Manufacturers) String() string {
	return lib.ObjectToString(e)
}

// @json
type Manufacturer struct {
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
	Email   string `json:"email,omitempty"`
	Phone   string `json:"phone_number,omitempty"`
}

func (e *Manufacturer) String() string {
	return lib.ObjectToString(e)
}

// @json
type ResponsiblePersonsResult struct {
	Code      int                     `json:"code"`
	Message   string                  `json:"message"`
	Data      ResponsiblePersonsData  `json:"data"`
	RequestId string                  `json:"request_id"`
}

func (e *ResponsiblePersonsResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type ResponsiblePersonsData struct {
	NextPageToken      string              `json:"next_page_token"`
	ResponsiblePersons ResponsiblePersons  `json:"responsible_persons"`
}

func (e *ResponsiblePersonsData) String() string {
	return lib.ObjectToString(e)
}

type ResponsiblePersons []ResponsiblePerson

func (e *ResponsiblePersons) String() string {
	return lib.ObjectToString(e)
}

// @json
type ResponsiblePerson struct {
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Email   string `json:"email,omitempty"`
	Phone   string `json:"phone_number,omitempty"`
	Address string `json:"address,omitempty"`
}

func (e *ResponsiblePerson) String() string {
	return lib.ObjectToString(e)
}

// @json
type CreateComplianceResult struct {
	Code      int                 `json:"code"`
	Message   string              `json:"message"`
	Data      CreateComplianceData `json:"data"`
	RequestId string              `json:"request_id"`
}

func (e *CreateComplianceResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type CreateComplianceData struct {
	Id string `json:"id"`
}

func (e *CreateComplianceData) String() string {
	return lib.ObjectToString(e)
}

// @json
type GlobalIdsRequest struct {
	GlobalProductIds []string `json:"global_product_ids,omitempty"`
}

func (e *GlobalIdsRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type PublishGlobalProductRequest struct {
	PublishTarget PublishTargets `json:"publish_target,omitempty"`
}

func (e *PublishGlobalProductRequest) String() string {
	return lib.ObjectToString(e)
}

type PublishTargets []PublishTarget

func (e *PublishTargets) String() string {
	return lib.ObjectToString(e)
}

// @json
type PublishTarget struct {
	Region string `json:"region,omitempty"`
	ShopId string `json:"shop_id,omitempty"`
}

func (e *PublishTarget) String() string {
	return lib.ObjectToString(e)
}
