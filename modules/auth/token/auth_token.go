package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/qnify/api-server/utils/errors"
)

var (
	accessSecret  []byte
	refreshSecret []byte
)

const (
	defaultAccessExpiry  = time.Minute * 60     // 1 hour
	defaultRefreshExpiry = time.Hour * 24 * 365 // 1 year
)

type (
	customJwtClaim struct {
		TokenData
		jwt.RegisteredClaims
	}
	TokenData struct {
		UserId int `json:"userId"`
	}
	TokenConfig struct {
		AccessSecret  string `yaml:"access_secret"`
		RefreshSecret string `yaml:"refresh_secret"`
	}
)

func InitConfig(config TokenConfig) {
	// panic if tokens already set
	if len(accessSecret) != 0 && len(refreshSecret) != 0 {
		panic(errors.New("token secret already set"))
	}

	if config.AccessSecret == "" || config.RefreshSecret == "" {
		panic(errors.New("invalid token secret found"))
	}
	accessSecret = []byte(config.AccessSecret)
	refreshSecret = []byte(config.RefreshSecret)
}

func generateToken(claim TokenData, secret []byte, expiry time.Duration) (string, error) {
	jwtClaim := customJwtClaim{
		TokenData: claim,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", errors.Wrap("error occured while generating jwt token", err)
	}
	return signedToken, nil
}

func GetTokens(claim TokenData) (string, string, error) {
	accessToken, err := generateToken(claim, accessSecret, defaultAccessExpiry)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := generateToken(claim, refreshSecret, defaultRefreshExpiry)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func verifytoken(token string, secret []byte) (TokenData, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &customJwtClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method while verifying jwt")
		}
		return secret, nil
	})
	if err != nil {
		return TokenData{}, errors.Wrap("jwt token verification failed", err)
	}

	if claims, ok := parsedToken.Claims.(*customJwtClaim); ok {
		return claims.TokenData, nil
	}
	return TokenData{}, errors.New("invalid claim type")
}

func VerifyAccessToken(token string) (TokenData, error) {
	return verifytoken(token, accessSecret)
}

func VerifyRefreshToken(token string) (TokenData, error) {
	return verifytoken(token, refreshSecret)
}
