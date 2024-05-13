package token

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/qnify/api-server/utils/helper"
)

const (
	otpLength = 5
	otpExpiry = time.Minute * 10
	hashSep   = "<"
)

var otpSecret = []byte("super secret")

func GenerateOtp() (string, error) {
	max := math.Pow10(otpLength)
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return "", err
	}

	// Format the random number as a string with leading zeros if necessary
	otp := fmt.Sprintf("%0*d", otpLength, randomNumber)
	return otp, nil
}

func GetOtpToken(phoneOrEmail string, otp string) string {
	expires := time.Now().Add(otpExpiry).Unix()
	data := phoneOrEmail + "." + otp + "." + strconv.FormatInt(expires, 10)
	hash := helper.GenerateHMAC(data, otpSecret)
	token := hash + hashSep + strconv.FormatInt(expires, 10)
	return token
}

func VerifyOtpToken(phoneOrEmail string, otp string, hash string) bool {
	parts := strings.Split(hash, hashSep)
	if len(parts) != 2 {
		return false
	}

	hashValue := parts[0]
	expires, err := strconv.Atoi(parts[1])
	now := time.Now().Unix()
	if err != nil || int(now) > expires {
		return false
	}

	data := phoneOrEmail + "." + otp + "." + strconv.Itoa(expires)
	newCalculatedHash := helper.GenerateHMAC(data, otpSecret)
	return newCalculatedHash == hashValue
}
