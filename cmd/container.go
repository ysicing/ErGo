// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package cmd

import "github.com/spf13/cobra"

var container = &cobra.Command{
	Use:     "container",
	Aliases: []string{"c"},
	Short:   "容器化",
}

func init() {
	container.AddCommand(installCmd, helmbase, pluginsCmd, composeCmd)
}
