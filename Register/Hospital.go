package Register

import(
	"../HttpCurl"
	"fmt"
	"encoding/json"
	"../Common"
)

type Hospital struct {
	areaList []map[string]string
	hospitalList map[string][]map[string]string
}

//从缓存文件中获取医院名称对应的医院id
func (Hospital *Hospital) GetHospitalByName(hosName string) {
	

}

//保存医院相关的信息
func (Hospital *Hospital) SaveHospitalInfo() {
	//获取所有的地区列表
	Hospital.getAreaList()

	//根据地区列表获取每个地区的医院
	Hospital.getHospitalList()

	//根据医院列表获取科室列表
	Hospital.getDeptList()

}

func (Hospital *Hospital) getAreaList() {
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

func (Hospital *Hospital) getHospitalList() {
	httpCurl := &HttpCurl.HttpCurl{}
	url := "http://www.zj12580.cn/hos/list/%s"

	var areaCode string
	res := make(map[string]map[string][]map[string]string)
	Hospital.hospitalList = make(map[string][]map[string]string)
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

func (Hospital *Hospital) getDeptList() {
	
}








