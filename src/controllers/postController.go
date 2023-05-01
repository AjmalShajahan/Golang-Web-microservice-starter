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

// func FindPostsByUser(c *gin.Context) {
// 	reqModel := &models.Users{}
// 	if err := c.ShouldBind(&reqModel); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
// 		return
// 	}
// 	if err := val.ValidateVariable(reqModel.ID, "required", "id"); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
// 		return
// 	}
// 	// Get user ID from URL parameter
// 	ID := &reqModel.ID
// 	fmt.Println(c)
// 	fmt.Println(ID)
// 	// fmt.Println(ID, err)
// 	// if err != nil {
// 	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
// 	// 	return
// 	// }

// 	// Find user in database
// 	var user models.Users
// 	result := database.DB.Where("id = ?", ID).First(&user)
// 	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	} else if result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
// 		return
// 	}

// 	// Find user's posts in database
// 	var posts []models.Posts
// 	result = database.DB.Where("id = ?", ID).Find(&posts)
// 	if result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user's posts"})
// 		return
// 	}

// 	// Return found posts
// 	c.JSON(http.StatusOK, gin.H{"posts": posts})
// }
