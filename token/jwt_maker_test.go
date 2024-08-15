package token

import (
    "testing"
    "time"
    "github.com/stretchr/testify/require"
    "github.com/techschool/simplebank/util"
    "github.com/dgrijalva/jwt-go"
)

func TestJWTMaker(t *testing.T) {
    maker, err := NewJWTMaker(util.RandomString(32))
    require.NoError(t, err)

    username := util.RandomOwner()
    duration := time.Minute

    issuedAt := time.Now()
    expiresAt := issuedAt.Add(duration)

    token, payload, err := maker.CreateToken(username, duration)
    require.NoError(t, err)
    require.NotEmpty(t, token)
    require.NotEmpty(t, payload)

    payload, err = maker.VerifyToken(token)
    require.NoError(t, err)
    require.NotEmpty(t, payload)

    require.NotZero(t, payload.ID)
    require.Equal(t, username, payload.Username)
    require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
    require.WithinDuration(t, expiresAt, payload.ExpiresAt, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
    maker, err := NewJWTMaker(util.RandomString(32))
    require.NoError(t, err)

    token, payload, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
    require.NoError(t, err)
    require.NotEmpty(t, token)
    require.NotEmpty(t, payload)

    payload, err = maker.VerifyToken(token)
    require.Error(t, err)
    require.EqualError(t, err, ErrExpiredToken.Error())
    require.Nil(t, payload)

}


func TestInvalidJWTTokenAlgNone(t *testing.T) {
    payload, err := NewPayload(util.RandomOwner(), time.Minute)
    require.NoError(t, err)

    jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
    token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
    require.NoError(t, err)

    maker, err := NewJWTMaker(util.RandomString(32))
    require.NoError(t, err)

    payload, err = maker.VerifyToken(token)
    require.Error(t, err)
    require.EqualError(t, err, ErrInvalidToken.Error())
    require.Nil(t, payload)
}