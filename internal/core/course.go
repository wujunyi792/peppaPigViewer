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

func (u *User) getCourseList(form *dto.FindClassReq, KklxdmArr []string) *dto.CourseListResp { //KklxdmArr表示不同种类课程的Kklxdm代码
	var res, tempRes dto.CourseListResp

	var err error
	for _, kklxdm := range KklxdmArr {
		form.Kklxdm = kklxdm
		_, err = u.client.R().SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8"). //设置多次查询
															SetResult(&tempRes).SetBody(form.MakeForm()).Post(JwBase + JwApiCourseList + u.getBaseQuery())
		if err != nil {
			log.Println(err)
			return nil
		}
		res.TmpList = append(res.TmpList, tempRes.TmpList...)
	}
	return &res
}

func (u *User) getCourseDetail(form *dto.GetCourseDetailReq, special map[string][]string) *[]dto.CourseDetail {
	var res, tempArr []dto.CourseDetail
	var err error
	for i := 0; i < ClassNumber; i++ {
		form.Kklxdm = special["firstKklxdmArr"][i]
		form.XkkzId = special["firstXkkzIdArr"][i]
		form.NjdmId = special["firstNjdmIdArr"][i]
		form.ZyhId = special["firstZyhIdArr"][i]
		if i == 0 {
			form.Rwlx = "1"
			form.Sfkknj = "1"
			form.Sfkkzy = "1"
			form.Xkxskcgskg = "0"
		} else if i == 1 {
			form.Rwlx = "2"
			form.Sfkknj = "0"
			form.Sfkkzy = "0"
			form.Xkxskcgskg = "1"
		} else {
			form.Rwlx = "2"
			form.Sfkknj = "0"
			form.Sfkkzy = "0"
			form.Xkxskcgskg = "0"
		}

		_, err = u.client.R().SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8").
			SetResult(&tempArr).SetBody(form.MakeForm()).Post(JwBase + JwApiCourseDetail + u.getBaseQuery()) //这里只能获取每一个大类中的第一个课程，所以在config.json中需要填写的是课程代码以保证唯一性
		if err != nil {
			log.Println(err)
			return nil
		}
		res = append(res, tempArr...) //TODO: 可能出现重复的课程

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
