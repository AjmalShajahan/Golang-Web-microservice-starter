package service

import (
	"web-api/src/models"
	"web-api/src/repository"
)

type PostService struct{}

// create post by user
func (c *PostService) CreatePost(insert *models.Post) (string, error) {

	if err := repository.Repo.Insert(insert); err != nil {
		return "Unable to insert post", err
	}

	return "Post inserted sucessfully", nil
}