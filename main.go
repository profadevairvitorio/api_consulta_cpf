package main

import (
	"cpf-lookup/api"
	"cpf-lookup/domain/service"
	"fmt"
	"net/http"
)

func main() {
	cpfService := service.NewCPFService()

	handler := api.NewCPFHandler(cpfService)

	port := 8080
	fmt.Printf("Server running at http://localhost:%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
}
