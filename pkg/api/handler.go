// pkg/api/handler.go

package api

import (
	"cpf-lookup/domain/service"
	"encoding/json"
	"net/http"
)

type CPFHandler struct {
	CPFService *service.CPFService
}

func NewCPFHandler(cpfService *service.CPFService) *CPFHandler {
	return &CPFHandler{
		CPFService: cpfService,
	}
}

func (h *CPFHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.handleCPFQuery(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *CPFHandler) handleCPFQuery(w http.ResponseWriter, r *http.Request) {
	var request CPFRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Error reading the request body", http.StatusBadRequest)
		return
	}

	response := h.CPFService.QueryCPF(request.CPF)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
