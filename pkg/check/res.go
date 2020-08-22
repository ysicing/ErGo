// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package check

import (
	"github.com/ysicing/ergo/pkg/logger"
	"os"
)

// CheckResError res err check
func CheckResError(err error) {
	if err != nil {
		logger.Error(err.Error())
		os.Exit(0)
	}
}
