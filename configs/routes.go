package configs

import (
	"../repository"
	"github.com/gin-gonic/gin"
	contactController "../controller"
)

func SetupRoutes(contactRepository *repository.ContactRepository) *gin.Engine {
	route := gin.Default()
	contactController.InitContactRoutes(route, contactRepository)

	return route
}