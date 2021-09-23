package main

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	// 可自定义盐值
	TokenSalt = "default_salt"
)

func MD5(data []byte) string {
	_md5 := md5.New()
	_md5.Write(data)
	return hex.EncodeToString(_md5.Sum([]byte("")))
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query("username") // 用户名
		ts := c.Query("ts")             // 时间戳
		token := c.Query("token")       // 访问令牌

		if strings.ToLower(MD5([]byte(username+ts+TokenSalt))) == strings.ToLower(token) {
			// 验证通过，会继续访问下一个中间件
			c.Next()
		} else {
			// 验证不通过，不再调用后续的函数处理
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"message": "访问未授权"})
			// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
			return
		}
	}
}

func ServiceWithoutAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "这是一个不用经过认证就能访问的接口"})
}

func ServiceWithAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "这是一个需要经过认证才能访问的接口，看到此信息说明验证已通过"})
}
