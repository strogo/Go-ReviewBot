package handler

import (
	"golang-starter-pack/article"
	"golang-starter-pack/user"
)

// type Handler struct {
// 	champStore  champ.Store
// 	originStore origin.Store
// }

// func NewHandler(chStore champ.Store, orStore origin.Store) *Handler {
// 	return &Handler{
// 		champStore:  chStore,
// 		originStore: orStore,
// 	}
// }

type Handler struct {
	userStore    user.Store
	articleStore article.Store
}

func NewHandler(us user.Store, as article.Store) *Handler {
	return &Handler{
		userStore:    us,
		articleStore: as,
	}
}
