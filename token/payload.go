package token

import (
	"errors"
	"time"

	uuid "github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrTokenExpired = errors.New("token has expired")
)

// Payload is the data that we will be encoding to the JWT token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// Valid method that implements jwt.Claims interface
func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrTokenExpired
	}
	return nil
}

// NewPayload creates a new payload for the JWT token
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Payload{
		ID:        tokenID,
		Username:  username,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}, nil
}
