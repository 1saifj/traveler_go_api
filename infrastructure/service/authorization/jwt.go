package authorization

import (
	"crypto/rsa"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog"
	"os"
	"taveler/infrastructure/model"
	"time"
)

var privatepem = "/certs/private.pem"
var publicpem = "/certs/public.pem"

type tokenService struct {
	logger                 zerolog.Logger
	accessTokenPrivateKey  *rsa.PrivateKey
	accessTokenPublicKey   *rsa.PublicKey
	refreshTokenPrivateKey *rsa.PrivateKey
	refreshTokenPublicKey  *rsa.PublicKey
}

type TokenService interface {
	GenerateAccessToken(*fiber.Ctx, *model.User) (string, error)
	ValidateAccessTokenFromRequest(*fiber.Ctx) (*Claims, error)
	GenerateRefreshToken(*fiber.Ctx, *model.User) (string, error)
	VerifyRefreshToken(*fiber.Ctx, string) (string, error)
	ParseToken(*fiber.Ctx, string) (*Claims, error)
}

func NewTokenService(logger zerolog.Logger) TokenService {
	return &tokenService{
		logger: logger,
	}
}

func (s *tokenService) init() *tokenService {
	b, err := os.ReadFile(privatepem)
	if err != nil {
		s.logger.Fatal().Err(err).Msg("could not read private.pem")
	}
	atprk, err := jwt.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		panic(err)
	}

	s.accessTokenPrivateKey = atprk

	// get public key from file
	b, err = os.ReadFile(publicpem)
	if err != nil {
		panic(err)
	}

	atpuk, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		panic(err)
	}

	s.accessTokenPublicKey = atpuk

	// get private key from file
	b, err = os.ReadFile(privatepem)
	if err != nil {
		panic(err)
	}

	rtprk, err := jwt.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		panic(err)
	}

	s.refreshTokenPrivateKey = rtprk

	// get public key from file
	b, err = os.ReadFile(publicpem)
	if err != nil {
		panic(err)
	}

	rtpuk, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		panic(err)
	}

	s.refreshTokenPublicKey = rtpuk

	return s
}

func (s *tokenService) GenerateAccessToken(ctx *fiber.Ctx, user *model.User) (string, error) {
	sub := user.ID
	if sub == "" {
		return "", errors.New("user id is empty")
	}
	now := time.Now().UTC()
	claims := Claims{
		Role: user.Role,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "taveler",
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(time.Minute * 15).Unix(),
			Subject:   sub,
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(s.accessTokenPrivateKey)
}
func (s *tokenService) ValidateAccessTokenFromRequest(ctx *fiber.Ctx) (*Claims, error) {
	token := ctx.Get("Authorization")
	if token == "" {
		return nil, errors.New("token is empty")
	}
	// remove bearer prefix
	token = token[7:]
	claims, err := s.ParseToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
func (s *tokenService) GenerateRefreshToken(ctx *fiber.Ctx, user *model.User) (string, error) {
	sub := user.ID
	if sub == "" {
		return "", errors.New("user id is empty")
	}
	now := time.Now().UTC()

	claims := jwt.MapClaims{
		"role": user.Role,
		"standard_claims": jwt.StandardClaims{
			Issuer:    "taveler",
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(time.Hour * 24 * 7).Unix(),
			Subject:   sub,
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(s.refreshTokenPrivateKey)
}
func (s *tokenService) VerifyRefreshToken(ctx *fiber.Ctx, token string) (string, error) {
	c, err := jwt.Parse(token, s.provideRefreshTokenPublicKey)
	if err != nil {
		return "", err
	}
	claims, ok := c.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("could not parse claims")
	}
	if !c.Valid {
		return "", errors.New("invalid token")
	}
	if !claims.VerifyIssuer("taveler", true) {
		return "", errors.New("invalid issuer")
	}
	if !claims.VerifyExpiresAt(time.Now().UTC().Unix(), true) {
		return "", errors.New("token expired")
	}
	sub := claims["sub"].(string)
	if sub == "" {
		return "", errors.New("subject is empty")
	}
	jti := claims["jti"].(string)
	if jti == "" {
		return "", errors.New("jti is empty")
	}
	//TODO check in redis
	return sub, nil

}
func (s *tokenService) ParseToken(ctx *fiber.Ctx, token string) (*Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, s.provideAccessTokenPublicKey)
	if err != nil {
		return nil, err
	}

	if tkn.Method.Alg() != jwt.SigningMethodRS256.Alg() {
		return nil, errors.New("invalid algorithm")
	}

	if !tkn.Valid {
		return nil, errors.New("token is invalid")
	}

	if !claims.VerifyIssuer("taveler", true) {
		return nil, errors.New("invalid issuer")
	}

	return claims, nil
}
func (s *tokenService) provideAccessTokenPublicKey(_ *jwt.Token) (interface{}, error) {
	return s.accessTokenPublicKey, nil
}

// provideRefreshTokenPublicKey provides access token public key to veriy token
func (s *tokenService) provideRefreshTokenPublicKey(_ *jwt.Token) (interface{}, error) {
	return s.refreshTokenPublicKey, nil
}
