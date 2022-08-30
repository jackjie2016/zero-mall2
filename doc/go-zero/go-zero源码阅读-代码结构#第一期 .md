go-zero 作为一个微服务框架，不仅给我们提供了很好的参考，而且核心代码量不多，我们可以在闲暇时间读读他的核心代码，来多多学习充电。

## rest 部分
## 代码结构
```bash
rest
├── handler // 自带中间件
│   ├── authhandler.go // 权限
│   ├── breakerhandler.go // 断路器
│   ├── contentsecurityhandler.go // 安全验证
│   ├── cryptionhandler.go // 加密解密
│   ├── gunziphandler.go // zip 压缩
│   ├── loghandler.go // 日志
│   ├── maxbyteshandler.go // 最大请求数据限制
│   ├── maxconnshandler.go // 最大请求连接数限制
│   ├── metrichandler.go // 请求指标统计
│   ├── prometheushandler.go // prometheus 上报
│   ├── recoverhandler.go // 错误捕获
│   ├── sheddinghandler.go // 过载保护
│   ├── timeouthandler.go // 超时控制
│   └── tracinghandler.go // 链路追踪
├── httpx
│   ├── requests.go
│   ├── responses.go
│   ├── router.go
│   ├── util.go
│   └── vars.go
├── internal
│   ├── cors // 跨域处理
│   │   └── handlers.go
│   ├── response
│   │   ├── headeronceresponsewriter.go
│   │   └── withcoderesponsewriter.go
│   ├── security // 加密处理
│   │   └── contentsecurity.go
│   ├── log.go
│   └── starter.go
├── pathvar // path 参数解析
│   └── params.go
├── router
│   └── patrouter.go
├── token
│   └── tokenparser.go
├── config.go // 配置
├── engine.go // 引擎
├── server.go
└── types.go
 
 ```
###服务启动流程
我们以 go-zero-example 项目 http/demo/main.go 代码来分析
![image.png](https://pic1.zhimg.com/80/v2-c382bf3001b15bc26dca256a85f0c37c_720w.jpg)

go-zero 给我们提供了如下组件与服务，我们来逐一阅读分析
 
```bash

http框架常规组件（路由、调度器、中间件、跨域）
权限控制
断路器
限流器
过载保护
prometheus
trace
cache
```
### http框架常规组件

### 路由

路由使用的是二叉查找树，高效的路由都会使用树形结构来构建

二叉查找树可参见源码

https://github.com/zeromicro/go-zero/tree/master/core/search

go-zero 路由实现了 http\server.go Handler interface 来拦截每个请求

入口源码地址: https://github.com/zeromicro/go-zero/rest/router/patrouter.go

```bash

func (pr *patRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reqPath := path.Clean(r.URL.Path) // 返回相当于path的最短路径名称
	if tree, ok := pr.trees[r.Method]; ok { // 查找对应 http method
		if result, ok := tree.Search(reqPath); ok { // 查找路由 path 
			if len(result.Params) > 0 {
				r = pathvar.WithVars(r, result.Params) // 获取路由参数并且添加到 *http.Request 中
			}
			result.Item.(http.Handler).ServeHTTP(w, r) // 调度方法
			return
		}
	}

	allows, ok := pr.methodsAllowed(r.Method, reqPath)
	if !ok {
		pr.handleNotFound(w, r)
		return
	}

	if pr.notAllowed != nil {
		pr.notAllowed.ServeHTTP(w, r)
	} else {
		w.Header().Set(allowHeader, allows)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
```
###调度器
go-zero 没有调度器，在上文 ServeHTTP 中已经使用了调度器，这归结于 golang 已经给我们实现了一个很好的 http 模块，如果是其他语言，我们在设计框架的时候往往要自己实现调度器。
###中间件
我们可以在 *.api 中添加如下代码来使用
```bash
@server(
	middleware: Example // 路由中间件声明
)
service User {
	@handler UserInfo
	post /api/user/userinfo returns (UserInfoResponse)
}
```
通过生成代码命令，生成的代码如下

```bash

package middleware

import (
	"log"
	"net/http"
)

type ExampleMiddleware struct{}

func NewExampleMiddleware() *ExampleMiddleware {
	return &ExampleMiddleware{}
}

func (m *ExampleMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		next(w, r)
	}
}
```
go-zero 给我们提供了一些常用的中间件，方便我们在开发时候使用


rest.WithCors() 跨域设置
```bash

// example
server := rest.MustNewServer(c.RestConf, rest.WithCors("localhost:8080"))

// 源码
func WithCors(origin ...string) RunOption {
	return func(server *Server) {
		server.router.SetNotAllowedHandler(cors.NotAllowedHandler(nil, origin...))
		server.Use(cors.Middleware(nil, origin...))
	}
}
```
###跨域
resrt.WithCustomCors() 自定义跨域方法

````bash
// example
var origins = []string{
	"localhost:8080",
}
server := rest.MustNewServer(c.RestConf,
	rest.WithCustomCors(
        // 设置 http header
		func(header http.Header) {
			header.Set("Access-Control-Allow-Origin", "Access-Control-Allow-Origin")
		},
        // 不允许地址返回指定数据
		func(writer http.ResponseWriter) {
			writer.Write([]byte("not allow"))
		},
        // 允许跨域地址
		origins...,
	),
)

// 源码
func WithCustomCors(middlewareFn func(header http.Header), notAllowedFn func(http.ResponseWriter),
	origin ...string) RunOption {
	return func(server *Server) {
		server.router.SetNotAllowedHandler(cors.NotAllowedHandler(notAllowedFn, origin...))
		server.Use(cors.Middleware(middlewareFn, origin...))
	}
}
````
rest.WithJwt() jwt

```bash
// example
rest.WithJwt("uOvKLmVfztaXGpNYd4Z0I1SiT7MweJhl")

// 源码
func WithJwt(secret string) RouteOption {
	return func(r *featuredRoutes) {
		validateSecret(secret)
		r.jwt.enabled = true
		r.jwt.secret = secret
	}
}
```
rest.WithJwtTransition() jwt token 转换，新老 token 可以同时使用
```bash
// example
rest.WithJwtTransition("uOvKLmVfztaXGpNYd4Z0I1SiT7MweJhl", "uOvKLmVfztaXGpNYd4Z0I1SiT7MweJh2")

// 源码
func WithJwtTransition(secret, prevSecret string) RouteOption {
	return func(r *featuredRoutes) {
		// why not validate prevSecret, because prevSecret is an already used one,
		// even it not meet our requirement, we still need to allow the transition.
		validateSecret(secret)
		r.jwt.enabled = true
		r.jwt.secret = secret
		r.jwt.prevSecret = prevSecret
	}
}
```