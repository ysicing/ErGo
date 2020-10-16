// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package command

import (
	"fmt"
	"github.com/koding/vagrantutil"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/spf13/cobra"
	"github.com/ysicing/ergo/utils/common"
	"github.com/ysicing/ergo/vm"
	"github.com/ysicing/ext/logger"
	"github.com/ysicing/ext/utils/convert"
	"github.com/ysicing/ext/utils/exfile"
	"github.com/ysicing/ext/utils/exmisc"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	vmCpu      int64
	vmMem      int64
	vmInstance int64
	vmIP       string
	vmPath     string
	IPs        []string
	IsLocal    bool
)

// NewVMCommand() vm of ergo
func NewVMCommand() *cobra.Command {
	vm := &cobra.Command{
		Use:     "vm",
		Short:   "管理vm环境，新建vm, 初始化vm或者安装一些常用工具或者执行shell，推荐MacOS使用",
		Aliases: []string{"debian", "vbox"},
	}
	vm.AddCommand(NewVmNewCommand())
	vm.AddCommand(NewVmInitCommand())
	vm.AddCommand(NewVmInstallCommand())
	vm.AddCommand(NewVmExecCommand())
	return vm
}

// NewVmNewCommand 创建vm
func NewVmNewCommand() *cobra.Command {
	vmnew := &cobra.Command{
		Use:    "new",
		Short:  "创建vm环境",
		PreRun: vmnewprecheckfunc,
		Run:    vmnewfunc,
	}
	vmnew.PersistentFlags().Int64Var(&vmCpu, "cpu", 1, "实例cpu核数")
	vmnew.PersistentFlags().Int64Var(&vmMem, "mem", 512, "实例内存数")
	vmnew.PersistentFlags().Int64Var(&vmInstance, "num", 1, "实例副本数")
	vmnew.PersistentFlags().StringVar(&vmIP, "ip", "11.11.11.0/24", "实例起始IP,不建议修改")
	vmnew.PersistentFlags().StringVar(&vmPath, "path", "~/vm", "配置文件路径")
	return vmnew
}

func NewVmInitCommand() *cobra.Command {
	vminit := &cobra.Command{
		Use:   "init",
		Short: "初始化debian或debian系环境",
		Run:   vminitfunc,
	}
	vminit.PersistentFlags().StringVar(&SSHConfig.User, "user", "root", "用户")
	vminit.PersistentFlags().StringVar(&SSHConfig.Password, "pass", "", "密码")
	vminit.PersistentFlags().StringVar(&SSHConfig.PkFile, "pk", "", "私钥")
	vminit.PersistentFlags().StringSliceVar(&IPs, "ips", nil, "机器IP")
	return vminit
}

func NewVmInstallCommand() *cobra.Command {
	vmins := &cobra.Command{
		Use:    "install",
		Short:  "debian系安装常用软件",
		PreRun: vmpreinstallfunc,
		Run:    vminstallfunc,
	}
	vmins.PersistentFlags().StringVar(&SSHConfig.User, "user", "root", "用户")
	vmins.PersistentFlags().StringVar(&SSHConfig.Password, "pass", "", "密码")
	vmins.PersistentFlags().StringVar(&SSHConfig.PkFile, "pk", "", "私钥")
	vmins.PersistentFlags().StringSliceVar(&IPs, "ips", nil, "机器IP")
	vmins.PersistentFlags().BoolVar(&IsLocal, "local", false, "本地模式")
	return vmins
}

func NewVmExecCommand() *cobra.Command {
	vmins := &cobra.Command{
		Use:   "exec",
		Short: "执行shell",
		Run:   vmexecfunc,
	}
	vmins.PersistentFlags().StringVar(&SSHConfig.User, "user", "root", "用户")
	vmins.PersistentFlags().StringVar(&SSHConfig.Password, "pass", "", "密码")
	vmins.PersistentFlags().StringVar(&SSHConfig.PkFile, "pk", "", "私钥")
	vmins.PersistentFlags().StringSliceVar(&IPs, "ips", nil, "机器IP")
	return vmins
}

func vmnewprecheckfunc(cmd *cobra.Command, args []string) {
	logger.Slog.Debugf("%v", exmisc.SGreen("check system res"))
	// CPU
	cputotal, _ := cpu.Counts(true)
	if int64(cputotal) <= vmCpu*vmInstance {
		logger.Slog.Exit0(exmisc.SRed("CPU资源不够"), " 调整CPU大小或者副本数")
	}
	// mem
	memtotal, _ := mem.VirtualMemory()
	if memtotal.Total <= uint64(vmMem*vmInstance*1024*1024) {
		logger.Slog.Exit0(exmisc.SRed("内存资源不够"), "请调整内存大小或者副本数")
	}
	logger.Slog.Debugf("check system res: %v", exmisc.SGreen("pass"))
	logger.Slog.Debugf("%v", exmisc.SGreen("check system tools"))
	if !common.WhichCmd("vagrant") || !common.WhichCmd("VirtualBoxVM") {
		logger.Slog.Exit0(exmisc.SRed("vagrant"), "或", exmisc.SRed("VirtualBox"), "未安装，请先安装")
	}
	logger.Slog.Debugf("check system tools: %v", exmisc.SGreen("pass"))
}

func vmnewfunc(cmd *cobra.Command, args []string) {
	// step 01 检查文件是否存在
	vmPath = common.GetPath(vmPath)
	vgfile := common.GetPath(vmPath + "/Vagrantfile")

	logger.Slog.Debugf("cpu: %v, mem: %v, 实例: %v, ip段: %v, Vagrantfile: %v", vmCpu, vmMem, vmInstance, vmIP, vgfile)
	vagrant, _ := vagrantutil.NewVagrant(vmPath)
	if exfile.CheckFileExistsv2(vgfile) {
		var rewritefile string
		logger.Slog.Info("vagrantfile exist, Are you sure you want to rewrite vagrantfile ? [y/N]")
		fmt.Scanln(&rewritefile)
		if strings.ToLower(rewritefile) == "y" || strings.ToLower(rewritefile) == "yes" {
			logger.Slog.Info("开始执行覆盖")
			status, _ := vagrant.Status()
			if status.String() == "Running" {
				logger.Slog.Info("Destroy VM")
				output, err := vagrant.Destroy()
				if err != nil {
					logger.Slog.Exit0f("Destroy VM err: %v", err.Error())
				}
				for line := range output {
					fmt.Println(line.Line)
					time.Sleep(30 * time.Second)
				}
			}
			vagrantfile := vm.NewVM(vm.MetaData{
				Cpus:     convert.Int642Str(vmCpu),
				Memory:   convert.Int642Str(vmMem),
				Instance: convert.Int642Str(vmInstance),
				IP:       vmIP,
			}).Template()
			exfile.WriteFile(vgfile, vagrantfile)
		} else {
			logger.Slog.Info("跳过此流程")
		}
	} else {
		vagrantfile := vm.NewVM(vm.MetaData{
			Cpus:     convert.Int642Str(vmCpu),
			Memory:   convert.Int642Str(vmMem),
			Instance: convert.Int642Str(vmInstance),
			IP:       vmIP,
		}).Template()
		exfile.WriteFile(vgfile, vagrantfile)
	}

	// step 02 存在，启动
	logger.Slog.Debugf("%v", exmisc.SGreen("StartUP VM"))
	output, err := vagrant.Up()
	for line := range output {
		fmt.Println(line.Line)
	}
	if err != nil {
		// vagrant.Destroy()
		logger.Slog.Exit0("启动虚拟机失败，清理失败数据")
	}
	logger.Slog.Infof("default user/password: %v", exmisc.SGreen("root/vagrant"))
	logger.Slog.Infof("销毁方式: cd %v, vagrant destroy -f ", vmPath)
	if vmInstance == 1 {
		logger.Slog.Infof("销毁方式: cd %v, vagrant ssh", vmPath)
	}
}

func vminitfunc(cmd *cobra.Command, args []string) {
	logger.Slog.Debug(SSHConfig, IPs)
	var wg sync.WaitGroup
	for _, ip := range IPs {
		wg.Add(1)
		go vm.RunInit(SSHConfig, ip, &wg)
	}
	wg.Wait()
}

func vmpreinstallfunc(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("目前支持如下工具包安装: redis, mysql, etcd, adminer, docker(不支持local)")
		os.Exit(0)
	}
	skipkey := []string{"docker", "go", "golang", "w"}
	if !convert.StringArrayContains(skipkey, args[0]) {
		// check docker
		var num int
		if !IsLocal {
			for _, ip := range IPs {
				if !vm.CheckCmd(SSHConfig, ip, "docker") {
					logger.Slog.Error(ip, " 需要安装docker")
					num++
				}
			}
		} else {
			if err := common.RunCmd("which", "docker"); err != nil {
				logger.Slog.Error("本机", " 需要安装docker")
				num++
			}
		}
		if num != 0 {
			os.Exit(0)
		}
	}
}

func vminstallfunc(cmd *cobra.Command, args []string) {
	var wg sync.WaitGroup
	if IsLocal {
		wg.Add(1)
		go vm.InstallPackage(SSHConfig, "", args[0], &wg, true)
	} else {
		for _, ip := range IPs {
			wg.Add(1)
			go vm.InstallPackage(SSHConfig, ip, args[0], &wg, false)
		}
	}
	wg.Wait()
}

func vmexecfunc(cmd *cobra.Command, args []string) {
	var wg sync.WaitGroup
	for _, ip := range IPs {
		wg.Add(1)
		vm.ExecSh(SSHConfig, ip, &wg, args...)
	}
	wg.Wait()
}