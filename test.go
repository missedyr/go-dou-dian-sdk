package main

import (
	"github.com/sirupsen/logrus"
	douDianSdk "github.com/xuexin520/go-dou-dian-sdk/src"
)

func main() {
	appKey := ""
	appSecret := ""
	shopId := ""
	logrus.Infof("1111111")

	douDianClient := douDianSdk.New(appKey, appSecret, shopId)

	// token
	douDianClient.GetAccessToken()

	// 订单详情
	douDianClient.OrderDetail("4871962930647831908")
}
