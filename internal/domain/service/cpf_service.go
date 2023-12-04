// internal/domain/service/cpf_service.go

package service

import "cpf-lookup/internal/domain/model"

type CPFService struct{}

func NewCPFService() *CPFService {
	return &CPFService{}
}

func (s *CPFService) QueryCPF(cpf string) *model.CPF {
	return &model.CPF{Value: cpf}
}
