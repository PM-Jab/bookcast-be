package controllers

import (
	"bookcast/modules/entities"

	"github.com/gin-gonic/gin"
)

type usersController struct {
	UsersSer entities.UsersService
}

func NewUsersController(router *gin.RouterGroup, usersService entities.UsersService) {
	controllers := &usersController{
		UsersSer: usersService,
	}
	router.POST("/api/v1/users/register", controllers.Register)
	router.POST("/api/v1/users/login", controllers.Login)
}

func (u *usersController) Register(c *gin.Context) {
	req := new(entities.UsersRegisterReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	res, err := u.UsersSer.Register(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, res)
}

func (u *usersController) Login(c *gin.Context) {
	req := new(entities.UsersLoginReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	res, err := u.UsersSer.Login(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, res)
}
