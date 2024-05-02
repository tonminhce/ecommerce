package models

import (
	"fmt"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	UserID   string `json:"userId,omitempty"`
	Username string `json:"username,omitempty"`
	Roles    []int  `json:"roles,omitempty"`
	jwt.StandardClaims
}

func (claims JwtClaims) Valid() error {
	now := time.Now().UTC().Unix()
	flags, _ := GetFlags()
	url, _ := flags.GetApplicationUrl()
	if claims.VerifyExpiresAt(now, true) && claims.VerifyIssuer(*url, true) {
		return nil
	}
	return fmt.Errorf("Token is invalid")
}

func (claims JwtClaims) VerifyAudience(origin string) bool {
	return strings.Compare(claims.Audience, origin) == 0
}
