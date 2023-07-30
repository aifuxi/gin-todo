> 学习 Gin 框架的练习，实现了一个 todo 的 CRUD

## 启动

1. 先安装 `air` ，用来支持开发中服务热重载

```shell
go install github.com/cosmtrek/air@latest
```

2. 安装好后，使用 `air init` 初始化 `air` 的配置文件，然后执行 `air` 就能启动服务

3. 浏览器打开 `localhost:9000`

## 部署

1. 在服务器上 git clone ，把代码下下来

2. build 一下

```shell
go build .
```

3. 使用 `nohup` 来让进程在后台挂起

```
GIN_MODE=release nohup ./gin-todo &
```

`GIN_MODE=release` 是设置 `Gin` 为 release 模式

## NGINX 反向代理

> 这里只贴一下核心配置

```conf
location /gin-todo/ {
  proxy_pass http://localhost:9000/;
}
```

这里踩了一个坑，使用 NGINX 进行反向代理的时候，HTML 发请求的路径不能以`/`开头，比如在 HTML 中，想要向`/get_todos`发请求，请求路径应该写`get_todos`，去掉开头的`/`。因为我们的是需要以当前访问的地址作为请求的前缀的，这样 NGINX 才能反代的上

比如现在访问的`https://demos.aifuxi.cool/gin-todo/`，在这个页面发送的请求，以`get_todos`举例

- `/get_todos`，请求路径以`/`开头时，请求的地址是`https://demos.aifuxi.cool/get_todos`，无法匹配 NGINX 的规则，❌
- `get_todos`，请求路径不以`/`开头时，请求的地址是以当前访问的地址最为基准，发出请求的地址是`https://demos.aifuxi.cool/gin-todo/get_todos`，可以匹配到 NGINX 的规则，✅
