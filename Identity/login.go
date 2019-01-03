package Identity

import (
	"../HttpCurl"
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"regexp"

)

//挂号登陆类
type RegLogin struct {
	websiteUrl string
	loginurl string
	username string
	password string
	token string
	captcha string
	tokenCookie string
}

func (RegLogin *RegLogin) SetUsername(username string) {
	RegLogin.username = username
}

func (RegLogin *RegLogin) SetPassword(password string) {
	RegLogin.password = password
}

func (RegLogin *RegLogin) init() {
	RegLogin.websiteUrl = "http://www.zj12580.cn/index?t=1545462434499"
	RegLogin.loginurl = "http://www.zj12580.cn/login"

	//校验用户身份证密码等信息是否正确
	if (!RegLogin.checkUser()) {
		panic("username or password is not correct")
	}

	//设置token和初始cookie
	RegLogin.setTokenAndCookie()
}

func (RegLogin *RegLogin) checkUser() bool {
	if (len(RegLogin.username) == 0 || len(RegLogin.password) == 0) {
		return false
	}
 
	//校验身份证格式 @TODO


	return true
}


func (RegLogin *RegLogin) setTokenAndCookie() {
	//首先访问12580官网
	httpCurl := &HttpCurl.HttpCurl{}
	httpCurl.SetUrl(RegLogin.websiteUrl)
	httpCurl.GetContentsFromUrl()
	//设置初始cookie
	RegLogin.tokenCookie = httpCurl.GetCookies()
	//获取初始token
	match := regexp.MustCompile("csrf_token=(.*); Expires")
	matchArr := match.FindStringSubmatch(RegLogin.tokenCookie)
	for _,param :=range matchArr {
		RegLogin.token = param
	}

}

func (RegLogin *RegLogin) SetCaptcha(captcha string) {
	RegLogin.captcha = captcha
}

func (RegLogin *RegLogin) InitCaptcha() {
	captcha := &Captcha{}
	base64Img := captcha.GetCaptchaImgBase64()
	captchaStr := captcha.GetYzmResultByImg(base64Img)

	RegLogin.SetCaptcha(captchaStr)
	// RegLogin.Login()
}

func (RegLogin *RegLogin) getMd5Passwd() string {
	md5Ctx := md5.New()
    md5Ctx.Write([]byte(RegLogin.password))
    cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr) + "5jMHTe"
}

func (RegLogin *RegLogin) getPostData() map[string]string {
	//获取base64img
	Captcha := &Captcha{}
	base64Img := Captcha.GetCaptchaImgBase64()
	//获取验证码
	RegLogin.captcha = Captcha.GetYzmResultByImg(base64Img)
	fmt.Println(RegLogin.captcha)
	md5pwd := RegLogin.getMd5Passwd()
	postData := make(map[string]string)
	postData["username"] = RegLogin.username
	postData["pwd"] = md5pwd
	postData["password"] = md5pwd
	postData["token"] = RegLogin.token
	postData["captcha"] = RegLogin.captcha
	postData["pwdIrregular"] = ""
	postData["pwdFourNum"] = ""
	postData["btnlogin"] = "登 录"
	postData["rememberMe"] = "1"
	return postData
}

func (RegLogin *RegLogin) getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["Upgrade-Insecure-Requests"] = "1"
	headers["Origin"] = "http://www.zj12580.cn"
	headers["Host"] = "www.zj12580.cn"
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["charset"] = "UTF-8"
	headers["Accept-Encoding"] = "gzip, deflate"
	headers["Cookie"] = RegLogin.tokenCookie

	return headers
} 

func (RegLogin *RegLogin) isAlreadyLogin() bool {
	//获取上次保存的cookie内容
	Cookie := &Cookie{}
	cookieStr := Cookie.GetCookie(RegLogin.username)
	if (len(cookieStr) == 0) {
		return false
	}

	str := RegLogin.getUserBaseInfo(cookieStr)
	//匹配字符串中是否有“退出系统”
	reg := regexp.MustCompile(`退出系统`)
	res := reg.FindAllString(str, -1)	
	if (len(res) > 0) {
		return true
	}

	return false
}

func (RegLogin *RegLogin) getUserBaseInfo(cookieStr string) string {
	httpCurl := &HttpCurl.HttpCurl{}
	httpCurl.SetUrl("http://www.zj12580.cn/user/baseInfo")

	headers := make(map[string]string)
	headers["Cookie"] = cookieStr
	httpCurl.SetHeaders(headers)

	str, _ := httpCurl.GetContentsFromUrl()

	return string(str)
}

//模拟登陆
func (RegLogin *RegLogin) Login() {
	//如果已经登陆了，则不重复登陆
	if (RegLogin.isAlreadyLogin()) {
		return
	}

	//参数初始化
	RegLogin.init()

	//获取验证码
	RegLogin.InitCaptcha()

	httpCurl := &HttpCurl.HttpCurl{}

	httpCurl.SetUrl(RegLogin.loginurl)

	//设置post参数
	postData := RegLogin.getPostData()
	httpCurl.SetPostData(postData)

	//设置header头
	headers := RegLogin.getHeaders()
	httpCurl.SetHeaders(headers)

	//登陆
	httpCurl.GetContentsFromUrl()

	//保存登陆后的cookie
	cookieStr := httpCurl.GetCookies()
	Cookie := &Cookie{}
	//需要把所有的cookie合到一块才能正常供后面的页面登陆
	cookieStr = cookieStr + ";" + RegLogin.tokenCookie + RegLogin.username + "=1;"
	//保存
	Cookie.SetCookie(cookieStr, RegLogin.username)
}