package repository

import (
	"fmt"

	domain "github.com/simple-blog/domain"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func (p *PostRepository) FindAll() ([]domain.Post, error) {
	var post []domain.Post
	err := p.db.Model(&post).Find(&post).Error
	if err != nil {
		fmt.Println("error when find all post", err)
		return nil, err
	}
	return post, nil
}

func (p *PostRepository) FindByID(id int) (domain.Post, error) {
	var post domain.Post
	err := p.db.Debug().Model(&domain.Post{}).Where("id = ?", id).First(&post).Error
	if err != nil {
		fmt.Println("error when find post by id", err)
		return domain.Post{}, err
	}
	return post, nil
}

func (p *PostRepository) Create(title, content string) (domain.Post, error) {
	post := domain.Post{
		Title:   title,
		Content: content,
	}
	err := p.db.Debug().Create(&post).Error
	if err != nil {
		fmt.Println("error when create post", err)
		return domain.Post{}, err
	}
	id := post.ID
	post, err = p.FindByID(id)
	if err != nil {
		fmt.Println("error when find post by id", err)
		return domain.Post{}, err
	}
	return post, nil
}

func (p *PostRepository) Update(id int, title, content string) (domain.Post, error) {
	err := p.db.Debug().Model(&domain.Post{}).Where("id = ?", id).Updates(&domain.Post{Title: title, Content: content}).Error
	if err != nil {
		fmt.Println("error when update post", err)
		return domain.Post{}, err
	}
	post, err := p.FindByID(id)
	if err != nil {
		fmt.Println("error when find post by id", err)
		return domain.Post{}, err
	}
	return post, nil
}

func (p *PostRepository) Delete(id int) error {
	return p.db.Debug().Model(&domain.Post{}).Where("id = ?", id).Delete(&domain.Post{}).Error
}

func NewPostRepository(db *gorm.DB) domain.PostRepository {
	return &PostRepository{
		db: db,
	}
}
