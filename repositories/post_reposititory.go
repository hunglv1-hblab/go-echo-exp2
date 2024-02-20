package repositories

import (
	"fund-aplly-back/models"
	"fund-aplly-back/requests"

	"gorm.io/gorm"
)

type PostRepositoryQ interface {
	GetPosts(posts *[]models.Post)
	GetPost(post *models.Post, id int)
	Create(post *models.Post)
	Delete(post *models.Post)
	Update(post *models.Post, updatePostRequest *requests.UpdatePostRequest)
}

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (postRepository *PostRepository) GetPosts(posts *[]models.Post) {
	postRepository.DB.Preload("User").Find(posts)
}

func (postRepository *PostRepository) GetPost(post *models.Post, id int) {
	postRepository.DB.Where("id = ? ", id).Find(post)
}

func (postService *PostRepository) Create(post *models.Post) {
	postService.DB.Create(post)
}
func (postService *PostRepository) Delete(post *models.Post) {
	postService.DB.Delete(post)
}

func (postService *PostRepository) Update(post *models.Post, updatePostRequest *requests.UpdatePostRequest) {
	post.Content = updatePostRequest.Content
	post.Title = updatePostRequest.Title
	postService.DB.Save(post)
}
