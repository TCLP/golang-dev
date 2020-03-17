package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/TCLP/golang-dev/gin/authware"
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 中间件
func JWTAuthMiddleWare() func(c *gin.Context) {
	return func(c *gin.Context) {

		// 获取头部中的token字段
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		// 按空格分割token字段
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "请求头中的token格式错误",
			})
			c.Abort()
			return
		}
		//解析token
		mc, err := authware.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的token",
			})
			c.Abort()
			return
		}
		c.Set("username", mc.Username)
		c.Next()
	}
}

func authHandler(c *gin.Context) {
	var user UserInfo

	// 检查参数是否正确
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}

	// 验证用户名密码正确性
	if user.Username == "TCLP" && user.Password == "TCLP" {
		tokenstring, err := authware.GenToken(user.Username)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{"token": tokenstring},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 2002,
			"msg":  "用户名密码错误",
		})
	}
	return
}

func homeHandler(c *gin.Context) {
	username := c.MustGet("username").(string)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{"username": username},
	})
}

func main() {
	r := gin.Default()

	r.POST("/auth", authHandler)
	r.POST("/home", JWTAuthMiddleWare(), homeHandler)

	r.Run()
}
