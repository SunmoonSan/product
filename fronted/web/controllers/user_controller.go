package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"product/datamodels"
	"product/services"
	"product/tool"
	"strconv"
)

type UserController struct {
	Ctx     iris.Context
	Service services.IUserService
	Session *sessions.Session
}

func (c *UserController) GetRegister() mvc.View {
	return mvc.View{
		Name: "user/register.html",
	}
}

func (c *UserController) PostRegister() {
	var (
		nickName = c.Ctx.FormValue("nickName")
		userName = c.Ctx.FormValue("userName")
		password = c.Ctx.FormValue("password")
	)

	user := &datamodels.User{
		UserName:     userName,
		Nickname:     nickName,
		HashPassword: password,
	}

	_, err := c.Service.AddUser(user)
	c.Ctx.Application().Logger().Debug(err)
	if err != nil {
		c.Ctx.Redirect("/user/error")
		return
	}
	c.Ctx.Redirect("/user/login")
	return
}

func (c *UserController) GetLogin() mvc.View {
	return mvc.View{
		Name: "user/login.html",
	}
}

func (c *UserController) PostLogin() mvc.Response {
	var (
		userName = c.Ctx.FormValue("userName")
		password = c.Ctx.FormValue("password")
	)
	user, isOK := c.Service.IsPwdSuccess(userName, password)
	if !isOK {
		return mvc.Response{
			Path: "/user/login",
		}
	}

	tool.GlobalCookie(c.Ctx, "uid", strconv.FormatInt(user.ID, 10))
	c.Session.Set("userID", strconv.FormatInt(user.ID, 10))
	return mvc.Response{
		Path: "/product/",
	}
}
