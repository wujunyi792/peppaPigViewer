package core

import (
	cas "github.com/wujunyi792/hdu-cas-helper"
	"newJwCourseHelper/internal/config"
)

func (u *User) LoginPW(username, password string) (*User, error) {
	ticket := cas.CasPasswordLogin(username, password)
	session := cas.NewJWLogin(ticket)
	if err := session.Error(); err != nil {
		return nil, err
	}
	user := u
	user.auth = session
	user.init()
	err := user.serviceInit()

	return user, err
}

func LoadConfig(c config.Config) *User {
	if c.Ua == "" {
		c.Ua = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36"
	}
	if c.Rate < 500 {
		c.Rate = 500
	}
	u := new(User)
	targetArr := make([]Target, 0)
	for _, each := range c.Target {
		targetArr = append(targetArr, Target{Name: each.Name, Type: each.Type})
	}
	// Adding retry module config
	u.retryConfig = &c
	u.config = &missionConfig{
		target: targetArr,
		errTag: c.ErrTag,
		rate:   c.Rate,
		ua:     c.Ua,
	}
	return u
}

func (u *User) Error() error {
	return u.e
}
