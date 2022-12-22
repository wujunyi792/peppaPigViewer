package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/robfig/cron"
	"io"
	"log"
	"newJwCourseHelper/internal/core"
	"os"
	"syscall"
)

var ChooseCourseLoggerBuffer = &bytes.Buffer{}
var MultiBuffer = io.MultiWriter(os.Stdout, ChooseCourseLoggerBuffer)
var ChooseCourseLogger = log.New(MultiBuffer, "[DebugMessage] ", log.LstdFlags)
var interactiveMode = flag.Bool("i", false, "Interactive mode")

// Auto ChooseCourse func will start with new go runtime
func ChooseCourse() {
	core.Init(ChooseCourseLogger)
	for _, user := range c {
		res, e := core.LoadConfig(user).LoginPW(user.User.StaffId, user.User.Password) //用户登录（TODO:不需要多次登录，可以直接传入Token）
		if e != nil {
			panic(e)
		}
		//res.PrintCourseChosenList()                                   //输出已选课程列表
		res.SetTarget(user.Target) //输出待选课程列表//继续debug，把config文件对应的结构体数组修改好

		// 立刻抢课
		core.Job(res)
		res.SetCorn(cron.New())
		var cornExpr = fmt.Sprintf("*/%d * * * * *", user.Interval)
		err := res.GetCorn().AddFunc(cornExpr, func() {
			core.Job(res)
		})
		if err != nil {
			log.Printf("用户 %s 定时任务添加失败: %v", user.User.StaffId, err)
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

		// Todo: Process command Should not be there. Need reform later.
		switch cmd {
		case "help":
			fmt.Println("help: show help")
			fmt.Println("log/dmesg: print debug message")
			fmt.Println("list/status: list all user status currently")
			fmt.Println("relogin: forced re auth to the server")
			fmt.Println("quit/exit: exit the program")
		case "log", "dmesg":
			fmt.Println(ChooseCourseLoggerBuffer)
		case "list", "status":
			for k, user := range users {
				log.Printf("This is the User #%v Chosen CourseList \n", k)
				user.PrintCourseChosenList()
				log.Printf("This is the User #%v Target CourseList \n", k)
				user.PrintFireCourseList()
			}
		case "relogin":
			for k, user := range users {
				var err error
				users[k], err = user.ForceRetryAuth()
				if err != nil {
					log.Printf("Re Login user[%v] got error: %v \n", k, err)
				}
			}
		case "geterror":
			for k, user := range users {
				log.Printf("This is user #%v Error: %v \n", k, user.Error())
			}
		case "quit", "exit":
			for i := 0; i < len(users); i++ {
				users[i].GetCorn().Stop()
			}
			quit <- syscall.SIGTERM
			os.Exit(0)
		}
	}
}
