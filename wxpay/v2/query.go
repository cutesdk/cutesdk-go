package wxpay

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/common/request"
)

// QueryOrder: query order
func (c *Client) QueryOrder(params map[string]interface{}) (request.Result, error) {
	uri := "/pay/orderquery"

	data, err := c.BuildParams(params)
	if err != nil {
		return nil, fmt.Errorf("build params failed: %v", err)
	}

	res, err := c.Post(uri, data)

	return res, err
}

// QueryOrderByOutTradeNo: query order by out_trade_no
func (c *Client) QueryOrderByOutTradeNo(outTradeNo string) (request.Result, error) {
	params := map[string]interface{}{
		"out_trade_no": outTradeNo,
	}

	return c.QueryOrder(params)
}

// QueryOrderByTransactionId: query order by transaction_id
func (c *Client) QueryOrderByTransactionId(transactionId string) (request.Result, error) {
	params := map[string]interface{}{
		"transaction_id": transactionId,
	}

	return c.QueryOrder(params)
}
