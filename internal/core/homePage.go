package core

import (
	"regexp"
)

func (u *User) getHomePage() string {
	resp, _ := u.client.R().Get(JwBase + new(JwApiHome).GetString())
	if resp.StatusCode() != 200 {
		return ""
	}
	page := resp.Body()
	u.getCache().SetDefault("homePage", string(page))
	return string(page)
}

func (u *User) getEntranceList() *entrance {
	body, exist := u.getCache().Get("homePage")
	if !exist {
		body = u.getHomePage()
	}

	var re = regexp.MustCompile(`(?m)clickMenu\('(?P<moduleId>.*?)','(?P<uri>.*?)','(?P<moduleName>.*?)','(?P<isFunc>.*?)'\)`)
	matches := re.FindAllStringSubmatch(body.(string), -1)
	var entries []entranceList
	for _, match := range matches {
		entries = append(entries, entranceList{
			name: match[3],
			uri:  match[2],
			id:   match[1],
		})
	}
	return &entrance{
		list: entries,
	}
}

func (u *User) getStaffId() string {
	if u.info.staffId != "" {
		return u.info.staffId
	}
	body, exist := u.getCache().Get("homePage")
	if !exist {
		body = u.getHomePage()
	}
	var re = regexp.MustCompile(`(?m)<input type="hidden" id="sessionUserKey" value="(?P<su>.*?)" />`)
	matches := re.FindAllStringSubmatch(body.(string), -1)
	if len(matches) == 0 {
		log.Println("Err when find core session id")
		log.Println("无法从首页获取用户学号，将使用默认用户名")
		u.info.staffId = u.loginUser
	} else {
		u.info.staffId = matches[0][1]
	}
	return u.info.staffId
}
