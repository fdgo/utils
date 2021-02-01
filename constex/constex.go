package constex

import (
	"github.com/spf13/viper"
)

var (
	GloCfg = &viper.Viper{}
)

const REDIS_USER_TOKEN = "redis:user:token:"   //登陆token
const REDIS_USER_INFO = "redis:user:info:"     //用户信息
const REDIS_USER_CITY = "redis:user:city:"     //只保存某一个城市一个用户位置，作为该城市参考坐标
const REDIS_USER_CHINA = "redis:user:city:中国"  //全国所有用户位置信息
const REDIS_USER_CITYEX = "redis:user:cityex:" //某一个城市所有用户位置信息
//--------------------------------------------------------------------------------------------
const ALI_PAY_APPID = "aaaabbbbccccdddd" //支付宝appid  16位
//--------------------------------------------------------------------------------------------
const WXPAY_COMPANY_NUM = "aaaaabbbbb" //商户编号  10位
const WXPAY_UNIORDER_URL = "https://api.mch.weixin.qq.com/pay/unifiedorder"
const WXPAY_PAY_APPID = "111111222222333333"               //微信appid  18位
const WXPAY_API_KEY = "aaaaabbbbbcccccdddddeeeeefffff88"   //微信api-key  32位
const WXPAY_APPSECRET = "aaaaabbbbbcccccdddddeeeeefffff88" //微信app-secret    32位
const WXPAY_ASYNC_CALLBACK_URL = "http://xx.yy.zz.mm:8088/api/v1/user/walletpay-callback"
const WXPAY_ORDER_CHECK = "https://api.mch.weixin.qq.com/pay/orderquery"
const WXPAY_BILL_DOWNLOAD = "https://api.mch.weixin.qq.com/pay/downloadbill"
const WXPAY_FUNDFLOW_DOWNLOAD = "https://api.mch.weixin.qq.com/pay/downloadfundflow"

//---------------------------------------------------------------------------------------------
const IM_APPKEY = "aaaaabbbbbcccccdddddeeeeefffff88" //网页云信 app_key    32位
const IM_NONCE = "aaaaabbbbbcccccddddd"              //随机值   20位
const IM_APPSECRET = "666666888888"                  //网页云信 app_secret 12位

//---------------------------------------------------------------------------------------------
const ALI_interview/accessKEYID = ""     //阿里人脸识别 interview/access_key
const ALI_interview/accessKEYSECRET = "" //阿里人脸识别 interview/access_secret

//---------------------------------------------------------------------------------------------
const IDENTIFY_PHOABLUM_COUNT = 10            //人脸识别认证之后，每天能够查看不同女性的次数
const IDENTIFY_PHOABLUM_ENDTIME = ONEDAY * 10 //人脸识别认证之后，有效的天数

const VIP_PHOABLUM_COUNT = 10            //vip认证之后，每天能够查看不同女性的次数
const VIP_PHOABLUM_ENDTIME = ONEDAY * 10 //vip认证之后，有效的天数

//单次支付业务不是很确定
const PAY_PHOABLUM_COUNT = 10
const PAY_PHOABLUM_EDNTIME = ONEDAY * 10

//普通用户每天可以免费查看女性的次数
const NORMAL_MAINPAGE_COUNT = 15

//有的女性主页设置了（要看其主页必须上传照片），上传完照片之后能看多少天
const NORMAL_MAINPAGE_ENDTIME = ONEDAY * 10

//---------------------------------------------------------------------------------------------
//[关于公园里的用户之间的距离]
//该模块基于redis专门针对距离的算法，它的思想是，根据同一个redis键，存入的值为一批用户信息(经纬度)，这样我们就可以算出任何两人之间的距离.
//另一种重要用法是以某一个人经纬度为参考点，方圆多少公里范围内查找一批用户。
const DEFAULTDISTANCE = 5000 //全国范围内，5000Km范围内
const DEFAULTNEARBY = 10000  //多少人,即人数上限
//---------------------------------------------------------------------------------------------
const HALFHOUR = 1800             //半小时
const ONEHOUR = HALFHOUR * 2      //一个小时
const FIVEHOUR = ONEHOUR * 5      //五个小时
const ONEDAY = ONEHOUR * 24       //24小时
const THREEDAY = ONEHOUR * 24 * 3 //三天

const VFCODE_TEXT = "【company网络】您的验证码是"
const VFCODE_APIKEY = "xxxx"
const VFCODE_MOBILE = "redis:vfcode:mobile:"

const AUTH_VIP_STR = "vip认证"
const AUTH_VIP_NUM = 5
const AUTH_GODDNESS_STR = "女人认证"
const AUTH_GODDNESS_NUM = 6
const AUTH_IDENTIFY_STR = "真人认证"
const AUTH_IDENTIFY_NUM = 7

const REDPACKET_STATUS_OPEN = "open"
const REDPACKET_STATUS_NOTOPEN = "not-open"

const WALLETPAY_BROCAST = 1 //支付用途类型   发布节目
const WALLETPAY_BROCAST_TITLE = "发布广播收费"

const WALLETPAY_PHO = 2 //支付用途类型   看别人照片
const WALLETPAY_PHO_TITLE = "查看相册照片收费"

const WALLETPAY_MIKE = 3 //支付用途类型   连麦
const WALLETPAY_MIKE_TITLE = "连麦收费"

const WALLETPAY_PHOABLUM = 4 //支付用途类型   解锁相册
const WALLETPAY_PHOABLUM_TITLE = "解锁相册收费"

const WALLETPAY_VIP = 5 //支付用途类型   vip
const WALLETPAY_VIP_TITLE = "Vip认证收费"

const WALLETPAY_PHOABLUM_LASTDAY = ONEDAY * 7 //支付解锁相册   有效天数

func GCfg(path string) error { //配置文件
	GloCfg = viper.New()
	//设置配置文件的名字
	GloCfg.SetConfigName("dev")
	//添加配置文件所在的路径,注意在Linux环境下%GOPATH要替换为$GOPATH
	GloCfg.AddConfigPath(path)
	//设置配置文件类型
	GloCfg.SetConfigType("yaml")
	if err := GloCfg.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
