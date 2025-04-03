package tiktokgo

import (
	tiktokConfig "github.com/wjp-letgo/tiktok-go/config"
	"github.com/wjp-letgo/tiktok-go/oauth"
	oauthentity "github.com/wjp-letgo/tiktok-go/oauth/entity"
	"github.com/wjp-letgo/tiktok-go/shop"
	shopentity "github.com/wjp-letgo/tiktok-go/shop/entity"
	"github.com/wjp-letgo/tiktok-go/authorization"
	authorizationentity "github.com/wjp-letgo/tiktok-go/authorization/entity"
	"github.com/wjp-letgo/tiktok-go/seller"
	sellerentity "github.com/wjp-letgo/tiktok-go/seller/entity"
	"github.com/wjp-letgo/tiktok-go/product"
	productentity "github.com/wjp-letgo/tiktok-go/product/entity"
	"github.com/wjp-letgo/tiktok-go/orders"
	ordersentity "github.com/wjp-letgo/tiktok-go/orders/entity"
	"github.com/wjp-letgo/tiktok-go/logistic"
	logisticentity "github.com/wjp-letgo/tiktok-go/logistic/entity"
	"github.com/wjp-letgo/tiktok-go/fulfillment"
	fulfillmententity "github.com/wjp-letgo/tiktok-go/fulfillment/entity"
)

//TikToker
type TikToker interface {
	//授权
	AuthorizationURL(state string) string
	GetAccessToken(code string) *oauthentity.GetAccessTokenResult
	RefreshToken(refreshToken string) *oauthentity.GetAccessTokenResult
	//店铺
	GetAuthorizedShop() *shopentity.GetAuthorizedShopResult
	//授权店铺相关
	Shops()*authorizationentity.ShopsResult
	//授权店铺相关
	CategoryAssets()*authorizationentity.CategoryAssetsResult

	//获得店铺信息
	SellerShops()*sellerentity.ShopsResult
	Permissions()*sellerentity.PermissionsResult

	//商品相关接口
	Prerequisites()*productentity.PrerequisitesResult
	Categories(locale,keyword,categoryVersion,listingPlatform string)*productentity.CategoriesResult
	RecommendCategory(data *productentity.RecommendCategoryRequest)*productentity.RecommendCategoryResult

	//订单相关接口
	OrdersSearch(pageSize int,sortOrder,pageToken,sortField string,body *ordersentity.OrdersRequest)*ordersentity.OrdersResult
	OrderDetail(ids []string)*ordersentity.OrderDetailResult
	PriceDetail(orderId string)*ordersentity.PriceDetailResult
	AddExternalOrder(body *ordersentity.AddExternalOrderRequest)*ordersentity.AddExternalOrderResult
	
	//物流
	WarehouseList()*logisticentity.WarehouseListResult
	GlobalSellerWarehouse()*logisticentity.GlobalSellerWarehouseResult
	WarehouseDeliveryOptions(warehouseId string)*logisticentity.WarehouseDeliveryOptionsResult
	ShippingProviders(deliveryOptionId string)*logisticentity.ShippingProvidersResult

	//发货
	OrderSplitAttributes(orderIds []string) *fulfillmententity.OrdersSplitAttributesResult
	SplitOrders(orderId string, body *fulfillmententity.SplitOrdersRequest) *fulfillmententity.SplitOrdersResult
	EligibleShippingService(orderId string,body *fulfillmententity.EligibleShippingServiceRequest)*fulfillmententity.EligibleShippingServiceResult
	FirstMileBundle(body *fulfillmententity.FirstmileBundleRequest)*fulfillmententity.FirstmileBundleResult
	CreatePackages(body *fulfillmententity.CreatePackagesRequest)*fulfillmententity.CreatePackagesResult
	SearchPackage(pageSize int,pageToken,sortField,sortOrder string,body *fulfillmententity.SearchPackageRequest)*fulfillmententity.SearchPackageResult
	PackageHandoverTimeSlots(packageId string)*fulfillmententity.PackageHandoverTimeSlotsResult
	ShipPackage(packageId string,body *fulfillmententity.ShipPackageRequest)*fulfillmententity.ShipPackageResult
	PackageShippingDocument(packageId, documentType, documentSize string) *fulfillmententity.PackageShippingDocumentResult
	PackageDetail(packageId string)*fulfillmententity.PackageDetailResult
	UpdateShippingInfo(orderId string,body *fulfillmententity.UpdateShippingInfoRequest)*fulfillmententity.UpdateShippingInfoResult
	BatchShipPackages(body *fulfillmententity.BatchShipPackagesRequest)*fulfillmententity.BatchShipPackagesResult
}

//TikTok
type TikTok struct {
	oauth.OAuth
	shop.Shop
	authorization.Authorization
	seller.Seller
	product.Product
	orders.Orders
	logistic.Logistic
	fulfillment.Fulfillment
}

//NewApi
func NewApi(cfg *tiktokConfig.Config) TikToker {
	return &TikTok{
		oauth.OAuth{Config: cfg},
		shop.Shop{Config: cfg},
		authorization.Authorization{Config: cfg},
		seller.Seller{Config: cfg},
		product.Product{Config: cfg},
		orders.Orders{Config: cfg},
		logistic.Logistic{Config: cfg},
		fulfillment.Fulfillment{Config: cfg},
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
	tiktokList[name] = &TikTok{
		oauth.OAuth{Config: cfg},
		shop.Shop{Config: cfg},
		authorization.Authorization{Config: cfg},
		seller.Seller{Config: cfg},
		product.Product{Config: cfg},
		orders.Orders{Config: cfg},
		logistic.Logistic{Config: cfg},
		fulfillment.Fulfillment{Config: cfg},
	}
}

//GetApi
func GetApi(name string) TikToker {
	return tiktokList[name]
}