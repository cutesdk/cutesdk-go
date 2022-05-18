package tests

import "testing"

func TestQueryOrderByOutTradeNo(t *testing.T) {
	client := getClient()

	outTradeNo := "202205181652863740781065"

	res, err := client.QueryOrderByOutTradeNo(outTradeNo)

	t.Error(res, err)
}
