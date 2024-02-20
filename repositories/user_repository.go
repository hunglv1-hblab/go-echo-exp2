package repositories

import (
	"fund-aplly-back/models"
	"fund-aplly-back/requests"
	"fund-aplly-back/server/builders"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepositoryQ interface {
	GetUserByEmail(user *models.User, email string)
	Create(post *models.Post)
	Delete(post *models.Post)
	Update(post *models.Post, updatePostRequest *requests.UpdatePostRequest)
	Register(request *requests.RegisterRequest) error
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (userRepository *UserRepository) GetUserByEmail(user *models.User, email string) {
	userRepository.DB.Where("email = ?", email).Find(user)
}

func (postService *UserRepository) Create(post *models.Post) {
	postService.DB.Create(post)
}

func (postService *UserRepository) Delete(post *models.Post) {
	postService.DB.Delete(post)
}

func (postService *UserRepository) Update(post *models.Post, updatePostRequest *requests.UpdatePostRequest) {
	post.Content = updatePostRequest.Content
	post.Title = updatePostRequest.Title
	postService.DB.Save(post)
}

func (userService *UserRepository) Register(request *requests.RegisterRequest) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user := builders.NewUserBuilder().SetEmail(request.Email).
		SetName(request.Name).
		SetPassword(string(encryptedPassword)).
		Build()

	userService.DB.Create(&user)

	return nil
}
