// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package rc

import (
	"bytes"
	"fmt"
	"github.com/ysicing/ergo/pkg/logger"
	"os"
	"os/exec"
)

// GetCmdStatus 获取状态
func GetCmdStatus(bin string) bool {
	cmd := exec.Command("which", bin)
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}

func CmdStatus(name string, arg ...string) bool {
	cmd := exec.Command(name, arg[:]...)
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}

func Cmd(name string, arg ...string) {
	cmd := exec.Command(name, arg[:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		logger.Exit(fmt.Sprintf("[os] run cmd %v %v, err: %v", name, arg, err))
	}
}

func CmdRes(name string, arg ...string) string {
	var b bytes.Buffer
	logger.Info(fmt.Sprintf("[os]exec cmd is : ", name, arg))
	cmd := exec.Command(name, arg[:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = &b
	cmd.Stdout = &b
	err := cmd.Run()
	if err != nil {
		logger.Exit(fmt.Sprintf("[os] run cmd %v %v err: %v", name, arg, err))
	}
	return b.String()
}

func Cmdv2(name string, arg []string) {
	cmd := exec.Command(name, arg...)
	logger.Info(fmt.Sprintf("[os]exec cmd is : ", name, cmd.Args))
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		logger.Exit(fmt.Sprintf("[os] run cmd %v %v, err: %v", name, arg, err))
	}
}

func CmdResv2(name string, arg []string) string {
	var b bytes.Buffer
	cmd := exec.Command(name, arg...)
	logger.Info(fmt.Sprintf("[os]exec cmd is : ", name, cmd.Args))
	cmd.Stdin = os.Stdin
	cmd.Stderr = &b
	cmd.Stdout = &b
	err := cmd.Run()
	if err != nil {
		logger.Exit(fmt.Sprintf("[os] run cmd %v %v err: %v", name, arg, err))
	}
	return b.String()
}
