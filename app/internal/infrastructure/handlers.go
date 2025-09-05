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
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		c.String(http.StatusOK, `body looking good!`)
	} else {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("errA: %v", errA))
		return
	}

	fmt.Println(objA)
}
