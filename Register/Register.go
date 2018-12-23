package Register

type Register struct {
	//用户身份证号
	Usercardno string
	//用户登陆密码
	Password string
	//地区
	Area string
	//医院名称
	HosName string
	//科室名称
	DeptName string
	//医生名称
	DocName string
	//预约日期
	OrderDate string
	//预约时段 (上午am或下午pm)
	OrderPeriod string
	//预约最早时间
	OrderStime string
	//预约最晚时间
	OrderEtime string

	//信息确认页url
	checkUrl string
	//预约url
	registerUrl string
}

func (Register *Register) init() {	
	Register.registerUrl = "http://www.zj12580.cn/order/save?yzmType=6&code="

	//参数校验
	Params := &Params{}
	Params.RegisterObj = Register
	if (!Params.CheckInfo()) {
		panic("register param is not correct")
	}

	//参数初始化
	Params.SetDefaultValues()
}

func (Register *Register) Register() {
	Register.init()
}

