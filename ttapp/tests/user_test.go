package tests

import "testing"

func TestCode2Session(t *testing.T) {
	ins := getIns()

	uri := "/api/apps/v2/jscode2session"

	code := "pox7TA-_rr3DrPLDaeuyWRGBnJsn9vWTxaLa6tPYJVMiBxaZ_cc47wGLVDlJyCsklQw5E3eRsqthmLVRzco_rBodQvGIUeujOgCjKcIkf-7HiqyPritSmlleOOA"
	anonymousCode := "KTqKa6x31on0X6bf7E3SlwU4MSF3BNizc64zsNBd6KGDfkkSvCjjysmCad04rfUqLLcUPUfMGyQvpkeGW1Jjt1Q6Yrjp5OCX7rkfKDtOduqf2Y_1UQmlIx4YMdA"

	params := map[string]interface{}{
		"appid":          ins.GetAppid(),
		"secret":         ins.GetSecret(),
		"code":           code,
		"anonymous_code": anonymousCode,
	}

	res, err := ins.Post(uri, params)

	t.Error(res.GetMap("data"), err)
}
