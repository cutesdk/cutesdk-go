package tests

// import (
// 	"testing"
// )

// func TestWxOpenVerifyNotifyData(t *testing.T) {
// 	w := getWxOpen()

// 	timestamp := "1640089955"
// 	nonce := "546879428"
// 	msgSignature := "ef133f5e2cf3ab68fb16852bf2acfd7067507f94"
// 	msgEncrypt := "eeZxFxNdy3ZLQeHu/c7HGpCf4pupLiLZlbSR5Ty5Pu+p31WifUqGpPJE7Z+/52/rvEOcJUk47p2abTFb4Ghs2RS0uc4yicpdKEqlyiVtUHcUNpfX6DcYNyduRkrvbmism4olO+1EuBOLzX7meaVjLvzImkZtFimMdj6dY7M0CJ7h4p+0sU6LK2e9odf8B25HZH0ytEHYsoFGHYRmUVDDkAMoBN7S0+Sn+yF/RZN1qcYS7pkAoKGWJt6H94FwXcKLFkx8Y69H1NGaXLAQBCGCubKTERyYJMUi9lhEbV/GFwF1JrBlyHhltivlXH+5DEdtdDWGLTESvWRLjEYCHHyekpS3xNC1L+sSkZkWAmOF6fe9Wue0LOYO19+uiFOOtxo5FZGke9OwzjfXYe+K/VKSd+FPxBSQeK/wZxdvMnA1rxXeFBGHJIPxSY0HLOf5iDDXlOJlHdUgrQfQ3R+mwOAEyw=="

// 	err := w.VerifyNotifyData(timestamp, nonce, msgSignature, msgEncrypt)
// 	if err != nil {
// 		t.Fatalf("verify notify data failed: %v\n", err)
// 	}
// }

// func TestWxOpenDecryptNotifyData(t *testing.T) {
// 	w := getWxOpen()

// 	msgEncrypt := "eeZxFxNdy3ZLQeHu/c7HGpCf4pupLiLZlbSR5Ty5Pu+p31WifUqGpPJE7Z+/52/rvEOcJUk47p2abTFb4Ghs2RS0uc4yicpdKEqlyiVtUHcUNpfX6DcYNyduRkrvbmism4olO+1EuBOLzX7meaVjLvzImkZtFimMdj6dY7M0CJ7h4p+0sU6LK2e9odf8B25HZH0ytEHYsoFGHYRmUVDDkAMoBN7S0+Sn+yF/RZN1qcYS7pkAoKGWJt6H94FwXcKLFkx8Y69H1NGaXLAQBCGCubKTERyYJMUi9lhEbV/GFwF1JrBlyHhltivlXH+5DEdtdDWGLTESvWRLjEYCHHyekpS3xNC1L+sSkZkWAmOF6fe9Wue0LOYO19+uiFOOtxo5FZGke9OwzjfXYe+K/VKSd+FPxBSQeK/wZxdvMnA1rxXeFBGHJIPxSY0HLOf5iDDXlOJlHdUgrQfQ3R+mwOAEyw=="

// 	data, err := w.DecryptNotifyData(msgEncrypt)
// 	if err != nil {
// 		t.Fatalf("get notify data error: %v\n", err)
// 	}

// 	if data.XmlParsed().Get("xml.InfoType").String() != "component_verify_ticket" {
// 		t.Fatalf("get notify data failed")
// 	}
// }

// func getWxOpen() *WxOpen {
// 	opts := Options{
// 		Debug:          true,
// 		Appid:          "wxf2f955ce09390e6a",
// 		AppSecret:      "d6e9032e5f5bcea2f96b66f2c4e1cab8",
// 		VerifyToken:    "OaNCoqFftJz7YkUD",
// 		EncodingAesKey: "MNfsPhrt28W4dksbARCANqIHqLmzdbZvQH8WtGgGzHv",
// 	}

// 	return New(opts)
// }
