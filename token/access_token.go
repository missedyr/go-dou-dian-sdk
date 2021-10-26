package douDianSdk

import (
	"encoding/json"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/silenceper/log"
	"github.com/xuexin520/go-dou-dian-sdk/openHttp"
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
	}
	// 序列化参数
	paramJson := douDianSdk.Marshal(body)
	fmt.Println("param_json:" + paramJson)

	// 计算签名
	signVal := douDianSdk.Sign(appKey, appSecret, method, timestamp, paramJson)
	fmt.Println("sign_val:" + signVal)
	respBody, _ := openHttp.Fetch(method, timestamp, appKey, "", signVal, paramJson, "get")

	var data SAccessTokenData
	_ = json.Unmarshal(respBody, &data)
	log.Infof("douDianSdk-->GetAccessToken data:%s", data)

	return data, nil
}
