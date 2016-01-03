package youhaosuda

import (
	"testing"
)

var conf = Config{
	ApiUrl:    "http://api.public.com/",
	TokenUrl:  "http://apps.localtest.com/oauth2/token/",
	AppKey:    "ab3217683c964c82a685c22d9440f240",
	AppSecret: "13516ce822b841ce8d5b91630d97d050",
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
	AssertResponseErr(res.status, "PrivateApp.Get()", t)
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
	AssertResponseErr(res.status, "PrivateApp.Post()", t)
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
	AssertResponseErr(res.status, "PrivateApp.Put()", t)
}

func TestPrivateDelete(t *testing.T) {
	res := privateApp.Delete("redirects/23")
	AssertResponseErr(res.status, "PrivateApp.Delete()", t)
}
