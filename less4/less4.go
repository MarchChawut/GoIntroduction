package main

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoginForm struct{
	Username string `form: "username" binding: "required"`
	Password string `form: "password" binding: "required"`
}

func main() {
	r := gin.Default()

	//อ่านค่า string query
	r.GET("/login", func(c *gin.Context) {
		username, password := c.Query("username"), c.Query("password")
		//                    hashmap{key: value}
		c.JSON(http.StatusOK, gin.H{"Result": "ok", "username": username, "password": password})
	})
	//-------------------------------------------------------------------------------------------------

	//Gin POST formUrlEncoded
	r.POST("/login", func(c *gin.Context) {
		//ประกาศ object
		var form LoginForm

		// ใช้ shouldbind เพื่อผูก object โดยใช้ &(address) ในการ pass value
		// ถ้า shouldbind แล้วไม่มี error == nil
		if c.ShouldBind(&form) == nil {
			if form.Username == "admin" && form.Password == "1234" {
				msg := fmt.Sprintf("Login Successful with: %s, %s", form.Username, form.Password)
				c.JSON(200, gin.H{"Status": msg})
			}else {
				c.JSON(401, gin.H{"Status": "User/Password Failed"})
			}
		}else {
			c.JSON(401, gin.H{"Status": "Enable to bind Data"})
		}
		
	})

	r.Run(":85")
}
