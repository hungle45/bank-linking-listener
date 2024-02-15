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

func (s *bankService) GetBankListByUserID(ctx context.Context, userID uint) ([]entity.Bank, entity.Error) {
	return s.bankRepo.FetchByUserID(ctx, userID)
}

func (s *bankService) LinkBank(ctx context.Context, userID uint, bankCode string) entity.Error {
	return s.bankRepo.LinkBank(ctx, userID, bankCode)
}

func (s *bankService) CreateBank(ctx context.Context, bank entity.Bank) (entity.Bank, entity.Error) {
	bank, rerr := s.bankRepo.Create(ctx, bank)
	if rerr != nil {
		return entity.Bank{}, rerr
	}
	return bank, nil
}
