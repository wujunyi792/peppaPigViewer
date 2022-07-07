package main

import (
	"github.com/parnurzeal/gorequest"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

type Entrance struct {
	host string
	list []EntranceList
}

type EntranceList struct {
	name string
	uri  string
	id   string
}

func getHomePage() string {
	res, body, _ := gorequest.New().Get(JWBASE+JWHOMEPAGEURI+strconv.Itoa(int(time.Now().UnixMilli()))).
		RedirectPolicy(func(req gorequest.Request, via []gorequest.Request) error {
			return http.ErrUseLastResponse
		}).
		Set("User-Agent", UA).Set("Cookie", getCookie()).End()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	if res.Request.URL.Path == "/jwglxt/xtgl/login_slogin.html" {
		log.Fatalf("cannot login newjw, the server redirect to login page")
	}
	setCache("homePage", body)
	return body
}

func getEntranceList() *Entrance {
	body, exist := getCache("homePage")
	if !exist {
		body = getHomePage()
	}

	var re = regexp.MustCompile(`(?m)clickMenu\('(?P<moduleId>.*?)','(?P<uri>.*?)','(?P<moduleName>.*?)','(?P<isFunc>.*?)'\)`)
	matches := re.FindAllStringSubmatch(body, -1)
	var entries []EntranceList
	for _, match := range matches {
		entries = append(entries, EntranceList{
			name: match[3],
			uri:  match[2],
			id:   match[1],
		})
	}
	return &Entrance{
		list: entries,
	}
}

func getUserID() string {
	body, exist := getCache("homePage")
	if !exist {
		body = getHomePage()
	}
	var re = regexp.MustCompile(`(?m)<input type="hidden" id="sessionUserKey" value="(?P<su>.*?)" />`)
	matches := re.FindAllStringSubmatch(body, -1)
	if len(matches) == 0 {
		log.Fatalln("err when get student number")
	}
	return matches[0][1]
}
