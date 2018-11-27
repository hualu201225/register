package main

import (
 	_"./Identity"
 	"./HttpCurl"
 	"fmt"
)

func main() {
	// Cookie := &Identity.Cookie{};
	// Cookie.SetCookie("432412341234发大水发射点发", "3625290027");

	// cookie := Cookie.GetCookie("3625290027")
	// fmt.Printf(cookie)

	httpCurl := &HttpCurl.HttpCurl{}
	httpCurl.SetUrl("http://www.zj12580.cn/captcha?yzmType=6")

	headers := make(map[string]string)
	headers["Origin"] = "http://www.zj12580.cn"
	headers["Host"] = "www.zj12580.cn"
	httpCurl.SetHeaders(headers)

	res := make(map[string]interface{})
	res, _ = httpCurl.GetContentsFromUrl()
	fmt.Println(res)
	//print_map(res)
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