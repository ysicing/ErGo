// MIT License
// Copyright (c) 2019 ysicing <i@ysicing.me>

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ysicing/ergo/config"
	"github.com/ysicing/ergo/pkg/check"
	"github.com/ysicing/ergo/utils"
	"github.com/ysicing/go-utils/exfile"
)

var cfgFile string
var debugMode bool

var rootCmd = &cobra.Command{
	Use:   "ergo",
	Short: "An awesome tool",
}

// Execute execute
func Execute() {
	check.CheckResError(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.config/ergo/config.yaml)")
	rootCmd.PersistentFlags().BoolVar(&debugMode, "debug", false, "debug mode (default: false)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.DisableSuggestions = false
}

func initConfig() {
	if cfgFile == "" {
		cfgFile = fmt.Sprintf("%v/%v/%v/%v", utils.GetHome(), ".config", "ergo", "config.yaml")
	}
	viper.Set("global.debug", debugMode)
	if !exfile.CheckFileExistsv2(cfgFile) {
		config.WriteDefaultCfg(cfgFile)
	}
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv()
	check.CheckResError(viper.ReadInConfig())
}
