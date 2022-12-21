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
	// warn: Patch There with Retry
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

func (u *User) getCourseList(form *dto.FindClassReq, targetArr []Target) *dto.CourseListResp { //KklxdmArr表示不同种类课程的Kklxdm代码
	var res, tempRes dto.CourseListResp
	classTypeMap := make(map[int][]Target)
	for i := 0; i < 3; i++ {
		classTypeMap[i] = make([]Target, 0)
	}

	for i := 0; i < len(targetArr); i++ {
		classTypeMap[targetArr[i].Type] = append(classTypeMap[targetArr[i].Type], targetArr[i])
	}

	for key, value := range classTypeMap {
		if len(value) == 0 {
			delete(classTypeMap, key)
		}
	}

	for classType, targets := range classTypeMap {
		form.Kklxdm = u.info.special["firstKklxdmArr"][classType]
		form.FilterList = []string{}
		for _, target := range targets {
			form.FilterList = append(form.FilterList, target.Name)
		}
		_, err := u.client.R().
			SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8"). //设置多次查询
			SetResult(&tempRes).SetBody(form.MakeForm()).Post(JwBase + JwApiCourseList + u.getBaseQuery())
		// force return nil
		if err != nil {
			log.Println(err)
			return nil
		}
		for i := 0; i < len(tempRes.TmpList); i++ {
			tempRes.TmpList[i].ClassType = classType
		}
		res.TmpList = append(res.TmpList, tempRes.TmpList...)
	}
	return &res
}

func (u *User) getCourseDetail(form *dto.GetCourseDetailReq, classType int) *[]dto.CourseDetail {
	var res, tempArr []dto.CourseDetail
	var err error

	form.Kklxdm = u.info.special["firstKklxdmArr"][classType]
	form.XkkzId = u.info.special["firstXkkzIdArr"][classType]
	form.NjdmId = u.info.special["firstNjdmIdArr"][classType]
	form.ZyhId = u.info.special["firstZyhIdArr"][classType]
	if classType == 0 {
		form.Rwlx = "1"
		form.Sfkknj = "1"
		form.Sfkkzy = "1"
		form.Xkxskcgskg = "0"
	} else if classType == 1 {
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
		SetResult(&tempArr).SetBody(form.MakeForm()).Post(JwBase + JwApiCourseDetail + u.getBaseQuery())
	// warn: Patch There when need relogin.
	// err = ERROR_NEED_RELOGIN
	if err != nil {
		log.Println(err)
		u.e = err
		if u.IsRetryAuth() {
			u, _ = u.ForceRetryAuth()
		}
		return nil // Todo: return nil Will breakdown the program May need handle this
	}
	res = append(res, tempArr...)
	//这里每次循环要删掉form.FilterList中的一个值，要不然之后的课程查不出来//TODO: 有冗余的过程
	form.FilterList = form.FilterList[1:]
	//}
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
