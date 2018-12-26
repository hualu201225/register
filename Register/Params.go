package Register

import(
	"fmt"
	_"time"
	"../HttpCurl"
	"regexp"
	_"strings"
)

type Params struct {
	RegisterObj *Register
	Hospital *Hospital
}

func (Params *Params) CheckInfo() bool {
	//不为空字段判断
	if (len(Params.RegisterObj.Usercardno) == 0 ||
		len(Params.RegisterObj.Password) == 0 ||
		len(Params.RegisterObj.HosName) == 0 ||
		len(Params.RegisterObj.DeptName) == 0 ||
		len(Params.RegisterObj.OrderDate) == 0 ||
		len(Params.RegisterObj.OrderPeriod) == 0 ) {
		return false
	}

	//挂号日期判断（最近7天的）
	// regDateTimeStamp, _ := time.Parse("2016-01-02 00:00", Params.RegisterObj.OrderDate + " 00:00")
	// fmt.Println(regDateTimeStamp.Unix())
	// return false

	return true
}

func (Params *Params) SetDefaultValues() {
	//挂号时间段如果没传则赋予初始值
	Params.SetDefaultValue()

	//设置医院相关参数
	Params.SetHospitalValue()

	//设置符合要求的可选的最优号码
	Params.SetRegisterNumEtc()
}

func (Params *Params) SetDefaultValue() {
	if (len(Params.RegisterObj.OrderStime) == 0) {
		if (Params.RegisterObj.OrderPeriod == "am") {
			Params.RegisterObj.OrderStime = "08:00"
		} else {
			Params.RegisterObj.OrderStime = "13:00"
		}
	}

	if (len(Params.RegisterObj.OrderEtime) == 0) {
		if (Params.RegisterObj.OrderPeriod == "am") {
			Params.RegisterObj.OrderEtime = "12:00"
		} else {
			Params.RegisterObj.OrderEtime = "17:00"	
		}
	}

	if (len(Params.RegisterObj.Area) == 0) {
		Params.RegisterObj.Area = "杭州"
	}
}

//获取医院、科室、医生等详细信息
func (Params *Params) SetHospitalValue() {
	Params.Hospital = &Hospital{}
	Params.Hospital.CurrentAreaName = Params.RegisterObj.Area
	Params.Hospital.CurrentHosName = Params.RegisterObj.HosName
	Params.Hospital.CurrentDeptName = Params.RegisterObj.DeptName
	Params.Hospital.CurrentDocName = Params.RegisterObj.DocName
	Params.Hospital.GetDetailInfo()
}

//获取符合要求的最优号码（有无指定医生的区别处理）
func (Params *Params) SetRegisterNumEtc() {
	//预约页面url
	selectUrl := fmt.Sprintf("http://www.zj12580.cn/dept/queryDepartInfo/%s/%s/",Params.Hospital.CurrentHosId,Params.Hospital.CurrentDeptName)

	httpCurl := &HttpCurl.HttpCurl{}
	httpCurl.SetUrl(selectUrl)
	str, _ := httpCurl.GetContentsFromUrl()
	// fmt.Println(string(str))

	resString := string(str)
	reg := regexp.MustCompile("\\s+")
    resString = reg.ReplaceAllString(resString, "")
	// resString = strings.Replace(resString, " ", "", -1)
	// resString = strings.Replace(resString, "\n", "", -1)
	// resString = strings.Replace(resString, "\n", "", -1)

	//获取预约列表
	match := regexp.MustCompile("<form(.*)</form>")
	matchArr := match.FindAllString(resString, -1)
	fmt.Println(len(matchArr))
	for _,param :=range matchArr {
		fmt.Println(param)
		fmt.Println("\n\r")
		return
	}
}

func (Params *Params) setNormalNumEtc() {

}

func (Params *Params) setDocSpecificNumEtc() {
	
}
