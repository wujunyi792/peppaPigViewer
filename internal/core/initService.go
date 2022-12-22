package core

import (
	"errors"
	"net/url"
	"newJwCourseHelper/internal/dto"
	"newJwCourseHelper/internal/util/field"
	"newJwCourseHelper/internal/util/idArr"
)

func (u *User) serviceInit() error { //TODO:在初始化函数中将其他的firstXkkzId、firstNjdmId、firstZyhId等初始化
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
		log.Println("无法找到选课入口，可能当前不在选课时间，也可能网站裂了，也可能你被ban了，将使用默认入口")
		reqUrl.Path += DEFAULTCHOOSECLASSURL
		query := make(url.Values)
		query.Add("gnmkdm", "N253512")
		query.Add("layout", "default")
		query.Add("su", u.getStaffId())
		reqUrl.RawQuery = query.Encode()
		lid = "N253512"
	}

	baseQuery := make(url.Values)
	baseQuery.Add("gnmkdm", lid)
	baseQuery.Add("su", u.getStaffId())
	u.info.baseQuery = baseQuery.Encode()

	p := u.getCoursePage(reqUrl.String()) //获取到了课程页面的源代码，可以在这里添加加入其他课程组的XkkzId等
	u.info.field = field.GetInputField(p, nil)
	u.info.special = idArr.FindIDArr(p, ClassNumber) //修改

	if u.info.special == nil {
		return errors.New("选课基本参数获取失败，看起来不在选课时间")
	}

	p = u.getDisplayPage(dto.MakeGetDisplayReq(u.getField()))
	u.info.field = field.GetInputField(p, u.getField())

	chosenCourse := u.getChosenCourse(dto.MakeGetChosenClassReq(u.getField()))
	if chosenCourse == nil {
		return errors.New("get user chosen course failed")
	}
	u.info.chosenCourse = chosenCourse

	return nil
}
