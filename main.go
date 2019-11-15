package main

import (
	"FirstGoProject/db"
	"FirstGoProject/handler"
	"FirstGoProject/router"
	"FirstGoProject/store"
)

func main() {

	r := router.New()
	v1 := r.Group("/api")

	d := db.New()
	db.AutoMigrate(d)

	us := store.NewUserStore(d)
	as := store.NewArticleStore(d)
	h := handler.NewHandler(us, as)
	h.Register(v1)
	r.Logger.Fatal(r.Start("127.0.0.1:8585"))
}
