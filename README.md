# lit-edu-go
 
一个集成查询成绩课表等功能的服务后端, 其本质为洛阳理工学院青果教务在线的中间件

![Go Report Card](https://goreportcard.com/badge/github.com/icepie/lit-edu-go)
![License](https://img.shields.io/github/license/icepie/lit-edu-go)

[![QQ Group](https://img.shields.io/badge/QQ%20Group-768887710-red.svg)](https://jq.qq.com/?_wv=1027&k=lz0XyN86)
[![TG Group](https://img.shields.io/badge/TG%20Group-lit_edu-blue.svg)](https://t.me/lit_edu)

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

## 接口 (Restful API)

### 检查青果登录状态

```js
GET    /api/v1/jw/status
```

### 通过学号获取成绩

```js
POST   /api/v1/jw/score
Content-type: application/json
```

```json
{
  "stuid": "B19071121",
  "password": "PASSWORD(jwauth: true)"
}
```

### 通过学号获取课表

```
TODO
```

## 许可证

[AGPLv3](https://github.com/icepie/lit-edu-go/blob/main/LICENSE)

