package tiktokgo

import (
	"github.com/wjp-letgo/tiktok-go/authorization"
	authorizationentity "github.com/wjp-letgo/tiktok-go/authorization/entity"
	tiktokConfig "github.com/wjp-letgo/tiktok-go/config"
	"github.com/wjp-letgo/tiktok-go/customerservice"
	customerserviceentity "github.com/wjp-letgo/tiktok-go/customerservice/entity"
	"github.com/wjp-letgo/tiktok-go/event"
	evententity "github.com/wjp-letgo/tiktok-go/event/entity"
	"github.com/wjp-letgo/tiktok-go/fbt"
	fbtentity "github.com/wjp-letgo/tiktok-go/fbt/entity"
	"github.com/wjp-letgo/tiktok-go/finance"
	financeentity "github.com/wjp-letgo/tiktok-go/finance/entity"
	"github.com/wjp-letgo/tiktok-go/fulfillment"
	fulfillmententity "github.com/wjp-letgo/tiktok-go/fulfillment/entity"
	"github.com/wjp-letgo/tiktok-go/logistic"
	logisticentity "github.com/wjp-letgo/tiktok-go/logistic/entity"
	"github.com/wjp-letgo/tiktok-go/oauth"
	oauthentity "github.com/wjp-letgo/tiktok-go/oauth/entity"
	"github.com/wjp-letgo/tiktok-go/orders"
	ordersentity "github.com/wjp-letgo/tiktok-go/orders/entity"
	"github.com/wjp-letgo/tiktok-go/product"
	productentity "github.com/wjp-letgo/tiktok-go/product/entity"
	"github.com/wjp-letgo/tiktok-go/promotion"
	promotionentity "github.com/wjp-letgo/tiktok-go/promotion/entity"
	"github.com/wjp-letgo/tiktok-go/returnrefund"
	returnrefundentity "github.com/wjp-letgo/tiktok-go/returnrefund/entity"
	"github.com/wjp-letgo/tiktok-go/seller"
	sellerentity "github.com/wjp-letgo/tiktok-go/seller/entity"
	"github.com/wjp-letgo/tiktok-go/shop"
	shopentity "github.com/wjp-letgo/tiktok-go/shop/entity"
	"github.com/wjp-letgo/tiktok-go/supplychain"
	supplychainentity "github.com/wjp-letgo/tiktok-go/supplychain/entity"
)

// TikToker
type TikToker interface {
	AuthorizationURL(state string) string
	GetAccessToken(code string) *oauthentity.GetAccessTokenResult
	RefreshToken(refreshToken string) *oauthentity.GetAccessTokenResult
	GetAuthorizedShop() *shopentity.GetAuthorizedShopResult
	Shops() *authorizationentity.ShopsResult
	CategoryAssets() *authorizationentity.CategoryAssetsResult
	SellerShops() *sellerentity.ShopsResult
	Permissions() *sellerentity.PermissionsResult

	GetShopWebhooks() *evententity.WebhooksResult
	UpdateShopWebhook(body *evententity.UpdateWebhookRequest) *evententity.UpdateWebhookResult
	DeleteShopWebhook(body *evententity.DeleteWebhookRequest) *evententity.DeleteWebhookResult

	Prerequisites() *productentity.PrerequisitesResult
	Categories(locale, keyword, categoryVersion, listingPlatform string) *productentity.CategoriesResult
	RecommendCategory(data *productentity.RecommendCategoryRequest) *productentity.RecommendCategoryResult
	CategoryRules(categoryId, locale, categoryVersion, listingPlatform string) *productentity.CategoryRulesResult
	Attributes(categoryId, locale, categoryVersion, listingPlatform string) *productentity.AttributesResult
	Brands(pageSize int, pageToken, categoryId, brandName, categoryVersion string, isAuthorized bool) *productentity.BrandsResult
	CreateBrand(body *productentity.CreateBrandRequest) *productentity.CreateBrandResult
	SearchProducts(pageSize int, pageToken string, body *productentity.SearchProductsRequest) *productentity.SearchProductsResult
	GetProduct(productId string) *productentity.ProductDetailResult
	CreateProduct(body *productentity.CreateProductRequest) *productentity.CreateProductResult
	EditProduct(productId string, body *productentity.CreateProductRequest) *productentity.CreateProductResult
	PartialEditProduct(productId string, body *productentity.CreateProductRequest) *productentity.CreateProductResult
	DeleteProducts(body *productentity.ProductIdsRequest) *productentity.ProductIdsResult
	ActivateProducts(body *productentity.ProductIdsRequest) *productentity.ProductIdsResult
	DeactivateProducts(body *productentity.ProductIdsRequest) *productentity.ProductIdsResult
	RecoverProducts(body *productentity.ProductIdsRequest) *productentity.ProductIdsResult
	UpdatePrice(productId string, body *productentity.UpdatePriceRequest) *productentity.EmptyProductResult
	UpdateInventory(productId string, body *productentity.UpdateInventoryRequest) *productentity.EmptyProductResult
	InventorySearch(body *productentity.InventorySearchRequest) *productentity.InventorySearchResult
	CheckProductListing(body *productentity.CreateProductRequest) *productentity.EmptyProductResult
	UploadProductImage(filePath, useCase string) *productentity.UploadImageResult
	UploadProductFile(filePath, name string) *productentity.UploadFileResult
	SearchSizeCharts(pageSize int, pageToken string, body *productentity.SearchSizeChartsRequest) *productentity.SearchSizeChartsResult
	OptimizeImages(body *productentity.OptimizeImagesRequest) *productentity.OptimizeImagesResult
	SearchManufacturers(pageSize int, pageToken string, body *productentity.ComplianceSearchRequest) *productentity.ManufacturersResult
	CreateManufacturer(body *productentity.Manufacturer) *productentity.CreateComplianceResult
	PartialEditManufacturer(manufacturerId string, body *productentity.Manufacturer) *productentity.EmptyProductResult
	SearchResponsiblePersons(pageSize int, pageToken string, body *productentity.ComplianceSearchRequest) *productentity.ResponsiblePersonsResult
	CreateResponsiblePerson(body *productentity.ResponsiblePerson) *productentity.CreateComplianceResult
	PartialEditResponsiblePerson(id string, body *productentity.ResponsiblePerson) *productentity.EmptyProductResult
	GlobalCategories(locale, keyword, categoryVersion string) *productentity.CategoriesResult
	RecommendGlobalCategory(data *productentity.RecommendCategoryRequest) *productentity.RecommendCategoryResult
	GlobalCategoryRules(categoryId string) *productentity.CategoryRulesResult
	GlobalAttributes(categoryId, locale string) *productentity.AttributesResult
	CreateGlobalProduct(body *productentity.CreateProductRequest) *productentity.CreateProductResult
	GetGlobalProduct(globalProductId string) *productentity.ProductDetailResult
	EditGlobalProduct(globalProductId string, body *productentity.CreateProductRequest) *productentity.CreateProductResult
	PublishGlobalProduct(globalProductId string, body *productentity.PublishGlobalProductRequest) *productentity.EmptyProductResult
	DeleteGlobalProducts(body *productentity.GlobalIdsRequest) *productentity.EmptyProductResult
	SearchGlobalProducts(pageSize int, pageToken string, body *productentity.SearchProductsRequest) *productentity.SearchProductsResult
	UpdateGlobalInventory(globalProductId string, body *productentity.UpdateInventoryRequest) *productentity.EmptyProductResult

	CreateActivity(body *promotionentity.CreateActivityRequest) *promotionentity.CreateActivityResult
	UpdateActivity(activityId string, body *promotionentity.UpdateActivityRequest) *promotionentity.UpdateActivityResult
	DeactivateActivity(activityId string) *promotionentity.DeactivateActivityResult
	GetActivity(activityId string) *promotionentity.ActivityDetailResult
	SearchActivities(pageSize int, pageToken string, body *promotionentity.SearchActivitiesRequest) *promotionentity.SearchActivitiesResult
	UpdateActivityProduct(activityId string, body *promotionentity.UpdateActivityProductRequest) *promotionentity.UpdateActivityProductResult
	RemoveActivityProduct(activityId string, body *promotionentity.RemoveActivityProductRequest) *promotionentity.RemoveActivityProductResult
	SearchCoupons(pageSize int, pageToken string, body *promotionentity.SearchCouponsRequest) *promotionentity.SearchCouponsResult
	GetCoupon(couponId string) *promotionentity.CouponDetailResult

	OrdersSearch(pageSize int, pageToken, sortOrder, sortField string, body *ordersentity.OrdersRequest) *ordersentity.OrdersResult
	OrderDetail(ids []string) *ordersentity.OrderDetailResult
	PriceDetail(orderId string) *ordersentity.PriceDetailResult
	AddExternalOrder(body *ordersentity.AddExternalOrderRequest) *ordersentity.AddExternalOrderResult
	ExternalOrders(orderId, platform string) *ordersentity.ExternalOrdersResult
	SearchExternalOrder(platform, externalOrderId string) *ordersentity.SearchExternalOrderResult

	WarehouseList() *logisticentity.WarehouseListResult
	GlobalSellerWarehouse() *logisticentity.GlobalSellerWarehouseResult
	WarehouseDeliveryOptions(warehouseId string) *logisticentity.WarehouseDeliveryOptionsResult
	ShippingProviders(deliveryOptionId string) *logisticentity.ShippingProvidersResult

	OrderSplitAttributes(orderIds []string) *fulfillmententity.OrdersSplitAttributesResult
	SplitOrders(orderId string, body *fulfillmententity.SplitOrdersRequest) *fulfillmententity.SplitOrdersResult
	EligibleShippingService(orderId string, body *fulfillmententity.EligibleShippingServiceRequest) *fulfillmententity.EligibleShippingServiceResult
	FirstMileBundle(body *fulfillmententity.FirstmileBundleRequest) *fulfillmententity.FirstmileBundleResult
	CreatePackages(body *fulfillmententity.CreatePackagesRequest) *fulfillmententity.CreatePackagesResult
	SearchPackage(pageSize int, pageToken, sortField, sortOrder string, body *fulfillmententity.SearchPackageRequest) *fulfillmententity.SearchPackageResult
	PackageHandoverTimeSlots(packageId string) *fulfillmententity.PackageHandoverTimeSlotsResult
	ShipPackage(packageId string, body *fulfillmententity.ShipPackageRequest) *fulfillmententity.ShipPackageResult
	PackageShippingDocument(packageId, documentType, documentSize string) *fulfillmententity.PackageShippingDocumentResult
	PackageDetail(packageId string) *fulfillmententity.PackageDetailResult
	UpdateShippingInfo(orderId string, body *fulfillmententity.UpdateShippingInfoRequest) *fulfillmententity.UpdateShippingInfoResult
	BatchShipPackages(body *fulfillmententity.BatchShipPackagesRequest) *fulfillmententity.BatchShipPackagesResult
	SearchCombinablePackages(pageSize int, pageToken string) *fulfillmententity.SearchCombinablePackageResult
	CombinePackage(body *fulfillmententity.CombinePackageRequest) *fulfillmententity.CombinePackageResult
	UncombinePackages(packageId string, body *fulfillmententity.UncombinePackagesRequest) *fulfillmententity.UncombinePackagesResult
	MarkPackageAsShipped(orderId string, body *fulfillmententity.MarkPackageAsShippedRequest) *fulfillmententity.MarkPackageAsShippedResult
	GetTracking(orderId string) *fulfillmententity.TrackingResult
	UpdatePackageShippingInfo(packageId string, body *fulfillmententity.UpdatePackageShippingInfoRequest) *fulfillmententity.UpdatePackageShippingInfoResult
	UpdatePackageDeliveryStatus(body *fulfillmententity.UpdatePackageDeliveryStatusRequest) *fulfillmententity.UpdatePackageDeliveryStatusResult

	MerchantOnboardedRegions() *fbtentity.OnboardedRegionsResult
	FbtWarehouseList() *fbtentity.WarehouseListResult
	InboundOrders(ids []string) *fbtentity.InboundOrdersResult
	SearchFbtInventory(pageSize int, pageToken string, body *fbtentity.SearchInventoryRequest) *fbtentity.SearchInventoryResult
	SearchFbtInventoryRecord(pageSize int, pageToken string, body *fbtentity.SearchInventoryRecordRequest) *fbtentity.SearchInventoryRecordResult
	SearchFbtGoods(pageSize int, pageToken string, body *fbtentity.SearchGoodsRequest) *fbtentity.SearchGoodsResult

	AftersaleEligibility(orderId string) *returnrefundentity.AftersaleEligibilityResult
	RejectReasons(returnOrCancelId, locale string) *returnrefundentity.RejectReasonsResult
	CreateReturn(idempotencyKey string, body *returnrefundentity.CreateReturnRequest) *returnrefundentity.CreateReturnResult
	ApproveReturn(returnId, idempotencyKey string, body *returnrefundentity.ApproveReturnRequest) *returnrefundentity.ApproveReturnResult
	RejectReturn(returnId, idempotencyKey string, body *returnrefundentity.RejectReturnRequest) *returnrefundentity.RejectReturnResult
	SearchReturns(pageSize int, pageToken, sortField, sortOrder string, body *returnrefundentity.SearchReturnsRequest) *returnrefundentity.SearchReturnsResult
	ReturnRecords(returnId string) *returnrefundentity.ReturnRecordsResult
	CancelOrder(body *returnrefundentity.CancelOrderRequest) *returnrefundentity.CancelOrderResult
	ApproveCancellation(cancelId, idempotencyKey string) *returnrefundentity.ApproveCancellationResult
	RejectCancellation(cancelId, idempotencyKey string, body *returnrefundentity.RejectCancellationRequest) *returnrefundentity.RejectCancellationResult
	SearchCancellations(pageSize int, pageToken, sortField, sortOrder string, body *returnrefundentity.SearchCancellationsRequest) *returnrefundentity.SearchCancellationsResult
	CalculateRefund(body *returnrefundentity.CalculateRefundRequest) *returnrefundentity.CalculateRefundResult

	GetStatements(pageSize int, pageToken, sortField, sortOrder, statementTimeGe, statementTimeLt, paymentStatus string) *financeentity.StatementsResult
	GetStatementTransactions(statementId string, pageSize int, pageToken, sortField, sortOrder string) *financeentity.StatementTransactionsResult
	GetOrderStatementTransactions(orderId string) *financeentity.OrderStatementTransactionsResult
	GetPayments(pageSize int, pageToken, sortField, sortOrder, createTimeGe, createTimeLt string) *financeentity.PaymentsResult
	GetWithdrawals(types string, pageSize int, pageToken, createTimeGe, createTimeLt string) *financeentity.WithdrawalsResult

	CreateConversation(body *customerserviceentity.CreateConversationRequest) *customerserviceentity.CreateConversationResult
	GetConversations(pageSize int, pageToken, locale string) *customerserviceentity.ConversationsResult
	GetConversationMessages(conversationId string, pageSize int, pageToken string) *customerserviceentity.MessagesResult
	SendMessage(conversationId string, body *customerserviceentity.SendMessageRequest) *customerserviceentity.SendMessageResult
	ReadMessage(conversationId string) *customerserviceentity.ReadMessageResult
	GetAgentSettings() *customerserviceentity.AgentSettingsResult
	UpdateAgentSettings(body *customerserviceentity.UpdateAgentSettingsRequest) *customerserviceentity.UpdateAgentSettingsResult
	UploadBuyerMessagesImage(filePath string) *customerserviceentity.UploadImageResult
	GetPerformance(startDate, endDate string) *customerserviceentity.PerformanceResult

	ConfirmPackageShipment(body *supplychainentity.ConfirmPackageShipmentRequest) *supplychainentity.ConfirmPackageShipmentResult
}

// TikTok
type TikTok struct {
	oauth.OAuth
	shop.Shop
	authorization.Authorization
	seller.Seller
	event.Event
	product.Product
	promotion.Promotion
	orders.Orders
	logistic.Logistic
	fulfillment.Fulfillment
	fbt.Fbt
	returnrefund.ReturnRefund
	finance.Finance
	customerservice.CustomerService
	supplychain.SupplyChain
}

func newTikTok(cfg *tiktokConfig.Config) *TikTok {
	return &TikTok{
		oauth.OAuth{Config: cfg},
		shop.Shop{Config: cfg},
		authorization.Authorization{Config: cfg},
		seller.Seller{Config: cfg},
		event.Event{Config: cfg},
		product.Product{Config: cfg},
		promotion.Promotion{Config: cfg},
		orders.Orders{Config: cfg},
		logistic.Logistic{Config: cfg},
		fulfillment.Fulfillment{Config: cfg},
		fbt.Fbt{Config: cfg},
		returnrefund.ReturnRefund{Config: cfg},
		finance.Finance{Config: cfg},
		customerservice.CustomerService{Config: cfg},
		supplychain.SupplyChain{Config: cfg},
	}
}

func NewApi(cfg *tiktokConfig.Config) TikToker {
	return newTikTok(cfg)
}

var tiktokList map[string]TikToker

func init() {
	tiktokList = make(map[string]TikToker)
}

func Register(name string, cfg *tiktokConfig.Config) {
	tiktokList[name] = newTikTok(cfg)
}

func GetApi(name string) TikToker {
	return tiktokList[name]
}
