package douDianSdk

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/silenceper/log"
	douDianSdk "github.com/xuexin520/go-dou-dian-sdk/sign"
	"time"
)

// https://op.jinritemai.com/docs/guide-docs/138/21
func GetAccessToken(appKey string, appSecret string, shopId string) (SAccessTokenData, error) {
	err := validation.Validate(appKey, validation.Required)
	err = validation.Validate(appSecret, validation.Required)

	if err != nil {
		log.Warnf("douDianSdk-->GetAccessToken validation err:", err)
		return SAccessTokenData{}, err
	}

	method := "token.create"
	timestamp := time.Now().Unix()
	body := map[string]interface{}{
		"app_id":     appKey,
		"app_secret": appSecret,
		"shop_id":    shopId,
		"grant_type": "authorization_self",
		"code":       "",
	}
	// 序列化参数
	paramJson := douDianSdk.Marshal(body)

	// 计算签名
	signVal := douDianSdk.Sign(appKey, appSecret, method, timestamp, paramJson)
	respBodyData, _ := Fetch(method, timestamp, appKey, "", signVal, paramJson, "get")

	var data SAccessTokenData
	_ = json.Unmarshal(respBodyData, &data)
	log.Infof("douDianSdk-->GetAccessToken data:%s", data)

	return data, nil
}
