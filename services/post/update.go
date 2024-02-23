package post

import (
	"echo-demo-project/models"
	"echo-demo-project/requests"
)

func (postService *Service) Update(post *models.Post, updatePostRequest *requests.UpdatePostRequest) error {
	post.Content = updatePostRequest.Content
	post.Title = updatePostRequest.Title
	return postService.DB.Save(&post).Error
}
