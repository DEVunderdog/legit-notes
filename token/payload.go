package token



import (
	"errors"
	"time"

	"github.com/google/uuid"

)

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("invalid token")
) // Mentioning the required errors

type Payload struct {
	ID uuid.UUID `json:"id"`
	Username string `json:"username"`
	IssuedAt time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
} // Describing the Payload for the token.

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	// Taking the username and duration as the argument
	// Returns Payload struct and error
	tokenID, err := uuid.NewRandom() // Generating the tokenID

	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID: tokenID,
		Username: username,
		IssuedAt: time.Now(),
		ExpiredAt: time.Now().Add(duration),
	} // Creating the Payload struct

	return payload, nil // Returns the created the Payload struct
}

func (payload *Payload) Valid() error {
	// Validating the payload in terms of Expiration
	if time.Now().After(payload.ExpiredAt){
		return ErrExpiredToken
	}

	return nil
}