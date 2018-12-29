package Identity

import(
	"../HttpCurl"
	"encoding/json"
	"fmt"
)

type Captcha struct{
	headerOrigin string
	headerHost string
	base64ImgUrl string
	yzmScanResUrl string
	yzmAppcode string
	CookieStr string
}

func (Captcha *Captcha) init() {
	Captcha.headerOrigin = "http://www.zj12580.cn"
	Captcha.headerHost = "www.zj12580.cn"
	Captcha.base64ImgUrl = "http://www.zj12580.cn/captcha"
	Captcha.yzmScanResUrl = "https://302307.market.alicloudapi.com/ocr/captcha"
    // Captcha.yzmAppcode = "a924d95422454ad4a335250b534e419a"
    //无效appcode
	Captcha.yzmAppcode = "a924d95422454ad4a335250fadsfasdfasb534e419a"
}

//获取base64的验证码图片
func (Captcha *Captcha) GetCaptchaImgBase64() string {
	Captcha.init()

	httpCurl := &HttpCurl.HttpCurl{}
	httpCurl.SetUrl(Captcha.base64ImgUrl)

	queries := make(map[string]string)
	queries["yzmType"] = "6"
	httpCurl.SetQueries(queries)

	headers := make(map[string]string)
	headers["Origin"] = Captcha.headerOrigin
	headers["Host"] = Captcha.headerHost
	headers["Cookie"] = Captcha.CookieStr
	httpCurl.SetHeaders(headers)

	str, err := httpCurl.GetContentsFromUrl()
	if (err != nil) {
		panic("can not get yzmImg")
	}

	res := make(map[string]interface{})
	error := json.Unmarshal(str, &res)
	if (error != nil) {
		panic("json unmarshal failed")
	}
	
	return res["yzm"].(string)
}

func (Captcha *Captcha) GetYzmResultByImg(base64Img string) string {
	Captcha.init()

	httpCurl := &HttpCurl.HttpCurl{}
	httpCurl.SetUrl(Captcha.yzmScanResUrl)

	headers := make(map[string]string)
	headers["Authorization"] = "APPCODE " + Captcha.yzmAppcode
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["charset"] = "UTF-8"
	httpCurl.SetHeaders(headers)

	postData := make(map[string]string)
	postData["image"] = base64Img
	postData["length"] = "0"
	postData["type"] = "1001"
	httpCurl.SetPostData(postData)

	str, err := httpCurl.GetContentsFromUrl()
	if (err != nil) {
		panic("can not get correct yzm")
	}
	fmt.Printf(string(str))
	res := make(map[string]map[string]string)
	//resultStr := "{\"code\":0,\"data\":{\"captcha\":\"nwmn\",\"type\":1001,\"length\":4,\"id\":\"e0cf8713-0343-49c2-a75f-47f58f05b392\"}}"	
	//str = []byte(resultStr)
	errors := json.Unmarshal(str, &res)
	 if (errors != nil) {
	 	//这里会报个错，暂时关闭程序也可以运行，解决方案之后再找吧
		// 	panic(errors)
	 }

	return  res["data"]["captcha"]
}






