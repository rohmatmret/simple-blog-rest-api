package usecase

import (
	domain "github.com/simple-blog/domain"
	repo "github.com/simple-blog/post/repository"
)

type postUseCase struct {
	postRepo repo.PostRepository
}

func (p *postUseCase) FindAll() ([]domain.Post, error) {
	return p.postRepo.FindAll()
}

func (p *postUseCase) FindByID(id int) (domain.Post, error) {
	return p.postRepo.FindByID(id)
}

func (p *postUseCase) Create(title, content string) (domain.Post, error) {
	return p.postRepo.Create(title, content)
}

func (p *postUseCase) Update(id int, title, content string) (domain.Post, error) {
	return p.postRepo.Update(id, title, content)
}

func (p *postUseCase) Delete(id int) error {
	return p.postRepo.Delete(id)
}

func NewPostUseCase(pr repo.PostRepository) domain.PostUseCase {
	return &postUseCase{}
}
