package product

import (
	"fmt"

	"github.com/wjp-letgo/letgo/lib"
	productentity "github.com/wjp-letgo/tiktok-go/product/entity"
)

func (a *Product) SearchProducts(pageSize int, pageToken string, body *productentity.SearchProductsRequest) *productentity.SearchProductsResult {
	var result productentity.SearchProductsResult
	params := lib.InRow{
		"page_size": pageSize,
	}
	if pageToken != "" {
		params["page_token"] = pageToken
	}
	err := a.Config.HttpS3("/product/202312/products/search", params, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) GetProduct(productId string) *productentity.ProductDetailResult {
	var result productentity.ProductDetailResult
	params := lib.InRow{}
	err := a.Config.GetS3(fmt.Sprintf("/product/202309/products/%s", productId), params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) CreateProduct(body *productentity.CreateProductRequest) *productentity.CreateProductResult {
	var result productentity.CreateProductResult
	err := a.Config.HttpS3("/product/202309/products", nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) EditProduct(productId string, body *productentity.CreateProductRequest) *productentity.CreateProductResult {
	var result productentity.CreateProductResult
	err := a.Config.PutS3(fmt.Sprintf("/product/202309/products/%s", productId), nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) PartialEditProduct(productId string, body *productentity.CreateProductRequest) *productentity.CreateProductResult {
	var result productentity.CreateProductResult
	err := a.Config.HttpS3(fmt.Sprintf("/product/202309/products/%s/partial_edit", productId), nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) DeleteProducts(body *productentity.ProductIdsRequest) *productentity.ProductIdsResult {
	var result productentity.ProductIdsResult
	err := a.Config.DeleteS3("/product/202309/products", nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) ActivateProducts(body *productentity.ProductIdsRequest) *productentity.ProductIdsResult {
	var result productentity.ProductIdsResult
	err := a.Config.HttpS3("/product/202309/products/activate", nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) DeactivateProducts(body *productentity.ProductIdsRequest) *productentity.ProductIdsResult {
	var result productentity.ProductIdsResult
	err := a.Config.HttpS3("/product/202309/products/deactivate", nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) RecoverProducts(body *productentity.ProductIdsRequest) *productentity.ProductIdsResult {
	var result productentity.ProductIdsResult
	err := a.Config.HttpS3("/product/202309/products/recover", nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) UpdatePrice(productId string, body *productentity.UpdatePriceRequest) *productentity.EmptyProductResult {
	var result productentity.EmptyProductResult
	err := a.Config.HttpS3(fmt.Sprintf("/product/202309/products/%s/prices/update", productId), nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) UpdateInventory(productId string, body *productentity.UpdateInventoryRequest) *productentity.EmptyProductResult {
	var result productentity.EmptyProductResult
	err := a.Config.HttpS3(fmt.Sprintf("/product/202309/products/%s/inventory/update", productId), nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) InventorySearch(body *productentity.InventorySearchRequest) *productentity.InventorySearchResult {
	var result productentity.InventorySearchResult
	err := a.Config.HttpS3("/product/202309/inventory/search", nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) CheckProductListing(body *productentity.CreateProductRequest) *productentity.EmptyProductResult {
	var result productentity.EmptyProductResult
	err := a.Config.HttpS3("/product/202309/products/listing_check", nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) UploadProductImage(filePath, useCase string) *productentity.UploadImageResult {
	var result productentity.UploadImageResult
	values := lib.InRow{
		"@data": filePath,
	}
	if useCase != "" {
		values["use_case"] = useCase
	}
	err := a.Config.MultipartS4("/product/202309/images/upload", nil, values, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) UploadProductFile(filePath, name string) *productentity.UploadFileResult {
	var result productentity.UploadFileResult
	values := lib.InRow{
		"@data": filePath,
	}
	if name != "" {
		values["name"] = name
	}
	err := a.Config.MultipartS4("/product/202309/files/upload", nil, values, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) SearchSizeCharts(pageSize int, pageToken string, body *productentity.SearchSizeChartsRequest) *productentity.SearchSizeChartsResult {
	var result productentity.SearchSizeChartsResult
	params := lib.InRow{
		"page_size": pageSize,
	}
	if pageToken != "" {
		params["page_token"] = pageToken
	}
	err := a.Config.HttpS3("/product/202309/sizecharts/search", params, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) OptimizeImages(body *productentity.OptimizeImagesRequest) *productentity.OptimizeImagesResult {
	var result productentity.OptimizeImagesResult
	err := a.Config.HttpS3("/product/202404/images/optimize", nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) SearchManufacturers(pageSize int, pageToken string, body *productentity.ComplianceSearchRequest) *productentity.ManufacturersResult {
	var result productentity.ManufacturersResult
	params := lib.InRow{
		"page_size": pageSize,
	}
	if pageToken != "" {
		params["page_token"] = pageToken
	}
	err := a.Config.HttpS4("/product/202309/compliance/manufacturers/search", params, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) CreateManufacturer(body *productentity.Manufacturer) *productentity.CreateComplianceResult {
	var result productentity.CreateComplianceResult
	err := a.Config.HttpS4("/product/202309/compliance/manufacturers", nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) PartialEditManufacturer(manufacturerId string, body *productentity.Manufacturer) *productentity.EmptyProductResult {
	var result productentity.EmptyProductResult
	err := a.Config.HttpS4(fmt.Sprintf("/product/202309/compliance/manufacturers/%s/partial_edit", manufacturerId), nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) SearchResponsiblePersons(pageSize int, pageToken string, body *productentity.ComplianceSearchRequest) *productentity.ResponsiblePersonsResult {
	var result productentity.ResponsiblePersonsResult
	params := lib.InRow{
		"page_size": pageSize,
	}
	if pageToken != "" {
		params["page_token"] = pageToken
	}
	err := a.Config.HttpS4("/product/202309/compliance/responsible_persons/search", params, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) CreateResponsiblePerson(body *productentity.ResponsiblePerson) *productentity.CreateComplianceResult {
	var result productentity.CreateComplianceResult
	err := a.Config.HttpS4("/product/202309/compliance/responsible_persons", nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) PartialEditResponsiblePerson(id string, body *productentity.ResponsiblePerson) *productentity.EmptyProductResult {
	var result productentity.EmptyProductResult
	err := a.Config.HttpS4(fmt.Sprintf("/product/202309/compliance/responsible_persons/%s/partial_edit", id), nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) GlobalCategories(locale, keyword, categoryVersion string) *productentity.CategoriesResult {
	var result productentity.CategoriesResult
	params := lib.InRow{}
	if locale != "" {
		params["locale"] = locale
	}
	if keyword != "" {
		params["keyword"] = keyword
	}
	if categoryVersion != "" {
		params["category_version"] = categoryVersion
	}
	err := a.Config.GetS2("/product/202309/global_categories", params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) RecommendGlobalCategory(data *productentity.RecommendCategoryRequest) *productentity.RecommendCategoryResult {
	var result productentity.RecommendCategoryResult
	err := a.Config.HttpS4("/product/202309/global_categories/recommend", nil, data, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) GlobalCategoryRules(categoryId string) *productentity.CategoryRulesResult {
	var result productentity.CategoryRulesResult
	params := lib.InRow{}
	err := a.Config.GetS2(fmt.Sprintf("/product/202309/categories/%s/global_rules", categoryId), params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) GlobalAttributes(categoryId, locale string) *productentity.AttributesResult {
	var result productentity.AttributesResult
	params := lib.InRow{}
	if locale != "" {
		params["locale"] = locale
	}
	err := a.Config.GetS2(fmt.Sprintf("/product/202309/categories/%s/global_attributes", categoryId), params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) CreateGlobalProduct(body *productentity.CreateProductRequest) *productentity.CreateProductResult {
	var result productentity.CreateProductResult
	err := a.Config.HttpS4("/product/202309/global_products", nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) GetGlobalProduct(globalProductId string) *productentity.ProductDetailResult {
	var result productentity.ProductDetailResult
	params := lib.InRow{}
	err := a.Config.GetS2(fmt.Sprintf("/product/202309/global_products/%s", globalProductId), params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) EditGlobalProduct(globalProductId string, body *productentity.CreateProductRequest) *productentity.CreateProductResult {
	var result productentity.CreateProductResult
	err := a.Config.PutS4(fmt.Sprintf("/product/202309/global_products/%s", globalProductId), nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) PublishGlobalProduct(globalProductId string, body *productentity.PublishGlobalProductRequest) *productentity.EmptyProductResult {
	var result productentity.EmptyProductResult
	err := a.Config.HttpS4(fmt.Sprintf("/product/202309/global_products/%s/publish", globalProductId), nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) DeleteGlobalProducts(body *productentity.GlobalIdsRequest) *productentity.EmptyProductResult {
	var result productentity.EmptyProductResult
	err := a.Config.DeleteS4("/product/202309/global_products", nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) SearchGlobalProducts(pageSize int, pageToken string, body *productentity.SearchProductsRequest) *productentity.SearchProductsResult {
	var result productentity.SearchProductsResult
	params := lib.InRow{
		"page_size": pageSize,
	}
	if pageToken != "" {
		params["page_token"] = pageToken
	}
	err := a.Config.HttpS4("/product/202309/global_products/search", params, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Product) UpdateGlobalInventory(globalProductId string, body *productentity.UpdateInventoryRequest) *productentity.EmptyProductResult {
	var result productentity.EmptyProductResult
	err := a.Config.HttpS4(fmt.Sprintf("/product/202309/global_products/%s/inventory/update", globalProductId), nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}
