package handler

import (
	"fmt"
	"net/http"

	"github.com/zhetkerbaevan/study-mongodb/internal/models"
	"github.com/zhetkerbaevan/study-mongodb/internal/utils"
)

type Handler struct {
	todoService models.TodoServiceInterface
}

func NewHandler(todoService models.TodoServiceInterface) *Handler {
	return &Handler{todoService: todoService}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/todo", h.handleCreateTodo)
	router.HandleFunc("/", h.handleGetTodos)
	router.HandleFunc("/delete/todo", h.handleDeletoTodo)
	router.HandleFunc("/todo/{_id}", h.handleUpdateTodo)
}

func (h *Handler) handleCreateTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { //Check that method is right
		http.Error(w, "INVALID REQUEST METHOD", http.StatusMethodNotAllowed)
		return
	}

	var payload models.TodoPayload
	if err := utils.ParseJSON(r, &payload); err != nil { //Get data from request body
		http.Error(w, "INVALID PAYLOAD", http.StatusBadRequest)
		return
	}

	err := h.todoService.InsertTodo(payload)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) handleGetTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet { //Check that method is right
		http.Error(w, "INVALID REQUEST METHOD", http.StatusMethodNotAllowed)
		return
	}

	todos, err := h.todoService.GetTodos()
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, todos)
}

func (h *Handler) handleDeletoTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete { //Check that method is right
		http.Error(w, "INVALID REQUEST METHOD", http.StatusMethodNotAllowed)
		return
	}

	var payload models.TodoPayload
	if err := utils.ParseJSON(r, &payload); err != nil { //Get data from request body
		http.Error(w, "INVALID PAYLOAD", http.StatusBadRequest)
		return
	}

	err := h.todoService.DeleteTodo(payload.Task)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleUpdateTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut { //Check that method is right
		http.Error(w, "INVALID REQUEST METHOD", http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("_id") //Get _id from URL parametres
	fmt.Println("Id from url", id)

	var payload models.TodoPayload
	if err := utils.ParseJSON(r, &payload); err != nil { //Get data from request body
		http.Error(w, "INVALID PAYLOAD", http.StatusBadRequest)
		return
	}

	err := h.todoService.UpdateTodo(id, payload)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)

}
