package Identity

import(
	"../HttpCurl"
)

type Captcha struct{
	headerOrigin string
	headerHost string
	base64ImgUrl string
	yzmScanResUrl string
	yzmAppcode string

	// headerOrigin = "http://www.zj12580.cn"
	// headerHost := "www.zj12580.cn"

	// //获取base64验证码图片地址
	// base64ImgUrl := "http://www.zj12580.cn/captcha?yzmType=6"
	// //验证码图片识别地址
	// yzmScanResUrl := "https://302307.market.alicloudapi.com/ocr/captcha"
	// //验证码识别appcode
	// yzmAppcode := "a924d95422454ad4a335250b534e419a"
	
}

func (Captcha *Captcha) init() {
	Captcha.headerOrigin = "http://www.zj12580.cn"
	Captcha.headerHost = "www.zj12580.cn"
	Captcha.base64ImgUrl = "http//www.zj12580.cn/captcha?yzmType=6"
	Captcha.yzmScanResUrl = "https//302307.market.alicloudapi.com/ocr/captcha"
	Captcha.yzmAppcode = "a924d95422454ad4a335250b534e419a"
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
	httpCurl.SetHeaders(headers)

	res := make(map[string]interface{})
	res, err := httpCurl.GetContentsFromUrl()
	if (err != nil) {
		panic("can not get yzmImg")
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

	postData := make(map[string]interface{})
	postData["image"] = base64Img
	postData["length"] = 0
	postData["type"] = 1001
	httpCurl.SetPostData(postData)

	res := make(map[string]interface{})
	res, err := httpCurl.GetContentsFromUrl()
	if (err != nil) {
		panic("can not get correct yzm")
	}

	return res["data"]["captcha"].(string)

}




