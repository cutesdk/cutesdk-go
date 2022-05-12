package tests

import (
	"testing"
)

func TestVerifyMessage(t *testing.T) {
	s := getServer()

	timestamp := "1652257715"
	nonce := "1935875759"
	msgSignature := "a3f08b678de9ba9d8e662ed5afce51da0dde46b0"
	msgEncrypt := "6RpNu8UGvmUhPonHwnPDbaKhZ//xnzd/JoMj9WaY0fG1bu44WdyyoAfokUUkvYNOKNI51yEQMnWvLeHvKRxyGHH6W+iQYJneblEBnIVAoekqvVeLrYDlx3Jf9jfmY181ldVQTtQ1TsYQLP33qHWGgMh9Z+lbkPCLkTc6+flVXicoMTPNesyOTXMj3R+/5UkARW9Mww37IxqwAJrAoccb38GlEHIpb4HMW5d+bRE5innJQq/dkiCyLE6ujvJeiOBZXJzr2r8oLdA98dKuq7Va2b+fZBsUSsfCp3QGfyxM4n+2JgkpPvQ9LlYrEohE+ppcgquuJOSunMy9b5F2MklZmCtsceGWdylwgDfxf6Z/mUPUPBBj64wJPTRQp4jH4Ckl8ptzLkdb/OC2wgvplo68lLnrI4FAZmyfSFpX5Em/GSE="

	err := s.VerifyMsg(timestamp, nonce, msgSignature, msgEncrypt)

	t.Error(err)
}

func TestDecryptMessage(t *testing.T) {
	s := getServer()

	encryptedMsg := "sR9g2NyStU80jYxj7qo+V6bxdNqOP6ElbZtfMYnQZ/xEkyRUI02TBuC/zKwpx/FqsE6HLA91fNjhBgTTFa0tyQ7ApKwtGU2RZqWEwPapm3JlJwZB+PIinFfrOVNlRHCAdvcoUZF7sC0PNUaSNL2STkHRzBoE+fn55exRl1GvOIis8XeU4sJwCSrDFuHHa82pkcCUtBE4u9y1VF7KxAfThAmIFY7tdYFPUBbsF4Sts2jLQwDNyLxCZlcTSKqr29z4RUSDCqN30fSBaXsfzUt7E4x1rzCCt4HCIQm9CitfY01NFE21Pi0yoGvLWkSMwnvogH63kVeyYoISb7AnBdJi3Fhhgq71evLIuMm38a+2PvQICVPQVUkrrhW4PH/X30gK1IFvS9NPhO17/507/RnKXw=="

	res, err := s.DecryptMsg(encryptedMsg)

	t.Errorf("%s, %v", res, err)
}

func TestEncryptMessage(t *testing.T) {
	s := getServer()

	rawMsg := `<xml>
	<ToUserName><![CDATA[gh_5a75ea150c7b]]></ToUserName>
	<FromUserName><![CDATA[orNyi07T5GeE2nwVV7b1dJ3xGnPM]]></FromUserName>
	<CreateTime>1652257715</CreateTime>
	<MsgType><![CDATA[text]]></MsgType>
	<Content><![CDATA[哦哦]]></Content>
	<MsgId>23654642931944714</MsgId>
</xml>`

	res, err := s.EncryptMsg([]byte(rawMsg))

	t.Errorf("%s, %v", res, err)
}
