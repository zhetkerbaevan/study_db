package api

import (
	"log"
	"net/http"

	"github.com/zhetkerbaevan/study-mongodb/internal/handler"
	"github.com/zhetkerbaevan/study-mongodb/internal/service"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	//Create new instance of APIServer struct
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	//create new router
	router := http.NewServeMux()

	todoService := service.NewTodoService()
	todoHandler := handler.NewHandler(todoService)
	todoHandler.RegisterRoutes(router)
	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
