package douDianSdk

type SGoodsInfo struct {
	ProductIdStr    string          `json:"product_id_str"`   // 商品ID
	ProductIdOuter  string          `json:"outer_product_id"` // 商品外部ID
	OpenUserId      string          `json:"open_user_id"`     // open应用id
	Name            string          `json:"name"`             // 商品名称
	MarketPrice     int64           `json:"market_price"`     // 划线价 单位分
	DiscountPrice   int64           `json:"discount_price"`   // 售卖价 单位分
	DeliveryMethod  int64           `json:"delivery_method"`  // 承诺发货时间，单位是天
	SpecId          int64           `json:"spec_id"`          // 规格id
	Mobile          string          `json:"mobile"`           // 手机号
	BrandId         int64           `json:"brand_id"`         // 品牌id
	PayType         int64           `json:"pay_type"`         // 支持的支付方式：0 货到付款 1在线支付 2两者都支持
	RecommendRemark string          `json:"recommend_remark"` // 商家推荐语
	Img             string          `json:"img"`              // 头图，主图第一张
	Pic             []string        `json:"pic"`              // 商品主图
	SpecPrices      []SGoodsSkuInfo `json:"spec_prices"`      // sku详情列表
	DraftStatus     string          `json:"draft_status"`     // 草稿状态：0 无草稿,1 未提审,2 待审核,3 审核通过,4 审核未通过
	CheckStatus     int64           `json:"check_status"`     // 商品审核状态：1未提审 2审核中 3审核通过 4审核驳回 5封禁 7审核通过，待上架状态
	Status          int64           `json:"status"`           // 商品上下架状态：0上架 1下架 2已删除
	CreateTime      string          `json:"create_time"`      // 创建时间
	UpdateTime      string          `json:"update_time"`      // 更新时间
}

type SGoodsSkuInfo struct {
	SkuId         int64    `json:"sku_id"`          // sku id
	SkuIdOuter    string   `json:"outer_sku_id"`    //外部sku id  字符串形式
	SkuType       int64    `json:"sku_type"`        // sku类型
	SpecDetailIds []string `json:"spec_detail_ids"` // 规格id
	StockNum      int64    `json:"stock_num"`       // 当前现货可售库存、当前预售可售库存（物理库存）
	Price         int64    `json:"price"`           // 价格 单位分
	Code          string   `json:"code"`            // 商家后台商品编码 -- 配置洋葱商品ID
	SupplierId    string   `json:"supplier_id"`     // 供应商编码
}
