# lemon

一个开箱即用，高度集成的微服务组件库,可以快速开发集grpc,http,graphql的云原生微服务
### quick start
`go get github.com/hopeio/lemon@main`
#### install env tools
`install protoc`
`go get github.com/hopeio/lemon/tools/protoc@main`
`go run $(go list -m -f {{.Dir}}  github.com/hopeio/lemon)/tools/protoc/install-env.go`
#### generate protobuf
`protogen go -e -w (-q) (-v) -p $proto_path -g $proto_output_path`
#### use docker
`docker run --rm -v $project:/work jybl/goprotoc protogen go -e -w -p $proto_path -g $proto_output_path`

## initialize
基于反射自动注入的配置及dao注入初始化，并暴露一个全局变量，记录模块信息
![initialize](_readme/assets/initialize.webp)

### 一个应用的启动，应该如此简单
#### config（配置）
支持nacos,local file,http请求作为配置中心,可扩展支持etcd,apollo,viper(获取配置代理，底层是其他配置中心)，支持toml格式的配置文件，
支持dev，test，prod环境本，启动命令区分
仅需配置配置中心,后续配置均从配置中心拉取及自动更新
```toml
Module = "user"

[dev]
configType = "local"
Watch  = true
NoInject = ["Apollo","Etcd", "Es"]

[dev.local]
Debug = true
ConfigPath = "local.toml"
ReloadType = "fsnotify"

[dev.http]
Interval = 100
Url = "http://localhost:6666/local.toml"

[dev.nacos]
DataId = "pro"
Group = "DEFAULT_GROUP"

[[dev.nacos.ServerConfigs]]
Scheme = "http"
IpAddr = "nacos"
Port = 9000
GrpcPort = 10000

[dev.nacos.ClientConfig]
NamespaceId = ""
username = "nacos"
password = "nacos"
LogLevel = "debug"

```
```go
type config struct {
	//自定义的配置
	Customize serverConfig
	Server    initialize.ServerConfig
	Mail      initialize.MailConfig
	GORMDB    initialize.DatabaseConfig
	Redis     initialize.RedisConfig
	Cache     initialize.CacheConfig
	Log       initialize.LogConfig
	Viper     *viper.Viper
}

var Conf = &config{}

func (c *config) Init() {
	if runtime.GOOS == "windows" {
	}

	c.Customize.TokenMaxAge = time.Second * 60 * 60 * 24 * c.Customize.TokenMaxAge
}

func main() {
    //配置初始化应该在第一位
    defer initialize.Start(Conf, nil)()
}
```
如果还有Dao要初始化
```go
import(
    "github.com/hopeio/lemon/initialize/basic_dao/gormdb/postgres"
    initredis "github.com/hopeio/lemon/initialize/basic_dao/redis"
)
// dao dao.
type dao struct {
	// GORMDB 数据库连接
	GORMDB   *postgres.DB
	StdDB    *sql.DB
	// RedisPool Redis连接池
	Redis *redis.Client
}


func (d *dao) Init() {
	db := d.GORMDB
	db.Callback().Create().Remove("gorm:save_before_associations")
	db.Callback().Create().Remove("gorm:save_after_associations")
	db.Callback().Update().Remove("gorm:save_before_associations")
	db.Callback().Update().Remove("gorm:save_after_associations")

	d.StdDB, _ = db.DB()
}
func main() {
//配置初始化应该在第一位
defer initialize.Start(Conf, nil)()
}
```
原生集成了redis,gormdb(mysql,postgressql,sqlite),kafka,pebbledb,apollo,badgerdb,etcd,elasticsearch,nsq,ristretto,viper等，并且非常简单的支持自定义扩展,不局限于Dao对象，任何对象都支持根据配置自动注入生成
## context
一个轻量却强大的上下文管理器,一个请求会生成一个context，贯穿整个请求，context记录原始请求上下文，请求时间，客户端信息，权限校验信息，及负责判断是否内部调用，
及附带唯一traceId的日志记录器
其中权限校验采用jwt，具体的校验模型采用接口，可供使用方自定义
支持http及fasthttp,并支持自定义的请求类型
![context](_readme/assets/context.webp)

## [pick框架](https://github.com/actliboy/pick)
一个简单的易于开发的http api服务器,灵感来自于grpc和springmvc,基于反射自动注入
```go

package main

import (
	"github.com/hopeio/lemon/pick"
	"github.com/hopeio/lemon/utils/log"
)

func init(){
    pick.RegisterService(&service.UserService{},&other.Service{})
}

type UserService struct{}
//需要实现Service方法，返回该服务的说明，url前缀，以及需要的中间件
func (*UserService) Service() (string, string, []http.HandlerFunc) {
return "用户相关", "/api/user", []http.HandlerFunc{middleware.Log}
}


func (*UserService) Add(ctx *model.Claims, req *model.SignupReq) (*model.User, error) {
	pick.Api(func() {
            pick.Method(http.MethodPost).//定义请求的方法
            Title("用户注册").//接口描述
            Middleware(nil).//中间件
            //接口迭代信息
            CreateLog("1.0.0", "jyb", "2019/12/16", "创建").//创建，唯一
            ChangeLog("1.0.1", "jyb", "2019/12/16", "修改测试").//变更，可有多个
            End()
	})

	return &model.User{Name: "测试"}, nil
}

func main() {
    router := pick.NewRouter(true)
    router.ServeFiles("/static", "E:/")
    log.Info("visit http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}

```
上面代码将会自动创建POST /api/user/add 路由

pick的路由基于httprouter改造,如不想使用，pick同时兼容gin,fiber(fasthttp),底层可选择。

## server
lemon服务器，各种服务接口的保留，集成支持，一个服务暴露grpc,http,graphql接口
- 集成opencensus实现调用链路跟踪记录，配合context及utils-log 实现完整的请求链路日志记录
- 集成prometheus及pprof实现性能监控及性能问题排查
- 支持框架自生成的由gin提供支持的grpc转http，也支持原生的grpc-gateway
![server](_readme/assets/server.webp)

```go
package main

import (
	"github.com/hopeio/lemon/pick"
	"github.com/hopeio/lemon/utils/net/http/gin/handler"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hopeio/lemon/server"
	"github.com/hopeio/lemon/initialize"
	"user/protobuf/user"
	uconf "user/conf"
	udao "user/dao"
	userservice "user/service"
	"github.com/hopeio/lemon/utils/log"

	"go.opencensus.io/examples/exporter"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"google.golang.org/grpc"
)

func main() {
	//配置初始化应该在第一位
	defer initialize.Start(uconf.Conf, udao.Dao)()
	view.RegisterExporter(&exporter.PrintExporter{})
	view.SetReportingPeriod(time.Second)
	// Register the view to collect gRPC client stats.
	if err := view.Register(ocgrpc.DefaultClientViews...); err != nil {
		log.Fatal(err)
	}
	pick.RegisterService(userservice.GetUserService())
	server.Start(&server.Server{
		//为了可以自定义中间件
		GRPCOptions: []grpc.ServerOption{
			grpc.ChainUnaryInterceptor(),
			grpc.ChainStreamInterceptor(),
			//grpc.StatsHandler(&ocgrpc.ServerHandler{})
		},
		GRPCHandle: func(gs *grpc.Server) {
			user.RegisterUserServiceServer(gs, userservice.GetUserService())
		},
		GinHandle: func(app *gin.Engine) {
			_ = user.RegisterUserServiceHandlerServer(app, userservice.GetUserService())
			app.Static("/static", "F:/upload")
			pick.Gin(app, true, initialize.GlobalConfig.Module)
		},
        /*	GraphqlResolve: model.NewExecutableSchema(model.Config{
                Resolvers: &model.GQLServer{
                UserService:  service.GetUserService(),
                OauthService: service.GetOauthService(),
            }}),*/
	})
}

```

## protobuf
库中用到的protobuf定义及生成文件
![protobuf](_readme/assets/protobuf.webp)

### 生成库protobuf代码
`go run protobuf/generate.go`

## tools/protoc
本项目需要用到的protobuf插件，`go run tools/protoc/install-env.go`或者运行`go generate tools/protoc/install-tools.go`或者运行`tools/protoc/install-tools.sh`，会自动安装

- protogen为go语言写的protobuf生成程序
  - go/dart 生成go/dart文件，E.g: protogen go -p xxx -g xxx
  - -p proto dir
  - -g generate dir
  - (-d) 指定lemon proto dir,如项目引用本项目或使用jybl/protogen image 可省略
  - -e 是否使用enum扩展插件
  - -w 是否使用grpc-gateway插件
  - -v 是否使用validators插件
  - -q 是否使用graphql插件
  - --patch 是否使用原生protopatch
- protoc-go-patch 支持通过ast重新生成自定义结构体tag,生成结构体方法等功能
- protoc-gen-grpc-gin github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway的gin版本，支持生成http路由代理转发到grpc sercvice中
- protoc-gin-enum 分为错误enum及普通enum，生成性能更高支持中文的`String()`,错误enum会额外生成`Error()string`，支持生成枚举的辅助方法,错误enum会额外生成`Error()string`
- 集成github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 用于生成swagger文档
- 集成github.com/mwitkow/go-proto-validators/protoc-gen-govalidators 用于生成请求的校验的代码
- 集成github.com/danielvladco/go-proto-gql 用于生成graphql schema 及 grahpql服务

### build docker image
```base
`tools/protoc/docker_build.sh`
```
### upgrade go
`docker build -t jybl/goprotoc -f $lemon_dir/tools/protoc/Dockerfile-upgrade .`
### template
user.model.proto
```protobuf
syntax = "proto3";
package user;
import "user/user.enum.proto";
import "lemon/protobuf/utils/patch/go.proto";
import "protoc-gen-openapiv2/options/annotations.proto";

option java_package = "xyz.hoper.protobuf.user";
option go_package = "github.com/hopeio/lemon/protobuf/user";
// 用户
message User {
  uint64 id = 1 [(go.field) = {tags:'gorm:"primaryKey;"'}];
  string name = 2 [(go.field) = {tags:'gorm:"size:10;not null" annotation:"昵称"'}];
    // 性别，0未填写，1男，2女
  Gender gender = 8 [(go.field) = {tags:'gorm:"type:int2;default:0"'}, (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) = {
    type:INTEGER
  }];
}
```
user.enum.proto
```protobuf
syntax = "proto3";
package user;
import "lemon/protobuf/utils/enum/enum.proto";
import "lemon/protobuf/utils/patch/go.proto";

option (enum.enum_gqlgen_all) = true;

option java_package = "xyz.hoper.protobuf.user";
option go_package = "github.com/hopeio/lemon/protobuf/user";

option (enum.enum_prefix_all) = false;
option (go.file) = {no_enum_prefix:true};
// 用户性别
enum Gender{
    option (go.enum) = {stringer_name: 'OrigString'};
    GenderPlaceholder = 0 [(enum.enumvalue_cn)= "占位"];
    GenderUnfilled = 1 [(enum.enumvalue_cn)= "未填"];
    GenderMale = 2 [(enum.enumvalue_cn)= "男"];
    GenderFemale = 3 [(enum.enumvalue_cn)= "女"];
}

```

user.service.proto
```protobuf
syntax = "proto3";
package user;
import "user/user.model.proto";
import "user/user.enum.proto";
import "protoc-gen-openapiv2/options/annotations.proto";
import "github.com/mwitkow/go-proto-validators/validator.proto";
import "google/api/annotations.proto";
import "lemon/protobuf/utils/empty/empty.proto";
import "lemon/protobuf/utils/response/response.proto";
import "lemon/protobuf/utils/request/param.proto";
import "lemon/protobuf/utils/proto/gogo/graphql.proto";
import "lemon/protobuf/utils/oauth/oauth.proto";
import "lemon/protobuf/utils/patch/go.proto";
import "google/protobuf/wrappers.proto";
option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_swagger) = {
  info: {
    version: "1.0"
  }
};

service UserService {

  option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_tag) = {
    description: "用户相关接口"
  };
    //获取用户信息
  rpc Info (request.Object) returns (User) {
    option (google.api.http) = {
      get: "/api/v1/user/{id}"
    };
    option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
      tags:["用户相关接口", "v1.0.0"]
      summary : "获取用户信息"
      description : "根据Id获取用户信息接口"
    };
    option (gogo.graphql_operation) = Query;
  }

}
```
```sh
# 安装执行
`go get github.com/hopeio/lemon/tools/protoc@main`
`go run $(go list -m -f {{.Dir}}  github.com/hopeio/lemon)/tools/protoc/install-env.go`
`protogen go -p $proto_path -g $proto_output_path`
```
### docker
`docker run --rm -v $project_path:/work jybl/goprotoc protogen go --proto=$proto_path --genpath=$proto_output_path`

## utils

各种工具库

以下是一些可以单独成库的工具
### conctrl
并发控制，一个任务调度框架，可以控制goroutine数量,任务失败重试，任务衍生子任务执行，任务检测，任务统计
#### crawler
爬虫框架，基于conctrl



