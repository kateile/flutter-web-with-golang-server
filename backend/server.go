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
		port = "8282"
	}

	router := chi.NewRouter()
	//router.Use(func(next http.Handler) http.Handler {
	//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		w.Header().Set("Content-Type", "application/javascript")
	//		next.ServeHTTP(w, r)
	//	})
	//})
	//Status router for testing if server is working
	router.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("It is working!!!!!!!!!!!!!!!!!"))
	})

	// SOLUTION 1
	//fs := http.FileServer(http.Dir("web"))
	//router.Handle("/*", http.StripPrefix("/", fs))

	// SOLUTION 2
	//AnotherFileServer(router)

	//SOLUTION 3
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "web"))
	FileServer(router, "/", filesDir)

	log.Printf("connect to http://localhost:%s for viewing flutter web", port)
	log.Fatal(http.ListenAndServe("127.0.0.1:"+port, router))
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
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

// AnotherFileServer FileServer is serving static files.
//func AnotherFileServer(router *chi.Mux) {
//	root := "./web"
//	fs := http.FileServer(http.Dir(root))
//
//	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
//		if _, err := os.Stat(root + r.RequestURI); os.IsNotExist(err) {
//			http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
//		} else {
//			fs.ServeHTTP(w, r)
//		}
//	})
//}
