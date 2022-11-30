package jwt

import (
	"crypto/rand"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"io"
	"time"
)

var keys map[string]key
var currentKid = ""

type UserClaims struct {
	jwt.StandardClaims
	SessionID int64
}

func createToken(c *UserClaims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	signedToken, err := t.SignedString(keys[currentKid].Key)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func parseToken(signedToken string) (*UserClaims, error) {
	claims := &UserClaims{}
	t, err := jwt.ParseWithClaims(signedToken, claims, func(unverified *jwt.Token) (interface{}, error) {
		if unverified.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, fmt.Errorf("invalid signing algorithm")
		}
		kid, ok := unverified.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("invalid key id")
		}

		k, ok := keys[kid]
		if !ok {
			return nil, fmt.Errorf("invalid key id")
		}
		return k, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error in parseToken while parsing token: %w", err)
	}

	if !t.Valid {
		return nil, fmt.Errorf("error in parseToken, token is not valid")
	}
	return t.Claims.(*UserClaims), nil
}

func generateNewKey() error {
	newKey := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, newKey)
	if err != nil {
		return err
	}
	uid := uuid.New()
	keys[uid.String()] = key{
		Key:     newKey,
		Created: time.Now(),
	}
	currentKid = uid.String()
	return err
}

type key struct {
	Key     []byte
	Created time.Time
}
