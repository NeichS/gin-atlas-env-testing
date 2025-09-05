package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type CreateUserDTO struct {
	Age  int    `form:"age" json:"age" binding:"required"`
	Name string `form:"name" json:"name" binding:"required"`
}

func HandleCreateUser(c *gin.Context) {
	objA := CreateUserDTO{}
	// This c.ShouldBind consumes c.Request.Body and it cannot be reused.
	if err := c.ShouldBindBodyWith(&objA, binding.JSON); err == nil {
		resp := "user created successfully"
		c.JSON(http.StatusCreated, gin.H{"status": resp})
	} else {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(objA)
}
