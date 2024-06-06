package configs

import (
	"encoding/base64"
	"os"
)

type Midtrans struct {
	ClientKey string
	ServerKey string
	BaseURL   string
	IsProd    bool
}

func MidtransConfig() *Midtrans {
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	serverKeyBase64 := base64.StdEncoding.EncodeToString([]byte(serverKey))
	return &Midtrans{
		ClientKey: os.Getenv("MIDTRANS_KEY"),
		ServerKey: serverKeyBase64,
		BaseURL:   os.Getenv("MIDTRANS_BASE_URL"),
		IsProd:    os.Getenv("MIDTRANS_ENV") == "production",
	}
}
