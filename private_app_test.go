package youhaosuda

import (
	"testing"
)

var conf = Config{
	ApiUrl:    "https://api.youhaosuda.com/",
	TokenUrl:  "https://apps.youhaosuda.com/oauth2/token/",
	AppKey:    "0249cb63872d43a28c0d9bf0ddfd6a9c",
	AppSecret: "45ab37bed9334b86b4cb6be2b4459cdf",
}
var privateApp = NewPrivateApp(&conf, "")

func TestGenerateToken(t *testing.T) {
	err := privateApp.GenerateToken()
	if err != nil {
		t.Errorf("GenerateToken() has error: %s", err.Error())
	}
	a := privateApp.AccessToken
	if len(a) == 0 {
		t.Errorf("GenerateToken() failed, token is empty.")
	}
}

func TestPrivateGet(t *testing.T) {
	res := privateApp.Get("shop")
	AssertResponseErr(res, "PrivateApp.Get()", t)
}

func TestPrivatePost(t *testing.T) {
	data := `
  {
        "redirect": {
          "path": "/123",
          "target": "/blogs"
        }
    }
  `

	res := privateApp.Post("redirects", data)
	AssertResponseErr(res, "PrivateApp.Post()", t)
}

func TestPrivatePut(t *testing.T) {
	data := `
  {
        "redirect": {
          "path": "/66",
          "target": "/blogs"
        }
    }
  `
	res := privateApp.Put("redirects/23", data)
	AssertResponseErr(res, "PrivateApp.Put()", t)
}

func TestPrivateDelete(t *testing.T) {
	res := privateApp.Delete("redirects/23")
	AssertResponseErr(res, "PrivateApp.Delete()", t)
}
