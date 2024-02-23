package repositories

import (
	"echo-demo-project/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostRepositoryQ interface {
	GetPosts(posts *[]models.Post)
	GetPost(post *models.Post, uuid uuid.UUID)
}

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (postRepository *PostRepository) GetPosts(posts *[]models.Post) error {
	return postRepository.DB.Preload("User").Find(&posts).Error
}

func (postRepository *PostRepository) GetPost(post *models.Post, uuid uuid.UUID) error {
	return postRepository.DB.Where("uuid = ? ", uuid).Find(&post).Error
}
