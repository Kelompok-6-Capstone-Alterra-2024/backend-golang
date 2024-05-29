package midtrans

import (
	"capstone/configs"
	"capstone/constants"
	midtransEntities "capstone/entities/midtrans"
	"capstone/entities/transaction"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransUseCase struct {
	midtransConfig *configs.Midtrans
	envi           midtrans.EnvironmentType
}

func NewMidtransUseCase(m *configs.Midtrans) midtransEntities.MidtransUseCase {
	envi := midtrans.Sandbox
	if m.IsProd {
		envi = midtrans.Production
	}
	return &MidtransUseCase{
		midtransConfig: m,
		envi:           envi,
	}
}

func (usecase *MidtransUseCase) GenerateSnapURL(transaction *transaction.Transaction) (*transaction.Transaction, error) {
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.ID.String(),
			GrossAmt: int64(transaction.Price),
		},
	}

	var client snap.Client
	client.New(usecase.midtransConfig.Key, usecase.envi)

	snapResponse, err := client.CreateTransaction(req)
	if err != nil {
		return nil, err
	}
	transaction.SnapURL = snapResponse.RedirectURL
	return transaction, nil
}

func (usecase *MidtransUseCase) VerifyPayment(orderID string) (string, error) {
	var client coreapi.Client
	client.New(usecase.midtransConfig.Key, usecase.envi)

	// 4. Check transaction to Midtrans with param orderId
	transactionStatusResp, e := client.CheckTransaction(orderID)
	if e != nil {
		return constants.Deny, e
	} else {
		if transactionStatusResp != nil {
			// 5. Do set transaction status based on response from check transaction status
			if transactionStatusResp.TransactionStatus == "capture" {
				if transactionStatusResp.FraudStatus == "challenge" {
					// TODO set transaction status on your database to 'challenge'
					// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
				} else if transactionStatusResp.FraudStatus == "accept" {
					//return true, nil
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				return constants.Success, nil
			} else if transactionStatusResp.TransactionStatus == "deny" {
				return constants.Deny, nil
			} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
				return constants.Cancel, nil
			} else if transactionStatusResp.TransactionStatus == "pending" {
				return constants.Pending, nil
			}
		}
	}
	return constants.Deny, e
}
