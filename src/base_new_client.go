package douDianSdk

import "github.com/sirupsen/logrus"

// DouDianClient 定义DouDianClient对象
type DouDianClient struct {
	AppKey    string // 应用key
	AppSecret string // 应用秘钥
	ShopId    string // 店铺ID，仅自用型应用有效； 若不传，则默认返回最早授权成功店铺对应的token信息
}

// New 创建一个新的操作对象
func New(appKey string, appSecret string, shopId string) *DouDianClient {
	if appKey == "" {
		logrus.Infof(`appKey can not be empty`)
		return nil
	}
	if appSecret == "" {
		logrus.Infof(`appSecret can not be empty`)
		return nil
	}

	return &DouDianClient{
		AppKey:    appKey,
		AppSecret: appSecret,
		ShopId:    shopId,
	}
}
