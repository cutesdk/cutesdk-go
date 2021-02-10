package wxpay

import (
	"fmt"
	"strings"

	"github.com/idoubi/goutils"
	"github.com/idoubi/goz"
	"github.com/tidwall/gjson"
)

// UnifiedOrder 统一下单
func (w *Wxpay) UnifiedOrder(params map[string]string) (*gjson.Result, error) {
	if _, ok := params["mch_id"]; !ok {
		params["mch_id"] = w.opts.MchID
	}
	if _, ok := params["sub_mch_id"]; !ok {
		params["sub_mch_id"] = w.opts.SubMchID
	}
	if _, ok := params["appid"]; !ok {
		params["appid"] = w.opts.Appid
	}
	if _, ok := params["sub_appid"]; !ok {
		params["sub_appid"] = w.opts.SubAppid
	}
	if _, ok := params["notify_url"]; !ok {
		params["notify_url"] = w.opts.NotifyURL
	}
	if _, ok := params["nonce_str"]; !ok {
		params["nonce_str"] = goutils.NonceStr(32)
	}
	if _, ok := params["trade_type"]; !ok {
		params["trade_type"] = "JSAPI"
	}

	// 生成签名
	if _, ok := params["sign"]; !ok {
		params["sign"] = strings.ToUpper(goutils.MD5(goutils.SortEncodeMap(params, "key", w.opts.APIKey)))
	}

	apiURL := fmt.Sprintf("%s/pay/unifiedorder", apiBase)
	resp, err := goz.Post(apiURL, goz.Options{
		XML:   params,
		Debug: w.opts.Debug,
	})

	body, err := resp.GetBody()

	if err != nil {
		return nil, err
	}

	jb, err := goutils.XML2JSON(body)
	if err != nil {
		return nil, err
	}

	d := gjson.ParseBytes(jb)
	xd := d.Get("xml")

	return &xd, nil
}

// GetPayParams 获取支付参数
// wx.requestPayment(OBJECT) 接口需要的参数
func (w *Wxpay) GetPayParams(prepayID string) map[string]string {
	params := map[string]string{
		"appId":     w.opts.Appid,
		"timeStamp": goutils.TimestampStr(),
		"nonceStr":  goutils.NonceStr(32),
		"package":   "prepay_id=" + prepayID,
		"signType":  "MD5",
	}

	params["paySign"] = strings.ToUpper(goutils.MD5(goutils.SortEncodeMap(params, "key", w.opts.APIKey)))

	return params
}
