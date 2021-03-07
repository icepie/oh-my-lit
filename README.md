# lit-edu-go
 
一个集成查询成绩课表等功能的服务后端, 其本质为洛阳理工学院青果教务在线的中间代理实现

[![Go Report Card](https://goreportcard.com/badge/github.com/icepie/lit-edu-go)](https://goreportcard.com/badge/github.com/icepie/lit-edu-go)
[![License](https://img.shields.io/github/license/icepie/lit-edu-go)](https://github.com/icepie/lit-edu-go/blob/main/LICENSE)
[![QQ Group](https://img.shields.io/badge/qq%20group-768887710-red.svg)](https://jq.qq.com/?_wv=1027&k=lz0XyN86)
[![TG Group](https://img.shields.io/badge/tg%20group-lit_edu-blue.svg)](https://t.me/lit_edu)
![Kingosoft Online Num](https://img.shields.io/badge/dynamic/json?color=brightgreen&label=Kingosoft%20Online%20Num&query=%24.data.online_number&url=https%3A%2F%2Flit.icepie.net%2Fapi%2Fv1%2Fjw%2Fstatus)

## 安装

### get

```bash
$ go get -u github.com/icepie/lit-edu-go
```

### build

```bash
$ git clone  http://github.com/icepie/lit-edu-go
$ cd lit-edu-go
$ go build # GO111MODULE=on
```

### release

前往 [release](https://github.com/icepie/lit-edu-go/releases) 页面, 下载与你适用的版本

## 配置

请自行查看并编辑配置文件 `conf.yaml`

## 使用

```bash
$ lit-edu-go
```

## 接口

> 如在配置中开启了 `jwauth` 则部分接口需要 `password` 参数

### 检查青果登录状态

```js
GET    /api/v1/jw/status
```

### 通过学号获取基本信息


```js
POST   /api/v1/jw/profile
```

```json
{
  "stuid": "B19071121",
  "password": "PASSWORD"
}
```

### 通过学号获取成绩

```js
POST   /api/v1/jw/score
```

```json
{
  "stuid": "B19071121",
  "password": "PASSWORD"
}
```

### 获取课表

```
TODO
```

## 许可证

[AGPLv3](https://github.com/icepie/lit-edu-go/blob/main/LICENSE)

