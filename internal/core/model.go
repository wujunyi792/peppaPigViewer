package core

import (
	"github.com/go-resty/resty/v2"
	"github.com/patrickmn/go-cache"
	"github.com/robfig/cron"
	cas "github.com/wujunyi792/hdu-cas-helper"
	"golang.org/x/time/rate"
	"newJwCourseHelper/internal/dto"
)

type baseInfo struct {
	staffId      string
	field        map[string]string
	chosenCourse *[]dto.CourseChosenResp
	baseQuery    string
	special      map[string][]string
}

type missionConfig struct {
	target     []Target
	errTag     []string
	bucketFull int
	rate       int
	ua         string
}

type Target struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type User struct {
	auth       *cas.NewJW
	formParam  map[string]string
	info       *baseInfo
	rateBucket *rate.Limiter
	config     *missionConfig
	client     *resty.Client
	cache      *cache.Cache

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
