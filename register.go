package main

import (
 	"./Identity"
 	"fmt"
)

func main() {
	// Cookie := &Identity.Cookie{};
	// Cookie.SetCookie("432412341234发大水发射点发", "3625290027");

	// cookie := Cookie.GetCookie("3625290027")
	// fmt.Printf(cookie)

	Captcha := &Identity.Captcha{}
	yzmImg := Captcha.GetCaptchaImgBase64()
	
	fmt.Println(yzmImg)
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