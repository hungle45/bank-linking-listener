package consumer_dto

type BankLinkRequest struct {
	UserID   uint   `json:"user_id"`
	BankCode string `json:"bank_code"`
}
