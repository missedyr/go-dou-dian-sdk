package douDianSdk

// SOrderSettleDetailParams 获取订单结算详情参数
type SOrderSettleDetailParams struct {
	Size       int64  `json:"size"`        // 非必须 查询页大小(默认100,最大为1000)
	ProductId  string `json:"product_id"`  // 非必须 商品id
	OrderId    string `json:"order_id"`    // 非必须 SKU单，子订单号
	PayType    string `json:"pay_type"`    // 非必须 结算账户 0:全部 1:微信（升级前） 2:微信 3:支付宝 4:合众支付 5:聚合账户
	FlowType   string `json:"flow_type"`   // 非必须 业务类型，不传则默认为0 0:全部 1:鲁班广告, 2:值点商城, 3:精选联盟 4:小店自卖
	TimeType   string `json:"time_type"`   // 非必须 时间类型 0:结算时间 1：下单时间
	StartTime  string `json:"start_time"`  // 非必须 查询开始时间。注:订单号未传的情况下，时间必须传  2021-05-18 00:00:00
	EndTime    string `json:"end_time"`    // 非必须 查询结束时间。注:订单号未传的情况下，时间必须传 2021-05-18 23:59:59
	StartIndex string `json:"start_index"` // 非必须 查询开始索引,值为上一次请求的结果next_start_index,第一次查询可以不填
}

// SOrderSettleDetail 订单结算详细信息
type SOrderSettleDetail struct {
	ShopId                 int64  `json:"shop_id"`                  // 店铺id
	ProductId              string `json:"product_id"`               // 商品id
	GoodsCount             int64  `json:"goods_count"`              // 商品数量
	ShopOrderId            string `json:"shop_order_id"`            // 店铺单号（主订单号）
	OrderId                string `json:"order_id"`                 // SKU单（子订单号）
	OrderTime              string `json:"order_time"`               // 下单时间 2021-04-07 11:05:56
	FlowType               string `json:"flow_type"`                // 业务类型 1:鲁班广告, 2:值点商城, 3:精选联盟 4:小店自卖
	FlowTypeDesc           string `json:"flow_type_desc"`           // 业务类型描述
	RequestNo              string `json:"request_no"`               // 结算单号,每一条订单流水明细结算单号唯一
	SettleTime             string `json:"settle_time"`              // 结算时间
	PayType                string `json:"pay_type"`                 // 结算账户类型 1:微信（升级前）, 2:微信，3:支付宝, 4:合众支付 5:聚合账户
	PayTypeDesc            string `json:"pay_type_desc"`            // 结算账户类型描述
	TradeType              string `json:"trade_type"`               // 结算状态 0:已结算 1:已退款
	TradeTypeDesc          string `json:"trade_type_desc"`          // 结算状态描述
	TotalIncome            int64  `json:"total_income"`             // 总收入(分)
	TotalOutcome           int64  `json:"total_outcome"`            // 总支出(分)
	SettleAmount           int64  `json:"settle_amount"`            // 实际结算金额(分)
	ActualSubsidyAmount    int64  `json:"actual_subsidy_amount"`    // 实际补贴金额(分)
	TotalAmount            int64  `json:"total_amount"`             // 订单总价(分)
	TotalGoodsAmount       int64  `json:"total_goods_amount"`       // 商品总价(分)
	PostAmount             int64  `json:"post_amount"`              // 运费(分)
	RealPayAmount          int64  `json:"real_pay_amount"`          // 订单实付(分)
	SettledPayAmount       int64  `json:"settled_pay_amount"`       // 订单实付应结(分)
	PlatformCoupon         int64  `json:"platform_coupon"`          // 平台补贴(分)
	AuthorCoupon           int64  `json:"author_coupon"`            // 达人券补贴(分)
	PayPromotion           int64  `json:"pay_promotion"`            // 支付补贴(分)
	ActualPlatformCoupon   int64  `json:"actual_platform_coupon"`   // 实际平台补贴(分)
	ActualAuthorCoupon     int64  `json:"actual_author_coupon"`     // 实际达人券补贴(分)
	ActualPayPromotion     int64  `json:"actual_pay_promotion"`     // 实际支付补贴(分)
	ShopCoupon             int64  `json:"shop_coupon"`              // 店铺优惠(分)
	PlatformServiceFee     int64  `json:"platform_service_fee"`     // 平台服务费(分)
	Refund                 int64  `json:"refund"`                   // 订单退款(分)
	RefundBeforeSettlement int64  `json:"refund_before_settlement"` // 结算前退款(分)
	ShopRefundLoss         int64  `json:"shop_refund_loss"`         // 退款扣佣金(退佣失败垫付金额)(分)
	Commission             int64  `json:"commission"`               // 佣金(分)
	ColonelServiceFee      int64  `json:"colonel_service_fee"`      // 招商服务费(分)
	ZtPayPromotion         int64  `json:"zt_pay_promotion"`         // 抖音支付支付补贴(分)
	ZrPayPromotion         int64  `json:"zr_pay_promotion"`         // DOU分期支付补贴(分)
	ActualZtPayPromotion   int64  `json:"actual_zt_pay_promotion"`  // 实际抖音支付补贴金额(分)
	ActualZrPayPromotion   int64  `json:"actual_zr_pay_promotion"`  // 实际DOU分期支付补贴金额(分)
	OtherSharingAmount     int64  `json:"other_sharing_amount"`     // 其他分成金额(分)
	Remark                 string `json:"remark"`                   //  备注
}
