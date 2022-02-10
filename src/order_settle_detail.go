package douDianSdk

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
)

func (cl *DouDianClient) OrderSettleDetailByOrderId(orderId string, printLog bool) ([]SOrderSettleDetail, error) {
	params := SOrderSettleDetailParams{OrderId: orderId}
	data, err := cl.OrderSettleDetail(params, printLog)
	return data.Data, err
}

func (cl *DouDianClient) OrderSettleDetail(params SOrderSettleDetailParams, printLog bool) (SOrderSettleDetailData, error) {
	var data SOrderSettleDetailData
	// 参数验证
	if len(params.OrderId) == 0 && (len(params.StartTime) == 0 || len(params.EndTime) == 0) {
		errMsg := "OrderSettleDetail--params--err: orderId or (StartTime || EndTime) at least one cannot be empty"
		logrus.Errorf(errMsg)
		return data, errors.New(errMsg)
	}

	paramJson := map[string]interface{}{}
	paramsM, _ := json.Marshal(params)
	_ = json.Unmarshal(paramsM, &paramJson)
	method := "order.getSettleBillDetailV2"
	respBodyData, err := cl.FetchSignAndToken(method, paramJson, "get")

	_ = json.Unmarshal(respBodyData, &data)
	if printLog {
		dataJson, _ := json.Marshal(data)
		logrus.Infof("douDianSdk-->OrderSettleDetail -- body:%v", string(dataJson))
	}
	logrus.Infof("douDianSdk-->OrderSettleDetail -- code:%d -- codeMsg: %s -- dataListCount: %d", data.Code, data.CodeMsg, len(data.Data))
	return data, err
}
