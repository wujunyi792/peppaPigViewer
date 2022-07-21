package core

import (
	"errors"
	"fmt"
	"log"
	"newJwCourseHelper/internal/dto"
)

func (u *User) FindCourse() *User {
	if len(u.config.target) == 0 {
		u.e = errors.New("empty target, please set")
		return u
	}

	findClassBaseField := dto.MakeFindClassReq(u.getField())
	findClassBaseField.FilterList = u.getTarget()
	list := u.getCourseList(findClassBaseField)
	getClassDetailField := dto.MakeGetClassDetailReq(u.getField())
	for i := 0; i < len(list.TmpList); {
		getClassDetailField.KchId = list.TmpList[i].KchId
		getClassDetailField.FilterList = u.getTarget()

		details := u.getCourseDetail(getClassDetailField)
		if *details == nil {
			id := list.TmpList[i].KchId
			for j := 0; j < len(list.TmpList); j++ {
				if list.TmpList[j].KchId == id {
					i++
				}
			}
			continue
		}
		for index, detail := range *details {
			for j := 0; j < len(list.TmpList); j++ {
				if list.TmpList[j].JxbId == detail.JxbId {
					list.TmpList[j].DetailList = &(*details)[index]
					list.TmpList[j].HaveSet = list.TmpList[j].Yxzrs < (*details)[index].Jxbrl
					i++
					break
				}
			}
		}
	}
	log.Printf("使用关键词 【 %s 】 共查询到 %d 门课程\n", u.getTarget(), len(list.TmpList))
	u.courses = list
	u.e = nil
	return u
}

// PrintFireCourseList 输出待选课的列表
func (u *User) PrintFireCourseList() {
	if u.Error() != nil {
		log.Printf("find an err: %v\n", u.Error())
		return
	}
	if u.courses == nil {
		u.e = errors.New("empty course list, please use FindCourse first")
		log.Printf("find an err: %v\n", u.Error())
		return
	}
	for i := 0; i < len(u.courses.TmpList) && u.courses.TmpList[i].DetailList != nil; i++ {
		fmt.Printf("【%02d】 %s 课程号 %s 班级号 %s    总容量 %s 已选 %s\n",
			i+1,
			u.courses.TmpList[i].Kcmc,
			u.courses.TmpList[i].Kch,
			u.courses.TmpList[i].Jxbmc,
			(*u.courses.TmpList[i].DetailList).Jxbrl,
			u.courses.TmpList[i].Yxzrs)
	}
}

func (u *User) FireCourses() ([]string, error) {
	if u.Error() != nil {
		log.Printf("find an err: %v", u.Error())
		return nil, u.Error()
	}
	if u.courses == nil {
		u.e = errors.New("empty course list, please use FindCourse first")
		log.Printf("find an err: %v", u.Error())
		return nil, u.Error()
	}

	fireList := u.courses.TmpList
	var success []string

	for i := 0; i < len(fireList) && fireList[i].DetailList != nil; i++ {
		// 跳过选课失败的课程 & 已选课程
		{
			if u.checkInErrList(fireList[i].Jxbmc) || u.checkChosen(fireList[i].Kch) {
				continue
			}
		}

		// 有余量则选课
		if fireList[i].HaveSet {

			prvChooseReq := dto.MakeChooseClassPrvReq(u.getField())
			prvChooseReq.JxbIds = (*fireList[i].DetailList).DoJxbId
			prvChooseReq.KchId = fireList[i].Kch
			prvChooseReq.Cxbj = fireList[i].Cxbj
			prvChooseReq.Xxkbj = fireList[i].Xxkbj

			err := u.prvChooseCourse(prvChooseReq)
			if err != nil {
				log.Printf("【err】 选择 %s 时发生错误： %v\n", fireList[i].Jxbmc, err.Error())
				u.config.errTag = append(u.config.errTag, fireList[i].Jxbmc)
			}
			success = append(success, fireList[i].Jxbmc)

			// 刷新已选课程
			c := u.getChosenCourse(dto.MakeGetChosenClassReq(u.getField()))
			if c == nil {
				u.e = errors.New("get user chosen course failed")
			} else {
				u.info.chosenCourse = c
			}
		}
	}
	return success, u.Error()
}

func (u *User) checkInErrList(m string) bool {
	for _, s := range u.config.errTag {
		if m == s {
			return true
		}
	}
	return false
}

func (u *User) checkChosen(m string) bool {
	for j := 0; j < len(*u.info.chosenCourse); j++ {
		if (*u.info.chosenCourse)[j].Kch == m {
			return true
		}
	}
	return false
}

func (u *User) PrintCourseChosenList() {
	if u.info.chosenCourse == nil || len(*u.info.chosenCourse) == 0 {
		u.e = errors.New("empty course list")
		log.Printf("find an err: %v\n", u.Error())
		return
	}
	cl := *u.info.chosenCourse
	for i := 0; i < len(cl); i++ {
		fmt.Printf("【%02d】 %s 课程号 %s 班级号 %s 教师 %s\n",
			i+1,
			cl[i].Kcmc,
			cl[i].Kch,
			cl[i].Jxbmc,
			cl[i].Sksj)
	}
}
