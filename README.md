# yhsd-api-go

友好速搭应用开发 Golang SDK

[![Build status](https://img.shields.io/travis/yeezon/yhsd-api-go.svg?style=flat-square)](https://travis-ci.org/yeezon/yhsd-api-go)
[![Coverage Status](https://img.shields.io/coveralls/yeezon/yhsd-api-go.svg?style=flat-square)](https://coveralls.io/repos/yeezon/yhsd-api-go)
[![Dependency Status](https://img.shields.io/david/yeezon/yhsd-api-go.svg?style=flat-square)](https://david-dm.org/yeezon/yhsd-api-go)

## 安装

```go
go get github.com/yeezon/yhsd-api-go
```

## 使用方法

```go
import "github.com/yeezon/yhsd-api-go"
```

### 私有应用

配置如下：
```go
var conf = youhaosuda.Config{
  // App Key
  AppKey:    "ab3217683c964c82a685c22d9440f240",
  // Shared Secret
  AppSecret: "13516ce822b841ce8d5b91630d97d050",
}
```

在首次使用时，需要先获取 access token：

```go
var privateApp = youhaosuda.NewPrivateApp(&conf, "")

err := privateApp.GenerateToken()

access_token := privateApp.AccessToken
```

首次获取 access token 后，可自行存储，之后直接使用：

```go
var privateApp = youhaosuda.NewPrivateApp(&conf, "应用的 access token")
```

调用友好速搭的开放 API：

```go
// 获取店铺数据
res := privateApp.Get("shop")


data := `
 {
        "redirect": {
          "path": "/123",
          "target": "/blogs"
        }
    }
`

// 创建 redirect
res := privateApp.Post("redirects", data)

// 修改 redirect
res := privateApp.Put("redirects/23", data)

// 删除 redirect
res := privateApp.Delete("redirects/23")
```

### 开放应用

配置如下：
```go
var conf = youhaosuda.Config{
  // 用 , 间隔的权限，默认：read_basic,write_basic
  Scope:    "read_basic,write_basic",
  // App Key
  AppKey:    "ab3217683c964c82a685c22d9440f240",
  // Shared Secret
  AppSecret: "13516ce822b841ce8d5b91630d97d050",
}
```

获取 access token：

```go
var publicApp = youhaosuda.NewPublicApp(&conf, "")

// 在收到客户的安装请求后，拼接安装确认地址
var confirmUrl := publicApp.AuthorizeUrl("应用回调地址","安装请求中的 shop key 参数值", "需要返回的参数")

// 客户确认安装后，会请求上面参数中的回调地址，并带 code 参数，通过 code 值来获取 access token
err := publicApp.GenerateToken("应用回调地址", "获取的 code 值")

access_toekn = publicApp.AccessToken
```

首次获取 access token 后，可自行存储，之后直接使用：

```go
var publicApp = youhaosuda.NewPublicApp(&conf, "应用的 access token")
```

调用友好速搭的开放 API：

```go
// 获取店铺数据
res := publicApp.Get("shop")

data := `
 {
        "redirect": {
          "path": "/123",
          "target": "/blogs"
        }
    }
`

// 创建 redirect
res := publicApp.Post("redirects", data)

// 修改 redirect
res := publicApp.Put("redirects/23", data)

// 删除 redirect
res := publicApp.Delete("redirects/23")
```
