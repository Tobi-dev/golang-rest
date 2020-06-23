package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../model"
	"../helpers"
	"../service"
	"../repository"
)

func InitContactRoutes(route *gin.Engine, contactRepository *repository.ContactRepository){
	route.POST("/create", func(context *gin.Context){
		var contact model.Contact

		err := context.ShouldBindJSON(&contact)

		if err != nil {
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		code := http.StatusOK

		response := service.CreateContact(&contact, *contactRepository)

		for i := 0; i < 1000; i++ {
			service.CreateContact(&contact, *contactRepository)
		}

		if !response.Success{
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/", func(context *gin.Context){
		code := http.StatusOK

		response := service.FindAllContacts(*contactRepository)

		if !response.Success{
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/show/:id", func(context *gin.Context){
		id := context.Param("id")

		code := http.StatusOK

		response := service.FindOneContactById(id, *contactRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.PUT("/update/:id", func(context *gin.Context) {
		id := context.Param("id")

		var contact model.Contact

		err := context.ShouldBindJSON(&contact)

		// validation errors
		if err != nil {
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		code := http.StatusOK

		response := service.UpdateContactById(id, &contact, *contactRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.DELETE("/delete/:id", func(context *gin.Context){
		id := context.Param("id")

		code := http.StatusOK

		response := service.DeleteContactById(id, *contactRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})
}
