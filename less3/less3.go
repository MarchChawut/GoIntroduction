package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/home", func(c *gin.Context) {
		c.Data(http.StatusOK, "text.html; charset=utf-8", []byte("1234"))
	})

	//อ่านค่า string query
	r.GET("/login", func(c *gin.Context) {
		username, password := c.Query("username"), c.Query("password")
		//                    hashmap{key: value}
		c.JSON(http.StatusOK, gin.H{"Result": "ok", "username": username, "password": password})
	})
//---------------------------------------------------------------------------------------------------------
	//gin get parameter (pass parameter ผ่าน url) && outside function
	r.GET("/flight/:from/:to/:vehicle", handleBookRequest)
	//http://localhost:85/flight/thai/japan/airplain

	r.Run(":85")
}

func handleBookRequest(c *gin.Context) {
	from, to, vehicle := c.Param("from"), c.Param("to"), c.Param("vehicle")
	c.JSON(200, gin.H{"Result": "OK", "From": from, "To": to, "Vehicle": vehicle})
}