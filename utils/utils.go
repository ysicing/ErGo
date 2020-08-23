// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package utils

import (
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"os"
	"runtime"
)

// DebugMode debug mode
func DebugMode() bool {
	return viper.GetBool("global.debug")
}

// RunLinux is linux
func RunLinux() bool {
	if runtime.GOOS == "linux" {
		return true
	}
	return false
}

// GetHome 获取home目录
func GetHome() string {
	if home, err := homedir.Dir(); err == nil {
		return home
	}
	os.Exit(0)
	return ""
}
