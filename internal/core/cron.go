package core

import "github.com/robfig/cron"

func (u *User) GetCorn() *cron.Cron {
	return u.cron
}

func (u *User) SetCorn(e *cron.Cron) {
	u.cron = e
}
