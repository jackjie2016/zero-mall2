package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	errorx "zero-mal/common/error"
	"zero-mal/common/initialize"
	"zero-mal/service/user/api/internal/config"
	"zero-mal/service/user/api/internal/handler"
	"zero-mal/service/user/api/internal/svc"
)

var configFile = flag.String("f", "service/user/api/etc/user-api-dev.yaml", "the config file")

//var configFile = flag.String("f", "etc/user-api-dev.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	//引入自定义的表单验证 并且加入翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err.Error())
	}

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
