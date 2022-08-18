package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "9090"
	}

	router := chi.NewRouter()
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/javascript")
			next.ServeHTTP(w, r)
		})
	})
	//Status router for testing if server is working
	router.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("It is working!!!!!!!!!!!!!!!!!"))
	})

	//Configuring frontend
	fs := http.FileServer(http.Dir("web"))

	router.Handle("/*", http.StripPrefix("/", fs))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe("127.0.0.1:"+port, router))
}
