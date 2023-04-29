package controllers

import (
	"net/http"
	"web-api/src/models"
	service "web-api/src/service"
	"web-api/utils/constant"
	"web-api/utils/response"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// CreatePost : API to Create data into the database
func CreatePost(c *gin.Context) {
	reqModel := models.Post{}
	if err := c.ShouldBind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.PostService{}
	saved, err := service.CreatePost(&reqModel)
	if err != nil {
		log.Error().Msgf("Error inserting data into the database: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	log.Info().Msgf("Post Inserted Succesfully in the database: %s", saved)
	c.JSON(http.StatusOK, response.SuccessResponse(saved))
}
