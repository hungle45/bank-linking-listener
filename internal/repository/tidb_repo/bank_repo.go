package tidb_repo

import "demo/bank-linking-listener/internal/repository"

type bankRepository struct {}

func NewBankRepository() repository.BankRepository {
	return &bankRepository{}
}