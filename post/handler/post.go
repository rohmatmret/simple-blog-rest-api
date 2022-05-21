package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	domain "github.com/simple-blog/domain"
)

type postHandler struct {
	PostUsecase domain.PostUseCase
}

func NewPostHandler(r chi.Router, p domain.PostUseCase) *postHandler {
	return &postHandler{
		PostUsecase: p,
	}
}

func (h *postHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	post, err := h.PostUsecase.FindAll()
	if err != nil {
		domain.SetErrResponse(w, domain.PostErrResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error when Find post",
			Error:      err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	jb, err_marshal := json.Marshal(post)
	if err_marshal != nil {
		fmt.Println("error when marshal post", err_marshal)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err_marshal.Error()))
		return
	}
	domain.SetResponse(w, jb, http.StatusOK)
}

func (h *postHandler) Create(w http.ResponseWriter, r *http.Request) {
	var post domain.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		domain.SetErrResponse(w, domain.PostErrResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error when create post",
			Error:      err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	result, err := h.PostUsecase.Create(post.Title, post.Content)
	if err != nil {
		domain.SetErrResponse(w, domain.PostErrResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error when create post",
			Error:      err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(result)
	if err != nil {
		domain.SetErrResponse(w, domain.PostErrResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error Marshal post",
			Error:      err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	domain.SetResponse(w, data, http.StatusCreated)
}

func (h *postHandler) FindById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	postId, _ := strconv.Atoi(id)
	post, err := h.PostUsecase.FindByID(postId)
	if err != nil {
		domain.SetErrResponse(w, domain.PostErrResponse{
			StatusCode: http.StatusNotFound,
			Message:    "Item Not Found",
			Error:      err.Error(),
		}, http.StatusNotFound)
		return
	}
	jb, err_marshal := json.Marshal(post)
	if err_marshal != nil {
		domain.SetErrResponse(w, domain.PostErrResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error when marshal post",
			Error:      err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	domain.SetResponse(w, jb, http.StatusOK)
}

func (h *postHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	postId, _ := strconv.Atoi(id)
	var post domain.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		domain.SetErrResponse(w, domain.PostErrResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error Get Body post",
			Error:      err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	result, err := h.PostUsecase.Update(postId, post.Title, post.Content)
	if err != nil {
		domain.SetErrResponse(w, domain.PostErrResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error when update post",
			Error:      err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(result)
	if err != nil {
		domain.SetErrResponse(w, domain.PostErrResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error when marshal post",
			Error:      err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	domain.SetResponse(w, data, http.StatusOK)
}

func (h *postHandler) Deleted(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	postId, _ := strconv.Atoi(id)
	post, err := h.PostUsecase.Delete(postId)
	if err != nil {
		domain.SetErrResponse(w, domain.PostErrResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error when update post",
			Error:      err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	data, _ := json.Marshal(post)
	domain.SetResponse(w, data, http.StatusOK)
}
