package youhaosuda

import (
	"testing"
)

var conf1 = Config{
	AuthUrl:   "http://apps.localtest.com/oauth2/authorize/",
	TokenUrl:  "http://apps.localtest.com/oauth2/token/",
	ApiUrl:    "http://api.public.com/",
	AppKey:    "d677ae82993a48fcaaf3c05ead9f46ea",
	AppSecret: "6e6d1e96f23f49a1a59f9ce87fed1763",
}
var publicApp = NewPublicApp(&conf1, "b66079ff889e463e8c583c2c3755bd2d")

func TestPublicGet(t *testing.T) {
	res := publicApp.Get("shop")
	AssertResponseErr(res.status, "PublicApp.Get()", t)
}

func TestPublicPost(t *testing.T) {
	data := `
  {
        "redirect": {
          "path": "/123",
          "target": "/blogs"
        }
    }
  `

	res := publicApp.Post("redirects", data)
	AssertResponseErr(res.status, "PublicApp.Post()", t)
}

func TestPublicPut(t *testing.T) {
	data := `
  {
        "redirect": {
          "path": "/66",
          "target": "/blogs"
        }
    }
  `
	res := publicApp.Put("redirects/23", data)
	AssertResponseErr(res.status, "PublicApp.Put()", t)
}

func TestPublicDelete(t *testing.T) {
	res := publicApp.Delete("redirects/23")
	AssertResponseErr(res.status, "PublicApp.Delete()", t)
}

func AssertResponseErr(status int, mehod string, t *testing.T) {
	if status != 200 && status != 422 {
		t.Errorf(mehod, " failed, response status:%d", status)
	}
}
