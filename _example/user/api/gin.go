package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hopeio/lemon/_example/protobuf/user"
	"github.com/hopeio/lemon/_example/user/conf"
	"github.com/hopeio/lemon/_example/user/service"
	"github.com/hopeio/lemon/initialize"
	"github.com/hopeio/lemon/pick"
	gin2 "github.com/hopeio/lemon/pick/gin"
	"net/http"
)

func GinRegister(app *gin.Engine) {
	_ = user.RegisterUserServiceHandlerServer(app, service.GetUserService())
	//oauth.RegisterOauthServiceHandlerServer(app, service.GetOauthService())
	app.StaticFS("/oauth/login", http.Dir("./static/login.html"))

	pick.RegisterService(service.GetUserService())
	gin2.Register(app, conf.Conf.Server.GenDoc, initialize.GlobalConfig.Module, conf.Conf.Server.OpenTracing)

}
