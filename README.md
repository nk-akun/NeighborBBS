# NeighborBBS

## 简介

NeighborBBS 是一款论坛项目，后端基于 Go 编写，前端基于 Vue 框架编写，前端页面参考了小码哥编写的[bbs-go](https://github.com/mlogclub/bbs-go)，是本人的练手项目之一，适合用来学习和使用。

## 特性

- [x] 注册/登陆模块(用户名或邮箱登陆)
- [x] 设置昵称、邮箱、用户名
- [x] 发表动态、文章
- [x] 评论系统
- [x] 动态/文章的点赞
- [x] 支持浏览器 token 记住登录
- [x] 支持文章或评论流式获取
- [x] 支持 markdown 语法发表文章或评论
- [ ] 站内信
- [ ] 用户资料编辑
- [ ] 文章标签管理

## 技术选型

- 后端：整体使用 golang 编写，用 Gin 框架搭建 API 部分
- 包管理：go-mod
- 配置文件：使用 viper 实现的 yaml 格式的配置文件
- 日志：基于 zap 实现的日志系统
- 数据库：使用 mysql-5.7，采用 gorm 库来操作数据库
- 前端：基于 Vue.js 编写，使用 Nuxt.js 快速构建和渲染前端

## 目录结构

```
.
├── LICENSE
├── api             (API文件夹)
├── bbs.yaml        (配置文件)
├── build.sh        (构建脚本，构建可在linux上运行的二进制文件)
├── config          (配置包)
├── logs            (日志包)
├── main.go         (main函数)
├── middleware      (中间件)
├── model           (结构体)
├── nbbs.service    (linux服务配置文件)
├── repository      (数据库层)
├── service         (服务层)
├── util            (通用工具)
├── site            (前端)
│   ├── Dockerfile  (docker文件)
│   ├── app.html    (app)
│   ├── assets      (静态文件)
│   ├── common      (通用工具)
│   ├── components        (通用组件)
│   ├── jsconfig.json     (配置)
│   ├── layouts           (布局)
│   ├── middleware        (中间件)
│   ├── nuxt.config.js    (nuxt配置)
│   ├── pages             (页面组件)
│   ├── plugins           (插件)
│   ├── start.sh          (运行脚本)
│   ├── static            (静态文件)
│   ├── store             (vuex状态管理仓)
│   └── utils             (通用工具)
```

## 安装说明

### 1.获取源码

```shell
git clone https://github.com/mlogclub/mlog.git
```

### 2.创建 mysql 中的数据库

在 mysql 中创建好 database，在步骤 3 中填入 database 的信息，无需创建数据表

示例：

```shell
ceate database neighborbbs;
```

### 3.修改配置

修改 bbs.yaml 文件，配置 mysql、服务端口、日志等信息

示例：

```yaml
mysql:
  host: 127.0.0.1
  port: 3306
  username: root
  password: 123456
  dbname: neighborbbs
```

### 4.启动后端

> 如果没有 go 环境，请先安装和配置 go 环境

####安装依赖

```shell
go mod download
```

#### 启动服务

**方式一**

```shell
go run main.go
```

**方式二**

```shell
go build        #编译项目
./NeighborBBS   #执行二进制
```

**方式三**

```shell
./build.sh      #编译成linux可执行文件
#上传到linux服务器运行
```

### 5.启动前端

> 如果没有 npm 环境，请先安装 npm 环境

#### (1) 进入 site 目录下

```shell
cd site
```

####(2) 在 nuxt.config.js 文件中配置启动端口等信息(可选项)

#### (3)安装依赖

```shell
npm install
```

#### (4)启动前端服务

```shell
npm run dev
```
