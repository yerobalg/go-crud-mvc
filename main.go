package main

import (
	"crud-user-mvc/bootstrap"
	"crud-user-mvc/controllers"
	"crud-user-mvc/helpers"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	///init env variable
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("failed to load env from local file")
		panic(err)
	}

	//init database
	db, err := bootstrap.InitMySQL()
	if err != nil {
		panic(err)
	}

	//init helper
	hp := helper.Init()

	//init controller
	ctrl := controller.Init(db, *hp)

	//init router
	router := gin.Default()
	router.Group("/api/v1")

	router.POST("/users", ctrl.CreateNewUser)
	router.GET("/users", ctrl.GetUserList)
	router.PUT("/users/:id", ctrl.UpdateUser)
	router.DELETE("/users/:id", ctrl.DeleteUser)

	router.Run(":8080")
}
