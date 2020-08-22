// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package vm

import (
	"github.com/spf13/cobra"
	"github.com/ysicing/ergo/core/vm"
	"github.com/ysicing/ergo/utils"
)

var Debian = &cobra.Command{
	Use:   "debian",
	Short: "debian环境",
}

var newdeb = &cobra.Command{
	Use:   "new",
	Short: "创建debian virtualbox虚拟机",
	PreRun: func(cmd *cobra.Command, args []string) {
		// 检查虚拟机

	},
	Run: func(cmd *cobra.Command, args []string) {
		utils.WarningOs()
		vm.VmInit()
	},
}

var initdeb = &cobra.Command{
	Use:   "init",
	Short: "初始化debian",
	Run: func(cmd *cobra.Command, args []string) {
		utils.WarningDocker()
		vm.InitDebian()
	},
}

var reinstallDebian = &cobra.Command{
	Use:   "reinstall",
	Short: "重装debian",
	Run: func(cmd *cobra.Command, args []string) {
		vm.ReinstallDebian()
	},
}

func init() {
	Debian.AddCommand(newdeb, initdeb, reinstallDebian)
	newdeb.PersistentFlags().StringVar(&vm.Name, "vmname", "", "虚拟机名")
	newdeb.PersistentFlags().StringVar(&vm.Cpus, "vmcpus", "2", "虚拟机CPU数")
	newdeb.PersistentFlags().StringVar(&vm.Memory, "vmmem", "4096", "虚拟机Mem MB数")
	newdeb.PersistentFlags().StringVar(&vm.Instance, "vmnum", "1", "虚拟机副本数")
	newdeb.PersistentFlags().StringVar(&vm.Path, "path", "", "Vagrantfile所在目录, $HOME/vm")

	initdeb.PersistentFlags().StringSliceVar(&vm.Hosts, "ip", []string{"11.11.11.111"}, "ssh ip")
	initdeb.PersistentFlags().StringVar(&vm.Port, "port", "22", "ssh端口")
	initdeb.PersistentFlags().StringVar(&vm.User, "user", "root", "管理员用户")
	initdeb.PersistentFlags().StringVar(&vm.Pass, "pass", "vagrant", "管理员密码")
	initdeb.PersistentFlags().BoolVar(&vm.DockerInstall, "docker", false, "是否安装docker")

	reinstallDebian.PersistentFlags().StringSliceVar(&vm.Hosts, "ip", []string{"11.11.11.111"}, "ssh ip")
	reinstallDebian.PersistentFlags().StringVar(&vm.SSHConfig.User, "user", "", "管理员用户")
	reinstallDebian.PersistentFlags().StringVar(&vm.SSHConfig.PkFile, "pk", "", "管理员私钥")
	reinstallDebian.PersistentFlags().StringVar(&vm.SSHConfig.Password, "pass", "", "管理员密码")
	reinstallDebian.PersistentFlags().BoolVar(&vm.Local, "local", false, "本地安装")
	reinstallDebian.PersistentFlags().StringVar(&vm.ReInstallPass, "repass", "vagrant", "默认重装密码")
	reinstallDebian.PersistentFlags().StringVar(&vm.ReInstallDisk, "redisk", "", "自定义硬盘,如/dev/sdb")
}
