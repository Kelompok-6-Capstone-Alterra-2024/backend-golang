package configs

import "os"

type Midtrans struct {
	Key    string
	IsProd bool
}

func (m *Midtrans) GetKey() *Midtrans {
	return &Midtrans{
		Key:    os.Getenv("MIDTRANS_KEY"),
		IsProd: os.Getenv("MIDTRANS_IS_PROD") == "production",
	}
}
