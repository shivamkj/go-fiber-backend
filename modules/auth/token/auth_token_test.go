package token

import (
	"testing"
	"time"

	. "github.com/qnify/api-server/utils/helper"
)

var (
	testConfig = TokenConfig{AccessSecret: "secret1", RefreshSecret: "secret2"}
	secretKey1 = []byte("secret3")
	secretKey2 = []byte("secret4")
	testClaim1 = TokenData{UserId: 765567}
	testClaim2 = TokenData{UserId: 886432}
)

func TestGenerateToken(t *testing.T) {
	// generate and validate a valid token
	token, err := generateToken(testClaim1, secretKey1, defaultAccessExpiry)
	_verifyToken(token, err, testClaim1, secretKey1, t)
	// check with invalid secret should throw err
	_, err = verifytoken(token, secretKey2)
	ShouldErr(err, t)
}

func TestExpiredToken(t *testing.T) {
	token, err := generateToken(testClaim2, secretKey2, time.Second*3)
	_verifyToken(token, err, testClaim2, secretKey2, t)
	time.Sleep(time.Second * 4)
	_, err = verifytoken(token, secretKey2)
	ShouldErr(err, t)
}

func TestAccessAndRefreshToken(t *testing.T) {
	InitConfig(testConfig)
	accessToken, refreshToken, err := GetTokens(testClaim1)
	_verifyToken(accessToken, err, testClaim1, []byte(testConfig.AccessSecret), t)
	claim, err := VerifyAccessToken(accessToken)
	NoErr(err, t)
	Check(claim.UserId == testClaim1.UserId, t)
	_verifyToken(refreshToken, err, testClaim1, []byte(testConfig.RefreshSecret), t)
	claim, err = VerifyRefreshToken(refreshToken)
	NoErr(err, t)
	Check(claim.UserId == testClaim1.UserId, t)
}

func TestInvalidCnfig(t *testing.T) {
	testConfig.AccessSecret = ""
	defer TestPanic(t)
	InitConfig(testConfig) // should panic
}

func _verifyToken(token string, err error, claimData TokenData, secret []byte, t *testing.T) {
	NoErr(err, t)
	tokenData, err := verifytoken(token, secret)
	NoErr(err, t)
	Check(tokenData.UserId == claimData.UserId, t)
}
