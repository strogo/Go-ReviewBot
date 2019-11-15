package handler

import (
	"FirstGoProject/champ"
	"FirstGoProject/origin"
)

// Handler
type Handler struct {
	champStore  champ.Store
	originStore origin.Store
}

// NewHandler
func NewHandler(chStore champ.Store, orStore origin.Store) *Handler {
	return &Handler{
		champStore:  chStore,
		originStore: orStore,
	}
}

// type Handler struct {
// 	userStore    user.Store
// 	articleStore article.Store
// }

// func NewHandler(us user.Store, as article.Store) *Handler {
// 	return &Handler{
// 		userStore:    us,
// 		articleStore: as,
// 	}
// }
