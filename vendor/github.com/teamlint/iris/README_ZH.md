<!-- # Iris Web Framework <a href="README_ZH.md"> <img width="20px" src="https://iris-go.com/images/flag-china.svg?v=10" /></a> <a href="README_RU.md"><img width="20px" src="https://iris-go.com/images/flag-russia.svg?v=10" /></a> <a href="README_ID.md"> <img width="20px" src="https://iris-go.com/images/flag-indonesia.svg?v=10" /></a> <a href="README_GR.md"><img width="20px" src="https://iris-go.com/images/flag-greece.svg?v=10" /></a> <a href="README_PT_BR.md"><img width="20px" src="https://iris-go.com/images/flag-pt-br.svg?v=10" /></a> <a href="README_JPN.md"><img width="20px" src="https://iris-go.com/images/flag-japan.svg?v=10" /></a> -->

# Iris <a href="README.md"> <img width="20px" src="https://iris-go.com/images/flag-unitedkingdom.svg?v=10" /></a> <a href="README_GR.md"><img width="20px" src="https://iris-go.com/images/flag-greece.svg?v=10" /></a> <a href="README_ES.md"><img width="20px" src="https://iris-go.com/images/flag-spain.png" /></a> <a href="README_KO.md"><img width="20px" src=""https://iris-go.com/images/flag-south-korea.svg" />

[![build status](https://img.shields.io/travis/kataras/iris/master.svg?style=for-the-badge)](https://travis-ci.org/kataras/iris) [![report card](https://img.shields.io/badge/report%20card-a%2B-ff3333.svg?style=for-the-badge)](https://goreportcard.com/report/github.com/kataras/iris)<!--[![godocs](https://img.shields.io/badge/go-%20docs-488AC7.svg?style=for-the-badge)](https://godoc.org/github.com/kataras/iris)--> [![view examples](https://img.shields.io/badge/learn%20by-examples-0077b3.svg?style=for-the-badge)](https://github.com/kataras/iris/tree/master/_examples) [![chat](https://img.shields.io/gitter/room/iris_go/community.svg?color=blue&logo=gitter&style=for-the-badge)](https://gitter.im/iris_go/community) [![release](https://img.shields.io/badge/release%20-v11.2-0077b3.svg?style=for-the-badge)](https://github.com/kataras/iris/releases)

Iris 是基于 Go 编写的一个快速，简单但功能齐全且非常高效的 Web 框架。 它为您的下一个网站或 API 提供了一个非常富有表现力且易于使用的基础。

看看 [其他人如何评价 Iris](https://iris-go.com/testimonials/)，同时欢迎各位点亮 **star**。

> 新版本 11.2 发布! [散布消息](https://www.facebook.com/iris.framework/posts/3276606095684693).

<<<<<<< HEAD
看看[别人是如何评价 Iris](#support)，同时欢迎各位点亮 Iris [Star](https://github.com/teamlint/iris/stargazers)，或者关注 [Iris facebook 主页](https://facebook.com/iris.framework)。
=======
## 学习 Iris
>>>>>>> upstream/master

<details>
<summary>快速入门</summary>

```sh
# 假设文件已经存在
$ cat example.go
```

```go
package main

import "github.com/teamlint/iris"

func main() {
    app := iris.Default()
    app.Get("/ping", func(ctx iris.Context) {
        ctx.JSON(iris.Map{
            "message": "pong",
        })
    })

    app.Run(iris.Addr(":8080"))
}
```

```sh
# 运行 example.go
# 在浏览器中访问 http://localhost:8080/ping
$ go run example.go
```

> 路由由 [muxie](https://github.com/kataras/muxie) 提供支持，muxie 是基于 Go 编写的最强大最快速的基于 trie 的路由

</details>

<<<<<<< HEAD
```sh
$ go get -u github.com/teamlint/iris
```
=======
Iris 包含详细而完整的 **[文档](https://github.com/kataras/iris/wiki)**，使你很容易开始使用该框架。
>>>>>>> upstream/master

要了解更多详细的技术文档，可以访问我们的 [godocs](https://godoc.org/github.com/kataras/iris)。对于可执行代码，可以随时访问示例代码，在仓库的 [_examples](_examples/)  目录下。

### 你喜欢在旅行中看书吗？

你现在可以 [获取](https://bit.ly/iris-req-book) PDF版本和在线访问我们的 **电子书** 并参与 Iris 的开发。

[![https://iris-go.com/images/iris-book-overview.png](https://iris-go.com/images/iris-book-overview.png)](https://bit.ly/iris-req-book)

## 贡献

我们很高兴看到你对 Iris Web 框架的贡献！有关为 Iris 做出贡献的更多信息，请查看 [CONTRIBUTING.md](CONTRIBUTING.md)。

[所有贡献者名单](https://github.com/kataras/iris/graphs/contributors)

<<<<<<< HEAD
- [更新记录](HISTORY_ZH.md#fr-11-january-2019--v1111) 是您最好的朋友，它包含有关最新功能和更改的信息
- 你碰巧找到了一个错误？ 请提交 [github issues](https://github.com/kataras/iris/issues)
- 您是否有任何疑问或需要与有经验的人士交谈以实时解决问题？ [加入我们的聊天](https://chat.iris-go.com)
- [点击这里完成我们基于表单的用户体验报告](https://docs.google.com/forms/d/e/1FAIpQLSdCxZXPANg_xHWil4kVAdhmh7EBBHQZ_4_xSZVDL-oCC_z5pA/viewform?usp=sf_link) 
- 你喜欢这个框架吗？ Twitter 上关于 Iris 的评价:

<a href="https://twitter.com/gelnior/status/769100480706379776"> 
    <img src="https://comments.iris-go.com/comment27_mini.png" width="350px">
</a>

<a href="https://twitter.com/MeAlex07/status/822799954188075008"> 
    <img src="https://comments.iris-go.com/comment28_mini.png" width="350px">
</a>

<a href="https://twitter.com/_mgale/status/818591490305761280"> 
    <img src="https://comments.iris-go.com/comment29_mini.png" width="350px">
</a>
<a href="https://twitter.com/VeayoX/status/813273328550973440"> 
    <img src="https://comments.iris-go.com/comment30_mini.png" width="350px">
</a>

<a href="https://twitter.com/pvsukale/status/745328224876408832"> 
    <img src="https://comments.iris-go.com/comment31_mini.png" width="350px">
</a>

<a href="https://twitter.com/blainsmith/status/745338092211560453"> 
    <img src="https://comments.iris-go.com/comment32_mini.png" width="350px">
</a>

<a href="https://twitter.com/tjbyte/status/758287014210867200"> 
    <img src="https://comments.iris-go.com/comment33_mini.png" width="350px">
</a>

<a href="https://twitter.com/tangzero/status/751050577220698112"> 
    <img src="https://comments.iris-go.com/comment34_mini.png" width="350px">
</a>

<a href="https://twitter.com/tjbyte/status/758287244947972096"> 
    <img src="https://comments.iris-go.com/comment33_2_mini.png" width="350px">
</a>

<a href="https://twitter.com/ferarias/status/902468752364773376"> 
    <img src="https://comments.iris-go.com/comment41.png" width="350px">
</a>


[如何贡献代码](CONTRIBUTING.md)

[贡献者列表](https://github.com/teamlint/iris/graphs/contributors)

## 学习

首先，从 Web 框架开始的最正确的方法是学习 Golang 标准库 [net/http](https://golang.org/pkg/net/http/ "net/http") 的基础知识，如果您的 web 应用程序是一个非常简单的个人项目，没有性能和可维护性要求，您可能只需使用标准库即可。 之后，遵循以下指导原则：

- 浏览 **100+** **[例子](_examples)** 和 我们提供的 [一些入门经验](#iris-starter-kits)
- 通过 [godocs](https://godoc.org/github.com/teamlint/iris) 阅读细节
- 准备一杯咖啡或茶，无论你喜欢什么，并阅读我们为你推荐的 [一些文章](#articles)

### Iris 入门

<!-- table form
| Description | Link |
| -----------|-------------|
| Hasura hub starter project with a ready to deploy golang helloworld webapp with IRIS! | https://hasura.io/hub/project/hasura/hello-golang-iris |
| A basic web app built in Iris for Go |https://github.com/gauravtiwari/go_iris_app |
| A mini social-network created with the awesome Iris💖💖 | https://github.com/iris-contrib/Iris-Mini-Social-Network |
| Iris isomorphic react/hot reloadable/redux/css-modules starter kit | https://github.com/iris-contrib/iris-starter-kit |
| Demo project with react using typescript and Iris | https://github.com/ionutvilie/react-ts |
| Self-hosted Localization Management Platform built with Iris and Angular | https://github.com/iris-contrib/parrot |
| Iris + Docker and Kubernetes | https://github.com/iris-contrib/cloud-native-go |
| Quickstart for Iris with Nanobox | https://guides.nanobox.io/golang/iris/from-scratch |
-->

1. [snowlyg/IrisApiProject: Iris + gorm + jwt + sqlite3](https://github.com/snowlyg/IrisApiProject) **NEW-Chinese**
2. [yz124/superstar: Iris + xorm to implement the star library](https://github.com/yz124/superstar) **NEW-Chinese**
3. [jebzmos4/Iris-golang: A basic CRUD API in golang with Iris](https://github.com/jebzmos4/Iris-golang)
4. [gauravtiwari/go_iris_app: A basic web app built in Iris for Go](https://github.com/gauravtiwari/go_iris_app)
5. [A mini social-network created with the awesome Iris💖💖](https://github.com/iris-contrib/Iris-Mini-Social-Network)
6. [Iris isomorphic react/hot reloadable/redux/css-modules starter kit](https://github.com/iris-contrib/iris-starter-kit)
7. [ionutvilie/react-ts: Demo project with react using typescript and Iris](https://github.com/ionutvilie/react-ts)
8. [Self-hosted Localization Management Platform built with Iris and Angular](https://github.com/iris-contrib/parrot)
9. [Iris + Docker and Kubernetes](https://github.com/iris-contrib/cloud-native-go)
10. [nanobox.io: Quickstart for Iris with Nanobox](https://guides.nanobox.io/golang/iris/from-scratch)
11. [hasura.io: A Hasura starter project with a ready to deploy Golang hello-world web app with IRIS](https://hasura.io/hub/project/hasura/hello-golang-iris)

> 如果你有类似的使用经验吗 [请提交给我们](https://github.com/teamlint/iris/pulls)!

### 中间件

Iris 拥有大量的中间件 [[1]](middleware/)[[2]](https://github.com/iris-contrib/middleware) 供您的 Web 应用程序使用。 不过，您并不局限于此，您可以自由使用与 [net/http](https://golang.org/pkg/net/http/) 包兼容的任何第三方中间件，相关示例 [_examples/convert-handlers](_examples/convert-handlers) 。

### 相关文章（英文）

* [CRUD REST API in Iris (a framework for golang)](https://medium.com/@jebzmos4/crud-rest-api-in-iris-a-framework-for-golang-a5d33652401e)
* [A Todo MVC Application using Iris and Vue.js](https://hackernoon.com/a-todo-mvc-application-using-iris-and-vue-js-5019ff870064)
* [A Hasura starter project with a ready to deploy Golang hello-world web app with IRIS](bit.ly/2lmKaAZ)
* [Top 6 web frameworks for Go as of 2017](https://blog.usejournal.com/top-6-web-frameworks-for-go-as-of-2017-23270e059c4b)
* [Iris Go Framework + MongoDB](https://medium.com/go-language/iris-go-framework-mongodb-552e349eab9c)
* [How to build a file upload form using DropzoneJS and Go](https://hackernoon.com/how-to-build-a-file-upload-form-using-dropzonejs-and-go-8fb9f258a991)
* [How to display existing files on server using DropzoneJS and Go](https://hackernoon.com/how-to-display-existing-files-on-server-using-dropzonejs-and-go-53e24b57ba19)
* [Iris, a modular web framework](https://medium.com/@corebreaker/iris-web-cd684b4685c7)
* [Go vs .NET Core in terms of HTTP performance](https://medium.com/@kataras/go-vs-net-core-in-terms-of-http-performance-7535a61b67b8)
* [Iris Go vs .NET Core Kestrel in terms of HTTP performance](https://hackernoon.com/iris-go-vs-net-core-kestrel-in-terms-of-http-performance-806195dc93d5)
* [How to Turn an Android Device into a Web Server](https://twitter.com/ThePracticalDev/status/892022594031017988)
* [Deploying a Iris Golang app in hasura](https://medium.com/@HasuraHQ/deploy-an-iris-golang-app-with-backend-apis-in-minutes-25a559bf530b)
* [A URL Shortener Service using Go, Iris and Bolt](https://medium.com/@kataras/a-url-shortener-service-using-go-iris-and-bolt-4182f0b00ae7)

### 视频教程（英文） - Youtube

* [Daily Coding - Web Framework Golang: Iris Framework]( https://www.youtube.com/watch?v=BmOLFQ29J3s) by WarnabiruTV
* [Tutorial Golang MVC dengan Iris Framework & Mongo DB](https://www.youtube.com/watch?v=uXiNYhJqh2I&list=PLMrwI6jIZn-1tzskocnh1pptKhVmWdcbS) (19 parts so far) by Musobar Media
* [Go/Golang 27 - Iris framework : Routage de base](https://www.youtube.com/watch?v=rQxRoN6ub78) by stephgdesign
* [Go/Golang 28 - Iris framework : Templating](https://www.youtube.com/watch?v=nOKYV073S2Y) by stephgdesignn
* [Go/Golang 29 - Iris framework : Paramètres](https://www.youtube.com/watch?v=K2FsprfXs1E) by stephgdesign
* [Go/Golang 30 - Iris framework : Les middelwares](https://www.youtube.com/watch?v=BLPy1So6bhE) by stephgdesign
* [Go/Golang 31 - Iris framework : Les sessions](https://www.youtube.com/watch?v=RnBwUrwgEZ8) by stephgdesign

### 工作机会
=======
## 安全漏洞
>>>>>>> upstream/master

如果你发现在 Iris 存在安全漏洞，请发送电子邮件至 [iris-go@outlook.com](mailto:iris-go@outlook.com)，所有安全漏洞都会被及时解决。

## 授权协议

项目名称 "Iris" 的灵感来自于希腊神话。

Iris Web 框架授权基于 [3-Clause BSD License](LICENSE) 许可的免费开源软件。
