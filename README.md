# go-dou-dian-sdk
抖店-后端服务对接SDK-go语言

###  使用方法
go get github.com/xuexin520/go-dou-dian-sdk

##
### 签名算法
* [func  Marshal](#FuncMarshal)
* [func  Sign](#FuncSign)

##
### 授权
* [func  GetAccessToken](#funcGetAccessToken)

##
### 订单API
* [func  OrderDetail](#funcOrderDetail)

### func 描述

###  <a name='FuncMarshal'></a> func FuncMarshal
```
func Marshal(o interface{}) string

Marshal 序列化参数
```

###  <a name='FuncSign'></a> func FuncSign
```
func Sign(appKey string, appSecret string, method string, timestamp int64, paramJson string) string

Sign 计算签名
```

###  <a name='funcGetAccessToken'></a> func GetAccessToken
```
func GetAccessToken(appKey string, appSecret string, shopId string) (SAccessTokenData, error)

Required appKey & appSecret
Return:  SAccessTokenData, error
```

###  <a name='funcOrderDetail'></a> func funcOrderDetail
```
func OrderDetail(appKey string, appSecret string, shopId string, orderIdParent string) (SDYOrderInfo, error)

note： 
    订单详情信息获取
    自动进行 sign 签名并获取填充 accessToken
Return:  SDYOrderInfo, error
```