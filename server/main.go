package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin"
	"tgoj/server/global"
	"tgoj/server/model"
)



func main() {
	//r := gin.Default()
	//
	//r.GET("/problems")

	r := gin.Default()
	r.GET("/problems", func(c *gin.Context) {
		var problems []model.Question
		global.DB.Limit(10).Find(&problems)
		fmt.Println(len(problems))
		c.JSON(200, problems)
	})

	r.GET("/problem/:id", func(c *gin.Context) {
		id := c.Param("id")
		var problem model.Question
		global.DB.Where("id = ?", id).First(&problem)
		c.JSON(200, problem)
	})

	r.Run(":8080")
}
