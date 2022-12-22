package main

import (
	"flag"
	"fmt"
	"os"
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
