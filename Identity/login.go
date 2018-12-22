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

func (RegLogin *RegLogin) init() {
	RegLogin.websiteUrl = "http://www.zj12580.cn/index?t=1545462434499"
	RegLogin.loginurl = "http://www.zj12580.cn/login"
	RegLogin.username = "362529199402120027"
	RegLogin.password = "huav587lu"
	RegLogin.setTokenAndCookie()
}

func (RegLogin *RegLogin) setTokenAndCookie() {
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

func (RegLogin *RegLogin) getMd5Passwd() string {
	// data := []byte(RegLogin.password)
	// has := md5.Sum(data)
	// md5str1 := fmt.Sprintf("%x", has)
	md5Ctx := md5.New()
    md5Ctx.Write([]byte(RegLogin.password))
    cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr) + "5jMHTe"
}

func (RegLogin *RegLogin) getPostData() map[string]string {
	md5pwd := RegLogin.getMd5Passwd()
	postData := make(map[string]string)
	postData["username"] = RegLogin.username
	postData["pwd"] = md5pwd
	postData["password"] = md5pwd
	postData["token"] = RegLogin.token
	postData["captcha"] = "fadsf" //RegLogin.captcha
	postData["pwdIrregular"] = ""
	postData["pwdFourNum"] = ""
	postData["btnlogin"] = "登 录"
	postData["rememberMe"] = "1"
	fmt.Println(RegLogin.getMd5Passwd())
	return postData
}

func (RegLogin *RegLogin) getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["Upgrade-Insecure-Requests"] = "1";
	headers["Origin"] = "http://www.zj12580.cn";
	headers["Host"] = "www.zj12580.cn";
	headers["Content-Type"] = "application/x-www-form-urlencoded;charset=UTF-8";
	headers["Accept-Encoding"] = "gzip, deflate";
	headers["Cookie"] = RegLogin.tokenCookie

	return headers
} 

//模拟登陆
func (RegLogin *RegLogin) Login() {
	RegLogin.init()
	httpCurl := &HttpCurl.HttpCurl{}

	httpCurl.SetUrl(RegLogin.loginurl)

	//设置post参数
	postData := RegLogin.getPostData()
	httpCurl.SetPostData(postData)

	//设置header头
	headers := RegLogin.getHeaders()
	httpCurl.SetHeaders(headers)

	//获取登陆结果
	result, _ := httpCurl.GetContentsFromUrl()

	fmt.Printf(string(result))

	//保存登陆后的cookie
	cookieStr := httpCurl.GetCookies()
	Cookie := &Cookie{}
	Cookie.SetCookie(cookieStr, RegLogin.username)
}