package main

import (
	"encoding/json"
	"flag"
	"log"
	"newJwCourseHelper/internal/config"
	"newJwCourseHelper/internal/core"
	"os"
	"os/signal"
	"syscall"
)

var quit = make(chan os.Signal)
var users []*core.User
var c []config.Config

func main() {
	flag.Parse()

	var configPath = determineConfigPath()
	content, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	if *interactiveMode {
		ChooseCourseLogger = log.New(ChooseCourseLoggerBuffer, "[DebugMessage] ", log.LstdFlags)
	} else {
		ChooseCourseLogger = log.New(os.Stdout, "", log.LstdFlags)
	}

	if err = json.Unmarshal(content, &c); err != nil {
		panic(err)
	}

	//c中的每一个元素都是一个config.Config结构体，保存了用户的配置、任务等

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go ChooseCourse()
	if *interactiveMode {
		Execute()
	}
	<-quit
}
