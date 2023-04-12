package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	runningDir, _ := os.Getwd()
	count := 0

	errlogfile, _ := os.OpenFile(fmt.Sprintf("%s/gin_error.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	accesslogfile, _ := os.OpenFile(fmt.Sprintf("%s/gin_access.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)

	gin.DefaultErrorWriter = errlogfile
	gin.DefaultWriter = accesslogfile

	r.Use(gin.Logger())

	// //custom format Logger
	// r.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {

	// }))

	r.GET("/home", func(c *gin.Context) {
		count = count + 1
		accesslogfile.WriteString(fmt.Sprintf("Count : %d\n", count))
		c.Data(http.StatusOK, "text.html; charset=utf-8", []byte("1234"))
	})

	//อ่านค่า string query
	r.GET("/login", func(c *gin.Context) {
		username, password := c.Query("username"), c.Query("password")
		//                    hashmap{key: value}  
		c.JSON(http.StatusOK, gin.H{"Result": "ok", "username": username, "password": password})
	})

	r.Run(":85")
}