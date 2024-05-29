package configs

import "os"

type Midtrans struct {
	Key    string
	IsProd bool
}

func MidtransConfig() *Midtrans {
	return &Midtrans{
		Key:    os.Getenv("MIDTRANS_KEY"),
		IsProd: os.Getenv("MIDTRANS_ENV") == "production",
	}
}
