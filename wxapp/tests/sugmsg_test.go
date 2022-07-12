package tests

import "testing"

func TestGetSubmsgTpls(t *testing.T) {
	ins := getIns()

	uri := "/wxaapi/newtmpl/gettemplate"

	res, err := ins.GetWithAccessToken(uri)

	t.Error(res, err)
}
