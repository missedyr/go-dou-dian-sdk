package douDianSdk

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

// GoodsDetail 获取商品详情
// goodsId    非必需 商品id
// goodsIdOut 非必需 外部商品id
// showDraft true：读取草稿数据 false：读取上架数据
// https://op.jinritemai.com/docs/api-docs/14    商品查询 /product/detail
func (cl *DouDianClient) GoodsDetail(goodsId, goodsIdOut string, showDraft bool) (SGoodsInfo, error) {
	paramJson := map[string]interface{}{
		"show_draft":    cast.ToString(showDraft),
	}
	if goodsId != "" {
		paramJson["product_id"] = goodsId
	}
	if goodsIdOut != "" {
		paramJson["out_product_id"] = goodsIdOut
	}

	method := "product.detail"
	respBody, err := cl.FetchSignAndToken(method, paramJson, "get")
	if err != nil {
		return SGoodsInfo{}, err
	}

	var body SGoodsInfo
	_ = json.Unmarshal(respBody, &body)
	bodyJson, _ := json.Marshal(body)
	logrus.Infof("douDianSdk-->GoodsDetail -- bodyJson:%v", string(bodyJson))
	return body, err
}
