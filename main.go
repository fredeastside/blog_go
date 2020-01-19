package main

import (
	"blog/pkg/db"
	"blog/pkg/post"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := db.NewDB(fmt.Sprintf(
		"%s:%s@/%s?parseTime=true",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
	),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	posts := post.NewPostsRepository(db)
	postsListHandler := post.NewListHandler(posts)

	router := httprouter.New()
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", r.Header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})
	router.GET("/api/post", postsListHandler.GetAll)
	router.GET("/api/post/:slug", postsListHandler.GetOne)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("APP_PORT"), router))
}
