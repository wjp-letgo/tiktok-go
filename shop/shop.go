package shop

import (
	"github.com/wjp-letgo/letgo/lib"
	tiktokConfig "github.com/wjp-letgo/tiktok-go/config"
	shopentity "github.com/wjp-letgo/tiktok-go/shop/entity"
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
