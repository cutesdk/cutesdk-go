package tests

import (
	"fmt"
	"testing"
)

func TestGetDailyVisit(t *testing.T) {
	ins := getIns()

	accessToken, err := ins.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/datacube/getweanalysisappiddailyvisittrend?access_token=%s", accessToken)

	params := map[string]interface{}{
		"begin_date": "20220601",
		"end_date":   "20220601",
	}

	res, err := ins.Post(uri, params)

	t.Error(res.GetArray("list"), err)
}
