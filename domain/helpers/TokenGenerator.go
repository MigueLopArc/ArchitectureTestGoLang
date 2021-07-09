package helpers

import (
	"errors"
	"fmt"
	"time"

	"github.com/MigueLopArc/ArchitectureTestGoLang/config"
	"github.com/MigueLopArc/ArchitectureTestGoLang/domain/models/auth"
	jwt "github.com/dgrijalva/jwt-go"
)

const tokenTTLSeconds = 3600

type TokenGenerator struct {
	JwtConfig config.JwtConfig
}

func NewTokenGenerator() *TokenGenerator {
	config := config.GetEnv()
	return &TokenGenerator{
		JwtConfig: config.Jwt,
	}
}

func (tokenGenerator *TokenGenerator) GenerateJwtToken(userIdentity *auth.JwtUserIdentity) (*auth.JsonWebToken, error) {
	claims := auth.Claims{
		StandardClaims: jwt.StandardClaims{
			Audience:  tokenGenerator.JwtConfig.Audience,
			Issuer:    tokenGenerator.JwtConfig.Issuer,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Second * tokenTTLSeconds).Unix(),
			Subject:   userIdentity.UserId,
		},
		Email:    userIdentity.Email,
		UserName: userIdentity.UserName,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(tokenGenerator.JwtConfig.SigningKey))

	if err != nil {
		return nil, errors.New("There was an error generating the JWT")
	}

	return &auth.JsonWebToken{
		Token:     tokenString,
		Email:     userIdentity.Email,
		UserId:    userIdentity.UserId,
		ExpiresAt: claims.ExpiresAt,
	}, nil
}

func (tokenGenerator *TokenGenerator) GetClaimsFromToken(tokenString string) (*auth.Claims, error) {
	token, err := jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return []byte(tokenGenerator.JwtConfig.SigningKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Invalid claim")
	}

	userId, ok := claims["sub"]
	if !ok {
		return nil, errors.New("User identity not found")
	}

	email, ok := claims["email"]
	if !ok {
		return nil, errors.New("User email not found")
	}

	return &auth.Claims{
		Email:    fmt.Sprintf("%v", email),
		UserName: claims["userName"].(string),
		StandardClaims: jwt.StandardClaims{
			Subject:   userId.(string),
			ExpiresAt: int64(claims["exp"].(float64)),
			IssuedAt:  int64(claims["iat"].(float64)),
			Audience:  claims["aud"].(string),
			Issuer:    claims["iss"].(string),
		},
	}, nil
}
