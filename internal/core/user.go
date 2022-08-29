package core

import (
	"github.com/patrickmn/go-cache"
	"time"
)

func (u *User) getCookie() string {
	if u.auth == nil {
		return ""
	}
	return u.auth.GetCookie()
}

func (u *User) getRequestTicket() *time.Ticker {
	return u.requestTicket
}

func (u *User) init() {
	u.formParam = make(map[string]string, 10)
	u.info = new(baseInfo)
	u.requestTicket = time.NewTicker(time.Duration(u.config.rate) * time.Millisecond)
	u.config = new(missionConfig)
	u.cache = cache.New(1*time.Hour, 24*time.Hour)
	u.client = u.newRequestClient()
}

func (u *User) getCache() *cache.Cache {
	return u.cache
}

func (u *User) getField() map[string]string {
	return u.info.field
}

func (u *User) getBaseQuery() string {
	return u.info.baseQuery
}

func (u *User) SetTarget(r []struct {
	Name string `json:"name"`
	Type int    `json:"type"`
}) *User {
	for _, target := range r {
		if target.Type > 2 || target.Type < 0 {
			panic("请检查课程类型是否正确")
		}
		u.config.target = append(u.config.target, target)
	}
	u.e = nil
	return u
}

func (u *User) getTarget() []Target {
	return u.config.target
}

func (u *User) removeTarget(name string) {
	for index, target := range u.config.target {
		if target.Name == name {
			u.config.target = append(u.config.target[:index], u.config.target[index+1:]...)
			break
		}
	}
}
