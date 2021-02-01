package cors

import "github.com/gin-gonic/gin"

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("interview/access-Control-Allow-Origin", "*")                            //跨域
		ctx.Header("interview/access-Control-Allow-Headers", "Token,Content-Type")          //必须的请求头
		ctx.Header("interview/access-Control-Allow-Methods", "OPTIONS,PUT,POST,GET,DELETE") //接收的请求方法
	}
}
