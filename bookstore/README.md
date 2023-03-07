# 用 Go 语言构建了一个最简单的 HTTP 服务，一个接近真实的图书管理 API 服务。



## 项目目录结构如下

```
├── cmd/
│   └── bookstore/         // 放置bookstore main包源码
│       └── main.go
├── go.mod                 // module bookstore的go.mod
├── go.sum
├── internal/              // 存放项目内部包的目录
│   └── store/
│       └── memstore.go     
├── server/                // HTTP服务器模块
│   ├── middleware/
│   │   └── middleware.go
│   └── server.go          
└── store/                 // 图书数据存储模块
    ├── factory/
    │   └── factory.go
    └── store.go
```

### 项目 main 包

main 包是主要包，为了搞清楚各个模块之间的关系，我在这里给出了 main 包的实现逻辑图

![图片](https://user-images.githubusercontent.com/90596113/223327568-86164824-47fb-41da-942d-c2312d33df5e.png)


### 图书数据存储模块（store)

图书数据存储模块的职责很清晰，就是用来**存储整个 bookstore 的图书数据**的。图书数据存储有很多种实现方式，最简单的方式莫过于在内存中创建一个 map，以图书 id 作为 key，来保存图书信息



### HTTP 服务模块（server）

HTTP 服务模块的职责是**对外提供 HTTP API 服务，处理来自客户端的各种请求，并通过 Store 接口实例执行针对图书数据的相关操作**



## 编译、运行与验证

```
go mod tidy

go build bookstore/cmd/bookstore

./main
//以下是日志输出
2022/11/15 16:08:36 web server start ok
```

如果你看到上面这个输出的日志，说明我们的程序启动成功了。



仿客户端向 bookstore 服务发起请求了，比如创建一个新书条目：

```
curl -X POST -H "Content-Type:application/json" -d '{"id": "978-7-111-55842-2", "name": "The Go Programming Language", "authors":["Alan A.A.Donovan", "Brian W. Kergnighan"],"press": "Pearson Education"}' localhost:8080/book
```

此时服务端会输出如下日志，表明我们的 bookstore 服务收到了客户端请求。

```
2022/11/15 16:09:10 recv a POST request from [::1]:58021
```

同时支持GET请求
