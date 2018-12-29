package Register

import(
	"fmt"
	"../HttpCurl"
	"../Identity"
	"../Common"
	_"encoding/json"
)

type Register struct {
	//用户身份证号
	Usercardno string
	//用户登陆密码
	Password string
	//地区
	Area string
	//医院名称
	HosName string
	//科室名称
	DeptName string
	//医生名称
	DocName string
	//预约日期
	OrderDate string
	//预约时段 (上午am或下午pm)
	OrderPeriod string
	//预约最早时间
	OrderStime string
	//预约最晚时间
	OrderEtime string

	//信息确认页url
	checkUrl string
	//预约url
	registerUrl string

	RegInfo map[string]string
	CookieStr string
	base64Img string
	captcha string
}

func (Register *Register) init() {	
	Register.registerUrl = "http://www.zj12580.cn/order/save?yzmType=6&code="

	//参数校验
	Params := &Params{}
	Params.RegisterObj = Register
	if (!Params.CheckInfo()) {
		panic("register param is not correct")
	}

	//参数初始化
	Params.SetDefaultValues()
}

func (Register *Register) Register() {
	//挂号信息初始化
	Register.init()

	//首先请求一次order/check页面
	Register.queryOrderCheck()

	//重置验证码
	Register.resetYzmImg()

	//预提交挂号信息
	Register.preReg()

	//获取有效验证码
	Register.setYesYzm()

	//提交挂号信息
	Register.saveRegInfo()

}

func (Register *Register) setYesYzm() {
	captcha := &Identity.Captcha{}
	Register.captcha = captcha.GetYzmResultByImg(Register.base64Img)
	fmt.Printf(Register.captcha)
}

//预提交挂号信息
func (Register *Register) preReg() {
	Register.captcha = "3213"
	str := Register.saveRegInfo()

	//提取页面上的验证码
	mustCompile := `inputtype="hidden"id="resultYzm"value="(?P<base64Img>.*)"><inputtype="hidden"id="resultCode"`
	CommonFunc := &Common.Func{}
	result := CommonFunc.ParseRegReturn(string(str), mustCompile)
	
	Register.base64Img = result[0]["base64Img"]
}

func (Register *Register) saveRegInfo() []byte {
	Register.registerUrl = fmt.Sprintf("http://www.zj12580.cn/order/save?code=%s&yzmType=6", Register.captcha)

	Register.RegInfo["flag"] = "-1"
	httpCurl := &HttpCurl.HttpCurl{}
	httpCurl.SetUrl(Register.registerUrl)
	httpCurl.SetPostData(Register.RegInfo)

	headers := make(map[string]string)
	headers["Cookie"] = Register.CookieStr
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Upgrade-Insecure-Requests"] = "1"
	httpCurl.SetHeaders(headers)
	str, _ := httpCurl.GetContentsFromUrl()

	return str
}

//重置验证码图片
func (Register *Register) resetYzmImg() {
	captcha := &Identity.Captcha{}
	captcha.CookieStr = Register.CookieStr
	captcha.GetCaptchaImgBase64()
}

//挂号信息确认页
func (Register *Register) queryOrderCheck() {
	Register.checkUrl = "http://www.zj12580.cn/order/check"	
	
	httpCurl := &HttpCurl.HttpCurl{}
	httpCurl.SetUrl(Register.checkUrl)
	httpCurl.SetPostData(Register.RegInfo)

	headers := make(map[string]string)
	headers["Cookie"] = Register.CookieStr
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	httpCurl.SetHeaders(headers)
	httpCurl.GetContentsFromUrl()
}

