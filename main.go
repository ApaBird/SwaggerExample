package main

import (
	"flag"
	"fmt"
	"net/http"
	"swaggerExample/service"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	_ "swaggerExample/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "PUT", "DELETE"}, // Allowing only get, just an example
	})

	r.HandleFunc("/Ping", service.Ping).Methods("GET")
	r.HandleFunc("/Even", service.Even).Methods("GET")
	r.HandleFunc("/Sum", service.SumSlice).Methods("POST")

	docs := flag.Bool("d", false, "Включить документацию")
	flag.Parse()

	fmt.Println(*docs)

	if *docs {
		SwaggerRouting(r)
		r.HandleFunc("/doc", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/docs/index.html", http.StatusSeeOther)
		})
	}

	fmt.Println("Сервер запущен")
	http.ListenAndServe(":9122", c.Handler(r))
}

func SwaggerRouting(router *mux.Router) {
	fmt.Println("Документация включена")
	prefix := "/docs"
	router.PathPrefix(prefix).Handler(httpSwagger.Handler(
		httpSwagger.URL("doc.json"),
	)).Methods(http.MethodGet)
}
