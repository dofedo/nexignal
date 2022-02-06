package user

import (
	f "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Sign-Up handler
func SignUp(c *gin.Context) {
	var su_form SUForm
	if err := c.ShouldBindJSON(&su_form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err, message := su_form.Validate()

	if err {
		c.JSON(http.StatusBadRequest, gin.H{"error": message})
		return
	}

	err, detail := su_form.Create()
	if err {
		f.Println("here")
		c.JSON(detail.code, gin.H{"error": detail.header})
		return
	}
	return
}

// Sign-Up handler
func Signin(c *gin.Context) {
	var si_form SIForm
	if err := c.ShouldBindJSON(&si_form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err, message := si_form.Validate()

	if err {
		c.JSON(http.StatusBadRequest, gin.H{"error": message})
		return
	}

	err, detail := si_form.Get()
	if err {
		f.Println("here")
		c.JSON(detail.code, gin.H{"error": detail.header})
		return
	}
	return
}
