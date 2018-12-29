package Common

import(
	_"strconv"
	"regexp"
	"strings"
	"fmt"
)

type Func struct {

}

// func (MapTrans *MapTrans) MapToString(mapNotTransed map[string]interface{}) (map[string]string) {
// 	mapTransed := make(map[string]string)
// 	for mapKey, mapVal := range mapNotTransed {
// 		//这里的类型转换有点问题 ？？？？？
// 		mapStr := strconv.Itoa(mapVal.(int))
// 		mapTransed[mapKey] = mapStr
// 	}

// 	return mapTransed
// }

func (Func *Func) ParseRegReturn(returnStr string, mustCompileStr string) []map[string]string {
	reg := regexp.MustCompile("\\s+")
    selectNumStr := reg.ReplaceAllString(returnStr, "")
    fmt.Println(selectNumStr)
    match := regexp.MustCompile(mustCompileStr)
    matchArr := match.FindAllStringSubmatch(selectNumStr, -1)

	groupNames := match.SubexpNames()
	var result []map[string]string
	//循环每行
	for _,param :=range matchArr {
		m := make(map[string]string)
		// 对每一行生成一个map
	    for j, name := range groupNames {
	        if j != 0 && name != "" {
	            m[name] = strings.TrimSpace(param[j])
	        }
	    }
	    result = append(result, m)
	}

	return result	
}
