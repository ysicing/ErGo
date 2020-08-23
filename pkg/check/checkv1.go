// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package check

import (
	"github.com/ysicing/ergo/pkg/rc"
	"runtime"
)

type CheckMeta struct {
	Type string
}

func (t CheckMeta) RunOnLinux() bool {
	if runtime.GOOS == "linux" {
		return true
	}
	return false
}

func (t CheckMeta) RunOnMac() bool {
	if runtime.GOOS == "darwin" {
		return true
	}
	return false
}

func (t CheckMeta) RunOs() bool {
	switch t.Type {
	case "linux":
		return t.RunOnLinux()
	case "macos":
		return t.RunOnMac()
	default:
		return t.RunOnMac()
	}
}

func (t CheckMeta) CheckBin(binpath string) bool {
	return rc.GetCmdStatus(binpath)
}
