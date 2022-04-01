package douDianSdk

// SAccessTokenData 获取抖店AccessToken 返回数据结构体
type SAccessTokenData struct {
	ShopId       int64  `json:"shop_id"`       // 店铺ID
	ShopName     string `json:"shop_name"`     // 店铺名称
	AccessToken  string `json:"access_token"`  // 用于调用API的access_token
	RefreshToken string `json:"refresh_token"` // 用于刷新access_token的刷新令牌（有效期：14 天）
	ExpiresIn    int64  `json:"expires_in"`    // 单位（秒），默认有效期：7天
	Scope        string `json:"scope"`         // 授权作用域，使用逗号,分隔。预留字段
}
