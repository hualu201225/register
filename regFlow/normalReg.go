package regFlow

import (
	"fmt"
)

func HandleNormal(param map[string]string){
	fmt.Println(param)
	for i,v:=range param {
		i += " :"
		fmt.Println(i,v)
	}
}
