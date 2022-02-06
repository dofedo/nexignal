package main

import (
	"nexignal/ss/user"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()

	// user component router
	u_router := router.Group("/u")
	{
		// Signing In, Out and Up
		// u_router.GET("/s", user.Ping)
		// u_router.GET("/so", user.Ping)
		u_router.POST("/si", user.Signin)
		u_router.POST("/su", user.SignUp)
	}

	router.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func main() {

	// v := user.SIForm{Identifier: "uuser2", Password: "user2pass"}
	// v.Get()
	user.GetRows()
	// Router starts the server and does routing
	// Router()
}
