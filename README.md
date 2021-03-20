# lit-edu-go
 
一个集成查询成绩课表等功能的服务后端, 其本质为洛阳理工学院青果教务在线的中间代理实现

![School Term](https://img.shields.io/badge/dynamic/json?color=blue&label=school%20trem&query=%24.data.jw_time.term&url=https%3A%2F%2Flit.icepie.net%2Fapi%2Fv2%2Fjw%2Fstatus)
![School Week](https://img.shields.io/badge/dynamic/json?color=9cf&label=school%20week&query=%24.data.jw_time.week&url=https%3A%2F%2Flit.icepie.net%2Fapi%2Fv2%2Fjw%2Fstatus)
![Kingosoft Online Num](https://img.shields.io/badge/dynamic/json?color=brightgreen&label=kingosoft%20online%20num&query=%24.data.online_number&url=https%3A%2F%2Flit.icepie.net%2Fapi%2Fv2%2Fjw%2Fstatus)

[![Go Report Card](https://goreportcard.com/badge/github.com/icepie/lit-edu-go)](https://goreportcard.com/badge/github.com/icepie/lit-edu-go)
[![License](https://img.shields.io/github/license/icepie/lit-edu-go)](https://github.com/icepie/lit-edu-go/blob/main/LICENSE)
[![QQ Group](https://img.shields.io/badge/qq%20group-768887710-red.svg)](https://jq.qq.com/?_wv=1027&k=lz0XyN86)
[![TG Group](https://img.shields.io/badge/tg%20group-lit_edu-blue.svg)](https://t.me/lit_edu)


## 安装

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

请查看文档

## 许可证

[AGPLv3](https://github.com/icepie/lit-edu-go/blob/main/LICENSE)

