package api

import (
	"higo/serializer"
	"higo/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
// @Summary 用户注册
// @Description 注册用户
// @Accept  json
// @Produce  json
// @Success 200
// @Router /user/register/ [post]
func UserRegister(c *gin.Context) {
	var service service.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
// @Summary 用户登录
// @Description 用户登录接口
// @Accept  json
// @Produce  json
// @Success 200
// @Router /user/login/ [post]
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserMe 用户详情
// @Summary 用户详情
// @Description 获取用户详情
// @Accept  json
// @Produce  json
// @Success 200
// @Router /user/me/ [get]
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.BuildUserResponse(*user)
	c.JSON(200, res)
}

// UserLogout 用户登出
// @Summary 用户登出
// @Description 登出用户
// @Accept  json
// @Produce  json
// @Success 200
// @Router /user/logout/ [delete]
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
