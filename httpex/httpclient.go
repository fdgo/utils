package httpex

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"interview/baiy/support/utils/constex"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout: 3 * time.Minute,
			//MaxConnsPerHost: 10000,
			TLSHandshakeTimeout: 10 * time.Second,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 10 * time.Minute,
				DualStack: true,
			}).DialContext,
		},
	}
}

//获取url对应的内容，返回信息：StatusCode，body，err
func Get(requestUrl string) (int, string, error) {
	reqest, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return 0, "", err
	}
	response, err := client.Do(reqest)
	if err != nil {
		return 0, "", err
	}
	defer response.Body.Close()
	if err != nil {
		return 0, "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, "", err
	}
	return response.StatusCode, string(body), nil
}

// 带上Bearer Token，发起一个get请求
func GetByToken(requestUrl, token string) (int, string, error) {
	reqest, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return 0, "", err
	}
	reqest.Header.Set("Authorization", token)
	response, err := client.Do(reqest)
	if err != nil {
		return 0, "", err
	}
	defer response.Body.Close()
	if err != nil {
		return 0, "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, "", err
	}
	return response.StatusCode, string(body), nil
}

func GetByTokenHead(requestUrl, token string, city string, lat, lgt string) (int, string, error) {
	reqest, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return 0, "", err
	}
	reqest.Header.Set("Authorization", token)
	reqest.Header.Set("city", city)
	reqest.Header.Set("lat", lat)
	reqest.Header.Set("lgt", lgt)
	response, err := client.Do(reqest)
	if err != nil {
		return 0, "", err
	}
	defer response.Body.Close()
	if err != nil {
		return 0, "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, "", err
	}
	return response.StatusCode, string(body), nil
}

type StringBuilder struct {
	buf bytes.Buffer
}

func NewStringBuilder() *StringBuilder {
	return &StringBuilder{buf: bytes.Buffer{}}
}
func (this *StringBuilder) Append(obj interface{}) *StringBuilder {
	this.buf.WriteString(fmt.Sprintf("%v", obj))
	return this
}

func (this *StringBuilder) ToString() string {
	return this.buf.String()
}

//获取url和参数列表对应的完整请求url
func BuildRequestUrl(requestUrl string, params url.Values) string {
	if len(params) <= 0 {
		return requestUrl
	}
	data := NewStringBuilder()
	data.Append(requestUrl)
	if !strings.Contains(requestUrl, "?") {
		data.Append("?")
	}
	has_param := false
	for k, v := range params {
		if len(k) > 0 && len(v[0]) > 0 {
			if has_param {
				data.Append("&")
			}
			data.Append(k)
			data.Append("=")
			data.Append(url.QueryEscape(v[0]))
			has_param = true
		}
	}
	return data.ToString()
}

func TimeStrToTimeStamp() (nTimeStamp int64) {
	now := time.Now()
	year, mon, day := now.UTC().Date()
	hour, min, sec := now.UTC().Clock()
	utc := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", year, mon, day, hour, min, sec)
	timestamp, _ := time.Parse("2006-01-02 15:04:05", utc)
	return timestamp.Unix() //- 3600*8
}
func Sha1(message []byte) []byte {
	h := sha1.New()
	h.Write(message)
	return h.Sum(nil)
}
func RemarkImUser(requestUrl string, jsonString string) (int, string, error) {
	cutime := strconv.FormatInt(TimeStrToTimeStamp(), 10)
	req := bytes.NewBuffer([]byte(jsonString))
	reqest, err := http.NewRequest("POST", requestUrl, req)
	if err != nil {
		return 0, "", err
	}
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	reqest.Header.Set("AppKey", constex.IM_APPKEY)
	reqest.Header.Set("Nonce", constex.IM_NONCE)
	reqest.Header.Set("CurTime", cutime)
	reqest.Header.Set("CheckSum", hex.EncodeToString(Sha1([]byte(constex.IM_APPSECRET+constex.IM_NONCE+cutime))))

	response, err := client.Do(reqest)
	if err != nil {
		return 0, "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, "", err
	}
	return response.StatusCode, string(body), nil
}
func AddFriendImUser(requestUrl string, jsonString string) (int, string, error) {
	cutime := strconv.FormatInt(TimeStrToTimeStamp(), 10)
	req := bytes.NewBuffer([]byte(jsonString))
	reqest, err := http.NewRequest("POST", requestUrl, req)
	if err != nil {
		return 0, "", err
	}
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	reqest.Header.Set("AppKey", constex.IM_APPKEY)
	reqest.Header.Set("Nonce", constex.IM_NONCE)
	reqest.Header.Set("CurTime", cutime)
	reqest.Header.Set("CheckSum", hex.EncodeToString(Sha1([]byte(constex.IM_APPSECRET+constex.IM_NONCE+cutime))))

	response, err := client.Do(reqest)
	if err != nil {
		return 0, "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, "", err
	}
	return response.StatusCode, string(body), nil
}
func UpdateImUserInfo(requestUrl string, jsonString string) (int, string, error) {
	cutime := strconv.FormatInt(TimeStrToTimeStamp(), 10)
	req := bytes.NewBuffer([]byte(jsonString))
	reqest, err := http.NewRequest("POST", requestUrl, req)
	if err != nil {
		return 0, "", err
	}
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	reqest.Header.Set("AppKey", constex.IM_APPKEY)
	reqest.Header.Set("Nonce", constex.IM_NONCE)
	reqest.Header.Set("CurTime", cutime)
	reqest.Header.Set("CheckSum", hex.EncodeToString(Sha1([]byte(constex.IM_APPSECRET+constex.IM_NONCE+cutime))))

	response, err := client.Do(reqest)
	if err != nil {
		return 0, "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, "", err
	}
	return response.StatusCode, string(body), nil
}
func GetImUserInfo(requestUrl string, jsonString string) (int, string, error) {
	cutime := strconv.FormatInt(TimeStrToTimeStamp(), 10)
	req := bytes.NewBuffer([]byte(jsonString))
	reqest, err := http.NewRequest("POST", requestUrl, req)
	if err != nil {
		return 0, "", err
	}
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	reqest.Header.Set("AppKey", constex.IM_APPKEY)
	reqest.Header.Set("Nonce", constex.IM_NONCE)
	reqest.Header.Set("CurTime", cutime)
	reqest.Header.Set("CheckSum", hex.EncodeToString(Sha1([]byte(constex.IM_APPSECRET+constex.IM_NONCE+cutime))))

	response, err := client.Do(reqest)
	if err != nil {
		return 0, "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, "", err
	}
	return response.StatusCode, string(body), nil
}

func SetBlack(requestUrl string, jsonString string) (int, string, error) {
	cutime := strconv.FormatInt(TimeStrToTimeStamp(), 10)
	req := bytes.NewBuffer([]byte(jsonString))
	reqest, err := http.NewRequest("POST", requestUrl, req)
	if err != nil {
		return 0, "", err
	}
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	reqest.Header.Set("AppKey", constex.IM_APPKEY)
	reqest.Header.Set("Nonce", constex.IM_NONCE)
	reqest.Header.Set("CurTime", cutime)
	reqest.Header.Set("CheckSum", hex.EncodeToString(Sha1([]byte(constex.IM_APPSECRET+constex.IM_NONCE+cutime))))

	response, err := client.Do(reqest)
	if err != nil {
		return 0, "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, "", err
	}
	return response.StatusCode, string(body), nil
}

//获取url对应的内容，返回信息：StatusCode，body，err
func GenImConPwd(requestUrl string, jsonString string) (int, string, error) {
	cutime := strconv.FormatInt(TimeStrToTimeStamp(), 10)
	req := bytes.NewBuffer([]byte(jsonString))
	reqest, err := http.NewRequest("POST", requestUrl, req)
	if err != nil {
		return 0, "", err
	}
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	reqest.Header.Set("AppKey", constex.IM_APPKEY)
	reqest.Header.Set("Nonce", constex.IM_NONCE)
	reqest.Header.Set("CurTime", cutime)
	reqest.Header.Set("CheckSum", hex.EncodeToString(Sha1([]byte(constex.IM_APPSECRET+constex.IM_NONCE+cutime))))

	response, err := client.Do(reqest)
	if err != nil {
		return 0, "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, "", err
	}
	return response.StatusCode, string(body), nil
}
func SetMessage(requestUrl string, jsonString string) (int, string, error)  {
	cutime := strconv.FormatInt(TimeStrToTimeStamp(), 10)
	req := bytes.NewBuffer([]byte(jsonString))
	reqest, err := http.NewRequest("POST", requestUrl, req)
	if err != nil {
		return 0, "", err
	}
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	reqest.Header.Set("AppKey", constex.IM_APPKEY)
	reqest.Header.Set("Nonce", constex.IM_NONCE)
	reqest.Header.Set("CurTime", cutime)
	reqest.Header.Set("CheckSum", hex.EncodeToString(Sha1([]byte(constex.IM_APPSECRET+constex.IM_NONCE+cutime))))

	response, err := client.Do(reqest)
	if err != nil {
		return 0, "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, "", err
	}
	return response.StatusCode, string(body), nil
}
func ToJson(data interface{}) string {
	b, _ := json.Marshal(data)
	return string(b)
}

//用Post方法获取url对应的内容，提交json，返回信息：StatusCode，body，err
func PostJson(requestUrl string, params map[string]string) (int, string, error) {
	req := bytes.NewBuffer([]byte(ToJson(params)))
	reqest, err := http.NewRequest("POST", requestUrl, req)
	if err != nil {
		fmt.Println(err, "**********")
		return 0, "", err
	}
	reqest.Header.Set("Content-Type", "application/json")
	response, err := client.Do(reqest)
	if err != nil {
		return 0, "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, "", err
	}
	return response.StatusCode, string(body), nil
}

// 带上Bearer Token，发起一个post请求，内容是json
func PostJsonByToken(requestUrl, token, jsonString string) (int, string, error) {
	req := bytes.NewBuffer([]byte(jsonString))
	reqest, err := http.NewRequest("POST", requestUrl, req)
	if err != nil {
		return 0, "", err
	}
	reqest.Header.Set("Authorization", token)
	reqest.Header.Set("Content-Type", "application/json")
	response, err := client.Do(reqest)
	if err != nil {
		return 0, "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, "", err
	}
	return response.StatusCode, string(body), nil
}

//用Post方法获取url对应的内容，提交body，返回信息：StatusCode，body，err
func PostBody(requestUrl string, reqBody string) (int, string, error) {
	req := bytes.NewBuffer([]byte(reqBody))
	reqest, err := http.NewRequest("POST", requestUrl, req)
	if err != nil {
		return 0, "", err
	}
	reqest.Header.Set("Content-Type", "application/json")
	response, err := client.Do(reqest)
	if err != nil {
		return 0, "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, "", err
	}
	return response.StatusCode, string(body), nil
}

//获取url对应的内容,同时上传文件，返回信息：StatusCode，body，err
func PostFile(requestUrl string, params url.Values, field_name, path string) (int, string, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile(field_name, path)
	if err != nil {
		return 0, "", err
	}

	//打开文件句柄操作
	fh, err := os.Open(path)
	if err != nil {
		return 0, "", err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return 0, "", err
	}

	//写入表单数据
	if len(params) > 0 {
		for k, v := range params {
			bodyWriter.WriteField(k, v[0])
		}
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(requestUrl, contentType, bodyBuf)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, "", err
	}
	return resp.StatusCode, string(resp_body), nil
}
