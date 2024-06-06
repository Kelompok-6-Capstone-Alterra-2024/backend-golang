package payment

import (
	"capstone/configs"
	midtransEntities "capstone/entities/midtrans"
	"capstone/entities/payment"
	transactionEntities "capstone/entities/transaction"
)

type PaymentUseCase struct {
	midtransConf    *configs.Midtrans
	midtransUseCase midtransEntities.MidtransUseCase
}

func NewPaymentUseCase(midtransConf *configs.Midtrans, midtransUseCase midtransEntities.MidtransUseCase) payment.Method {
	return &PaymentUseCase{
		midtransConf:    midtransConf,
		midtransUseCase: midtransUseCase,
	}
}

func (usecase *PaymentUseCase) EWallet(transaction *transactionEntities.Transaction) (*transactionEntities.Transaction, error) {
	// TODO implement me
	panic("implement me")
}

func (usecase *PaymentUseCase) BankTransfer(transaction *transactionEntities.Transaction) (*transactionEntities.Transaction, error) {
	transaction, err := usecase.midtransUseCase.BankTransfer(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
