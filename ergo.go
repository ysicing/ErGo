// MIT License
// Copyright (c) 2019 ysicing <i@ysicing.me>

package main

import (
	"github.com/ysicing/ergo/cmd"
	"github.com/ysicing/ergo/pkg/logger"
)

func init() {
	logger.InitLogger()
}

func main() {
	cmd.Execute()
}
