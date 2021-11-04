package douDianSdk

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"strings"
)

// Sign 计算签名
func Sign(appKey string, appSecret string, method string, timestamp int64, paramJson string) string {
	// 按给定规则拼接参数
	paramPattern := "app_key" + appKey + "method" + method + "param_json" + paramJson + "timestamp" + strconv.FormatInt(timestamp, 10) + "v2"
	signPattern := appSecret + paramPattern + appSecret
	return Hmac(signPattern, appSecret)
}

// Hmac 计算hmac
func Hmac(s string, appSecret string) string {
	h := hmac.New(sha256.New, []byte(appSecret))
	_, _ = h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Marshal 序列化参数
func Marshal(o interface{}) string {
	// 序列化一次
	raw, _ := json.Marshal(o)

	// 反序列化为map
	m := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	decode := json.NewDecoder(reader)
	decode.UseNumber()
	_ = decode.Decode(&m)

	// 重新做一次序列化，并禁用Html Escape
	buffer := bytes.NewBufferString("")
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	_ = encoder.Encode(m)

	marshal := strings.TrimSpace(buffer.String()) // Trim掉末尾的换行符
	return marshal
}
