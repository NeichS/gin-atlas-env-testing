package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"entdemo/ent"
)

type CreateUserDTO struct {
	Age  int    `form:"age" json:"age" binding:"required"`
	Name string `form:"name" json:"name" binding:"required"`
}

type Handlers struct {
	entClient *ent.Client
}

func NewHandler(client *ent.Client) *Handlers {
	return &Handlers{
		entClient: client,
	}
}

func (h *Handlers) HandleCreateUser(c *gin.Context) {
	objA := CreateUserDTO{}
	// This c.ShouldBind consumes c.Request.Body and it cannot be reused.
	if err := c.ShouldBindBodyWith(&objA, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUsr, err := h.entClient.User.Create().SetName(objA.Name).SetAge(objA.Age).Save(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"usr": newUsr})
}
