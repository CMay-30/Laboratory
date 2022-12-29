package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api/auth/register", func(c *gin.Context) {
		// 获取参数
		name := context.PostForm("name")
		telephone := context.PostForm("telephone")
		password := context.PostForm("password")
		// 数据验证
		if len(telephone) != 11 {
			context.JSON(http.StatusUnprocessableEntity, map[string]interface{}{"code": 422, "msg": "手机号必须为11位"})
		}
		if len(password) != 6 {
			context.JSON(http.StatusUnprocessableEntity, map[string]interface{}{"code": 422, "msg": "密码不能少于6位"})
		}

		// 如果名称没有传，给一个10位的随机字符串
		if len(name) == 0 {
			name = RandomString(10)
		}

		log.Println(name, telephone, password)
		// 判断手机号是否存在

		// 创建用户

		// 返回结果
		context.JSON(200, gin.H{
			"msg": "注册成功",
		})
	})
	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}

func RandomString(n int) {
	var letters = []byte("aasdfghjklqwertyuiopzxcvbnmASDFGHJKLQWERTYUIOPZXCVBNM")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
}
