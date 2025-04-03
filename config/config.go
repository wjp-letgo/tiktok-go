package config

import (
	"errors"
	"fmt"
	"github.com/wjp-letgo/letgo/encry"
	"github.com/wjp-letgo/letgo/httpclient"
	"github.com/wjp-letgo/letgo/lib"
	"sort"
	"strings"
)

// Config
type Config struct {
	BaseURL     string `json:"baseURL"`
	AppKey      string `json:"app_key"`
	AppSecret   string `json:"app_secret"`
	RedirectURL string `json:"redirect_url"`
	AccessToken string `json:"access_token"`
	ShopCipher  string `json:"shop_cipher"`
}

// String
func (c *Config) String() string {
	return lib.ObjectToString(c)
}

// GetApiURL
func (c *Config) GetApiURL(apiPath string) string {
	return fmt.Sprintf("%s%s", c.BaseURL, apiPath)
}

// GetCommonParam
func (c *Config) GetCommonParam(method string) lib.InRow {
	ti := lib.Time()
	param := lib.InRow{
		"app_key":      c.AppKey,
		"timestamp":    ti,
		"access_token": c.AccessToken,
	}
	if c.ShopCipher != "" {
		param["shop_cipher"] = c.ShopCipher
	}
	return param
}

// HttpGet
func (c *Config) HttpGet(method string, data interface{}, out interface{}) error {
	return c.Http("GET", method, data, out)
}

// HttpPost
func (c *Config) HttpPost(method string, data interface{}, out interface{}) error {
	return c.Http("POST", method, data, out)
}

// HttpPostFile
func (c *Config) HttpPostFile(method string, data interface{}, out interface{}) error {
	return c.Http("POSTFILE", method, data, out)
}

// HttpS2
func (c *Config) GetS2(method string, data interface{}, out interface{}) error {
	param := lib.InRow{
		"app_key":   c.AppKey,
		"timestamp": lib.Time(),
	}
	inputParam := data.(lib.InRow)
	allParam := lib.MergeInRow(param, inputParam)
	param["sign"] = Sign(c.AppSecret, method, allParam, "")
	apiUrl := c.BaseURL
	fullURL := fmt.Sprintf("%s%s?%s", apiUrl, method, httpclient.HttpBuildQuery(param))
	//fmt.Println(fullURL, inputParam)
	ihttp := httpclient.New().WithTimeOut(120).WithHeader("x-tts-access-token", c.AccessToken).WithHeader("content-type", "application/json")
	result := ihttp.Get(fullURL, inputParam)
	// fmt.Println(result.Dump)
	if result.Err != "" {
		return errors.New(result.Err)
	}
	/*
		if result.Code != 200 {
			return errors.New("请求失败")
		}
	*/
	lib.StringToObject(result.Body(), out)
	return nil
}

// HttpS2
func (c *Config) GetS3(method string, data interface{}, out interface{}) error {
	param := lib.InRow{
		"app_key":     c.AppKey,
		"timestamp":   lib.Time(),
		"shop_cipher": c.ShopCipher,
	}
	inputParam := data.(lib.InRow)
	allParam := lib.MergeInRow(param, inputParam)
	param["sign"] = Sign(c.AppSecret, method, allParam, "")
	apiUrl := c.BaseURL
	fullURL := fmt.Sprintf("%s%s?%s", apiUrl, method, httpclient.HttpBuildQuery(param))
	//fmt.Println(fullURL, inputParam)
	ihttp := httpclient.New().WithTimeOut(120).WithHeader("x-tts-access-token", c.AccessToken).WithHeader("content-type", "application/json")
	result := ihttp.Get(fullURL, inputParam)
	// fmt.Println(result.Dump)
	if result.Err != "" {
		return errors.New(result.Err)
	}
	/*
		if result.Code != 200 {
			return errors.New("请求失败")
		}
	*/
	lib.StringToObject(result.Body(), out)
	return nil
}

// HttpS3
func (c *Config) HttpS3(method string, query lib.InRow, data interface{}, out interface{}) error {
	param := lib.InRow{
		"app_key":      c.AppKey,
		"timestamp":    lib.Time(),
		"shop_cipher":  c.ShopCipher,
	}
	allParam := lib.MergeInRow(param, query)
	body := ""
	if data!=nil{
		body=fmt.Sprintf("%s", data)
	}
	allParam["sign"] = Sign(c.AppSecret, method, allParam, body)
	apiUrl := c.BaseURL
	fullURL := fmt.Sprintf("%s%s?%s", apiUrl, method, httpclient.HttpBuildQuery(allParam))
	// fmt.Println("fullURL:",fullURL)
	ihttp := httpclient.New().WithTimeOut(120).WithHeader("x-tts-access-token", c.AccessToken).WithHeader("content-type", "application/json")
	result := ihttp.PostJson(fullURL, data)
	if result.Err != "" {
		return errors.New(result.Err)
	}
	// fmt.Println(result.Dump)
	/*
		if result.Code != 200 {
			return errors.New("请求失败")
		}
	*/
	lib.StringToObject(result.Body(), out)
	return nil
}

// HttpS3
func (c *Config) HttpS4(method string, query lib.InRow, data interface{}, out interface{}) error {
	param := lib.InRow{
		"app_key":      c.AppKey,
		"timestamp":    lib.Time(),
	}
	allParam := lib.MergeInRow(param, query)
	body := ""
	if data!=nil{
		body=fmt.Sprintf("%s", data)
	}
	allParam["sign"] = Sign(c.AppSecret, method, allParam, body)
	apiUrl := c.BaseURL
	fullURL := fmt.Sprintf("%s%s?%s", apiUrl, method, httpclient.HttpBuildQuery(allParam))
	// fmt.Println("fullURL:",fullURL)
	ihttp := httpclient.New().WithTimeOut(120).WithHeader("x-tts-access-token", c.AccessToken).WithHeader("content-type", "application/json")
	result := ihttp.PostJson(fullURL, data)
	if result.Err != "" {
		return errors.New(result.Err)
	}
	// fmt.Println(result.Dump)
	/*
		if result.Code != 200 {
			return errors.New("请求失败")
		}
	*/
	lib.StringToObject(result.Body(), out)
	return nil
}
// HttpS2
func (c *Config) HttpS2(method string, data interface{}, out interface{}) error {
	param := lib.InRow{
		"app_key":   c.AppKey,
		"timestamp": lib.Time(),
	}
	inputParam := data.(lib.InRow)
	allParam := lib.MergeInRow(param, inputParam)
	param["sign"] = Sign(c.AppSecret, method, allParam, "")
	apiUrl := c.BaseURL
	fullURL := fmt.Sprintf("%s%s?%s", apiUrl, method, httpclient.HttpBuildQuery(param))
	//fmt.Println(fullURL, inputParam)
	ihttp := httpclient.New().WithTimeOut(120).WithHeader("x-tts-access-token", c.AccessToken).WithHeader("content-type", "application/json")
	result := ihttp.Post(fullURL, inputParam)
	fmt.Println(result.Dump)
	if result.Err != "" {
		return errors.New(result.Err)
	}
	/*
		if result.Code != 200 {
			return errors.New("请求失败")
		}
	*/
	lib.StringToObject(result.Body(), out)
	return nil
}

// HttpS
func (c *Config) HttpS(method string, data interface{}, out interface{}) error {
	param := lib.InRow{
		"app_key":    c.AppKey,
		"app_secret": c.AppSecret,
	}
	inputParam := data.(lib.InRow)
	allParam := lib.MergeInRow(param, inputParam)
	param["sign"] = Sign(c.AppSecret, method, allParam, "")
	apiUrl := c.BaseURL
	fullURL := fmt.Sprintf("%s%s?%s", apiUrl, method, httpclient.HttpBuildQuery(param))
	//fmt.Println(fullURL, inputParam)
	ihttp := httpclient.New().WithTimeOut(120)
	result := ihttp.Post(fullURL, inputParam)
	//fmt.Println(result.Dump)
	if result.Err != "" {
		return errors.New(result.Err)
	}
	/*
		if result.Code != 200 {
			return errors.New("请求失败")
		}
	*/
	lib.StringToObject(result.Body(), out)
	return nil
}

// Http 请求
func (c *Config) Http(requestMethod, method string, data interface{}, out interface{}) error {
	param := c.GetCommonParam(method)
	inputParam := data.(lib.InRow)
	allParam := lib.MergeInRow(param, inputParam)
	param["sign"] = Sign(c.AppSecret, method, allParam, "")
	//fmt.Println(param)
	apiUrl := c.BaseURL
	fullURL := fmt.Sprintf("%s%s?%s", apiUrl, method, httpclient.HttpBuildQuery(param))
	ihttp := httpclient.New().WithTimeOut(120)
	var result *httpclient.HttpResponse
	if requestMethod == "GET" {
		//fmt.Println(fullURL)
		result = ihttp.Get(fullURL, inputParam)
	} else {
		result = ihttp.Post(fullURL, inputParam)
	}
	//fmt.Println(result.Dump)
	if result.Err != "" {
		return errors.New(result.Err)
	}
	/*
		if result.Code != 200 {
			return errors.New("请求失败")
		}
	*/
	lib.StringToObject(result.Body(), out)
	return nil
}

// New
func New(apiURL, AppKey, AppSecret, AccessToken, shopCipher string, redirectURL string) *Config {
	return &Config{
		apiURL,
		AppKey,
		AppSecret,
		redirectURL,
		AccessToken,
		shopCipher,
	}
}

// Sign
func Sign(appSecret, method string, param lib.InRow, body string) string {
	tmpParam := lib.InRow{}
	//排除掉文件上传
	for k, v := range param {
		if k[0] != '@' {
			tmpParam[k] = v
		}
	}
	query := ""
	sortedKeys := make([]string, 0)
	for k, _ := range tmpParam {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)
	for _, k := range sortedKeys {
		if vs, ok := tmpParam[k].(string); ok {
			query += fmt.Sprintf("%s%s", k, vs)
		} else if vs, ok := tmpParam[k].(int); ok {
			query += fmt.Sprintf("%s%d", k, vs)
		} else if vs, ok := tmpParam[k].(int64); ok {
			query += fmt.Sprintf("%s%d", k, vs)
		} else if vs, ok := tmpParam[k].(int32); ok {
			query += fmt.Sprintf("%s%d", k, vs)
		} else if vs, ok := tmpParam[k].(float32); ok {
			query += fmt.Sprintf("%s%f", k, vs)
		} else if vs, ok := tmpParam[k].(float64); ok {
			query += fmt.Sprintf("%s%f", k, vs)
		} else if vs, ok := tmpParam[k].(bool); ok {
			query += fmt.Sprintf("%s%t", k, vs)
		}

	}
	input := appSecret + method + query + appSecret
	if body != "" {
		input = appSecret + method + query + body + appSecret
		// fmt.Println("input:",input)
	}
	// fmt.Println("input:", input)
	ret := strings.ToLower(encry.HmacHex(input, appSecret))
	// fmt.Println("sign:", ret)
	return ret
}
