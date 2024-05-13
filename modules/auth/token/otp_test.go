package token

import (
	"slices"
	"testing"

	. "github.com/qnify/api-server/utils/helper"
)

func TestGenerateOTP(t *testing.T) {
	otpLengths := []int{3, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 6, 6, 6}

	var otps = make([]string, len(otpLengths))

	for index := range otpLengths {
		otp, err := GenerateOtp()
		NoErr(err, t)
		Check(!slices.Contains(otps, otp), t) // to check uniqueness
		otps[index] = otp
		Check(len(otp) == otpLength, t)
	}
}

func TestOtpVerification(t *testing.T) {
	const testMail, otp = "test@test.com", "39933"
	token := GetOtpToken(testMail, otp)
	Check(VerifyOtpToken(testMail, otp, token), t)
}
