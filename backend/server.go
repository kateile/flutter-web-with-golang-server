package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "9090"
	}

	router := chi.NewRouter()

	//Status router for testing if server is working
	router.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("It is working!!!!!!!!!!!!!!!!!"))
	})

	//Configuring frontend
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "web")) //using "test" here works
	FileServer(router, "/", filesDir)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe("127.0.0.1:"+port, router))
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}

	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
