package tiktokgo

import (
	tiktokConfig "github.com/wjpxxx/tiktokgo/config"
	"github.com/wjpxxx/tiktokgo/oauth"
	oauthentity "github.com/wjpxxx/tiktokgo/oauth/entity"
	"github.com/wjpxxx/tiktokgo/shop"
	shopentity "github.com/wjpxxx/tiktokgo/shop/entity"
)

//TikToker
type TikToker interface {
	//授权
	AuthorizationURL(state string) string
	GetAccessToken(code string) *oauthentity.GetAccessTokenResult
	RefreshToken(refreshToken string) *oauthentity.GetAccessTokenResult
	//店铺
	GetAuthorizedShop() *shopentity.GetAuthorizedShopResult
}

//TikTok
type TikTok struct {
	oauth.OAuth
	shop.Shop
}

//NewApi
func NewApi(cfg *tiktokConfig.Config) TikToker {
	return &TikTok{
		oauth.OAuth{Config: cfg},
		shop.Shop{Config: cfg},
	}
}
