### Go-Start

对Gin的二次封装，快速开发 Web 服务

### 介绍

> 1. 这是一个基于 Go 语言、Gin 框架的 Web 项目骨架，对常用的库进行封装，开发者可以更快速搭建 Web 服务，更关注属于自己的业务
> 2. 项目要求 Go 版本 >= 1.15
> 3. 项目已经完成了 用户注册、用户登录、用户登出、用户信息查询 四个接口，拉取本项目骨架，在此基础上就可以快速开发自己的项目

### 快速上手

#### 前期准备

1. 安装go环境，version = 1.15
2. 安装Mysql（如果需要）
3. 安装Redis（如果需要）
4. 安装Mongodb（如果需要）

#### 1. 克隆项目

```shell
$ git clone git@github.com:mittacy/goods.git
```

##### 重命名项目

```shell
$ git clone git@github.com:mittacy/goods.git myProjectName
$ sed -i '' "s/goods/myProjectName/g" `grep goods -rl myProjectName`
```

#### 2. 数据库创建

```shell
$ cd myProjectName
$ mysql -u root -p < ./database/test.sql
```

#### 3. 修改配置文件

打开 `myProjectName/config/config.yaml`

修改对应的配置信息：

+ server 服务端口
+ mysql 地址、用户名、密码

#### 4. 启动服务

```shell
$ cd myProjectName
$ go mod download
$ go run cmd/api/api.go
初始化工作...
配置文件初始化成功
{"level":"INFO","ts":"2021-02-14T10:16:34.476+0800","caller":"logger/init.go:61","msg":"日志初始化成功"}
{"level":"INFO","ts":"2021-02-14T10:16:34.477+0800","caller":"utils/validator.go:14","msg":"validator校验器初始化成功"}
{"level":"INFO","ts":"2021-02-14T10:16:34.522+0800","caller":"database/init.go:51","msg":"Mysql初始化成功"}
{"level":"INFO","ts":"2021-02-14T10:16:34.523+0800","caller":"database/init.go:64","msg":"Redis初始化成功"}
{"level":"INFO","ts":"2021-02-14T10:16:34.523+0800","caller":"utils/jwt.go:26","msg":"JWT初始化成功"}
初始化工作完成
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

```

#### 5. 测试接口

浏览器打开 `http://localhost:5201/api/v1/user/1`，查看id为1的用户详细信息

注意

+ `5201` 为端口，具体根据配置文件变更
+ `v1` 为版本号，具体根据配置文件变更

#### 6. 生成API接口文档

安装 apidoc

```shell
$ npm install apidoc -g
```

修改controller控制器的注释后，更新 API 文档

```shell
# 生成 apidoc
$ cd myProjectName
$ apidoc -i app/ -o apidoc/
```

打开 `myProjectName/apidoc/index.html` 即可查看API接口文档

### 封装

| 技术          | 框架          | 版本     | 网址                                       |
| ------------- | ------------- | -------- | ------------------------------------------ |
| 配置文件读写  | viper         | v1.7.1   | https://github.com/spf13/viper             |
| 日志封装      | zap           | v1.16.0  | https://github.com/uber-go/zap             |
| 日志压缩      | lumberjack.v2 | v2.0.0   | https://github.com/natefinch/lumberjack    |
| Mysql         | gorm          | v1.20.12 | https://gorm.io/zh_CN/docs/index.html      |
| Redis         | redigo        | v1.8.3   | https://github.com/gomodule/redigo         |
| Mongodb         | qmgo        | v0.9.1   | https://github.com/qiniu/qmgo              |
| JWT认证       | jwt-go        | v3.2.0   | https://github.com/dgrijalva/jwt-go        |
| 数据校验      | validator     | v10.4.1  | https://github.com/go-playground/validator |
| 响应封装      | Gin           | v1.6.3   | https://github.com/gin-gonic/gin           |
| API文档生成器 | apidoc        | -        | https://apidocjs.com/#install              |

### 项目结构解析

![](http://static.mittacy.com/blog/20210214095536.jpg)

### 服务架构

