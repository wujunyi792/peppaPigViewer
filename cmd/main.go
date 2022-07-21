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

	for _, user := range c {
		res, e := core.LoadConfig(user).LoginPW(user.User.StaffId, user.User.Password)
		if e != nil {
			panic(e)
		}
		res.PrintCourseChosenList()
		res.SetTarget(user.Target).FindCourse().PrintFireCourseList()
		courses, e := res.FireCourses()
		if e != nil {
			panic(err)
		}
		fmt.Println(courses)
	}
	select {}
}
