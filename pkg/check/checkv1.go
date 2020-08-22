// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package check

import "runtime"

type CheckMeta struct {
	Type string
}

func (t CheckMeta) RunOnLinux() bool {
	if runtime.GOOS == "linux" {
		return true
	}
	return false
}
