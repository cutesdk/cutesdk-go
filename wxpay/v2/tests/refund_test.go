package tests

import "testing"

func TestRefund(t *testing.T) {
	ins := getPayIns()

	uri := "/secapi/pay/refund"

	params := map[string]interface{}{
		"out_trade_no":  "1587467834760368128",
		"out_refund_no": "RF1587467834760368128",
		"total_fee":     3,
		"refund_fee":    3,
	}

	res, err := ins.PostWithCert(uri, params)

	t.Error(res.GetString("trade_state"), err)
}
