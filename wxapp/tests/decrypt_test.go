package tests

import (
	"testing"
)

func TestDecryptUserInfo(t *testing.T) {
	client := getClient()

	sessionKey := "Of+1DKyM6qM0dhglGc6tNA=="
	encryptedData := "8zRDtmDSpcxoBRv1u8Uxg8l35S4Ihb1M1M9FEMw4CuB7kKbpiIVaN+igb4QgbTA+4te1+uEl/YzMCE4bs60R2twylUr1WYW1tpDvLp1zgoCSJSjB6MoW1fbnpdhDKREPBK3D/Jjh9MAi9tAtpH0F8T55bNr+G/WoyqtoAK8LkuIzVkv4yvemzvUyEjW9clWt4nDkQisYh3iQ+4umiB2Tmnyqcd0hn2OPmZIriplT/nt9n7KIcYDV7SCy5G15pJ9OwZjzRwoYa6iBjwXDfaWmZ2F0L3rSEJJ7aIeDnbgOmixNQ/UHgzW+LF8pxU7pp0yn2od2i4x/p/EajJYf9QGnTBAXiUiHh30TJUdldaT9b7QsqxhKGMUo0bJGEs8+tN4HWfQxR20YlYlE0Rxm6EoEZQ=="
	iv := "CxMLrVcn3wGaZsvORcMDrA=="

	res, err := client.DecryptUserInfo(sessionKey, encryptedData, iv)

	if err != nil {
		t.Fatalf("decrypt user info error: %v\n", err)
	}

	t.Error(res, err)
}

func TestDecryptPhone(t *testing.T) {
	client := getClient()

	sessionKey := "Of+1DKyM6qM0dhglGc6tNA=="
	encryptedData := "nTjSfq4Bvf4Hhp5VbsdatD9BstIzmLNe3OBrK9q4razl4WYlcGMte2u4jFh/UlwaGnvraZQOip4pjnFgLkj+TGDqCcp7ffEcmuQY7/CAQT8VUa6Ms2DYmY16DRjOcbtCRjRW6YvJfx0upjdDsAUqXmlcokmQXaNhmcw2467FmuvDl047oNUhG2ckMVtkILbYqQ6st75bvzUPt9AcxQOExg=="
	iv := "6UZpnno3DTCyfBqKbSaJiA=="

	res, err := client.DecryptPhone(sessionKey, encryptedData, iv)

	if err != nil {
		t.Fatalf("decrypt phone error: %v\n", err)
	}

	t.Error(res, err)
}
