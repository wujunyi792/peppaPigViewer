package core

import (
	"github.com/go-resty/resty/v2"
	"github.com/patrickmn/go-cache"
	"github.com/robfig/cron"
	cas "github.com/wujunyi792/hdu-cas-helper"
	"newJwCourseHelper/internal/config"
	"newJwCourseHelper/internal/dto"
	"time"
)

type baseInfo struct {
	staffId      string
	field        map[string]string
	chosenCourse *[]dto.CourseChosenResp
	baseQuery    string
	special      map[string][]string
}

type missionConfig struct {
	target []Target
	errTag []string
	rate   int
	ua     string
}

type Target struct {
	Name string `json:"name"`
	Type int    `json:"type"`
}

type User struct {
	auth        *cas.NewJW
	loginUser   string
	retryConfig *config.Config

	formParam     map[string]string
	info          *baseInfo
	requestTicket *time.Ticker
	config        *missionConfig
	client        *resty.Client
	cache         *cache.Cache

	cron *cron.Cron

	courses *dto.CourseListResp

	e error
}

type entrance struct {
	host string
	list []entranceList
}

type entranceList struct {
	name string
	uri  string
	id   string
}
