package douDianSdk

import (
	"encoding/json"
	"github.com/silenceper/log"
	douDianSdk "github.com/xuexin520/go-dou-dian-sdk/public"
)

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
	log.Infof("douDianSdk-->OrderDetail data:%s", data)
	return data.ShopOrderDetail, err
}
