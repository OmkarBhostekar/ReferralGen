package token

import (
	"time"
)

type Maker interface {
	CreateToken(string, int64, time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
