package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"param/input"
	"param/validation"
)
func main() {
	r := gin.Default()
	r.POST("/regist", Regist)
	r.POST("/pay", Pay)
	r.Run()
}
func Regist(c *gin.Context) {
	var in input.User
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 400, "data": nil, "msg": "请输入正确的参数类型!"})
		return
	}
	valid := validation.Validation{}
	isok_param, _ := valid.Valid(&in)
	if !isok_param {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, gin.H{"code": 400, "data": nil, "msg": err.Message})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": in, "msg": "good!"})
}
func Pay(c *gin.Context) {
	var in input.Config
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusOK, gin.H{"code":400, "data": nil, "msg": "请输入正确的参数类型!"})
		return
	}
	valid := validation.Validation{}
	isok_param, _ := valid.Valid(&in)
	if !isok_param {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, gin.H{"code": 400, "data": nil, "msg": err.Message})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": in, "msg": "good!"})
}
