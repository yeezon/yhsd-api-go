package youhaosuda

import ()

type PrivateApp struct {
	Config      *Config
	AccessToken string
}

func NewPrivateApp(conf *Config, token string) *PrivateApp {
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
	if len(conf.ApiVersion) == 0 {
		conf.ApiVersion = DefaultConf.ApiVersion
	}

	return &PrivateApp{conf, token}
}

func (p *PrivateApp) GenerateToken() error {
	auth := CalAuthorization(p.Config.AppKey, p.Config.AppSecret)
	v := make(map[string]string)
	v["grant_type"] = "client_credentials"
	data := URLEncode(v)
	h := make(map[string]string)
	h["Content-Type"] = "application/x-www-form-urlencoded"
	h["Authorization"] = auth
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

func (p *PrivateApp) Get(path string) response {
	return p.Config.Get(p.AccessToken, path)
}

func (p *PrivateApp) Post(path, data string) response {
	return p.Config.Post(p.AccessToken, path, data)
}

func (p *PrivateApp) Put(path, data string) response {
	return p.Config.Put(p.AccessToken, path, data)
}

func (p *PrivateApp) Delete(path string) response {
	return p.Config.Delete(p.AccessToken, path)
}
