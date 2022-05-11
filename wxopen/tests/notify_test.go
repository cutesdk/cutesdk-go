package tests

import (
	"testing"
)

func TestVerifyNotifyData(t *testing.T) {
	s := getServer()

	timestamp := "1640089955"
	nonce := "546879428"
	msgSignature := "ef133f5e2cf3ab68fb16852bf2acfd7067507f94"
	msgEncrypt := "eeZxFxNdy3ZLQeHu/c7HGpCf4pupLiLZlbSR5Ty5Pu+p31WifUqGpPJE7Z+/52/rvEOcJUk47p2abTFb4Ghs2RS0uc4yicpdKEqlyiVtUHcUNpfX6DcYNyduRkrvbmism4olO+1EuBOLzX7meaVjLvzImkZtFimMdj6dY7M0CJ7h4p+0sU6LK2e9odf8B25HZH0ytEHYsoFGHYRmUVDDkAMoBN7S0+Sn+yF/RZN1qcYS7pkAoKGWJt6H94FwXcKLFkx8Y69H1NGaXLAQBCGCubKTERyYJMUi9lhEbV/GFwF1JrBlyHhltivlXH+5DEdtdDWGLTESvWRLjEYCHHyekpS3xNC1L+sSkZkWAmOF6fe9Wue0LOYO19+uiFOOtxo5FZGke9OwzjfXYe+K/VKSd+FPxBSQeK/wZxdvMnA1rxXeFBGHJIPxSY0HLOf5iDDXlOJlHdUgrQfQ3R+mwOAEyw=="

	err := s.VerifyNotifyData(timestamp, nonce, msgSignature, msgEncrypt)

	t.Error(err)
}

func TestDecryptNotifyData(t *testing.T) {
	s := getServer()

	msgEncrypt := "eeZxFxNdy3ZLQeHu/c7HGpCf4pupLiLZlbSR5Ty5Pu+p31WifUqGpPJE7Z+/52/rvEOcJUk47p2abTFb4Ghs2RS0uc4yicpdKEqlyiVtUHcUNpfX6DcYNyduRkrvbmism4olO+1EuBOLzX7meaVjLvzImkZtFimMdj6dY7M0CJ7h4p+0sU6LK2e9odf8B25HZH0ytEHYsoFGHYRmUVDDkAMoBN7S0+Sn+yF/RZN1qcYS7pkAoKGWJt6H94FwXcKLFkx8Y69H1NGaXLAQBCGCubKTERyYJMUi9lhEbV/GFwF1JrBlyHhltivlXH+5DEdtdDWGLTESvWRLjEYCHHyekpS3xNC1L+sSkZkWAmOF6fe9Wue0LOYO19+uiFOOtxo5FZGke9OwzjfXYe+K/VKSd+FPxBSQeK/wZxdvMnA1rxXeFBGHJIPxSY0HLOf5iDDXlOJlHdUgrQfQ3R+mwOAEyw=="

	res, err := s.DecryptNotifyData(msgEncrypt)

	t.Error(res, err)
}
