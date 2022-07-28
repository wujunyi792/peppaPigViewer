package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"newJwCourseHelper/internal/config"
	"newJwCourseHelper/internal/core"
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
		res.SetTarget(user.Target).FindCourse().PrintFireCourseList() //输出待选课程列表
		courses, e := res.FireCourses()
		if e != nil {
			panic(err)
		}
		if len(courses) == 0 {
			fmt.Println("暂时无课程")
		} else {
			fmt.Println(courses)
		}
	}
	select {}
}
