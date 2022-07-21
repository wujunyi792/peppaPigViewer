package core

import "github.com/patrickmn/go-cache"

type userFunc interface {
	init()
	getCookie() string
	getStaffId() string
	getCache() *cache.Cache
	getField() map[string]string
	getBaseQuery() string
	getTarget() []string

	SetTarget(r []string) *User
	FindCourse() *User
	PrintFireCourseList()
	PrintCourseChosenList()
}
