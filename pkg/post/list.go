package post

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ListHandler struct {
	posts Posts
}

func NewListHandler(posts Posts) *ListHandler {
	return &ListHandler{posts}
}

func (h *ListHandler) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	posts, err := h.posts.GetAll()
	if err != nil {
		log.Println(err)
	}

	postsByYear := make(map[string][]*Post)
	for _, post := range posts {
		year := post.GetYear()
		if _, ok := postsByYear[year]; !ok {
			postsByYear[year] = []*Post{}
		}
		postsByYear[year] = append(postsByYear[year], post)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *ListHandler) GetOne(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	post, err := h.posts.GetBySlug(ps.ByName("slug"))
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(post); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
