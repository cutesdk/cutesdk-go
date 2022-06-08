package tests

import (
	"time"

	"github.com/cutesdk/cutesdk-go/wxopen"
)

var (
	componentAppid     string
	componentAppsecret string
	verifyToken        string
	encodingAesKey     string
)

func getIns() *wxopen.Instance {
	opts := &wxopen.Options{
		ComponentAppid:     componentAppid,
		ComponentAppsecret: componentAppsecret,
		VerifyToken:        verifyToken,
		EncodingAesKey:     encodingAesKey,
		Debug:              true,
	}

	ins, err := wxopen.New(opts)

	if err != nil {
		panic(err)
	}

	componentVerifyTicket := `ticket@@@qL9Rn3sBOGRzqZdjhCXPi27B0EqGJIvrpupPhjV-epEpkHrCrFRiCjwDAAIbfRRmMhis24sB9Ax06BaROLkjNQ`
	ins.SetComponentVerifyTicket(componentVerifyTicket, 5*time.Second)

	return ins
}

func init() {
	componentAppid = "wxf2f955ce09390e6a"
	componentAppsecret = "d6e9032e5f5bcea2f96b66f2c4e1cab8"
	verifyToken = "OaNCoqFftJz7YkUD"
	encodingAesKey = "MNfsPhrt28W4dksbARCANqIHqLmzdbZvQH8WtGgGzHv"
}
