package main

import (
 	_"./Identity"
 	"fmt"
 	_"./HttpCurl"
 	"./Register"
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

	// Register := &Register.Register{}
	// Register.Usercardno = "362529199402120027"
	// Register.Password = "huav587lu"
	// Register.HosName = "杭州西溪医院"
	// Register.DeptName = "呼吸内科"
	// Register.DocName = ""
	// Register.OrderDate = "20181224"
	// Register.OrderPeriod = "am"
	// Register.OrderStime = "08:30"
	// Register.OrderEtime = "11:00"
	// Register.Register()

	Hospital := &Register.Hospital{}
	Hospital.SaveHospitalInfo()

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