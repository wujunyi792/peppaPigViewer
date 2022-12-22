package core

import (
	"github.com/robfig/cron"
)

func (u *User) GetCorn() *cron.Cron {
	return u.cron
}

func (u *User) SetCorn(e *cron.Cron) {
	u.cron = e
}

func Job(user *User) {
	user.e = nil
	courses, e := user.FindCourse().PrintFireCourseList().FireCourses()
	if e != nil {
		log.Print(e)
		return
	}
	if len(courses) == 0 {
		log.Println("暂时无可选课程")
	} else {
		log.Println("已选到如下课程：")
		log.Println(courses)
	}
	log.Println()
}
