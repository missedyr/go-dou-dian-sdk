# go-dou-dian-sdk
抖店-后端服务对接SDK-go语言

###  使用方法
go get github.com/missedyr/go-dou-dian-sdk

##
### 使用示例
```go
AppKey    string // 应用key
AppSecret string // 应用秘钥
ShopId    string // 店铺ID 仅自用型应用有效 若不传-则默认返回最早授权成功店铺对应的token信息

douDianClient := douDianSdk.New(appKey, appSecret, shopId)
douDianClient.OrderDetail("父订单号") // 获取订单详情
douDianClient.GoodsDetail("商品id", "外部商品id", true) // 获取商品详情
```

##
### 签名算法
* [func Sign](#FuncSign)
* [func SignMarshal](#FuncSignMarshal)

##
### 授权
* [func GetAccessToken](#funcGetAccessToken)

##
### 商品API
* [func GoodsDetail](#funcGoodsDetail)

##
### 订单API
* [func OrderDetail](#funcOrderDetail)

##
### 订单-结算-API
* [func OrderSettleDetail](#funcOrderSettleDetail)
* [func OrderSettleDetailByOrderId](#funcOrderSettleDetailByOrderId)

##
### func 描述

###  <a name='FuncSign'></a> func Sign
```
func Sign(appKey string, appSecret string, method string, timestamp int64, paramJson string) string

Sign 计算签名
```

###  <a name='FuncSignMarshal'></a> func SignMarshal
```
func SignMarshal(o interface{}) string

Marshal 序列化参数
```

###  <a name='funcGetAccessToken'></a> func GetAccessToken
```
note: 获取 accessToken
      该方法当前仅支持  自用型应用
func GetAccessToken() (SAccessTokenData, error)
Return:  SAccessTokenData, error
```

###  <a name='funcGoodsDetail'></a> func GoodsDetail
```
note： 商品信息查询
func GoodsDetail(goodsId, goodsIdOut string, showDraft bool)
Return:  SDYOrderInfo, error
```

###  <a name='funcOrderDetail'></a> func OrderDetail
```
note： 订单详情信息获取
func OrderDetail(orderIdParent string)
Return:  SDYOrderInfo, error
```

###  <a name='funcOrderSettleDetail'></a> func OrderSettleDetail
```
note： 订单-结算-详情-根据不同条件查询
func OrderSettleDetail(params SOrderSettleDetailParams, printLog bool)
Return: SOrderSettleDetailData, error
```

###  <a name='funcOrderSettleDetailByOrderId'></a> func OrderSettleDetailByOrderId
```
note： 订单-结算-详情-根据子订单号查询
func OrderSettleDetailByOrderId(orderId string, printLog bool)
Return:  []SOrderSettleDetail, error
```