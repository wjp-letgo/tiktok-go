package shop

import (
	"github.com/wjpxxx/letgo/lib"
	tiktokConfig "github.com/wjpxxx/tiktokgo/config"
	shopentity "github.com/wjpxxx/tiktokgo/shop/entity"
)

//Shop
type Shop struct {
	Config *tiktokConfig.Config
}

//GetAuthorizedShop 店铺列表
func (s *Shop) GetAuthorizedShop() *shopentity.GetAuthorizedShopResult {
	var result shopentity.GetAuthorizedShopResult
	params := lib.InRow{}
	s.Config.HttpGet("/api/shop/get_authorized_shop", params, &result)
	return &result
}
