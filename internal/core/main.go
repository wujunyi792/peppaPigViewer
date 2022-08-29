package core

import (
	cas "github.com/wujunyi792/hdu-cas-helper"
	"newJwCourseHelper/internal/config"
	"newJwCourseHelper/internal/util/ua"
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
		c.Ua = ua.GetUA()
	}
	if c.Rate < 1 {
		c.Rate = 3
	}
	if c.BucketFull > 10 {
		c.BucketFull = 10
	}
	u := new(User)
	targetArr := make([]Target, 0)
	for _, each := range c.Target {
		targetArr = append(targetArr, Target{Name: each.Name, Type: each.Type})
	}
	u.config = &missionConfig{
		target:     targetArr,
		errTag:     c.ErrTag,
		bucketFull: c.BucketFull,
		rate:       c.Rate,
		ua:         c.Ua,
	}
	return u
}

func (u *User) Error() error {
	return u.e
}
