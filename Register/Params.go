package Register

import(
	"fmt"
	_"time"
	"../HttpCurl"
	"regexp"
	"strings"
	_"encoding/json"
	"strconv"
	"../Identity"
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

	//设置符合要求的挂号信息
	Params.SetRegisterNumEtc()

	//设置最优的号码
	Params.SetBestNum()
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

	//获取页面信息
	httpCurl := &HttpCurl.HttpCurl{}
	httpCurl.SetUrl(selectUrl)
	str, _ := httpCurl.GetContentsFromUrl()

	//页面信息格式化
	resString := string(str)
	reg := regexp.MustCompile("\\s+")
    resString = reg.ReplaceAllString(resString, "")

    //正则匹配挂号信息
	match := regexp.MustCompile(`data-idx="(?P<idx>\d+)"data-type="per"><formaction="/order/num"method="get"name="orderInfo"><inputtype="hidden"name="hisSchemeId"value=""><inputtype="hidden"name="schemeId"value="(?P<schemeId>\d+)"><inputtype="hidden"name="orderDate"value="(?P<orderDate>\d+)"><inputtype="hidden"name="hosId"value="(?P<hosId>\d+)"><inputtype="hidden"name="hosName"value="(?P<hosName>[\p{Han}|(|)]+)"><inputtype="hidden"name="deptId"value="(?P<deptId>\d+)"><inputtype="hidden"name="deptName"value="(?P<deptName>[\p{Han}]+)"><inputtype="hidden"name="docTitle"value="(?P<docTitle>[\p{Han}]{0,})"><inputtype="hidden"name="docId"value="(?P<docId>\d{0,})"><inputtype="hidden"name="docName"value="(?P<docName>[\p{Han}]{0,})"><inputtype="hidden"name="regFee"value="(?P<regFee>\d+)"><inputtype="hidden"name="takeNumAddr"value="(?P<takeNumAddr>\d{0,})"><inputtype="hidden"name="resTimeSign"value="(?P<resTimeSign>\d{0,})"><inputtype="submit"class="btnyy"value="&#13;&#10;预约&#13;&#10;\d+"title="总放号数(?P<totalNum>\d+)人次,剩余(?P<remainNum>\d+)`)
	matchArr := match.FindAllStringSubmatch(resString, -1)

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

	//结果格式化输出
	// prettyResult, _ := json.MarshalIndent(result, "", "  ")
	// fmt.Println(string(prettyResult))

	//解析结果获取正确的信息
	Params.getAvailableRegInfo(result)
}

func (Params *Params) getAvailableRegInfo(result []map[string]string) {
	var beFixed int
	if (Params.RegisterObj.OrderPeriod == "am") {
		beFixed = 1
	} else {
		beFixed = 2
	}
	var isPeriodMatch bool
	var left int
	isPeriodMatch = false
	for _, v := range result {
		//判断上午/下午是否匹配
		idx, _ := strconv.Atoi(v["idx"])
		left = idx % beFixed
		if ( left == 0) {
			isPeriodMatch = true
		}

		//如果日期和上午/下午时间段都能对上，则设置挂号信息
		if (isPeriodMatch == true && v["orderDate"] == Params.RegisterObj.OrderDate) {
			Params.RegisterObj.RegInfo = v
			return
		}
	}

	panic("there is no available register num")
}

func (Params *Params) SetBestNum() {
	orderUrl := "http://www.zj12580.cn/order/num"
	httpCurl := &HttpCurl.HttpCurl{}
	httpCurl.SetUrl(orderUrl)
	httpCurl.SetQueries(Params.RegisterObj.RegInfo)
	httpCurl.SetNeedUrlparse(true)

	Cookie := &Identity.Cookie{}
	cookieStr := Cookie.GetCookie(Params.RegisterObj.Usercardno)

	headers := make(map[string]string)
	headers["Cookie"] = cookieStr
	httpCurl.SetHeaders(headers)

	str, _ := httpCurl.GetContentsFromUrl()
	fmt.Println(string(str))

}

func (Params *Params) setDocSpecificNumEtc() {
	
}
