package youhaosuda

import (
	"strings"
)

type PublicApp struct {
	Config      *Config
	AccessToken string
}

func NewPublicApp(conf *Config, token string) *PublicApp {
	if len(conf.TokenUrl) == 0 {
		conf.TokenUrl = DefaultConf.TokenUrl
	} else {
		conf.TokenUrl = RemoveSuffix(conf.TokenUrl, "/")
	}
	if len(conf.ApiUrl) == 0 {
		conf.ApiUrl = DefaultConf.ApiUrl
	} else {
		conf.ApiUrl = RemoveSuffix(conf.ApiUrl, "/")
	}
	if len(conf.AuthUrl) == 0 {
		conf.AuthUrl = DefaultConf.AuthUrl
	} else {
		conf.AuthUrl = RemoveSuffix(conf.AuthUrl, "/")
	}
	if len(conf.ApiVersion) == 0 {
		conf.ApiVersion = DefaultConf.ApiVersion
	}
	if len(conf.Scope) == 0 {
		conf.Scope = DefaultConf.Scope
	}

	return &PublicApp{conf, token}
}

func (p *PublicApp) AuthorizeUrl(redirect_url, shop_key, state string) string {
	v := make(map[string]string)
	v["response_type"] = "code"
	v["client_id"] = p.Config.AppKey
	v["shop_key"] = shop_key
	v["scope"] = p.Config.Scope
	v["redirect_uri"] = redirect_url
	v["state"] = state

	para := URLEncode(v)

	u := []string{p.Config.AuthUrl, para}
	return strings.Join(u, "?")
}

func (p *PublicApp) GenerateToken(redirect_url, code string) error {
	v := make(map[string]string)
	v["grant_type"] = "authorization_code"
	v["code"] = code
	v["client_id"] = p.Config.AppKey
	v["redirect_uri"] = redirect_url
	data := URLEncode(v)
	h := make(map[string]string)
	h["Content-Type"] = "application/x-www-form-urlencoded"
	r := request{
		url:     p.Config.TokenUrl,
		method:  "POST",
		data:    data,
		headers: h,
	}
	res, err := r.request()
	if err != nil {
		return err
	}
	t := AccessToken{}
	res.parseJson(&t)
	p.AccessToken = t.AccessToken
	return nil
}

func (p *PublicApp) Get(path string) response {
	return p.Config.Get(p.AccessToken, path)
}

func (p *PublicApp) Post(path, data string) response {
	return p.Config.Post(p.AccessToken, path, data)
}

func (p *PublicApp) Put(path, data string) response {
	return p.Config.Put(p.AccessToken, path, data)
}

func (p *PublicApp) Delete(path string) response {
	return p.Config.Delete(p.AccessToken, path)
}
