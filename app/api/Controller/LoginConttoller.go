package Controller

import (
	"gin_demo/Common"
	"gin_demo/Model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginController 定义接收数据的结构体
type (
	LoginController struct{}
	form            struct {
		ID       uint   `form:"id" json:"id"`
		Name     string `form:"username" json:"user" uri:"user" xml:"user" `
		Tel      string `form:"tel" json:"tel" uri:"tel" xml:"tel" `
		Password string `form:"password" json:"password" uri:"password"`
	}
)

// Login
// @Description: 登录接口
// @receiver l
// @param c
func (l *LoginController) Login(c *gin.Context) {
	// 接收数据
	var form form
	// 绑定form表单
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"绑定": err.Error()})
		return
	}
	UserModel := Model.NewUserModel()
	user, err := UserModel.FindInfoByTel(form.Tel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err, "user": user})
		return
	}
	// 校验用户是否正确
	if form.Tel != user.Tel || Common.MD5(form.Password) != user.Password {
		c.JSON(http.StatusOK, gin.H{"status": "200", "msg": "账号或密码错误！"})
		return
	}
	// 重定向
	c.JSON(http.StatusOK, gin.H{"status": 200, "msg": "登录成功"})
	return
}

// Register 注册
// @Description: 用户注册
// @receiver l
// @param c
func (l *LoginController) Register(c *gin.Context) {
	var form Model.User
	// 绑定form表单
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	UserModel := Model.NewUserModel()
	count := UserModel.FindCountByTel(form.Tel)
	if count >= 1 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "该电话已注册，请换个号码"})
		return
	}
	form.Password = Common.MD5(form.Password)
	err := UserModel.Insert(form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "恭喜注册成功！", "code": "200"})

}

func (l *LoginController) Success(c *gin.Context) {

	c.Redirect(http.StatusMovedPermanently, "https://www.5lmh.com")
}

func (l *LoginController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "恭喜注册成功！", "code": "200"})
}
