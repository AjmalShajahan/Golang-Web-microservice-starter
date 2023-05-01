package controllers

import (
	"net/http"
	"web-api/src/models"
	service "web-api/src/service"
	"web-api/utils/constant"
	"web-api/utils/response"
	val "web-api/utils/validator"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// CreatePost : API to Create data into the database
func CreatePost(c *gin.Context) {
	reqModel := models.Posts{}
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

func ReadPost(c *gin.Context) {
	reqModel := &models.Posts{}
	if err := c.ShouldBind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	// if err := val.ValidateVariable(reqModel.ID, "required", "id"); err != nil {
	// 	c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
	// 	return
	// }
	var service = service.PostService{}
	readPost, err := service.FetchPosts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(readPost))
}

func FindPostsByUser(c *gin.Context) {
	reqModel := &models.Users{}
	if err := c.ShouldBind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.ValidateVariable(reqModel.ID, "required", "id"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.PostService{}
	readPosts, err := service.FindPostsByUser(reqModel.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(readPosts))
}
func CreateComment(c *gin.Context) {
	reqModel := models.Comments{}
	if err := c.ShouldBind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.CommentService{}
	saved, err := service.CreateComment(&reqModel)
	if err != nil {
		log.Error().Msgf("Error inserting data into the database: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	log.Info().Msgf("Comment Inserted Succesfully in the database: %s", saved)
	c.JSON(http.StatusOK, response.SuccessResponse(saved))
}
