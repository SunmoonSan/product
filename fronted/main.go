package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"product/common"
	"product/fronted/middleware"
	"product/fronted/web/controllers"
	"product/repositories"
	"product/services"
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

	//sess := sessions.New(sessions.Config{
	//	Cookie:                      "AdminCookie",
	//	Expires:                     600 * time.Minute,
	//})

	ctx, cancel := context

	user := repositories.NewUserRepository("user", db)
	userService := services.NewService(user)
	userPro := mvc.New(app.Party("/user"))
	userPro.Register(userService, ctx,)
	userPro.Handle(new(controllers.UserController))

	// 注册product控制器
	product := repositories.NewProductManager("product", db)
	productService := services.NewProductService(product)
	order := repositories.NewOrderManagerRepository("order", db)
	orderService := services.NewOrderService(order)
	proProduct := app.Party("/product")
	pro := mvc.New(proProduct)
	proProduct.Use(middleware.AuthConProduct)
	pro.Register(productService, orderService)
	pro.Handle(new(controllers.ProductController))

	app.Run(
		iris.Addr("0.0.0.0:8082"),
		iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
