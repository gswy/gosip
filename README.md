# GoSIP

[![](https://img.shields.io/badge/license-MIT-green.svg)](https://github.com/gswy/GB28181-simulation/blob/main/LICENSE.txt)

## 项目简介
本项目是Go实现的SIP协议解析库（目前只针对于`GB/T 28181-2016`中`UDP`传输方式开发）。在国标协议UDP传输中采用一问一答的设计方式，所以参考Gin框架的使用方式，简化使用方式。


## 快速使用
```go
package main

import (
    "github.com/gswy/gosip"
    "log"
)

func main() { 
    // 创建一个SIP服务器 
    engine := gosip.NewEngine("192.168.0.121", 9999) 
    // `RFC 3261`中规定的请求方法 
    engine.Register(func(ctx *gosip.Context) {
        ctx.RESP(200, *gosip.Response{})
    }) 
    // 启动服务器 
    err := engine.Run()
    if err != nil {
        log.Fatalf("服务器启动失败: %v", err)
    }
}
```
