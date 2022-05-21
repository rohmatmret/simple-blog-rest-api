package usecase

import (
	"log"
	"testing"

	domain "github.com/simple-blog/domain"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("../../post_db.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

type postMock struct {
	mock.Mock
}

func (p *postMock) FindAll() ([]domain.Post, error) {
	args := p.Called()
	return args.Get(0).([]domain.Post), args.Error(1)
}

func (p *postMock) FindByID(id int) (domain.Post, error) {
	args := p.Called()
	return args.Get(0).(domain.Post), args.Error(1)
}

func Test_postUseCase_FindAll(t *testing.T) {
	postUsecase := new(postMock)
	postUsecase.On("FindAll").Return([]domain.Post{}, nil)

}

func Test_postUseCase_FindByID(t *testing.T) {
	postUsecase := new(postMock)
	postUsecase.On("FindByID", 1).Return(domain.Post{}, nil)
}
