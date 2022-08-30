package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	errorx "zero-mal/common/error"
	"zero-mal/common/initialize"

	"zero-mal/service/common/api/internal/config"
	"zero-mal/service/common/api/internal/handler"
	"zero-mal/service/common/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "service/common/api/etc/common-api.yaml", "the config file")

//var configFile = flag.String("f", "etc/common-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	//引入自定义的表单验证 并且加入翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err.Error())
	}
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

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
