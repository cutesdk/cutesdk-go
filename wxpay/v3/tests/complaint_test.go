package tests

import (
	"fmt"
	"testing"

	"github.com/cutesdk/cutesdk-go/wxpay/v3"
)

func TestGetComplaints(t *testing.T) {
	ins := getPayIns()

	beginDate := "2022-10-01"
	endDate := "2022-10-30"
	uri := fmt.Sprintf("/v3/merchant-service/complaints-v2?begin_date=%s&end_date=%s", beginDate, endDate)

	res, err := ins.Get(uri)

	t.Error(res, err)
}

func TestGetComplaintInfo(t *testing.T) {
	ins := getPayIns()

	complaintId := "200000020221007210030055214"

	uri := fmt.Sprintf("/v3/merchant-service/complaints-v2/%s", complaintId)

	res, err := ins.Get(uri)

	t.Error(res, err)
}

func TestGetComplaintNotifyUrl(t *testing.T) {
	ins := getPayIns()

	uri := "/v3/merchant-service/complaint-notifications"

	res, err := ins.Get(uri)

	wxerr := wxpay.UnwrapError(err)

	t.Error(res, wxerr)
}

func TestCreateComplaintNotifyUrl(t *testing.T) {
	ins := getPayIns()

	uri := "/v3/merchant-service/complaint-notifications"

	params := map[string]interface{}{
		"url": "https://testapi.idoustudio.com/wxpay/complaint-notify",
	}

	res, err := ins.Post(uri, params)

	wxerr := wxpay.UnwrapError(err)

	t.Error(res, wxerr)
}
