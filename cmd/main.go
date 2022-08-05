package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/robfig/cron"
	"io/ioutil"
	"newJwCourseHelper/internal/config"
	"newJwCourseHelper/internal/core"
	"strconv"
	"time"
)

var configFile = flag.String("f", "config.json", "Specify the config file")

func main() {
	flag.Parse()
	var c []config.Config

	content, err := ioutil.ReadFile(*configFile)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(content, &c); err != nil {
		panic(err)
	}

	//c中的每一个元素都是一个config.Config结构体，保存了用户的配置、任务等
	for _, user := range c {
		res, e := core.LoadConfig(user).LoginPW(user.User.StaffId, user.User.Password) //用户登录（TODO:不需要多次登录，可以直接传入Token）
		if e != nil {
			panic(e)
		}
		res.PrintCourseChosenList()                                   //输出已选课程列表
		res.SetTarget(user.Target).FindCourse().PrintFireCourseList() //输出待选课程列表//继续debug，把cofig文件对应的结构体数组修改好

		timer := cron.New()
		err := timer.AddFunc("*/"+strconv.Itoa(user.Interval)+" * * * * *", func() {
			courses, e := res.FindCourse().FireCourses()
			if e != nil {
				panic(err)
			}
			if len(courses) == 0 {
				fmt.Println("暂时无可选课程")
			} else {
				fmt.Println("已选到如下课程：")
				fmt.Println(courses)
			}
		})
		go timer.Start()
		defer timer.Stop()
		if err != nil {
			fmt.Println("function Error!")
		} //每3秒打印一次已选课程列表

	}
	select {
	case <-time.After(time.Second * 1000):
		return
	}
}
