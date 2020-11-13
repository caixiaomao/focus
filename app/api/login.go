package api

import (
	"focus/app/model"
	"focus/app/service"
	"focus/library/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var Login = new(loginApi)

type loginApi struct{}

// @summary 展示登录页面
// @tags    登录
// @produce html
// @router  /login [GET]
// @success 200 {string} html "页面HTML"
func (a *loginApi) Index(r *ghttp.Request) {
	service.View.Render(r, model.View{})
}

// @summary 提交登录
// @description 前面5次不需要验证码，同一个IP登录失败5次之后将会启用验证码校验。
// @description 注意提交的密码是明文。
// @description 登录成功后前端引导页面跳转。
// @tags    用户
// @produce json
// @param   passport    formData string true "账号"
// @param   password    formData string true "密码"
// @param   verify_code formData string false "验证码"
// @router  /login/do [POST]
// @success 200 {object} response.JsonRes "执行结果"
func (a *userApi) Do(r *ghttp.Request) {
	var (
		data            *model.UserApiLoginReq
		serviceLoginReq *model.UserServiceLoginReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(data, &serviceLoginReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.User.Login(r.Context(), serviceLoginReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "")
	}
}
