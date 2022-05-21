package repository

import (
	"log"
	"testing"

	domain "github.com/simple-blog/domain"
	"github.com/stretchr/testify/assert"
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
func TestPostRepository_FindAll(t *testing.T) {

	conn := Connection()

	a := NewPostRepository(conn)
	list, err := a.FindAll()

	assert.NotEmpty(t, list)
	assert.NoError(t, err)
	for _, v := range list {
		assert.NotNil(t, v.Title)
		assert.NotNil(t, v.Content)
	}
}

func TestPostRepository_FindByID(t *testing.T) {
	conn := Connection()

	a := NewPostRepository(conn)
	list, err := a.FindByID(1)

	assert.NotEmpty(t, list)
	assert.NoError(t, err)
	assert.NotNil(t, list.Title)
}

func TestPostRepository_Create(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		title   string
		content string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Post
		wantErr bool
	}{
		{
			name: "create post",
			fields: fields{
				db: Connection(),
			},
			args: args{
				title:   "test insert post",
				content: "test insert post content",
			},
			want: domain.Post{
				Title:   "test insert post",
				Content: "test insert post content",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PostRepository{
				db: tt.fields.db,
			}
			got, err := p.Create(tt.args.title, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want.Title, got.Title)
			assert.Equal(t, tt.want.Content, got.Content)
		})
	}
}

func TestPostRepository_Update(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id      int
		title   string
		content string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Post
		wantErr bool
	}{
		{
			name: "Update post",
			fields: fields{
				db: Connection(),
			},
			args: args{
				title:   "test update post",
				content: "content",
			},
			want: domain.Post{
				Title:   "test update post",
				Content: "content",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PostRepository{
				db: tt.fields.db,
			}
			got, err := p.Update(tt.args.id, tt.args.title, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want.Title, got.Title)
			assert.Equal(t, tt.want.Content, got.Content)
		})
	}
}

func TestPostRepository_Delete(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Post
		wantErr bool
	}{
		{
			name: "Delete post",
			fields: fields{
				db: Connection(),
			},
			args: args{
				id: 1,
			},
			want: domain.Post{
				ID: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PostRepository{
				db: tt.fields.db,
			}
			got, err := p.Delete(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want.ID, got.ID)
		})
	}
}
