package examples

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/cutesdk/cutesdk-go/common/request"
)

func ExampleApply() {
	cli := getClient()

	sc, err := cli.GetSensitiveCryptor()
	if err != nil {
		log.Fatalf("get cert cryptor failed: %v\n", err)
	}

	subjectLicenseCopyMediaId := "xxx"
	legalIdCardCopyMediaId := "xxx"
	legalIdCardNationalMediaId := "xxx"

	businessCode := fmt.Sprintf("%s_%s", cli.GetMchId(), time.Now().Format("20060102150405"))

	data := map[string]interface{}{
		"subject": map[string]interface{}{
			"subject_type":       "SUBJECT_TYPE_ENTERPRISE",
			"merchant_name":      "xxx",
			"merchant_shortname": "xxx",
			"license_copy":       subjectLicenseCopyMediaId,
			"license_number":     "xxx",
			"service_phone":      "xxx",
			"settlement_id":      "716",
			"qualification_type": "休闲娱乐/旅游服务",
		},
		"legal": map[string]interface{}{
			"id_card_copy":      legalIdCardCopyMediaId,
			"id_card_national":  legalIdCardNationalMediaId,
			"id_card_name":      "xxx",
			"id_card_number":    "xxx",
			"id_card_address":   "xxx",
			"card_period_begin": "2017-09-30",
			"card_period_end":   "2027-09-30",
			"mobile_phone":      "xxx",
			"contact_email":     "xxx",
		},
		"settle": map[string]interface{}{
			"bank_account_type": "BANK_ACCOUNT_TYPE_CORPORATE",
			"account_name":      "xxx",
			"account_bank":      "xxx",
			"bank_address_code": "xxx",
			"bank_name":         "",
			"account_number":    "xxx",
		},
	}

	data["provider"] = map[string]interface{}{
		"mp_appid": "xxx",
	}

	b, _ := json.Marshal(data)
	info := request.NewResult(b)

	uri := "/v3/applyment4sub/applyment/"
	params := map[string]interface{}{
		"business_code": businessCode,
		"contact_info": map[string]interface{}{
			"contact_type":  "LEGAL",
			"contact_name":  sc.Encrypt(info.GetString("legal.id_card_name")),
			"mobile_phone":  sc.Encrypt(info.GetString("legal.mobile_phone")),
			"contact_email": sc.Encrypt(info.GetString("legal.contact_email")),
		},
		"subject_info": map[string]interface{}{
			"subject_type": info.GetString("subject.subject_type"),
			"business_license_info": map[string]interface{}{
				"license_copy":   info.GetString("subject.license_copy"),
				"license_number": info.GetString("subject.license_number"),
				"merchant_name":  info.GetString("subject.merchant_name"),
				"legal_person":   info.GetString("legal.id_card_name"),
			},
			"identity_info": map[string]interface{}{
				"id_doc_type": "IDENTIFICATION_TYPE_IDCARD",
				"id_card_info": map[string]interface{}{
					"id_card_copy":      info.GetString("legal.id_card_copy"),
					"id_card_national":  info.GetString("legal.id_card_national"),
					"id_card_name":      sc.Encrypt(info.GetString("legal.id_card_name")),
					"id_card_number":    sc.Encrypt(info.GetString("legal.id_card_number")),
					"id_card_address":   sc.Encrypt(info.GetString("legal.id_card_address")),
					"card_period_begin": info.GetString("legal.card_period_begin"),
					"card_period_end":   info.GetString("legal.card_period_end"),
				},
				"owner": true,
			},
		},
		"business_info": map[string]interface{}{
			"merchant_shortname": info.GetString("subject.merchant_shortname"),
			"service_phone":      info.GetString("subject.service_phone"),
			"sales_info": map[string]interface{}{
				"sales_scenes_type": []string{"SALES_SCENES_MP"},
				"mp_info": map[string]interface{}{
					"mp_appid": info.GetString("provider.mp_appid"),
					"mp_pics":  []string{info.GetString("subject.license_copy")},
				},
			},
		},
		"settlement_info": map[string]interface{}{
			"settlement_id":      info.GetString("subject.settlement_id"),
			"qualification_type": info.GetString("subject.qualification_type"),
		},
		"bank_account_info": map[string]interface{}{
			"bank_account_type": "BANK_ACCOUNT_TYPE_CORPORATE",
			"account_name":      sc.Encrypt(info.GetString("settle.account_name")),
			"account_bank":      info.GetString("settle.account_bank"),
			"bank_address_code": info.GetString("settle.bank_address_code"),
			// "bank_name":         "",
			"account_number": sc.Encrypt(info.GetString("settle.account_number")),
		},
	}

	headers := map[string]interface{}{
		"Wechatpay-Serial": sc.GetSerialNo(),
	}
	res, err := cli.Post(uri, params, headers)
	if err != nil {
		if wxerr := cli.UnwrapError(err); wxerr != nil {
			log.Fatalf("request api failed: %s, %v\n", wxerr.Code, wxerr)
		}
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res)
	// Output: xxxx
}
