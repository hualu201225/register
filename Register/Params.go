package Register

import(
	_"fmt"
)

type Params struct {
	RegisterObj *Register
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

	return true
}

func (Params *Params) SetDefaultValues() {
	//挂号时间段如果没传则赋予初始值
	Params.SetDefaultValue()

	//设置医院相关参数
	Params.SetHospitalValue()

	//设置科室相关参数
	Params.SetDeptValue()

	// //设置医生相关参数
	Params.SetDoctorValue()
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

//根据医院名称获取医院id
func (Params *Params) SetHospitalValue() {

}

//根据医院、科室名称获取科室id
func (Params *Params) SetDeptValue() {

}

//根据医院、科室、医生名称获取医生id
func (Parmas *Params) SetDoctorValue() {

}