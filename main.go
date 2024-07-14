package main

import (
	"log"
	"main/middleware"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/item", func(w http.ResponseWriter, r *http.Request) {
		log.Println("here")
	})

	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.IsAuthenticated,
	)

	server := http.Server {
		Addr: ":8000",
		Handler: stack(router),
	}

	log.Println("Starting server on port :8080")
	server.ListenAndServe()
}
