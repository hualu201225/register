package Register

import(
	"../HttpCurl"
	"fmt"
	"encoding/json"
	"../Common"
)

type Hospital struct {
	//所有地区列表
	areaList []map[string]string
	//所有医院列表
	hospitalList map[string][]map[string]string
	//一个医院的科室列表
	sinDeptList []map[string]interface{}
	//一个医院的医生列表
	sinDocList []map[string]interface{}

	CurrentAreaName string
	CurrentAreaId string
	CurrentHosName string
	CurrentHosId string
	CurrentDeptName string
	CurrentDeptId string
	CurrentDocName string
	CurrentDocId string
}

func (Hospital *Hospital) GetDetailInfo() {
	Hospital.GetAreaIdByName()
	Hospital.GetHosIdByName()
	Hospital.GetDeptIdByName()
	Hospital.GetDocIdByName()
}

//根据地区名称获取地区id
func (Hospital *Hospital) GetAreaIdByName() {
	if (len(Hospital.CurrentAreaName) == 0) {
		panic("area is needed")
	}
	CacheFile := &Common.CacheFile{}
	CacheFile.SetFileName("areaList")
	areaStr := CacheFile.GetContentFromFile()

	json.Unmarshal([]byte(areaStr), &Hospital.areaList)

	for _, v := range Hospital.areaList {
		if (Hospital.CurrentAreaName == v["areaName"]) {
			Hospital.CurrentAreaId = v["areaCode"]
			return
		}
	}

	panic("can not get area")
}

//根据地区id&医院名称获取医院id
func (Hospital *Hospital) GetHosIdByName(){
	if (len(Hospital.CurrentAreaId) == 0 || len(Hospital.CurrentHosName) == 0) {
		panic("hospital name is needed")
	}
	CacheFile := &Common.CacheFile{}
	CacheFile.SetFileName("hospitalList")
	areaStr := CacheFile.GetContentFromFile()

	json.Unmarshal([]byte(areaStr), &Hospital.hospitalList)

	if (len(Hospital.hospitalList[Hospital.CurrentAreaId]) == 0) {
		panic("can not find this hospital in this area")
	}

	for _, v := range Hospital.hospitalList[Hospital.CurrentAreaId] {
		if (Hospital.CurrentHosName == v["hosName"] || Hospital.CurrentHosName == v["aliasName"]) {
			Hospital.CurrentHosId = v["hosCode"]
			return
		}
	}

	panic("can not get hospital")
}

//根据医院id&科室名称获取科室id
func (Hospital *Hospital) GetDeptIdByName(){
	if (len(Hospital.CurrentHosId) == 0 || len(Hospital.CurrentDeptName) == 0) {
		panic("dept name is needed")
	}
	CacheFile := &Common.CacheFile{}
	CacheFile.SetFilePath("./Storage/Files/Depts/")
	CacheFile.SetFileName(Hospital.CurrentHosId)
	deptStr := CacheFile.GetContentFromFile()

	json.Unmarshal([]byte(deptStr), &Hospital.sinDeptList)

	for _, v := range Hospital.sinDeptList {
		if (Hospital.CurrentDeptName == v["deptName"]) {
			Hospital.CurrentDeptId = fmt.Sprintf("%v", v["deptId"])
			return
		}
	}

	panic("can not get dept")
}

//根据医院id&医生名字获取医生id
func (Hospital *Hospital) GetDocIdByName(){
	if (len(Hospital.CurrentHosId) == 0 || len(Hospital.CurrentDocName) == 0) {
		Hospital.CurrentDocName = "普通"
		Hospital.CurrentDocId = ""
		return
	}
	CacheFile := &Common.CacheFile{}
	CacheFile.SetFilePath("./Storage/Files/Docs/")
	CacheFile.SetFileName(Hospital.CurrentHosId)
	docStr := CacheFile.GetContentFromFile()

	json.Unmarshal([]byte(docStr), &Hospital.sinDocList)

	for _, v := range Hospital.sinDocList {
		if (Hospital.CurrentDocName == v["docName"]) {
			Hospital.CurrentDocId = fmt.Sprintf("%v", v["docId"])
			return
		}
	}

	panic("can not get doctor")
}

func (Hospital *Hospital) GetHospitals() map[string][]map[string]string {
	if (len(Hospital.hospitalList) > 0) {
		return Hospital.hospitalList
	}

	CacheFile := &Common.CacheFile{}
	CacheFile.SetFileName("hospitalList")
	str := CacheFile.GetContentFromFile()
	json.Unmarshal([]byte(str), &Hospital.hospitalList)

	return Hospital.hospitalList
}

//保存医院相关的信息
func (Hospital *Hospital) SaveHospitalInfo() {
	//获取所有的地区列表
	// Hospital.setAreaList()

	//根据地区列表获取每个地区的医院
	// Hospital.setHospitalList()

	//根据医院列表获取科室列表
	// Hospital.setDeptList()

	//根据医院列表获取医生列表
	Hospital.setDocList()
}

func (Hospital *Hospital) setAreaList() {
	httpCurl := &HttpCurl.HttpCurl{}
	httpCurl.SetUrl("http://www.zj12580.cn/area")

	str, _ := httpCurl.GetContentsFromUrl()

	res := make(map[string]map[string][]map[string]string)
	json.Unmarshal(str, &res)
	Hospital.areaList = res["data"]["areaList"]

	//保存地区列表
	areaStr, _ := json.Marshal(Hospital.areaList)
	CacheFile := &Common.CacheFile{}
	CacheFile.SetFileName("areaList")
	CacheFile.SetContentToFile(string(areaStr))
}

func (Hospital *Hospital) setHospitalList() {
	httpCurl := &HttpCurl.HttpCurl{}
	url := "http://www.zj12580.cn/hos/list/%s"

	var areaCode string
	res := make(map[string]map[string][]map[string]string)
	Hospital.hospitalList = Hospital.GetHospitals()
	for i, _ := range Hospital.areaList {
		areaCode = Hospital.areaList[i]["areaCode"]
		httpCurl.SetUrl(fmt.Sprintf(url, areaCode))
		str, _ := httpCurl.GetContentsFromUrl()	
		json.Unmarshal(str, &res)
		
		Hospital.hospitalList[areaCode] = res["data"]["hos"]
	}

	//保存医院列表
	hospitalStr, _ := json.Marshal(Hospital.hospitalList)
	CacheFile := &Common.CacheFile{}
	CacheFile.SetFileName("hospitalList")
	CacheFile.SetContentToFile(string(hospitalStr))
}

func (Hospital *Hospital) setDeptList() {
	httpCurl := &HttpCurl.HttpCurl{}
	CacheFile := &Common.CacheFile{}
	url := "http://www.zj12580.cn/dept/list/%s"
	res := make(map[string]map[string][]map[string]interface{})

	Hospital.hospitalList = Hospital.GetHospitals()

	var result []map[string]interface{}
	for _,v := range Hospital.hospitalList {
		for _,hospital := range v {
			httpCurl.SetUrl(fmt.Sprintf(url, hospital["hosCode"]))
			str, _ := httpCurl.GetContentsFromUrl()
			fmt.Printf(string(str))
			json.Unmarshal(str, &res)
			result = res["data"]["dept"] 

			//保存科室列表
			deptStr, _ := json.Marshal(result)
			CacheFile.SetFilePath("./Storage/Files/Depts")
			CacheFile.SetFileName(hospital["hosCode"])
			CacheFile.SetContentToFile(string(deptStr))
		}
	}
}

func (Hospital *Hospital) setDocList() {
	httpCurl := &HttpCurl.HttpCurl{}
	CacheFile := &Common.CacheFile{}
	url := "http://www.zj12580.cn/doc/list/%s"
	res := make(map[string]map[string][]map[string]interface{})

	Hospital.hospitalList = Hospital.GetHospitals()

	var result []map[string]interface{}
	for _,v := range Hospital.hospitalList {
		for _,hospital := range v {
			httpCurl.SetUrl(fmt.Sprintf(url, hospital["hosCode"]))
			str, _ := httpCurl.GetContentsFromUrl()
			
			json.Unmarshal(str, &res)
			result = res["data"]["docs"] 

			//保存科室列表
			deptStr, _ := json.Marshal(result)
			CacheFile.SetFilePath("./Storage/Files/Docs")
			CacheFile.SetFileName(hospital["hosCode"])
			CacheFile.SetContentToFile(string(deptStr))
		}
	}
}








