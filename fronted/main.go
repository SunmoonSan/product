package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
	"product/common"
	"time"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	tmplate := iris.HTML("./fronted/web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(tmplate)
	app.StaticWeb("/public", "./fronted/web/public")
	app.StaticWeb("/html", "./fronted/web/htmlProductShow")
	app.OnAnyErrorCode(func(ctx context.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错!"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})
	db, err := common.NewMysqlConn()
	if err != nil {

	}

	sess := sessions.New(sessions.Config{
		Cookie:                      "AdminCookie",
		Expires:                     600 * time.Minute,
	})

	//ctx, cancel := context.
}


