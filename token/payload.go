package token

import (
    "time"
    "github.com/google/uuid"
    "errors"
)

var (
    ErrExpiredToken = errors.New("token has expired")
    ErrInvalidToken = errors.New("token is invalid")
)


type Payload struct {
    ID uuid.UUID `json:"id"`
    Username string `json:"username"`
    IssuedAt time.Time `json:"issued_at"`
    ExpiresAt time.Time `json:"expired_at"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
    tokenID, err := uuid.NewRandom()
    if err != nil {
        return nil, err
    }

    payload := &Payload{
        ID: tokenID,
        Username: username,
        IssuedAt: time.Now(),
        ExpiresAt: time.Now().Add(duration),
    }

    return payload, nil
}

func (payload *Payload) Valid() error {
    if time.Now().After(payload.ExpiresAt) {
        return ErrExpiredToken
    }
    return nil
}