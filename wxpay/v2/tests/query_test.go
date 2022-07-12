package tests

import "testing"

func TestQueryOrderByOutTradeNo(t *testing.T) {
	ins := getPayIns()

	uri := "/pay/orderquery"

	params := map[string]interface{}{
		"out_trade_no": "xxx",
	}

	res, err := ins.Post(uri, params)

	t.Error(res.GetString("trade_state"), err)
}
