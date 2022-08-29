package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/robfig/cron"
	"newJwCourseHelper/internal/config"
	"newJwCourseHelper/internal/core"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var configFile = flag.String("f", "config.json", "Specify the config file")

func main() {
	flag.Parse()
	var c []config.Config
	var users []*core.User

	content, err := os.ReadFile(*configFile)
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
		res.SetTarget(user.Target).FindCourse().PrintFireCourseList() //输出待选课程列表//继续debug，把config文件对应的结构体数组修改好

		res.SetCorn(cron.New())
		err := res.GetCorn().AddFunc("*/"+strconv.Itoa(user.Interval)+" * * * * *", func() {
			core.Job(res)
		})
		if err != nil {
			fmt.Printf("用户 %s 定时任务添加失败: %v", user.User.StaffId, err)
		}
		users = append(users, res)
	}

	for i := 0; i < len(users); i++ {
		users[i].GetCorn().Start()
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	for i := 0; i < len(users); i++ {
		users[i].GetCorn().Stop()
	}
}
