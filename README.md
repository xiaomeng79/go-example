# go-micro-example

[![Build Status](https://travis-ci.org/xiaomeng79/go-example.svg?branch=master)](https://travis-ci.org/xiaomeng79/go-example) [![codecov](https://codecov.io/gh/xiaomeng79/go-example/branch/master/graph/badge.svg)](https://codecov.io/gh/xiaomeng79/go-example)
[![GitHub license](https://img.shields.io/github/license/xiaomeng79/go-example.svg)](https://github.com/xiaomeng79/go-example/blob/master/LICENSE)

## 这只是一个技术使用的示例项目



## 使用的技术

|功能|描述|
|---|---|
|框架|go-micro(微服务) + echo(web框架)|
|配置|默认值->yaml->env|
|日志|可选插件(zap logors),集成了链路跟踪[go-log](https://github.com/xiaomeng79/go-log)|
|链路跟踪|OpenTracing [Jaeger](https://github.com/jaegertracing/jaeger)|
|监控||
|打包|[bindata](https://github.com/jteeuwen/go-bindata)|
|编码|[protoc-gen-micro](https://github.com/micro/protoc-gen-micro)|
|部署|docker docker-compose k8s|
|其他|JWT|

## 依赖安装

-  安装 protoc protoc-gen-micro  protoc-gen-go 

[安装说明](https://github.com/micro/protoc-gen-micro)

-  安装docker和docker-compose


- 安装bindata打包

[安装说明](https://github.com/jteeuwen/go-bindata#installation)

## 编译
```go
make allbuild
```
## 本地运行

```go
make compose
```

## 请求

```go
//登录,返回token在响应头中:Authorization
curl -X POST   http://127.0.0.1:8888/common/v1/login   -H 'Cache-Controlapplication/json'    -d '{"username":"xiaomeng01","password":"123456"}' -i   
```

## 查看效果

1. 链路跟踪:http://127.0.0.1:16686 [本地效果](http://127.0.0.1:16686)