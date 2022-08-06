package core

import (
	"github.com/pkg/errors"
	"log"
	"newJwCourseHelper/internal/dto"
	"strconv"
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

func (u *User) getCourseList(form *dto.FindClassReq, KklxdmArr []string, targetArr []Target) (*dto.CourseListResp, []int) { //KklxdmArr表示不同种类课程的Kklxdm代码
	var res, tempRes dto.CourseListResp
	var eachLen []int

	var err error
	tempFilterList := form.FilterList
	for i, eachTarget := range targetArr {
		tempInt, _ := strconv.Atoi(eachTarget.Type) //这边暂时设定classNumber只能有一个值且一定为整数：[0,2]
		form.Kklxdm = KklxdmArr[tempInt]
		form.FilterList = []string{tempFilterList[i]}                                                       //防止有的target使用错误的参数进行查询(这里要保证targetArr和form.FilterList的长度一致)
		_, err = u.client.R().SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8"). //设置多次查询
															SetResult(&tempRes).SetBody(form.MakeForm()).Post(JwBase + JwApiCourseList + u.getBaseQuery())
		if err != nil {
			log.Println(err)
			return nil, []int{}
		}
		res.TmpList = append(res.TmpList, tempRes.TmpList...)
		eachLen = append(eachLen, len(tempRes.TmpList))
	}
	return &res, eachLen
}

func (u *User) getCourseDetail(form *dto.GetCourseDetailReq, special map[string][]string, classNumber string) *[]dto.CourseDetail {
	var res, tempArr []dto.CourseDetail
	var err error

	//for _, eachTarget := range targetArr {
	i, _ := strconv.Atoi(classNumber) //这边暂时设定classNumber只能有一个值且一定为整数：[0,2]
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
		SetResult(&tempArr).SetBody(form.MakeForm()).Post(JwBase + JwApiCourseDetail + u.getBaseQuery())
	if err != nil {
		log.Println(err)
		return nil
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
