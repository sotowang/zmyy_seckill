package consts

import "zmyy_seckill/limit"

var RequestLimitRate limit.LimitRate

const (
	UserAgent  = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36 MicroMessenger/7.0.9.501 NetType/WIFI MiniProgramEnv/Windows WindowsWechat"
	Refer      = "https://servicewechat.com/wx2c7f0f3c30d99445/72/page-frame.html"
	Connection = "keep-alive"
	Host       = "https://cloud.cn2030.com"
	//某地区医院列表URL
	CustomerListUrl = "https://cloud.cn2030.com/sc/wx/HandlerSubscribe.ashx?act=CustomerList"
	//授权URL
	AuthUrl = "https://cloud.cn2030.com/sc/wx/HandlerSubscribe.ashx?act=auth&code=061H55000QOs8L1yHN100Ba0N43H550I"
	//某医院内HPV疫苗情况URL
	CustomerProductURL = "https://cloud.cn2030.com/sc/wx/HandlerSubscribe.ashx?act=CustomerProduct"
	//预约用户信息
	UserInfoURL          = "https://cloud.cn2030.com/sc/wx/HandlerSubscribe.ashx?act=User"
	CustSubscribeDateUrl = "https://cloud.cn2030.com/sc/wx/HandlerSubscribe.ashx?act=GetCustSubscribeDateAll"
	CaptchaVerifyUrl     = "https://cloud.cn2030.com/sc/wx/HandlerSubscribe.ashx?act=CaptchaVerify"
	GetCaptchaUrl        = "https://cloud.cn2030.com/sc/wx/HandlerSubscribe.ashx?act=GetCaptcha"
	SaveUrl              = "https://cloud.cn2030.com/sc/wx/HandlerSubscribe.ashx?act=Save20"
	//获取订单状态
	OrderStatusUrl             = "https://cloud.cn2030.com/sc/wx/HandlerSubscribe.ashx?act=GetOrderStatus"
	CustSubscribeDateDetailUrl = "https://cloud.cn2030.com/sc/wx/HandlerSubscribe.ashx?act=GetCustSubscribeDateDetail"
)
