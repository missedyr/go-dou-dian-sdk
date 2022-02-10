package douDianSdk

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
)

func (cl *DouDianClient) OrderSettleDetailByOrderId(orderId string, printLog bool) ([]SOrderSettleDetail, error) {
	params := SOrderSettleDetailParams{OrderId: orderId}
	return cl.OrderSettleDetail(params, printLog)
}

func (cl *DouDianClient) OrderSettleDetail(params SOrderSettleDetailParams, printLog bool) ([]SOrderSettleDetail, error) {
	var list []SOrderSettleDetail
	// 参数验证
	if len(params.OrderId) == 0 && (len(params.StartTime) == 0 || len(params.EndTime) == 0) {
		errMsg := "OrderSettleDetail--params--err: orderId or (StartTime || EndTime) at least one cannot be empty"
		logrus.Errorf(errMsg)
		return list, errors.New(errMsg)
	}

	paramJson := map[string]interface{}{}
	paramsM, _ := json.Marshal(params)
	_ = json.Unmarshal(paramsM, &paramJson)
	method := "order.getSettleBillDetailV2"
	respBodyData, err := cl.FetchSignAndToken(method, paramJson, "get")

	type RespBody struct {
		Code    int64                `json:"code"`
		CodeMsg string               `json:"code_msg"`
		Data    []SOrderSettleDetail `json:"data"`
	}
	var body RespBody
	_ = json.Unmarshal(respBodyData, &body)
	if printLog {
		dataJson, _ := json.Marshal(body.Data)
		logrus.Infof("douDianSdk-->OrderSettleDetail -- list:%v", string(dataJson))
	}
	logrus.Infof("douDianSdk-->OrderSettleDetail -- code:%d -- codeMsg: %s -- dataListCount: %d", body.Code, body.CodeMsg, len(body.Data))
	return body.Data, err
}
