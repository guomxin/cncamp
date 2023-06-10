package main

import (
	"fmt"
	"module3/pkg/apis"
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	pprof.Register(r)

	r.GET("/", func(c *gin.Context) {
		c.Writer.Write([]byte("你好, gin!"))
	})

	r.GET("/pinfos", func(c *gin.Context) {
		c.JSON(200, []apis.PersonalInformation{
			{
				Name:   "小强",
				Sex:    "男",
				Tall:   1.7,
				Weight: 71,
				Age:    35,
			},

			{
				Name:   "大强",
				Sex:    "男",
				Tall:   1.73,
				Weight: 65,
				Age:    40,
			},
		})
	})

	r.POST("/register", func(c *gin.Context) {
		pi := &apis.PersonalInformation{}
		if err := c.BindJSON(pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "无法解析信息",
			})
			return
		}

		fmt.Println(pi)
		c.JSON(http.StatusOK, pi)
	})

	r.GET("/history/:name", func(c *gin.Context) {
		name, _ := c.Params.Get("name")
		c.Writer.Write([]byte(name))
	})

	r.Run(":8088")
}
