package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"goweb计算器/Calculate"
	"io/ioutil"
	"net/http"
)

func calculate1(c *gin.Context){
	buf := make([]byte, 1024)
	n,_ := c.Request.Body.Read(buf)
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(buf[:n]))
	j := buf[0:n]                         		//j就是Response payload中的值
	//fmt.Print(j)								//打印j的字符串形式
	fmt.Println(string(j))         				//打印出j的值，便于测试
	result := Calculate.Calculate(string(j))	//将j的字符串形式传入计算函数当中
	c.JSON(http.StatusOK , result)          		//将结果以json形式返回
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		fmt.Println(method)
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
}
}

func main(){
	r := gin.Default()//返回默认的路由引擎

	r.Use(Cors())//开启中间件  允许使用跨域请求

	r.POST("/calculator",calculate1)//处理POST请求

	_ = r.Run(":8800")//启动服务
}