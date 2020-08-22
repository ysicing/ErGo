// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ysicing/ergo/ops/vm"
)

var op = &cobra.Command{
	Use:     "ops",
	Aliases: []string{"op", "o", "devops"},
	Short:   "传统运维命令工具",
}

func init() {
	rootCmd.AddCommand(op)
	op.AddCommand(vm.Debian) // debian
	op.AddCommand(devops)    // op
	op.AddCommand(osCmd)     // os
	op.AddCommand(netCmd)
	op.AddCommand(systemdCmd)
	op.AddCommand(infoCmd)
	op.AddCommand(webCmd)
}
