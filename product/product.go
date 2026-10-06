package product

import (
	"fmt"

	"github.com/wjp-letgo/letgo/lib"
	tiktokConfig "github.com/wjp-letgo/tiktok-go/config"
	productentity "github.com/wjp-letgo/tiktok-go/product/entity"
)

//Product
type Product struct {
	Config *tiktokConfig.Config
}
//每家商店都需要满足一系列TikTok商店要求，然后才能开始上架产品。在您开始列出产品之前，请使用此API检查您的商店是否满足所有要求。
func (a *Product) Prerequisites()*productentity.PrerequisitesResult{
	var result productentity.PrerequisitesResult
	params := lib.InRow{
	}
	err := a.Config.GetS3("/product/202312/prerequisites", params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

//产品类别更新频繁，因此建议实时调用API，以确保您使用的是最新的类别数据。在本地缓存类别数据可能会导致使用过时的信息，从而导致创建产品时出错。
func (a *Product) Categories(locale,keyword,categoryVersion,listingPlatform string)*productentity.CategoriesResult{
	var result productentity.CategoriesResult
	params := lib.InRow{
	}
	if locale!=""{
		params["locale"]=locale
	}
	if keyword!=""{
		params["keyword"]=keyword
	}
	if categoryVersion!=""{
		params["category_version"]=categoryVersion
	}
	if listingPlatform!=""{
		params["listing_platform"]=listingPlatform
	}
	err := a.Config.GetS3("/product/202309/categories", params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}


//根据候选产品的标题、描述和图片检索其推荐类别。
func (a *Product)RecommendCategory(data *productentity.RecommendCategoryRequest)*productentity.RecommendCategoryResult{
	var result productentity.RecommendCategoryResult
	err := a.Config.HttpS3("/product/202309/categories/recommend",nil, data, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

// 获取指定类目的上架规则（认证、尺码表、COD 等）
func (a *Product) CategoryRules(categoryId, locale, categoryVersion, listingPlatform string) *productentity.CategoryRulesResult {
	var result productentity.CategoryRulesResult
	params := lib.InRow{}
	if locale != "" {
		params["locale"] = locale
	}
	if categoryVersion != "" {
		params["category_version"] = categoryVersion
	}
	if listingPlatform != "" {
		params["listing_platform"] = listingPlatform
	}
	err := a.Config.GetS3(fmt.Sprintf("/product/202309/categories/%s/rules", categoryId), params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

// 获取指定类目的属性列表
func (a *Product) Attributes(categoryId, locale, categoryVersion, listingPlatform string) *productentity.AttributesResult {
	var result productentity.AttributesResult
	params := lib.InRow{}
	if locale != "" {
		params["locale"] = locale
	}
	if categoryVersion != "" {
		params["category_version"] = categoryVersion
	}
	if listingPlatform != "" {
		params["listing_platform"] = listingPlatform
	}
	err := a.Config.GetS3(fmt.Sprintf("/product/202309/categories/%s/attributes", categoryId), params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

// 检索品牌列表
func (a *Product) Brands(pageSize int, pageToken, categoryId, brandName, categoryVersion string, isAuthorized bool) *productentity.BrandsResult {
	var result productentity.BrandsResult
	params := lib.InRow{
		"page_size": pageSize,
	}
	if pageToken != "" {
		params["page_token"] = pageToken
	}
	if categoryId != "" {
		params["category_id"] = categoryId
	}
	if brandName != "" {
		params["brand_name"] = brandName
	}
	if categoryVersion != "" {
		params["category_version"] = categoryVersion
	}
	if isAuthorized {
		params["is_authorized"] = true
	}
	err := a.Config.GetS3("/product/202309/brands", params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

// 创建自定义品牌
func (a *Product) CreateBrand(body *productentity.CreateBrandRequest) *productentity.CreateBrandResult {
	var result productentity.CreateBrandResult
	err := a.Config.HttpS3("/product/202309/brands", nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}