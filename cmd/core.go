// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package cmd

import "github.com/spf13/cobra"

var core = &cobra.Command{
	Use:   "core",
	Short: "核心模块",
}

func init() {
	rootCmd.AddCommand(core)
	core.AddCommand(debianCmd) // debian
	core.AddCommand(devops)    // op
	core.AddCommand(osCmd)     // os
	core.AddCommand(netCmd)
	core.AddCommand(systemdCmd)
	core.AddCommand(infoCmd)
	core.AddCommand(webCmd)
}
