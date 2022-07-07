package main

import (
	"errors"
	"github.com/parnurzeal/gorequest"
	"log"
)

func getChosenCourse(query string, form *GetChosenCourseReq) *[]CourseChosenResp {
	var res []CourseChosenResp
	log.Println("【request】", JWBASE+"/xsxk/zzxkyzb_cxZzxkYzbChoosedDisplay.html?"+query)
	gorequest.New().Post(JWBASE+"/xsxk/zzxkyzb_cxZzxkYzbChoosedDisplay.html?"+query).Set("User-Agent", UA).Set("Cookie", getCookie()).
		Type("form-data").
		Send(form.makeForm()).EndStruct(&res)
	return &res
}

func getCoursePage(url string) string {
	body, exist := getCache("courseHome")
	if !exist {
		log.Println("【request】", url)
		res, body, _ := gorequest.New().Get(url).Set("User-Agent", UA).Set("Cookie", getCookie()).End()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}
		setCache("courseHome", body)
		return body
	}
	return body
}

// https://newjw.hdu.edu.cn/jwglxt/xsxk/zzxkyzb_cxZzxkYzbDisplay.html
func getDisplayPage(url string, form *GetDisplayReq) string {
	body, exist := getCache("courseDisplay")
	if !exist {
		log.Println("【request】", url)
		res, body, _ := gorequest.New().Post(url).Set("User-Agent", UA).Set("Cookie", getCookie()).
			Type("form-data").
			Send(form.makeForm()).End()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}
		setCache("courseDisplay", body)
		return body
	}
	return body
}

func getCourseList(url string, form *FindClassReq) *CourseListResp {
	var res CourseListResp
	log.Println("【request】", url)
	gorequest.New().Post(url).Set("User-Agent", UA).Set("Cookie", getCookie()).
		Type("form-data").
		Send(form.makeForm()).
		EndStruct(&res)
	return &res
}

func getCourseDetail(url string, form *GetCourseDetailReq) *[]CourseDetail {
	var res []CourseDetail
	log.Println("【request】", url)
	gorequest.New().Post(url).Set("User-Agent", UA).Set("Cookie", getCookie()).
		Type("form-data").
		Send(form.makeForm()).EndStruct(&res)
	return &res
}

func prvChooseCourse(url string, form *ChooseClassPrvReq) error {
	type flag struct {
		Flag string `json:"flag"`
		Msg  string `json:"msg"`
	}
	var res flag
	log.Println("【request】", url)
	gorequest.New().Post(url).Set("User-Agent", UA).Set("Cookie", getCookie()).
		Type("form-data").
		Send(form.makeForm()).EndStruct(&res)
	switch res.Flag {
	case "1":
		return nil
	case "-1":
		return errors.New("容量超出")
	default:
		return errors.New(res.Msg)
	}
}
