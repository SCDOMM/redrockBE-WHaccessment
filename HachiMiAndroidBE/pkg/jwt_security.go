package pkg

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var StandardHeader = []byte(`{"alg":"HS256","typ":"JWT"}`)
var (
	accessSecret  = []byte("八百标兵奔北坡")
	refreshSecret = []byte("北坡炮兵并排跑")
	issuer        = "曾姐姐批发"
	accessTTL     = 1 * time.Minute
	refreshTTL    = 7 * 24 * time.Hour
)

type Claims struct {
	Account string `json:"user_id"`
	Role    uint   `json:"role"`
	Type    string `json:"type"`
	jwt.RegisteredClaims
}

func CreateAccessToken(account string, role uint) (accessToken string, err error) {
	now := time.Now()
	accessClaims := Claims{
		Account: account,
		Role:    role,
		Type:    "access",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			Subject:   fmt.Sprintf("%s", account),
			Audience:  []string{"user"},
			ExpiresAt: jwt.NewNumericDate(now.Add(accessTTL)),
			NotBefore: jwt.NewNumericDate(now.Add(-5 * time.Second)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	accessTok := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err = accessTok.SignedString(accessSecret)
	if err != nil {
		log.Println("fail to create access token!" + err.Error())
		return "", fmt.Errorf("sign access token: %w", err)
	}
	return accessToken, nil
}
func CreateRefreshToken(account string, role uint) (refreshToken string, err error) {
	now := time.Now()

	refreshClaims := Claims{
		Account: account,
		Role:    role,
		Type:    "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			Subject:   fmt.Sprintf("%s", account),
			Audience:  []string{"user"},
			ExpiresAt: jwt.NewNumericDate(now.Add(refreshTTL)),
			NotBefore: jwt.NewNumericDate(now.Add(-5 * time.Second)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	refreshTok := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err = refreshTok.SignedString(refreshSecret)
	if err != nil {
		log.Println("fail to create refresh token!" + err.Error())
		return "", fmt.Errorf("sign refresh token: %w", err)
	}
	return refreshToken, nil
}

func VerifyAccessToken(accessT string) (*Claims, error) {
	raw := strip(accessT)
	token, err := jwt.ParseWithClaims(raw, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return accessSecret, nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid access token")
	}
	if claims.Type != "access" {
		return nil, errors.New("not an access token")
	}
	return claims, nil
}

func VerifyRefreshToken(RefreshT string) (*Claims, error) {

	raw := strip(RefreshT)

	token, err := jwt.ParseWithClaims(raw, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return refreshSecret, nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}
	if claims.Type != "refresh" {
		return nil, errors.New("not an refresh token")
	}
	return claims, nil
}

func strip(s string) string {
	if strings.HasSuffix(strings.ToLower(strings.TrimSpace(s)), "bearer") {
		return strings.TrimSpace(s[len("bearer"):])
	}
	return strings.TrimSpace(s)
}
