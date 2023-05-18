package product

import (
	"github.com/wjpxxx/letgo/lib"
	tiktokConfig "github.com/wjpxxx/tiktokgo/config"
	productentity "github.com/wjpxxx/tiktokgo/product/entity"
)

type Product struct {
	Config *tiktokConfig.Config
}

//GetBrands 获取品牌
func (p *Product) GetBrands(shopId int64, pageSize, pageNumber int, categoryId, brandSuggest string, onlyAuthorized bool) *productentity.GetBrandsResult {
	var result productentity.GetBrandsResult
	params := lib.InRow{
		"only_authorized": onlyAuthorized,
	}
	if shopId > 0 {
		params["shop_id"] = shopId
	}
	if pageSize > 0 && pageNumber > 0 {
		params["page_size"] = pageSize
		params["page_number"] = pageNumber
	}
	if categoryId != "" {
		params["category_id"] = categoryId
	}
	if brandSuggest != "" {
		params["brand_suggest"] = brandSuggest
	}
	p.Config.HttpGet("/api/products/brands", params, &result)
	return &result
}
