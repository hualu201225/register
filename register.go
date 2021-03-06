package main

import (
 	"./Identity"
 	"fmt"
 	_"./HttpCurl"
 	_"./Register"
)

func main() {
	//获取当前用户的cookie
	// Cookie := &Identity.Cookie{};
	// cookie := Cookie.GetCookie("362529199402120027")

	// httpCurl := &HttpCurl.HttpCurl{}
	// httpCurl.SetUrl("http://www.zj12580.cn/order/queryOrderRecord?t=0.6084146469394638")

	// headers := make(map[string]string)
	// headers["Cookie"] = cookie
	// httpCurl.SetHeaders(headers)

	// result, _ := httpCurl.GetContentsFromUrl()
	// fmt.Printf(string(result))

	//登陆操作=====================
	// captcha := &Identity.Captcha{}
	// base64Img := captcha.GetCaptchaImgBase64()
	// captchaStr := captcha.GetYzmResultByImg(base64Img)

	// fmt.Printf(captchaStr)

	// login := &Identity.RegLogin{}
	// login.SetUsername("362529199402120027")
	// login.SetPassword("huav587lu")
	// // login.SetCaptcha(captchaStr)
	// login.Login()
	//===============================

	//挂号操作=============================
	Register := &Register.Register{}
	Register.Usercardno = "362529199402120027"
	Register.Password = "huav587lu"
	Register.HosName = "杭州市西溪医院"
	Register.DeptName = "呼吸内科"
	Register.DocName = ""
	Register.OrderDate = "20190104"
	Register.OrderPeriod = "am"
	Register.OrderStime = "08:30"
	Register.OrderEtime = "11:00"
	Register.Register()
	//======================================

	// Hospital := &Register.Hospital{}
	// // Hospital.SaveHospitalInfo()
	// Hospital.CurrentAreaName = "杭州"
	// Hospital.CurrentHosName = "杭州大同中医门诊部"
	// Hospital.CurrentDeptName = "中医妇科"
	// Hospital.CurrentDocName = "包汝中"
	// Hospital.GetDetailInfo()
	// fmt.Printf(Hospital.CurrentDocId)
}

//解析 map[string]interface{} 数据格式
func print_map(m map[string]interface{}) {
    for k, v := range m {
        switch value := v.(type) {
        case nil:
            fmt.Println(k, "is nil:", "null")
        case string:
            fmt.Println(k, "is string:", value)
        case int:
            fmt.Println(k, "is int:", value)
        case float64:
            fmt.Println(k, "is float64:", value)
        case []interface{}:
            fmt.Println(k, "is an array:")
            for i, u := range value {
                fmt.Println(i, u)
            }
        case map[string]interface{}:
            fmt.Println(k, "is an map:")
            print_map(value)
        default:
            fmt.Println(k, "is unknown type:", fmt.Sprintf("%T", v))
        }
    }
}