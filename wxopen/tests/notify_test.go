package tests

import (
	"testing"
)

func TestVerifyNotifyData(t *testing.T) {
	s := getServer()

	timestamp := "1652368452"
	nonce := "ahmxgnSMzUDZjNXT"
	msgSignature := "f5d0bd3a4b21b73a7d2b20f458af89c1335ad687"
	msgEncrypt := "sR9g2NyStU80jYxj7qo+V6bxdNqOP6ElbZtfMYnQZ/xEkyRUI02TBuC/zKwpx/FqsE6HLA91fNjhBgTTFa0tyQ7ApKwtGU2RZqWEwPapm3JlJwZB+PIinFfrOVNlRHCAdvcoUZF7sC0PNUaSNL2STkHRzBoE+fn55exRl1GvOIis8XeU4sJwCSrDFuHHa82pkcCUtBE4u9y1VF7KxAfThAmIFY7tdYFPUBbsF4Sts2jLQwDNyLxCZlcTSKqr29z4RUSDCqN30fSBaXsfzUt7E4x1rzCCt4HCIQm9CitfY01NFE21Pi0yoGvLWkSMwnvogH63kVeyYoISb7AnBdJi3Fhhgq71evLIuMm38a+2PvQICVPQVUkrrhW4PH/X30gK1IFvS9NPhO17/507/RnKXw=="

	err := s.VerifyMsg(timestamp, nonce, msgSignature, msgEncrypt)

	t.Error(err)
}

func TestDecryptNotifyInfo(t *testing.T) {
	s := getServer()

	msgEncrypt := "eeZxFxNdy3ZLQeHu/c7HGpCf4pupLiLZlbSR5Ty5Pu+p31WifUqGpPJE7Z+/52/rvEOcJUk47p2abTFb4Ghs2RS0uc4yicpdKEqlyiVtUHcUNpfX6DcYNyduRkrvbmism4olO+1EuBOLzX7meaVjLvzImkZtFimMdj6dY7M0CJ7h4p+0sU6LK2e9odf8B25HZH0ytEHYsoFGHYRmUVDDkAMoBN7S0+Sn+yF/RZN1qcYS7pkAoKGWJt6H94FwXcKLFkx8Y69H1NGaXLAQBCGCubKTERyYJMUi9lhEbV/GFwF1JrBlyHhltivlXH+5DEdtdDWGLTESvWRLjEYCHHyekpS3xNC1L+sSkZkWAmOF6fe9Wue0LOYO19+uiFOOtxo5FZGke9OwzjfXYe+K/VKSd+FPxBSQeK/wZxdvMnA1rxXeFBGHJIPxSY0HLOf5iDDXlOJlHdUgrQfQ3R+mwOAEyw=="

	res, err := s.DecryptMsg(msgEncrypt)

	t.Error(string(res), err)
}
