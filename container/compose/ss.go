// MIT License
// Copyright (c) 2019 ysicing <i@ysicing.me>

package compose

import (
	"github.com/ysicing/ergo/pkg/check"
	"github.com/ysicing/ergo/pkg/logger"
)

type Ss struct {
	cfg ComposeConfig
}

func (s Ss) Check() {
	logger.Debug("check ss")
	if s.cfg.DeployLocal {
		t := check.CheckMeta{}
		if !(t.CheckBin("docker") && t.CheckBin("docker-compose")) {
			logger.Exit("docker or docker-compose not found")
		}
	} else {
		for _, ip := range s.cfg.Hosts {
			SSHConfig.Cmd(ip, "which docker")
			SSHConfig.Cmd(ip, "which docker-compose")
		}
	}
}

func (s Ss) Write() {
	logger.Debug("write docker-compose")
}

func (s Ss) Up() {
	logger.Debug("up docker-compose")
}

func (s Ss) Down() {
	logger.Debug("down docker-compose")
}

const sscompose = `

`
