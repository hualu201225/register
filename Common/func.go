package Common

import(
	"strconv"
)

type MapTrans struct {

}

func (MapTrans *MapTrans) MapToString(mapNotTransed map[string]interface{}) (map[string]string) {
	mapTransed := make(map[string]string)
	for mapKey, mapVal := range mapNotTransed {
		//这里的类型转换有点问题 ？？？？？
		mapStr := strconv.Itoa(mapVal.(int))
		mapTransed[mapKey] = mapStr
	}

	return mapTransed
}
