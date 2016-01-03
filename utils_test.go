package youhaosuda

import (
	"testing"
)

func TestCalAuthorization(t *testing.T) {
	key := "b89bb09ba3994ab08906f876ad8ebaa1"
	secret := "5d2c04a456994fafa5d1741fb3d3415f"
	result := "Basic Yjg5YmIwOWJhMzk5NGFiMDg5MDZmODc2YWQ4ZWJhYTE6NWQyYzA0YTQ1Njk5NGZhZmE1ZDE3NDFmYjNkMzQxNWY="
	r := CalAuthorization(key, secret)
	if r != result {
		t.Errorf("Authorization(key, secret) failed, Got %s, expected %s", r, result)
	}
}

func TestCalHMAC(t *testing.T) {
	v := make(map[string]string)
	result := "a2a3e2dcd8a82fd9070707d4d921ac4cdc842935bf57bc38c488300ef3960726"
	v["shop_key"] = "a94a110d86d2452eb3e2af4cfb8a3828"
	v["code"] = "a84a110d86d2452eb3e2af4cfb8a3828"
	v["account_id"] = "1"
	v["time_stamp"] = "2013-08-27T13:58:35Z"
	secret := "hush"
	h := CalHMAC(secret, v)
	if h != result {
		t.Errorf("CalHMAC(secret, data) failed, Got %s, expected %s", h, result)
	}
}

func TestCalBase64HMAC(t *testing.T) {
	data := "data"
	webhook_token := "token"
	hmac := "aGJj4CYlgxhU9ZcxwdZ8tAMzgKugNK6pE0qfnt8sAWI="
	v := CalBase64HMAC(webhook_token, data)
	if v != hmac {
		t.Errorf("CalBase64HMAC(secret, data) failed, Got %s, expected %s", v, hmac)
	}
}

func TestCalBase64Aes(t *testing.T) {
	data := "data"
	secret := "906155047ff74a14a1ca6b1fa74d3390"
	hmac := "M4ewEmU8DPnivdJcxRggyg=="
	v := CalBase64Aes(secret, data)
	if v != hmac {
		t.Errorf("CalBase64Aes(secret, data) failed, Got %s, expected %s", v, hmac)
	}
}

func TestRemoveSuffix(t *testing.T) {
	s := "http://baidu.com/"
	res := "http://baidu.com"
	s = RemoveSuffix(s, "/")
	if s != res {
		t.Errorf("RemoveSuffix(s, suffix string) failed, Got %s, expected %s", s, res)
	}
}
