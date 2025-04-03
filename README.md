# tiktok go language API

API interface for tiktok go language version

my email :474790700@qq.com

## Contents

- [tiktok go language API](#tiktok-go)
  - [Installation](#installation)
  - [Quick start](#quick-start)

## Installation

To install tiktokgo package, you need to install Go and set your Go workspace first.

1. The first need [Go](https://golang.org/) installed (**version 1.12+ is required**), then you can use the below Go command to install tiktokgo.

```sh
$ go get -u github.com/wjp-letgo/tiktok-go
```

2. Import it in your code:

```go
import (
	"github.com/wjp-letgo/tiktok-go"
)
```
## Quick start

## API call

```go
package main

import (
    "fmt"
	"github.com/wjp-letgo/tiktok-go"
	
	"github.com/wjp-letgo/letgo/lib"
	tiktokConfig "github.com/wjp-letgo/tiktok-go/config"

	// productentity "github.com/wjp-letgo/tiktok-go/product/entity"
	// ordersentity "github.com/wjp-letgo/tiktok-go/orders/entity"
	fulfillmententity "github.com/wjp-letgo/tiktok-go/fulfillment/entity"
)

func main() {
	appKey := "234234234"
	appSecret := "234234"
	accessToken := "234234-23423-2342"
	shopCipher := "ROW_234234324"
	// Register("tiktok-api", tiktokConfig.New("https://auth.tiktok-shops.com", appKey, appSecret,"","", "/tiktokcallbackpush"))
	// fmt.Println(GetApi("tiktok-api").GetAccessToken("ROW_NBAewwAAAAC1bXwARPZYV-0M4wTgxSsQTi3TE53-52D4BTQM1Akoa6Gs0zhN-WHWxTCw4ejILQQz8PUNy0m9bbuUOFHRiVZUq_FWusQzbN84Mb8BwUdbq_xICk9eVD7R1spjvHf0y0a81XvDG4zMWXytHmE-Yeaa"))
	// Register("tiktok-api", tiktokConfig.New("https://auth.tiktok-shops.com", appKey,appSecret,accessToken,"", "/tiktokcallbackpush"))
	// fmt.Println(GetApi("tiktok-api").RefreshToken("ROW_p-1_tAAAAAAx9x1lb0WTEiqSJ1PcfWeHJRwIJdWCOMehui_-TYRIdngewdGNVRY97vDefg2na8I"))
	// Register("tiktok-api", tiktokConfig.New("https://open-api.tiktokglobalshop.com", appKey, appSecret,accessToken,"", "/tiktokcallbackpush"))
	// fmt.Println(GetApi("tiktok-api").Shops())
	Register("tiktok-api", tiktokConfig.New("https://open-api.tiktokglobalshop.com", appKey, appSecret, accessToken, shopCipher, "/tiktokcallbackpush"))
	// fmt.Println(GetApi("tiktok-api").Categories("","","",""))
	//https://p16-oec-sg.ibyteimg.com/tos-alisg-i-aphluv4xwc-sg/fb9f2980292849f4aaa51ae12077f944~tplv-aphluv4xwc-origin-jpeg.jpeg?dr=15568&from=1432613627&idc=my&ps=933b5bde&shcp=9cd7d13a&shp=5563f2fb&t=555f072dwidth=2500&height=2500
	// var images productentity.RecommendCategoryImages
	// images=append(images, productentity.RecommendCategoryImage{
	// 	Uri: "tos-alisg-i-aphluv4xwc-sg/fb9f2980292849f4aaa51ae12077f944",
	// })
	// data:=productentity.RecommendCategoryRequest{
	// 	ProductTitle: "wjp1231231312312312wwwwwerrrrrrreewwww",
	// 	Description: "wjp",
	// 	ListingPlatform: "TIKTOK_SHOP",
	// 	Images: images,
	// 	CategoryVersion: "v1",
	// }
	// fmt.Println( GetApi("tiktok-api").RecommendCategory(&data))
	// body:=&ordersentity.OrdersRequest{
	// 	OrderStatus: "AWAITING_SHIPMENT",
	// }
	// fmt.Println(GetApi("tiktok-api").OrdersSearch(1,"aDV5MHEyS1ZiRXRNUzJCTzh5MkliaGxpSWhEYXIvK3gyNjY0NmZJZFQzRVdBZz09","","",nil))
	// fmt.Println(GetApi("tiktok-api").OrderDetail([]string{"580147945690204170"}))
	// fmt.Println(GetApi("tiktok-api").PriceDetail("580144174201735178"))
	// var orders ordersentity.AddExternalOrders
	// var items ordersentity.ExternalLineItems
	// items=append(items, ordersentity.ExternalLineItem{
	// 	Id: "456",
	// 	OriginId: "580141452555028490",
	// })
	// orders=append(orders, ordersentity.AddExternalOrder{
	// 	Id: "580141452554897418",
	// 	ExternalOrder: ordersentity.ExternalOrder{
	// 		Id: "123456",
	// 		Platform: "ERP_SYSTEM",
	// 		LineItems: items,
	// 	},
	// })
	// body:=&ordersentity.AddExternalOrderRequest{
	// 	Orders: orders,
	// }
	// fmt.Println(GetApi("tiktok-api").AddExternalOrder(body))
	// fmt.Println(GetApi("tiktok-api").GlobalSellerWarehouse())
	// fmt.Println(GetApi("tiktok-api").WarehouseList())
	// fmt.Println(GetApi("tiktok-api").WarehouseDeliveryOptions("7487878078019880726"))
	// fmt.Println(GetApi("tiktok-api").ShippingProviders("7062652491816503041"))
	// fmt.Println(GetApi("tiktok-api").OrderSplitAttributes([]string{"580147944597981194"}))
	// var group fulfillmententity.SplittableGroups
	// group=append(group, fulfillmententity.SplittableGroup{
	// 	Id: "123456",
	// 	OrderLineItemIds: []string{"580147944598112266"},
	// })
	// group=append(group, fulfillmententity.SplittableGroup{
	// 	Id: "78910",
	// 	OrderLineItemIds: []string{"580147944598177802"},
	// })
	// body:=&fulfillmententity.SplitOrdersRequest{
	// 	SplittableGroups:group,
	// }
	// fmt.Println(GetApi("tiktok-api").SplitOrders("580147944597981194",body))
	// body:=&fulfillmententity.FirstmileBundleRequest{
	// 	OrderIds: []string{"580147944597981194"},
	// 	HandoverMethod: "PICKUP",
	// }
	// fmt.Println(GetApi("tiktok-api").FirstMileBundle(body))
	// body2:=&fulfillmententity.EligibleShippingServiceRequest{
	// 	OrderLineItemIds: []string{"580147944598177802","580147944598112266"},
	// 	Dimension: fulfillmententity.Dimension{
	// 		Length: "1.2",
	// 		Width: "1",
	// 		Height: "1",
	// 		Unit: "CM",
	// 	},
	// 	Weight: fulfillmententity.Weight{
	// 		Value: "1.2",
	// 		Unit: "POUND",
	// 	},
	// }
	// fmt.Println(GetApi("tiktok-api").EligibleShippingService("580147944597981194",body2))
	// body:=&fulfillmententity.CreatePackagesRequest{
	// 	OrderId: "580147944597981194",
	// 	OrderLineItemIds: []string{"580147944598177802","580147944598112266"},
	// 	Dimension: fulfillmententity.Dimension{
	// 		Length: "1.2",
	// 		Width: "1",
	// 		Height: "1",
	// 		Unit: "CM",
	// 	},
	// 	ShippingServiceId: "6617675021119438849",
	// 	Weight: fulfillmententity.Weight{
	// 		Value: "1.2",
	// 		Unit: "POUND",
	// 	},
	// }
	// fmt.Println(GetApi("tiktok-api").CreatePackages(body))
	// fmt.Println(GetApi("tiktok-api").SearchPackage(10,"","","",nil))
	// fmt.Println(GetApi("tiktok-api").PackageHandoverTimeSlots("1159836321565541386"))
	// body:=&fulfillmententity.ShipPackageRequest{
	// 	HandoverMethod: "PICKUP",
	// 	PickupSlot: fulfillmententity.ShipPickupSlot{
	// 		StartTime: lib.Time(),
	// 		EndTime: lib.Time(),
	// 	},
	// }
	// fmt.Println(GetApi("tiktok-api").ShipPackage("1161544026031491082",body))
	var packages fulfillmententity.BatchShipPackages
	packages=append(packages, fulfillmententity.BatchShipPackage{
		Id: "1161544026031491082",
		HandoverMethod: "PICKUP",
		PickupSlot: fulfillmententity.ShipPickupSlot{
			StartTime: lib.Time(),
			EndTime: lib.Time(),
		},
	})
	body := &fulfillmententity.BatchShipPackagesRequest{
		Packages: packages,
	}
	fmt.Println(GetApi("tiktok-api").BatchShipPackages(body))
	// fmt.Println(GetApi("tiktok-api").PackageShippingDocument("1161544026031491082","SHIPPING_LABEL_PICTURE","").Data.DocUrl)
	// fmt.Println(GetApi("tiktok-api").PackageDetail("1161544026031491082"))
	// body:=&fulfillmententity.UpdateShippingInfoRequest{
	// 	TrackingNumber: "173988430702626",
	// 	ShippingProviderId: "6617675021119438849",
	// }
	// fmt.Println(GetApi("tiktok-api").UpdateShippingInfo("580147945690204170",body))
}

```
