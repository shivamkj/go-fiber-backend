package auth

import (
	"database/sql"

	"github.com/qnify/api-server/utils/errors"
	. "github.com/qnify/api-server/utils/helper"
)

const Unspecified = iota

const (
	PhoneOtp = iota
	EmailOtp
	Google
	Apple
)

const (
	Female = iota
	Male
)

func (m *authModule) findUser(phone string, email string) (int, string, error) {
	var (
		userId     int
		phoneEmail string
		err        error
	)

	if email != "" {
		if !IsValidEmail(email) {
			return userId, "", errors.BadRequest("invalid email")
		}
		row := m.db.QueryRow(`SELECT id FROM users WHERE email = $1`, email)
		err = row.Scan(&userId)
		phoneEmail = email
	} else if phone != "" {
		row := m.db.QueryRow(`SELECT id FROM users WHERE phone = $1`, phone)
		err = row.Scan(&userId)
		phoneEmail = phone
	} else {
		return userId, "", errors.BadRequest("invalid phone/email")
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return userId, "", errors.NotFound("user not found")
		}
		return userId, "", errors.InternalError("error occured while searching user", err)
	}

	return userId, phoneEmail, err
}
