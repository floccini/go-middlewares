package filesHandler

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct{}

func (handler *Handler) Create(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusCreated)
	log.Println("Received a request to create")
	writer.Write([]byte("Created"))
}

func (handler *Handler) FindById(writer http.ResponseWriter, request *http.Request) {
	log.Println("Received a find by id request - Method:", request.Method)
	file, exists := getFiles()[request.PathValue("id")]

	if !exists {
		writer.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(writer).Encode(file)
}
