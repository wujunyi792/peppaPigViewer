package core

import (
	"errors"
	"net/url"
	"newJwCourseHelper/internal/dto"
	"newJwCourseHelper/internal/util/field"
)

func (u *User) serviceInit() error {
	e := u.getEntranceList()
	reqUrl := url.URL{
		Scheme: "https",
		Host:   JwHost,
		Path:   "/jwglxt",
	}
	flag := false
	lid := ""
	for _, list := range e.list {
		if list.name == "自主选课" {
			reqUrl.Path += list.uri
			query := make(url.Values)
			query.Add("gnmkdm", list.id)
			query.Add("layout", "default")
			query.Add("su", u.getStaffId())
			reqUrl.RawQuery = query.Encode()
			flag = true
			lid = list.id
			break
		}
	}
	if !flag {
		return errors.New("look like not in correct time to choose a course")
	}

	baseQuery := make(url.Values)
	baseQuery.Add("gnmkdm", lid)
	baseQuery.Add("su", u.getStaffId())
	u.info.baseQuery = baseQuery.Encode()

	p := u.getCoursePage(reqUrl.String())
	u.info.field = field.GetInputField(p, nil)

	p = u.getDisplayPage(dto.MakeGetDisplayReq(u.getField()))
	u.info.field = field.GetInputField(p, u.getField())

	chosenCourse := u.getChosenCourse(dto.MakeGetChosenClassReq(u.getField()))
	if chosenCourse == nil {
		return errors.New("get user chosen course failed")
	}
	u.info.chosenCourse = chosenCourse

	return nil
}
