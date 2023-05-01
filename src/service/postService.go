package service

import (
	"web-api/src/models"
	"web-api/src/repository"
	"web-api/utils/database"
)

type PostService struct{}

// create post by user
func (c *PostService) CreatePost(insert *models.Posts) (string, error) {

	if err := repository.Repo.Insert(insert); err != nil {
		return "Unable to insert post", err
	}

	return "Post inserted sucessfully", nil
}

func (c *PostService) FetchPosts() ([]models.Posts, error) {
	var posts []models.Posts
	if err := database.DB.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
func (s *PostService) FindPostsByUser(userID int) ([]models.Posts, error) {
	var posts []models.Posts
	err := database.DB.Where("user_id = ?", userID).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (c *PostService) CreateComment(insert *models.Comments) (string, error) {

	if err := repository.Repo.Insert(insert); err != nil {
		return "Unable to insert comment", err
	}

	return "Comment inserted sucessfully", nil
}

func (c *PostService) FindComPostsByUser(userID int) ([]models.Comments, error) {
	var comments []models.Comments
	err := database.DB.Where("user_id = ?", userID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}
