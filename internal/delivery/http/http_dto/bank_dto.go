package http_dto

import "demo/bank-linking-listener/internal/service/entity"

type BankCreateRequest struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (r *BankCreateRequest) ToEntity() *entity.Bank {
	return &entity.Bank{
		Code: r.Code,
		Name: r.Name,
	}
}
