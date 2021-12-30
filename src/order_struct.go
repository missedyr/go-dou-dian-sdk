package douDianSdk

// SDYOrderInfo 抖音订单详情
type SDYOrderInfo struct {
	ShopId       int64             `json:"shop_id"`        // 店铺ID
	ShopName     string            `json:"shop_name"`      // 商户名称
	BuyerWords   string            `json:"buyer_words"`    // 买家留言
	SellerWords  string            `json:"seller_words"`   // 商家备注
	SkuOrderList []SDYSkuOrderInfo `json:"sku_order_list"` // 单商品订单信息
}

type SDYSkuOrderInfo struct {
	OrderId         string `json:"order_id"`          // 子订单号 (商品订单号)
	ParentOrderId   string `json:"parent_order_id"`   // 父订单号（店铺订单号）
	ItemNum         int64  `json:"item_num"`          // 订单商品数量
	ProductIdStr    string `json:"product_id_str"`    // 商品ID，字符串类型
	ProductIdOut    string `json:"out_product_id"`    // 商品外部ID
	Code            string `json:"code"`              // 商家后台商品编码 -- 配置洋葱商品ID
	OriginAmount    int64  `json:"origin_amount"`     // 商品现价 (单件价格) 单位 分
	OrderAmount     int64  `json:"order_amount"`      // 订单金额（单位 分）
	PayAmount       int64  `json:"pay_amount"`        // 支付金额（单位 分）
	OrderStatus     int64  `json:"order_status"`      // 订单状态
	OrderStatusDesc string `json:"order_status_desc"` // 订单状态描述
	CreateTime      int64  `json:"create_time"`       // 下单时间，时间戳，秒
	PayTime         int64  `json:"pay_time"`          // 支付时间，时间戳，秒
	BType           int64  `json:"b_type"`            // 下单端
	BTypeDesc       string `json:"b_type_desc"`       // 下单端描述
	SendPay         int64  `json:"send_pay"`          // 流量来源
	SendPayDesc     string `json:"send_pay_desc"`     // 流量来源描述
	AuthorId        int64  `json:"author_id"`         // 主播id
	AuthorName      string `json:"author_name"`       // 直播主播名称
	RoomId          int64  `json:"room_id"`           // 直播间id
	VideoId         string `json:"video_id"`          // 视频id
	Cid             int64  `json:"cid"`               // 广告id
	CBiz            int64  `json:"c_biz"`             // C端流量来源
	CBizDesc        string `json:"c_biz_desc"`        // C端流量来源描述
}
