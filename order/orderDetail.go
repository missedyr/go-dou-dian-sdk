package douDianSdk

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	douDianSdk "github.com/xuexin520/go-dou-dian-sdk/public"
)

// OrderDetail 获取订单详情
// appKey    		应用key
// appSecret 		应用秘钥
// shopId  			店铺ID
// orderIdParent 	父订单号
func OrderDetail(appKey string, appSecret string, shopId string, orderIdParent string) (SDYOrderInfo, error) {
	paramJson := map[string]interface{}{
		"shop_order_id": orderIdParent,
	}
	method := "order.orderDetail"
	respBodyData, err := douDianSdk.FetchSignAndToken(method, appKey, appSecret, shopId, paramJson, "get")
	type sRespBodyD struct {
		ShopOrderDetail SDYOrderInfo `json:"shop_order_detail"`
	}

	var data sRespBodyD
	_ = json.Unmarshal(respBodyData, &data)
	dataJson, _ := json.Marshal(data)
	logrus.Infof("douDianSdk-->OrderDetail data:%v", string(dataJson))
	return data.ShopOrderDetail, err
}
