package Controller

import (
	"gin_demo/Model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserController 定义接收数据的结构体
type UserController struct {
}

func (l *UserController) Login(c *gin.Context) {
	// 接收数据
	var form Model.User
	// 绑定form表单
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//User := Model.NewUserModel()


	// 校验用户是否正确
	if form.Name != "root" || form.Password != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	// 重定向
	c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	//c.JSON(http.StatusOK, gin.H{"status": 200})

	// fmt.Println(UserController,password)
}

func (l *UserController) Register(c *gin.Context)  {
	var form Model.User
	// 绑定form表单
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var datas []Model.User
	UserModel := Model.NewUserModel()
	datas = append(datas,form)
	ids, err := UserModel.Inserts(datas)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"msg":"存储失败！"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ids": ids})
	return
	c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")

}






func (l *UserController) Success(c *gin.Context) {

	c.Redirect(http.StatusMovedPermanently, "http://www.5lmh.com")
}

func (l *UserController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是测试", "ce": "123456"})
}
