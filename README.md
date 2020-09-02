你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
Golang 支持在一个平台下生成另一个平台可执行程序的交叉编译功能。

1. Mac 下编译 Linux, Windows 平台的 64 位可执行程序：

```bash
$ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build test.go
$ CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build test.go
```

1. Linux 下编译 Mac, Windows 平台的 64 位可执行程序：

```bash
$ CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build test.go
$ CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build test.go
```

1. Windows 下编译 Mac, Linux 平台的 64 位可执行程序：

```bash
$ SET CGO_ENABLED=0SET GOOS=darwin3 SET GOARCH=amd64 go build test.go
$ SET CGO_ENABLED=0 SET GOOS=linux SET GOARCH=amd64 go build test.go
```

注：如果编译 web 等工程项目，直接 cd 到工程目录下直接执行以上命令

`GOOS：目标可执行程序运行操作系统，支持 darwin，freebsd，linux，windows`
`GOARCH：目标可执行程序操作系统构架，包括 386，amd64，arm`
