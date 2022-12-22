package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/robfig/cron"
	"log"
	"newJwCourseHelper/internal/config"
	"newJwCourseHelper/internal/core"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var configFile = flag.String("f", "config.json", "Specify the config file")

var quit = make(chan os.Signal)
var users []*core.User
var c []config.Config

// Auto ChooseCourse func will start with new go runtime
func ChooseCourse() {
	for _, user := range c {
		res, e := core.LoadConfig(user).LoginPW(user.User.StaffId, user.User.Password) //用户登录（TODO:不需要多次登录，可以直接传入Token）
		if e != nil {
			panic(e)
		}
		//res.PrintCourseChosenList()                                   //输出已选课程列表
		res.SetTarget(user.Target) //输出待选课程列表//继续debug，把config文件对应的结构体数组修改好

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

	// if quit Print Information
	<-quit
	for i := 0; i < len(users); i++ {
		users[i].GetCorn().Stop()
	}
}

// Command Line Execute for relogin and list Current Status
func Execute() {
	for {
		var cmd string
		fmt.Printf("Console> ")
		_, _ = fmt.Scanln(&cmd)

		switch cmd {
		case "list":
			for i := 0; i < len(users); i++ {
				log.Printf("This is the User #%v Chosen CourseList \n")
				users[i].PrintCourseChosenList()
			}
		case "relogin":
			for k, user := range users {
				var err error
				users[k], err = user.ForceRetryAuth()
				if err != nil {
					log.Printf("Re Login user[%v] got error: %v \n", k, err)
				}
			}
		case "quit":
			quit <- syscall.SIGTERM
			os.Exit(0)
		}
	}
}

func main() {
	flag.Parse()

	content, err := os.ReadFile(*configFile)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(content, &c); err != nil {
		panic(err)
	}

	//c中的每一个元素都是一个config.Config结构体，保存了用户的配置、任务等

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go ChooseCourse()
	Execute()
}
