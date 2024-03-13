package token

import "time"

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)

	VerifyToken(token string) (*Payload, error)
}

// Interface describes that it should have two function
// CreateToken for creating token which takes username, duration as argument and return string, error
// VerifyToken takes token string and returns Pointer to struct Payload and error