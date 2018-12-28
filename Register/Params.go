package Register

import(
	"fmt"
	_"time"
	"../HttpCurl"
	"regexp"
	"strings"
	"encoding/json"
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
	// resString := "<tr><tdclass=\"td_t\">普通</td><tddata-1=\"\"data-type=\"per\">&nbsp;</td><tddata-2=\"\"data-type=\"per\">&nbsp;</td><tddata-idx=\"3\"data-type=\"per\"title=\"上午下午&nbsp;上午\"><formaction=\"/order/num\"method=\"get\"name=\"orderInfo\"><inputtype=\"hidden\"name=\"hisSchemeId\"value=\"\"><inputtype=\"hidden\"name=\"schemeId\"value=\"0\"><inputtype=\"hidden\"name=\"orderDate\"value=\"20181229\"><inputtype=\"hidden\"name=\"hosId\"value=\"057166\"><inputtype=\"hidden\"name=\"hosName\"value=\"杭州市西溪医院(市六医院)\"><inputtype=\"hidden\"name=\"deptId\"value=\"10238\"><inputtype=\"hidden\"name=\"deptName\"value=\"呼吸内科\"><inputtype=\"hidden\"name=\"docTitle\"value=\"\"><inputtype=\"hidden\"name=\"docId\"value=\"\"><inputtype=\"hidden\"name=\"docName\"value=\"普通\"><inputtype=\"hidden\"name=\"regFee\"value=\"10\"><inputtype=\"hidden\"name=\"takeNumAddr\"value=\"\"><spanclass=\"cz\"onclick=\"alert('暂未放号或已过预约时间');\"title=\"总放号数11人次,剩余10人次,诊金10元\">预约</span></form></td><tddata-idx=\"4\"data-type=\"per\"title=\"上午下午&nbsp;下午\"><formaction=\"/order/num\"method=\"get\"name=\"orderInfo\"><inputtype=\"hidden\"name=\"hisSchemeId\"value=\"\"><inputtype=\"hidden\"name=\"schemeId\"value=\"0\"><inputtype=\"hidden\"name=\"orderDate\"value=\"20181229\"><inputtype=\"hidden\"name=\"hosId\"value=\"057166\"><inputtype=\"hidden\"name=\"hosName\"value=\"杭州市西溪医院(市六医院)\"><inputtype=\"hidden\"name=\"deptId\"value=\"10238\"><inputtype=\"hidden\"name=\"deptName\"value=\"呼吸内科\"><inputtype=\"hidden\"name=\"docTitle\"value=\"\"><inputtype=\"hidden\"name=\"docId\"value=\"\"><inputtype=\"hidden\"name=\"docName\"value=\"普通\"><inputtype=\"hidden\"name=\"regFee\"value=\"10\"><inputtype=\"hidden\"name=\"takeNumAddr\"value=\"\"><spanonclick=\"alert('暂未放号或已过预约时间');\"class=\"cz\"title=\"总放号数11人次,剩余11人次,诊金10元\"alt=\"总放号数11人次,剩余11人次,诊金10元\">预约</span></form></td><tddata-idx=\"5\"data-type=\"per\"><formaction=\"/order/num\"method=\"get\"name=\"orderInfo\"><inputtype=\"hidden\"name=\"hisSchemeId\"value=\"\"><inputtype=\"hidden\"name=\"schemeId\"value=\"12815339\"><inputtype=\"hidden\"name=\"orderDate\"value=\"20181230\"><inputtype=\"hidden\"name=\"hosId\"value=\"057166\"><inputtype=\"hidden\"name=\"hosName\"value=\"杭州市西溪医院(市六医院)\"><inputtype=\"hidden\"name=\"deptId\"value=\"10238\"><inputtype=\"hidden\"name=\"deptName\"value=\"呼吸内科\"><inputtype=\"hidden\"name=\"docTitle\"value=\"\"><inputtype=\"hidden\"name=\"docId\"value=\"\"><inputtype=\"hidden\"name=\"docName\"value=\"普通\"><inputtype=\"hidden\"name=\"regFee\"value=\"10\"><inputtype=\"hidden\"name=\"takeNumAddr\"value=\"\"><inputtype=\"hidden\"name=\"resTimeSign\"value=\"0\"><inputtype=\"submit\"class=\"btnyy\"value=\"预约13\"title=\"总放号数13人次,剩余13人次,诊金10元\"alt=\"总放号数13人次,剩余13人次,诊金10元\"></form></td><tddata-idx=\"6\"data-type=\"per\"><formaction=\"/order/num\"method=\"get\"name=\"orderInfo\"><inputtype=\"hidden\"name=\"hisSchemeId\"value=\"\"><inputtype=\"hidden\"name=\"schemeId\"value=\"12815339\"><inputtype=\"hidden\"name=\"orderDate\"value=\"20181230\"><inputtype=\"hidden\"name=\"hosId\"value=\"057166\"><inputtype=\"hidden\"name=\"hosName\"value=\"杭州市西溪医院(市六医院)\"><inputtype=\"hidden\"name=\"deptId\"value=\"10238\"><inputtype=\"hidden\"name=\"deptName\"value=\"呼吸内科\"><inputtype=\"hidden\"name=\"docTitle\"value=\"\"><inputtype=\"hidden\"name=\"docId\"value=\"\"><inputtype=\"hidden\"name=\"docName\"value=\"普通\"><inputtype=\"hidden\"name=\"regFee\"value=\"10\"><inputtype=\"hidden\"name=\"takeNumAddr\"value=\"\"><inputtype=\"hidden\"name=\"resTimeSign\"value=\"1\"><inputtype=\"submit\"class=\"btnyy\"value=\"预约11\"title=\" 总放号数11人次,剩余11人次,诊金10元\"alt=\"总放号数11人次,剩余11人次,诊金10元\"></form></td><tddata-idx=\"7\"data-type=\"per\"><formaction=\"/order/num\"method=\"get\"name=\"orderInfo\"><inputtype=\"hidden\"name=\"hisSchemeId\"value=\"\"><inputtype=\"hidden\"name=\"schemeId\"value=\"12820484\"><inputtype=\"hidden\"name=\"orderDate\"value=\"20181231\"><inputtype=\"hidden\"name=\"hosId\"value=\"057166\"><inputtype=\"hidden\"name=\"hosName\"value=\"杭州市西溪医院(市六医院)\"><inputtype=\"hidden\"name=\"deptId\"value=\"10238\"><inputtype=\"hidden\"name=\"deptName\"value=\"呼吸内科\"><inputtype=\"hidden\"name=\"docTitle\"value=\"\"><inputtype=\"hidden\"name=\"docId\"value=\"\"><inputtype=\"hidden\"name=\"docName\"value=\"普通\"><inputtype=\"hidden\"name=\"regFee\"value=\"10\"><inputtype=\"hidden\"name=\"takeNumAddr\"value=\"\"><inputtype=\"hidden\"name=\"resTimeSign\"value=\"0\"><inputtype=\"submit\"class=\"btnyy\"value=\"预约14\"title=\"总放号数14人次,剩余14人次,诊金10元\"alt=\"总放号数14人次,剩余14人次,诊金10元\"></form></td><tddata-idx=\"8\"data-type=\"per\"><formaction=\"/order/num\"method=\"get\"name=\"orderInfo\"><inputtype=\"hidden\"name=\"hisSchemeId\"value=\"\"><inputtype=\"hidden\"name=\"schemeId\"value=\"12820484\"><inputtype=\"hidden\"name=\"orderDate\"value=\"20181231\"><inputtype=\"hidden\"name=\"hosId\"value=\"057166\"><inputtype=\"hidden\"name=\"hosName\"value=\"杭州市西溪医院(市六医院)\"><inputtype=\"hidden\"name=\"deptId\"value=\"10238\"><inputtype=\"hidden\"name=\"deptName\"value=\"呼吸内科\"><inputtype=\"hidden\"name=\"docTitle\"value=\"\"><inputtype=\"hidden\"name=\"docId\"value=\"\"><inputtype=\"hidden\"name=\"docName\"value=\"普通\"><inputtype=\"hidden\"name=\"regFee\"value=\"10\"><inputtype=\"hidden\"name=\"takeNumAddr\"value=\"\"><inputtype=\"hidden\"name=\"resTimeSign\"value=\"1\"><inputtype=\"submit\"class=\"btnyy\"value=\"预约12\"title=\"总放号数12人次,剩余12人次,诊金10元\"alt=\"总放号数12人次,剩余12人次,诊金10元\"></form></td><tddata-idx=\"9\"data-type=\"per\"><formaction=\"/order/num\"method=\"get\"name=\"orderInfo\"><inputtype=\"hidden\"name=\"hisSchemeId\"value=\"\"><inputtype=\"hidden\"name=\"schemeId\"value=\"12829130\"><inputtype=\"hidden\"name=\"orderDate\"value=\"20190101\"><inputtype=\"hidden\"name=\"hosId\"value=\"057166\"><inputtype=\"hidden\"name=\"hosName\"value=\"杭州市西溪医院(市六医院)\"><inputtype=\"hidden\"name=\"deptId\"value=\"10238\"><inputtype=\"hidden\"name=\"deptName\"value=\"呼吸内科\"><inputtype=\"hidden\"name=\"docTitle\"value=\"\"><inputtype=\"hidden\"name=\"docId\"value=\"\"><inputtype=\"hidden\"name=\"docName\"value=\"普通\"><inputtype=\"hidden\"name=\"regFee\"value=\"10\"><inputtype=\"hidden\"name=\"takeNumAddr\"value=\"\"><inputtype=\"hidden\"name=\"resTimeSign\"value=\"0\"><inputtype=\"submit\"class=\"btnyy\"value=\"预约10\"title=\"总放号数10人次,剩余10人次,诊金10元\"alt=\"总放号数10人次,剩余10人次,诊金10元\"></form></td><tddata-idx=\"10\"data-type=\"per\"><formaction=\"/order/num\"method=\"get\"name=\"orderInfo\"><inputtype=\"hidden\"name=\"hisSchemeId\"value=\"\"><inputtype=\"hidden\"name=\"schemeId\"value=\"12829130\"><inputtype=\"hidden\"name=\"orderDate\"value=\"20190101\"><inputtype=\"hidden\"name=\"hosId\"value=\"057166\"><inputtype=\"hidden\"name=\"hosName\"value=\"杭州市西溪医院(市六医院)\"><inputtype=\"hidden\"name=\"deptId\"value=\"10238\"><inputtype=\"hidden\"name=\"deptName\"value=\"呼吸内科\"><inputtype=\"hidden\"name=\"docTitle\"value=\"\"><inputtype=\"hidden\"name=\"docId\"value=\"\"><inputtype=\"hidden\"name=\"docName\"value=\"普通\"><inputtype=\"hidden\"name=\"regFee\"value=\"10\"><inputtype=\"hidden\"name=\"takeNumAddr\"value=\"\"><inputtype=\"hidden\"name=\"resTimeSign\"value=\"1\"><inputtype=\"submit\"class=\"btnyy\"value=\"预约11\"title=\"总放号数11人次,剩余11人次,诊金10元\"alt=\"总放号数11人次,剩余11人次,诊金10元\"></form></td><tddata-idx=\"11\"data-type=\"per\"><formaction=\"/order/num\"method=\"get\"name=\"orderInfo\"><inputtype=\"hidden\"name=\"hisSchemeId\"value=\"\"><inputtype=\"hidden\"name=\"schemeId\"value=\"12837447\"><inputtype=\"hidden\"name=\"orderDate\"value=\"20190102\"><inputtype=\"hidden\"name=\"hosId\"value=\"057166\"><inputtype=\"hidden\"name=\"hosName\"value=\"杭州市西溪医院(市六医院)\"><inputtype=\"hidden\"name=\"deptId\"value=\"10238\"><inputtype=\"hidden\"name=\"deptName\"value=\"呼吸内科\"><inputtype=\"hidden\"name=\"docTitle\"value=\"\"><inputtype=\"hidden\"name=\"docId\"value=\"\"><inputtype=\"hidden\"name=\"docName\"value=\"普通\"><inputtype=\"hidden\"name=\"regFee\"value=\"10\"><inputtype=\"hidden\"name=\"takeNumAddr\"value=\"\"><inputtype=\"hidden\"name=\"resTimeSign\"value=\"0\"><inputtype=\"submit\"class=\"btnyy\"value=\"预约12\"title=\"总放号数12人次,剩余12人次,诊金10元\"alt=\"总放号数12人次,剩余12人次,诊金10元\"></form></td><tddata-idx=\"12\"data-type=\"per\"><formaction=\"/order/num\"method=\"get\"name=\"orderInfo\"><inputtype=\"hidden\"name=\"hisSchemeId\"value=\"\"><inputtype=\"hidden\"name=\"schemeId\"value=\"12837447\"><inputtype=\"hidden\"name=\"orderDate\"value=\"20190102\"><inputtype=\"hidden\"name=\"hosId\"value=\"057166\"><inputtype=\"hidden\"name=\"hosName\"value=\"杭州市西溪医院(市六医院)\"><inputtype=\"hidden\"name=\"deptId\"value=\"10238\"><inputtype=\"hidden\"name=\"deptName\"value=\"呼吸内科\"><inputtype=\"hidden\"name=\"docTitle\"value=\"\"><inputtype=\"hidden\"name=\"docId\"value=\"\"><inputtype=\"hidden\"name=\"docName\"value=\"普通\"><inputtype=\"hidden\"name=\"regFee\"value=\"10\"><inputtype=\"hidden\"name=\"takeNumAddr\"value=\"\"><inputtype=\"hidden\"name=\"resTimeSign\"value=\"1\"><inputtype=\"submit\"class=\"btnyy\"value=\"预约11\"title=\"总放号数11人次,剩余11人次,诊金10元\"alt=\"总放号数11人次,剩余11人次,诊 金10元\"></form></td><tddata-idx=\"13\"data-type=\"per\"><formaction=\"/order/num\"method=\"get\"name=\"orderInfo\"><inputtype=\"hidden\"name=\"hisSchemeId\"value=\"\"><inputtype=\"hidden\"name=\"schemeId\"value=\"12852345\"><inputtype=\"hidden\"name=\"orderDate\"value=\"20190103\"><inputtype=\"hidden\"name=\"hosId\"value=\"057166\"><inputtype=\"hidden\"name=\"hosName\"value=\"杭州市西溪医院(市六 医院)\"><inputtype=\"hidden\"name=\"deptId\"value=\"10238\"><inputtype=\"hidden\"name=\"deptName\"value=\"呼吸内科\"><inputtype=\"hidden\"name=\"docTitle\"value=\"\"><inputtype=\"hidden\"name=\"docId\"value=\"\"><inputtype=\"hidden\"name=\"docName\"value=\"普通\"><inputtype=\"hidden\"name=\"regFee\"value=\"10\"><inputtype=\"hidden\"name=\"takeNumAddr\"value=\"\"><inputtype=\"hidden\"name=\"resTimeSign\"value=\"0\"><inputtype=\"submit\"class=\"btnyy\"value=\"预约12\"title=\"总放号数12人次,剩余12人次,诊金10元\"alt=\"总放号数12人次,剩 余12人次,诊金10元\"></form></td><tddata-idx=\"14\"data-type=\"per\"><formaction=\"/order/num\"method=\"get\"name=\"orderInfo\"><inputtype=\"hidden\"name=\"hisSchemeId\"value=\"\"><inputtype=\"hidden\"name=\"schemeId\"value=\"12852345\"><inputtype=\"hidden\"name=\"orderDate\"value=\"20190103\"><inputtype=\"hidden\"name=\"hosId\"value=\"057166\"><inputtype=\"hidden\"name=\"hosName\"value=\"杭州市西 溪医院(市六医院)\"><inputtype=\"hidden\"name=\"deptId\"value=\"10238\"><inputtype=\"hidden\"name=\"deptName\"value=\"呼吸内科\"><inputtype=\"hidden\"name=\"docTitle\"value=\"\"><inputtype=\"hidden\"name=\"docId\"value=\"\"><inputtype=\"hidden\"name=\"docName\"value=\"普通\"><inputtype=\"hidden\"name=\"regFee\"value=\"10\"><inputtype=\"hidden\"name=\"takeNumAddr\"value=\"\"><inputtype=\"hidden\"name=\"resTimeSign\"value=\"1\"><inputtype=\"submit\"class=\"btnyy\"value=\"预约12\"title=\"总放号数12人次,剩余12人次,诊金10元\"alt=\"总放号 数12人次,剩余12人次,诊金10元\"></form></td><tddata-idx=\"15\"data-type=\"per\"><formaction=\"/order/num\"method=\"get\"name=\"orderInfo\"><inputtype=\"hidden\"name=\"hisSchemeId\"value=\"\"><inputtype=\"hidden\"name=\"schemeId\"value=\"12858729\"><inputtype=\"hidden\"name=\"orderDate\"value=\"20190104\"><inputtype=\"hidden\"name=\"hosId\"value=\"057166\"><inputtype=\"hidden\"name=\"hosName\"value=\"杭州市西溪医院(市六医院)\"><inputtype=\"hidden\"name=\"deptId\"value=\"10238\"><inputtype=\"hidden\"name=\"deptName\"value=\"呼吸 内科\"><inputtype=\"hidden\"name=\"docTitle\"value=\"\"><inputtype=\"hidden\"name=\"docId\"value=\"\"><inputtype=\"hidden\"name=\"docName\"value=\"普通\"><inputtype=\"hidden\"name=\"regFee\"value=\"10\"><inputtype=\"hidden\"name=\"takeNumAddr\"value=\"\"><inputtype=\"hidden\"name=\"resTimeSign\"value=\"0\"><inputtype=\"submit\"class=\"btnyy\"value=\"预约12\"title=\"总放号数12人次,剩余12人次,诊金10元\"alt=\"总放号数12人次,剩余12人次,诊金10元\"></form></td><tddata-idx=\"16\"data-type=\"per\"><formaction=\"/order/num\"method=\"get\"name=\"orderInfo\"><inputtype=\"hidden\"name=\"hisSchemeId\"value=\"\"><inputtype=\"hidden\"name=\"schemeId\"value=\"12858729\"><inputtype=\"hidden\"name=\"orderDate\"value=\"20190104\"><inputtype=\"hidden\"name=\"hosId\"value=\"057166\"><inputtype=\"hidden\"name=\"hosName\"value=\"杭州市西溪医院(市六医院)\"><inputtype=\"hidden\"name=\"deptId\"value=\"10238\"><inputtype=\"hidden\"name=\"deptName\"value=\"呼吸内科\"><inputtype=\"hidden\"name=\"docTitle\"value=\"\"><inputtype=\"hidden\"name=\"docId\"value=\"\"><inputtype=\"hidden\"name=\"docName\"value=\"普通\"><inputtype=\"hidden\"name=\"regFee\"value=\"10\"><inputtype=\"hidden\"name=\"takeNumAddr\"value=\"\"><inputtype=\"hidden\"name=\"resTimeSign\"value=\"1\"><inputtype=\"submit\"class=\"btnyy\"value=\"预约9\"title=\"总放号数9人次,剩余9人次,诊金10元\"alt=\"总放号数9人次,剩余9人次,诊金10元\"></form></td></tr>"
								 
	reg := regexp.MustCompile("\\s+")
    resString = reg.ReplaceAllString(resString, "")
    fmt.Println(resString)

	// resString = strings.Replace(resString, " ", "", -1)
	// resString = strings.Replace(resString, "\n", "", -1)
	// resString = strings.Replace(resString, "\n", "", -1)

	//获取预约列表
	// match := regexp.MustCompile(`data-idx="(?P<idx>\d+)(.*)name="schemeId"value="(?P<schemeId>\d+)">(.*)name="orderDate"value="(?P<orderDate>\d+)">`)
	match := regexp.MustCompile(`data-idx="(?P<idx>\d+)"data-type="per"><formaction="/order/num"method="get"name="orderInfo"><inputtype="hidden"name="hisSchemeId"value=""><inputtype="hidden"name="schemeId"value="(?P<schemeId>\d+)"><inputtype="hidden"name="orderDate"value="(?P<orderDate>\d+)"><inputtype="hidden"name="hosId"value="(?P<hosId>\d+)"><inputtype="hidden"name="hosName"value="(?P<hosName>[\p{Han}|(|)]+)"><inputtype="hidden"name="deptId"value="(?P<deptId>\d+)"><inputtype="hidden"name="deptName"value="(?P<deptName>[\p{Han}]+)"><inputtype="hidden"name="docTitle"value="(?P<docTitle>[\p{Han}]{0,})"><inputtype="hidden"name="docId"value="(?P<docId>\d{0,})"><inputtype="hidden"name="docName"value="(?P<docName>[\p{Han}]{0,})"><inputtype="hidden"name="regFee"value="(?P<regFee>\d+)"><inputtype="hidden"name="takeNumAddr"value="(?P<takeNumAddr>\d{0,})"><inputtype="hidden"name="resTimeSign"value="(?P<resTimeSign>\d{0,})"><inputtype="submit"class="btnyy"value="&#13;&#10;预约&#13;&#10;\d+"title="总放号数(?P<totalNum>\d+)人次,剩余(?P<remainNum>\d+)`)
	matchArr := match.FindAllStringSubmatch(resString, -1)
	fmt.Println(len(matchArr))

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

	for i, v := range result {
		if (v["orderDate"] == Params.RegisterObj.OrderDate) {
			
		}
	}

	//结果格式化输出
	// prettyResult, _ := json.MarshalIndent(result, "", "  ")
	// fmt.Println(string(prettyResult))
	return
}

func (Params *Params) setNormalNumEtc() {

}

func (Params *Params) setDocSpecificNumEtc() {
	
}
