package pay

import (
	"bytes"
	"encoding/xml"
	"errors"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lithammer/shortuuid"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"interview/baiy/model/mysql_mdl"
	"interview/baiy/model/param_in"
	"interview/baiy/support/utils/constex"
	"interview/baiy/support/utils/stringex"
	"interview/baiy/support/utils/timex"
	"net/http"
	"strconv"
	"strings"
)

//首先定义一个UnifyOrderReq用于填入我们要传入的参数。
type UnifyOrderReq struct {
	Appid            string `xml:"appid"`
	Body             string `xml:"body"`
	Mch_id           string `xml:"mch_id"`
	Nonce_str        string `xml:"nonce_str"`
	Notify_url       string `xml:"notify_url"`
	Trade_type       string `xml:"trade_type"`
	Spbill_create_ip string `xml:"spbill_create_ip"`
	Total_fee        int    `xml:"total_fee"`
	Out_trade_no     string `xml:"out_trade_no"`
	Sign             string `xml:"sign"`
}
type UnifyOrderCheck struct {
	Return_msg   string `xml:"return_msg"`
	Appid        string `xml:"appid"`
	Mch_id       string `xml:"mch_id"`
	Nonce_str    string `xml:"nonce_str"`
	Out_trade_no string `xml:"out_trade_no"`
	Sign         string `xml:"sign"`
	Result_code  string `xml:"result_code"`
}
type DownLoadBill struct {
	Return_msg  string `xml:"return_msg"`
	Appid       string `xml:"appid"`
	BillDate    string `xml:"bill_date"`
	BillType    string `xml:"bill_type"`
	Mch_id      string `xml:"mch_id"`
	Nonce_str   string `xml:"nonce_str"`
	Sign        string `xml:"sign"`
	Result_code string `xml:"result_code"`
}
type Downloadfundflow struct {
	Return_msg  string `xml:"return_msg"`
	Appid       string `xml:"appid"`
	BillDate    string `xml:"bill_date"`
	AccountType string `xml:"account_type"`
	Mch_id      string `xml:"mch_id"`
	Nonce_str   string `xml:"nonce_str"`
	Sign        string `xml:"sign"`
	Result_code string `xml:"result_code"`
}
type UnifyOrderResp struct {
	Return_code string `xml:"return_code"`
	Return_msg  string `xml:"return_msg"`
	Appid       string `xml:"appid"`
	Mch_id      string `xml:"mch_id"`
	Nonce_str   string `xml:"nonce_str"`
	Sign        string `xml:"sign"`
	Result_code string `xml:"result_code"`
	Prepay_id   string `xml:"prepay_id"`
	Trade_type  string `xml:"trade_type"`
}

func UnifiedOrder(tx *gorm.DB, title string, user_self string, in *param_in.WalletPayIn, product_id string) (error, string, string, string, string, string, string, string, string) { //统一下单
	//请求UnifiedOrder的代码
	Out_trade_no := shortuuid.New()
	var order mysql_mdl.TbOrder
	order.OrderID = uuid.New().String()
	order.ConsumeTime = timex.GetCurrentTime()
	order.TradeNo = Out_trade_no
	order.Status = 0
	order.ConsumeType = in.ConsumeType
	order.Title = title
	order.UseradvID = user_self
	err := tx.Create(&order).Error
	if err != nil {
		tx.Rollback()
		return errors.New("创建订单记录失败!"), "", "", "", "", "", "", "", ""
	}
	money, _ := strconv.Atoi(in.Money)
	logrus.Debug("真正的金额为:", money)
	Nonce_str := stringex.GetRandomString(12)
	var yourReq UnifyOrderReq
	yourReq.Appid = constex.WXPAY_PAY_APPID //微信开放平台我们创建出来的app的app id
	yourReq.Body = title
	yourReq.Mch_id = product_id
	yourReq.Nonce_str = Nonce_str
	yourReq.Notify_url = constex.WXPAY_ASYNC_CALLBACK_URL
	yourReq.Trade_type = "APP"
	yourReq.Spbill_create_ip = "124.236.24.85"
	yourReq.Total_fee = 1 //单位是分，这里是1分
	yourReq.Out_trade_no = Out_trade_no
	logrus.Debug("Out_trade_no为:", Out_trade_no)
	var m map[string]interface{}
	m = make(map[string]interface{}, 0)
	m["appid"] = yourReq.Appid
	m["body"] = yourReq.Body
	m["mch_id"] = yourReq.Mch_id
	m["notify_url"] = yourReq.Notify_url
	m["trade_type"] = yourReq.Trade_type
	m["spbill_create_ip"] = yourReq.Spbill_create_ip
	m["total_fee"] = yourReq.Total_fee
	m["out_trade_no"] = yourReq.Out_trade_no
	m["nonce_str"] = yourReq.Nonce_str
	yourReq.Sign = WxpayCalcSign(m, constex.WXPAY_API_KEY) //这个是计算wxpay签名的函数上面已贴出
	bytes_req, err := xml.Marshal(yourReq)
	if err != nil {
		tx.Rollback()
		return errors.New("以xml形式编码发送错误, 原因:" + err.Error()), "", "", "", "", "", "", "", ""
	}
	str_req := string(bytes_req)
	//wxpay的unifiedorder接口需要http body中xmldoc的根节点是<xml></xml>这种，所以这里需要replace一下
	str_req = strings.Replace(str_req, "UnifyOrderReq", "xml", -1)
	bytes_req = []byte(str_req)

	//发送unified order请求.
	req, err := http.NewRequest("POST", constex.WXPAY_UNIORDER_URL, bytes.NewReader(bytes_req))
	if err != nil {
		tx.Rollback()
		return errors.New("New Http Request发生错误，原因:" + err.Error()), "", "", "", "", "", "", "", ""
	}
	req.Header.Set("Accept", "application/xml")
	//这里的http header的设置是必须设置的.
	req.Header.Set("Content-Type", "application/xml;charset=utf-8")
	c := http.Client{}
	resp, _err := c.Do(req)
	if _err != nil {
		tx.Rollback()
		return errors.New("请求微信支付统一下单接口发送错误, 原因:" + _err.Error()), "", "", "", "", "", "", "", ""
	}
	tmpresp, _ := ioutil.ReadAll(resp.Body)
	xmlResp := UnifyOrderResp{}
	_err = xml.Unmarshal(tmpresp, &xmlResp)
	//处理return code.
	if xmlResp.Return_code == "FAIL" {
		err = tx.Model(&mysql_mdl.TbOrder{}).Where("trade_no=? and useradv_id=?", Out_trade_no, user_self).Update("status", -1).Error
		if err != nil {
			tx.Rollback()
			return err, "", "", "", "", "", "", "", ""
		}
		err = tx.Commit().Error
		if err != nil {
			tx.Rollback()
			return err, "", "", "", "", "", "", "", ""
		}
		return errors.New("微信支付统一下单不成功，原因:" + xmlResp.Return_msg), "", "", "", "", "", "", "", ""
	}
	err = tx.Model(&mysql_mdl.TbOrder{}).Where("trade_no=? and useradv_id=?", Out_trade_no, user_self).Update("status", 1).Error
	if err != nil {
		tx.Rollback()
		return err, "", "", "", "", "", "", "", ""
	}
	//这里已经得到微信支付的prepay id，需要返给客户端，由客户端继续完成支付流程
	logrus.Debug("微信支付统一下单成功，总消息:", xmlResp)
	t := strconv.FormatInt(timex.GetCurrentTimeStamp(), 10)
	mf := make(map[string]interface{}, 0)
	mf["appid"] = yourReq.Appid
	mf["partnerid"] = constex.WXPAY_COMPANY_NUM
	mf["prepayid"] = xmlResp.Prepay_id
	mf["package"] = "Sign=WXPay"
	mf["noncestr"] = yourReq.Nonce_str
	mf["timestamp"] = t
	sign := WxpayCalcSign(mf, constex.WXPAY_API_KEY)
	return nil, yourReq.Appid, constex.WXPAY_COMPANY_NUM, xmlResp.Prepay_id, Nonce_str, t, "Sign=WXPay", sign, Out_trade_no
}
