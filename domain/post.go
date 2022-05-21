package domain

import "time"

type Post struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PostRepository interface {
	FindAll() ([]Post, error)
	FindByID(id int) (Post, error)
	Create(title, content string) (Post, error)
	Update(id int, title, content string) (Post, error)
	Delete(id int) error
}

type PostUseCase interface {
	FindAll() ([]Post, error)
	FindByID(id int) (Post, error)
	Create(title, content string) (Post, error)
	Update(id int, title, content string) (Post, error)
	Delete(id int) error
}
