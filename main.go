package main

import (
	"fmt"
	"net/http"

	"github.com/profadevairvitorio/cpf-lookup/internal/domain/service"
	"github.com/profadevairvitorio/cpf-lookup/pkg/api"
)

func main() {
	cpfService := service.NewCPFService()

	handler := api.NewCPFHandler(cpfService)

	port := 8080
	fmt.Printf("Server running at http://localhost:%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
}
