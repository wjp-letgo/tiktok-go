package tiktokgo

import (
	tiktokConfig "github.com/wjp-letgo/tiktok-go/config"
	"github.com/wjp-letgo/tiktok-go/oauth"
	oauthentity "github.com/wjp-letgo/tiktok-go/oauth/entity"
	"github.com/wjp-letgo/tiktok-go/shop"
	shopentity "github.com/wjp-letgo/tiktok-go/shop/entity"
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

//tiktokList 接口列表
var tiktokList map[string]TikToker

//init
func init() {
	tiktokList = make(map[string]TikToker)
}

//Register
func Register(name string, cfg *tiktokConfig.Config) {
	tiktokList[name] = &TikToker{
		oauth.OAuth{Config: cfg},
		shop.Shop{Config: cfg},
	}
}

//GetApi
func GetApi(name string) TikToker {
	return tiktokList[name]
}