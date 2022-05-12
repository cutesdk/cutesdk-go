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

	encryptedMsg := "os29685SSuiYgpWvz+z1IUWBY/VPPhnrWOUaan90x/5qBZ49ySemoMcQDB9cDy95mCZwr1gNhKU8ig6lMSXZwpVSY9uo8qWXeEHq6h76Hw3ul8a+olgGPiGFF/yF6t9AnFf7m8TQqqspQRJTZzU6o0awAh4Q50AFUo23ecBE8bHfvP81P52pEgc4GCzbhdQt3fXP/Aa2g6OVqzr/h1RgBsVhKihnmsIxLRGeF7PMRILDwFrnASz3M8FJHDGvLmbFXpymUlIqGUQmOE819lqxvMPp8d/l2x3wFmh394TfZWSbB1mpRmc1LXcYWdbfQ/YNad5gzS+os1jKlpxeUM5u7qn2rqXY/v4xlZfTni4QTZHoBfkUZnKNQSwe8wuNZvxvT1qLhiHmxLMw5L8GBcUUKG7HYe/2/NoatF1cEXPx44teXXoQpRC2bNNXYLpsSN9Z"

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
