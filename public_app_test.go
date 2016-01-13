package youhaosuda

import (
	"testing"
)

var conf1 = Config{
	AuthUrl:   "https://apps.youhaosuda.com/oauth2/authorize/",
	TokenUrl:  "https://apps.youhaosuda.com/oauth2/token/",
	ApiUrl:    "https://api.youhaosuda.com/",
	AppKey:    "57a7a5aeeb6b4db78f776e3add846e67",
	AppSecret: "ca5c91b5ea1f48b78e8bf88ce8d8a6b2",
}
var publicApp = NewPublicApp(&conf1, "44e8d8f52062453b8fe7342c618d1aef")

func TestPublicGet(t *testing.T) {
	res := publicApp.Get("shop")
	AssertResponseErr(res, "PublicApp.Get()", t)
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
	AssertResponseErr(res, "PublicApp.Post()", t)
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
	AssertResponseErr(res, "PublicApp.Put()", t)
}

func TestPublicDelete(t *testing.T) {
	res := publicApp.Delete("redirects/23")
	AssertResponseErr(res, "PublicApp.Delete()", t)
}

func AssertResponseErr(response response, mehod string, t *testing.T) {
	if response.status != 200 && response.status != 422 {
		t.Errorf(mehod, " failed, response status:%d, body:%s", response.status, response.body)
	}
}
