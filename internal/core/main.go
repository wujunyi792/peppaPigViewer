package core

import (
	cas "github.com/wujunyi792/hdu-cas-helper"
)

func LoginPW(username, password string) (*User, error) {
	ticket := cas.CasPasswordLogin(username, password)
	session := cas.NewJWLogin(ticket)
	if err := session.Error(); err != nil {
		return nil, err
	}
	user := new(User)
	user.auth = session
	user.init()
	err := user.serviceInit()

	return user, err
}

func (u *User) Error() error {
	return u.e
}
