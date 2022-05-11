package tests

import (
	"testing"
)

func TestVerifyNotifyMessage(t *testing.T) {
	s := getServer()

	timestamp := "1652257715"
	nonce := "1935875759"
	msgSignature := "a3f08b678de9ba9d8e662ed5afce51da0dde46b0"
	msgEncrypt := "6RpNu8UGvmUhPonHwnPDbaKhZ//xnzd/JoMj9WaY0fG1bu44WdyyoAfokUUkvYNOKNI51yEQMnWvLeHvKRxyGHH6W+iQYJneblEBnIVAoekqvVeLrYDlx3Jf9jfmY181ldVQTtQ1TsYQLP33qHWGgMh9Z+lbkPCLkTc6+flVXicoMTPNesyOTXMj3R+/5UkARW9Mww37IxqwAJrAoccb38GlEHIpb4HMW5d+bRE5innJQq/dkiCyLE6ujvJeiOBZXJzr2r8oLdA98dKuq7Va2b+fZBsUSsfCp3QGfyxM4n+2JgkpPvQ9LlYrEohE+ppcgquuJOSunMy9b5F2MklZmCtsceGWdylwgDfxf6Z/mUPUPBBj64wJPTRQp4jH4Ckl8ptzLkdb/OC2wgvplo68lLnrI4FAZmyfSFpX5Em/GSE="

	err := s.VerifyNotifyData(timestamp, nonce, msgSignature, msgEncrypt)

	t.Error(err)
}

func TestDecryptNotifyMessage(t *testing.T) {
	s := getServer()

	msgEncrypt := "6RpNu8UGvmUhPonHwnPDbaKhZ//xnzd/JoMj9WaY0fG1bu44WdyyoAfokUUkvYNOKNI51yEQMnWvLeHvKRxyGHH6W+iQYJneblEBnIVAoekqvVeLrYDlx3Jf9jfmY181ldVQTtQ1TsYQLP33qHWGgMh9Z+lbkPCLkTc6+flVXicoMTPNesyOTXMj3R+/5UkARW9Mww37IxqwAJrAoccb38GlEHIpb4HMW5d+bRE5innJQq/dkiCyLE6ujvJeiOBZXJzr2r8oLdA98dKuq7Va2b+fZBsUSsfCp3QGfyxM4n+2JgkpPvQ9LlYrEohE+ppcgquuJOSunMy9b5F2MklZmCtsceGWdylwgDfxf6Z/mUPUPBBj64wJPTRQp4jH4Ckl8ptzLkdb/OC2wgvplo68lLnrI4FAZmyfSFpX5Em/GSE="

	res, err := s.DecryptNotifyData(msgEncrypt)

	t.Errorf("%+v, %v", res, err)

	t.Error(res.ToUserName, res.FromUserName, res.CreateTime)
}
