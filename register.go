package main   
import (
	_"fmt"
	"flag"
	"./regFlow"
)

func getTransferedParam() (param map[string]string) {
        userid := flag.String("userid", "123", "register_userid")
        username := flag.String("username", "张三", "register_name")
        identityNo := flag.String("identityNo", "32132432", "register_idNo")
	regType := flag.String("type", "1", "regFlow_type") 
	
	param = map[string]string{"userid":*userid, "username":*username, "identityNo":*identityNo, "type":*regType}

	return param
}

func pushToChannel(param map[string]string) {
	regType := param["type"]
	switch regType {
		case REGFLOW_TYPE_NORMAL:
			normalFlowChannel <- param
		case REGFLOW_TYPE_SPECIAL:
			specialFlowChannel <- param
		default : 
	}
}

func registerFromChannel() {
	normalReginfo := <-normalFlowChannel
	 regFlow.HandleNormal(normalReginfo)
}


var normalFlowChannel = make(chan map [string]string, 2)
var specialFlowChannel = make(chan map [string]string, 2)
const REGFLOW_TYPE_NORMAL string = "1" //正常挂号模式
const REGFLOW_TYPE_SPECIAL string = "2" //特殊挂号模式
//const REGFLOWS_CHANNEL_NAME string = "regflow" //挂号队列前置名  

func main() {
	//参数处理
	param := getTransferedParam()
	
	//挂号请求分发
	pushToChannel(param)
		
	//挂号请求处理
	registerFromChannel()
}


