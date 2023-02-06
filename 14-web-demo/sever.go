package main

import (
	"LearnGoProject/14-web-demo/handler"
	"LearnGoProject/14-web-demo/repository"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	err := Init("14-web-demo/data/") //14-web-demo/data/post
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	r.GET("ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong"})
	})
	r.GET("community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := handler.QueryPageInfo(topicId)
		c.JSON(200, data)
	})
	err = r.Run()
	if err != nil {
		return
	}

}
func Init(filePath string) error {
	err := repository.Init(filePath)
	if err != nil {
		return err
	}
	return nil

}
