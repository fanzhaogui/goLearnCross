package controllers

import (
	"csdn/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUser(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("接收的参数值为：",id)
	var person model.Person
	//person = model.FetchSingleUser(id)
	person.FetchSingleUser(id)
	c.AbortWithStatusJSON(200, person)
	return
}
