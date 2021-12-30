package douDianSdk

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sirupsen/logrus"
	"time"
)

// GetAccessToken
// https://op.jinritemai.com/docs/guide-docs/138/21
func (cl *DouDianClient) GetAccessToken() (SAccessTokenData, error) {
	err := validation.Validate(cl.AppKey, validation.Required)
	err = validation.Validate(cl.AppSecret, validation.Required)

	if err != nil {
		logrus.Warnf("douDianSdk-->GetAccessToken validation err:%s", err)
		return SAccessTokenData{}, err
	}

	method := "token.create"
	timestamp := time.Now().Unix()
	body := map[string]interface{}{
		"app_id":     cl.AppKey,
		"app_secret": cl.AppSecret,
		"shop_id":    cl.ShopId,
		"grant_type": "authorization_self",
		"code":       "",
	}
	// 序列化参数
	paramJson := cl.SignMarshal(body)

	// 计算签名
	signVal := cl.Sign(cl.AppKey, cl.AppSecret, method, timestamp, paramJson)
	respBodyData, _ := cl.Fetch(method, "", signVal, paramJson, "get", timestamp)

	var data SAccessTokenData
	_ = json.Unmarshal(respBodyData, &data)
	dataByte, _ := json.Marshal(data)
	logrus.Infof("douDianSdk-->GetAccessToken data:%s", string(dataByte))

	return data, nil
}
