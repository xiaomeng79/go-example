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
|监控|[go-metrics](https://github.com/rcrowley/go-metrics)|
|打包|[bindata](https://github.com/jteeuwen/go-bindata)|
|编码|[protoc-gen-micro](https://github.com/micro/protoc-gen-micro)|
|部署|docker docker-compose k8s|
|文档生成|swagger|
|其他|JWT|

## 目录结构

```go
.
├── api //restful接口
├── cinit //公共配置和初始化
├── cmd //服务入口
├── data //测试数据
├── deployments //部署目录,docker docker-compose k8s配置文件,自动化生成
├── go.mod //go1.11包管理
├── go.sum
├── internal //内部公共组件
├── LICENSE
├── Makefile 
├── README.md
├── scripts //makefile使用的脚本
├── srv //服务目录
└── third_party //第三方包目录

```
## 依赖安装

-  安装 protoc protoc-gen-micro  protoc-gen-go 

[安装说明](https://github.com/micro/protoc-gen-micro)

-  安装docker和docker-compose


- 安装bindata打包

[安装说明](https://github.com/jteeuwen/go-bindata#installation)

## 编译
```go
make vendor 
make allbuild
```
## 本地docker-compose运行

```go
make compose
```

## 请求

```go
//登录,返回token在响应头中:Authorization
curl -X POST http://127.0.0.1:8888/common/v1/login -H 'Content-Type: application/json' -d '{"username":"xiaomeng01","password":"123456"}' -i  
```

## 查看效果

1. 链路跟踪:http://127.0.0.1:16686 [本地效果](http://127.0.0.1:16686)

1. 监控:http://127.0.0.1:3000 [本地效果](http://127.0.0.1:3000) 用户名:admin 密码:admin