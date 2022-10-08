package tests

import "github.com/cutesdk/cutesdk-go/wxpay/v3"

func getPayIns() *wxpay.Instance {
	opts := &wxpay.Options{
		MchId:      "xxx",
		ApiKey:     "xxx",
		SerialNo:   "xxx",
		PrivateKey: `xxx`,
		Debug:      true,
	}

	ins, _ := wxpay.New(opts)

	return ins
}
