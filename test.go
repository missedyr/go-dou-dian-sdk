package test

import (
	"github.com/sirupsen/logrus"
	douDianSdk "github.com/xuexin520/go-dou-dian-sdk/src"
)

func test() {
	logrus.Infof("1111111")
	appKey := ""
	appSecret := ""
	shopId := ""
	douDianClient := douDianSdk.New(appKey, appSecret, shopId)

	// token
	//douDianClient.GetAccessToken()

	// s商品详情
	// douDianClient.GoodsDetail("", "0", false)

	// 订单详情
	//douDianClient.OrderDetail("4871962930647831908")

	// 订单结算信息
	douDianClient.OrderSettleDetailByOrderId("", true)

}
