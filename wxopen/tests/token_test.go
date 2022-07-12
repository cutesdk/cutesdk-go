package tests

import (
	"testing"
)

func TestFetchComponentAccessToken(t *testing.T) {
	ins := getIns()

	res, err := ins.FetchComponentAccessToken()

	t.Error(res, err)
}
