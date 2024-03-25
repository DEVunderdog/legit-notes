package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto *paseto.V2
	symmetricKey []byte
} // Struct of type of paseto.V2 and symmetric key required for paseto

func NewPasetoMaker(symmetricKey string)(Maker, error) {
	// Takes argument of symmetrickey which is of type string
	// Return type belongs to Maker interface it should return things which should
	// implement the Maker interface
	if len(symmetricKey) != chacha20poly1305.KeySize{
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}// Checking the len of symmetric such that it matches the condition of chacha20poly1305 condition

	maker := &PasetoMaker{
		paseto: paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	} // Creating struct named maker implementing PasetoMaker struct.

	return maker, nil // It returns maker
	// One thing to notice if we didn't implement Maker interface on PasetoMaker struct it will
	// mention the error because then we are just returning PasetoMaker struct reference
}

func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, *Payload,error) {
	// The capability of using the pointer is that the instance on which this method is called
	// can manipulate the state.
	payload, err := NewPayload(username, duration) // Creating the payload based on username and duration provided

	if err != nil {
		return "", payload,err
	}

	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	// The method on the instance it is called has paseto's instance hence calling
	// the Encrypt on it and encrypting the and returning token with payload provided

	return token, payload, err
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	// Empty instance of Payload struct
	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	// Decrypting the the token we got with payload provided
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid() // Calling the Valid() function to validate the token expiration

	if err != nil {
		return nil, err
	}

	return payload, nil // Returning the payload.
}