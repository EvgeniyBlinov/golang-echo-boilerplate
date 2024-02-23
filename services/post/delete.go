package post

import "echo-demo-project/models"

func (postService *Service) Delete(post *models.Post) error {
	return postService.DB.Where("uuid = ?", post.UUID).Delete(&post).Error
}
