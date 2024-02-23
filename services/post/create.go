package post

import "echo-demo-project/models"

func (postService *Service) Create(post *models.Post) error {
	return postService.DB.Create(post).Error
}
