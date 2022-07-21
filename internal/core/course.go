package core

import (
	"github.com/pkg/errors"
	"log"
	"newJwCourseHelper/internal/dto"
)

func (u *User) getChosenCourse(form *dto.GetChosenCourseReq) *[]dto.CourseChosenResp {
	var res []dto.CourseChosenResp
	_, err := u.client.R().SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8").
		SetBody(form.MakeForm()).SetResult(&res).Post(JwBase + JwApiCourseChosen + u.getBaseQuery())
	if err != nil {
		log.Println(err)
		return nil
	}
	return &res
}

func (u *User) getCoursePage(url string) string {
	body, exist := u.getCache().Get("courseHome")
	if !exist {
		res, _ := u.client.R().Get(url)
		body = string(res.Body())
		u.getCache().SetDefault("courseHome", body)
		return body.(string)
	}
	return body.(string)
}

// https://newjw.hdu.edu.cn/jwglxt/xsxk/zzxkyzb_cxZzxkYzbDisplay.html
func (u *User) getDisplayPage(form *dto.GetDisplayReq) string {
	body, exist := u.getCache().Get("courseDisplay")
	if !exist {
		res, _ := u.client.R().SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8").
			SetBody(form.MakeForm()).Post(JwBase + JwDisplayPage + u.getBaseQuery())
		body = string(res.Body())
		u.getCache().SetDefault("courseDisplay", body)
	}
	return body.(string)
}

func (u *User) getCourseList(form *dto.FindClassReq) *dto.CourseListResp {
	var res dto.CourseListResp
	_, err := u.client.R().SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8").
		SetResult(&res).SetBody(form.MakeForm()).Post(JwBase + JwApiCourseList + u.getBaseQuery())
	if err != nil {
		log.Println(err)
		return nil
	}
	return &res
}

func (u *User) getCourseDetail(form *dto.GetCourseDetailReq) *[]dto.CourseDetail {
	var res []dto.CourseDetail
	_, err := u.client.R().SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8").
		SetResult(&res).SetBody(form.MakeForm()).Post(JwBase + JwApiCourseDetail + u.getBaseQuery())
	if err != nil {
		log.Println(err)
		return nil
	}
	return &res
}

func (u *User) prvChooseCourse(form *dto.ChooseClassPrvReq) error {
	type flag struct {
		Flag string `json:"flag"`
		Msg  string `json:"msg"`
	}
	var res flag
	_, err := u.client.R().SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8").
		SetResult(&res).SetBody(form.MakeForm()).Post(JwBase + JwApiChooseCourse + u.getBaseQuery())
	if err != nil {
		return err
	}
	switch res.Flag {
	case "1":
		return nil
	case "-1":
		return errors.New("容量超出")
	default:
		return errors.New(res.Msg)
	}
}
