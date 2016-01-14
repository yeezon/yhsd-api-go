package youhaosuda

import (
	"time"
)

var DefaultConf = &Config{
	Scope:      "read_basic,write_basic",
	AuthUrl:    "https://apps.youhaosuda.com/oauth2/authorize",
	TokenUrl:   "https://apps.youhaosuda.com/oauth2/token",
	ApiUrl:     "https://api.youhaosuda.com",
	ApiVersion: ApiV1Version,
}

type Config struct {
	AppKey     string
	AppSecret  string
	Scope      string
	AuthUrl    string
	TokenUrl   string
	ApiUrl     string
	ApiVersion string
}

type AccessToken struct {
	AccessToken string `json:"token"`
}

const (
	AuthUrl        = "https://apps.youhaosuda.com/oauth2/authorize/"
	TokenUrl       = "https://apps.youhaosuda.com/oauth2/token/"
	ApiUrl         = "https://api.youhaosuda.com/"
	ApiV1Version   = "v1"
	DefaultTimeOut = 60 * time.Second
)
