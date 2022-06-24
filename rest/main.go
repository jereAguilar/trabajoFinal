package main

import (
	"rest/services/users"

	"https://github.com/gin-gonic/gin"
)

func main() {
	srv := service.NewUserService()
	router := gin.Default()

	router.GET("/users", service.GetUser)

	router.POST("/users", func(c *gin.Context){
		var u service.User
		if err := c.ShouldBindJSON(&u); err != nil{
			c.String(http.StatusBadRequest, "error al crear user")
		}
		u = srv.Add(u)
		c.JSON(http.StatusCreated, u)
	})
	router.PUT("/users", service.UpdateUser)
	router.Run()
}
