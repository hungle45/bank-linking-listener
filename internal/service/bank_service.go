package service

import (
	"context"
	"demo/bank-linking-listener/internal/repository"
	"demo/bank-linking-listener/internal/service/entity"
)

type bankService struct {
	bankRepo repository.BankRepository
}

func NewBankService(bankRepo repository.BankRepository) BankService {
	return &bankService{
		bankRepo: bankRepo,
	}
}

func (s *bankService) GetBankListByUserID(ctx context.Context, userID uint) ([]entity.Bank, error) {
	return s.bankRepo.FetchByUserID(ctx, userID)
}

func (s *bankService) LinkBank(ctx context.Context, userID uint, bankCode string) error {
	return s.bankRepo.LinkBank(ctx, userID, bankCode)
}

func (s *bankService) CreateBank(ctx context.Context, bank entity.Bank) (entity.Bank, error) {
	bank, err := s.bankRepo.Create(ctx, bank)
	if err != nil {
		return entity.Bank{}, err
	}
	return bank, nil
}
