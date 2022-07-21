package core

import (
	"github.com/patrickmn/go-cache"
	"golang.org/x/time/rate"
	"newJwCourseHelper/internal/util/ua"
	"time"
)

func (u *User) getCookie() string {
	if u.auth == nil {
		return ""
	}
	return u.auth.GetCookie()
}

func (u *User) getRequestRate() *rate.Limiter {
	return u.rateBucket
}

func (u *User) init() {
	u.formParam = make(map[string]string, 10)
	u.info = new(baseInfo)
	u.rateBucket = rate.NewLimiter(rate.Every(5*time.Second), 3)
	u.config = new(missionConfig)
	u.cache = cache.New(1*time.Hour, 24*time.Hour)
	u.client = u.NewRequestClient(ua.GetUA())
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

func (u *User) SetTarget(r []string) *User {
	u.config.target = r
	u.e = nil
	return u
}

func (u *User) getTarget() []string {
	return u.config.target
}
