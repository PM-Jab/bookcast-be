package servers

import (
	_userHttp "bookcast/modules/users/controllers"
	_usersRepository "bookcast/modules/users/repositories"
	_usersService "bookcast/modules/users/services"

	"github.com/gin-gonic/gin"
)

func (s *Server) MapHandlers() {
	v1 := s.App.Group("/v1")

	usersGroup := v1.Group("/users")
	usersRepository := _usersRepository.NewUsersRepository(s.Db)
	usersService := _usersService.NewUsersService(usersRepository)
	_userHttp.NewUsersController(usersGroup, usersService)

	// Not Found response
	s.App.Use(func(c *gin.Context) {
		c.AbortWithStatusJSON(404, gin.H{
			"error": "Not Found",
		})
	})

}
