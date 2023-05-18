package entity

import "github.com/wjpxxx/letgo/lib"

type GetBrandsResult struct {
	Code      int                 `json:"code"`
	Message   string              `json:"message"`
	Data      GetBrandsResultData `json:"data"`
	RequestId string              `json:"request_id"`
}

type GetBrandsResultData struct {
	BrandList []BrandList `json:"brand_list"`
	TotalNum  int         `json:"total_num"`
}

type BrandList struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	AuthorizedStatus int    `json:"authorized_status"`
}

func (getbrandsresult *GetBrandsResult) String() string {
	return lib.ObjectToString(getbrandsresult)
}
func (getbrandsresultdata *GetBrandsResultData) String() string {
	return lib.ObjectToString(getbrandsresultdata)
}
func (getbrandlist *BrandList) String() string {
	return lib.ObjectToString(getbrandlist)
}
