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
	"syscall"
)

var pConfigFile = flag.String("f", "config.json", "Specify the config file")

const (
	ConfigUsingNotFound = iota
	ConfigUsingArg      = iota + 1
	ConfigUsingEnv      = iota + 2
)

func determineConfigPath() (configPath string) {
	var configMode = ConfigUsingNotFound
	var envConfigPath = os.Getenv("HELPER_CONFIG_PATH")
	if _, err := os.Stat(envConfigPath); err != nil {
		configMode = ConfigUsingEnv
	}
	//if arg and env are both presented, use arg with priority
	if _, err := os.Stat(*pConfigFile); err == nil {
		configMode = ConfigUsingArg
	}
	switch configMode {
	case ConfigUsingArg:
		configPath = *pConfigFile
	case ConfigUsingEnv:
		configPath = envConfigPath
	default:
		panic("cannot get config path")
	}
	fmt.Println("config mode:", configMode, " path:", configPath)
	return configPath
}

func main() {
	flag.Parse()
	var c []config.Config
	var users []*core.User
	var configPath = determineConfigPath()
	content, err := os.ReadFile(configPath)
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
