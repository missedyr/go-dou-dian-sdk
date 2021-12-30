package douDianSdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (cl *DouDianClient) FetchSignAndToken(method string, paramJson map[string]interface{}, httpMethod string) ([]byte, error) {
	// 序列化参数
	paramJsonStr := cl.SignMarshal(paramJson)

	// 计算签名
	timestamp := time.Now().Unix()
	signVal := cl.Sign(cl.AppKey, cl.AppSecret, method, timestamp, paramJsonStr)

	// 获取 accessToken
	accessTokenData, err := cl.GetAccessToken()
	if err != nil || len(accessTokenData.AccessToken) == 0 {
		logrus.Warnf("douDianSdk-->openHttp-->FetchSignAndToken 请求 method:%s 前置获取accessToken 失败")
	}

	// 执行API调用
	return cl.Fetch(method, accessTokenData.AccessToken, signVal, paramJsonStr, httpMethod, timestamp)
}

// Fetch 调用Open Api，取回数据
func (cl *DouDianClient) Fetch(method, accessToken, sign, paramJson, httpMethod string, timestamp int64) ([]byte, error) {
	methodPath := strings.Replace(method, ".", "/", -1)

	u := "https://openapi-fxg.jinritemai.com/" + methodPath
	headers := map[string]string{
		"Accept":       "*/*",
		"Content-Type": "application/json;charset=UTF-8",
	}
	params := map[string]string{
		"v":            "2",
		"method":       method,
		"app_key":      cl.AppKey,
		"access_token": accessToken,
		"timestamp":    strconv.FormatInt(timestamp, 10),
		"sign":         sign,
		"sign_method":  "hmac-sha256",
		"param_json":   paramJson,
	}

	client := resty.New()
	resp := &resty.Response{}
	var err error
	if httpMethod == "get" || httpMethod == "GET" {
		resp, err = client.R().SetHeaders(headers).SetQueryParams(params).EnableTrace().Get(u)
	} else {
		body := strings.NewReader(paramJson)
		resp, err = client.R().SetHeaders(headers).SetQueryParams(params).SetBody(body).EnableTrace().Post(u)
	}

	if err != nil || resp.StatusCode() != http.StatusOK || resp.Body() == nil {
		logrus.Warnf("douDianSdk-->openHttp-->Fetch Err url:%s, params: %s, resp: %s, errOrigin: %s", u, params, resp, err)
		return resp.Body(), err
	}

	type respBody struct {
		ErrNo   int64       `json:"err_no"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	var data respBody
	_ = json.Unmarshal(resp.Body(), &data)
	dataDataByte, _ := json.Marshal(data.Data)

	if data.ErrNo != 0 {
		errMsg := fmt.Sprintf("douDianSdk-->openHttp-->Fetch respBodyErr ErrNo:%s, message:%s", data.ErrNo, data.Message)
		logrus.Warn(errMsg)
		return dataDataByte, errors.New(errMsg)
	}

	return dataDataByte, nil
}
