// package Controller
// @author: congyukun
// @since: 2021/12/15
// @desc: UserController

package Controller

import (
	"gin_demo/Model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	UserController struct{}
)

func (user *UserController) GetUserInfo(c *gin.Context) {
	// tel := c.DefaultPostForm("tel", "")
	tel := c.PostForm("tel")

	if tel == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "tel不能为空"})
		return
	}
	// var info Model.User
	newUser := Model.NewUserModel()
	where := make(map[string]interface{}, 1)
	where["tel"] = tel
	info, err := newUser.FindOne(where)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": err})
	}
	if info.ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "该账户不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登录成功"})
	return

}
