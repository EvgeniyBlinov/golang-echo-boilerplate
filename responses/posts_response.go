package responses

import (
	"echo-demo-project/models"

	"github.com/google/uuid"
)

type PostResponse struct {
	Title    string    `json:"title" example:"Echo"`
	Content  string    `json:"content" example:"Echo is nice!"`
	Username string    `json:"username" example:"John Doe"`
	UUID     uuid.UUID `json:"uuid" example:"018de611-20ed-71c5-9b18-863a7851eef2"`
}

func NewPostResponse(posts []models.Post) *[]PostResponse {
	postResponse := make([]PostResponse, 0)

	for i := range posts {
		postResponse = append(postResponse, PostResponse{
			Title:    posts[i].Title,
			Content:  posts[i].Content,
			Username: posts[i].User.Name,
			UUID:     posts[i].UUID,
		})
	}

	return &postResponse
}
