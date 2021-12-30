package douDianSdk

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

// OrderDetail 获取订单详情
// orderIdParent 	父订单号
func (cl *DouDianClient) OrderDetail(orderIdParent string) (SDYOrderInfo, error) {
	paramJson := map[string]interface{}{
		"shop_order_id": orderIdParent,
	}
	method := "order.orderDetail"
	respBodyData, err := cl.FetchSignAndToken(method, paramJson, "get")
	type sRespBodyD struct {
		ShopOrderDetail SDYOrderInfo `json:"shop_order_detail"`
	}

	var data sRespBodyD
	_ = json.Unmarshal(respBodyData, &data)
	dataJson, _ := json.Marshal(data)
	logrus.Infof("douDianSdk-->OrderDetail data:%v", string(dataJson))
	return data.ShopOrderDetail, err
}
