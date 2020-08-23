// MIT License
// Copyright (c) 2019 ysicing <i@ysicing.me>

package vm

import (
	"fmt"
	"github.com/koding/vagrantutil"
	"github.com/spf13/viper"
	"github.com/ysicing/ergo/pkg/logger"
	"github.com/ysicing/ergo/utils"
	"io/ioutil"
	"os"
)

const (
	DEBIAN             = "debian"
	CENTOS             = "centos"
	DefaultCpus        = "2"
	DefaultMemory      = "4096"
	DefaultInstance    = "1"
	DefaultVmDir       = "/vm"
	DefaultVagrantfile = "/Vagrantfile"
)

var (
	Cpus     string
	Memory   string
	Instance string
	Name     string
	Path     string
)

type MetaData struct {
	Cpus     string
	Memory   string
	Instance string
	Name     string
}

type Os interface {
	Osmode() string
	Template() string
}

type Vm struct{}

func NewVM(data MetaData) Os {
	return &Debian{metadata: data}
}

func VmInit() {
	i := Vm{}
	// 检查资源是否满足
	i.CheckSystem()
	// 写vagrantfile
	i.WriteVagrant()
	// 启动
	i.VmStartUP()
}

func (v *Vm) CheckSystem() {
	logger.Info("[vm] check system")
	if !utils.SysCmpOk(Cpus, Instance, utils.GetTotalCpuNum()) {
		logger.Exit("CPU资源不够，请调整CPU大小或者副本数")
	}
	if !utils.SysCmpOk(Memory, Instance, utils.GetTotalMemNum()) {
		logger.Exit("内存资源不够，请调整内存大小或者副本数")
	}
	logger.Info("[vm] check system done. It looks good")
}

func (v *Vm) WriteVagrant() {
	vagranfile := NewVM(MetaData{
		Name:     Name,
		Cpus:     Cpus,
		Memory:   Memory,
		Instance: Instance,
	}).Osmode()
	if Path == "" {
		Path = viper.GetString("debian.vagrantfile.path")
	}

	err := os.MkdirAll(Path, os.ModePerm)
	if err != nil {
		logger.Exit(fmt.Sprintf("create vagrantfile dir failed: %s", err))
	}
	cfgpath := fmt.Sprintf("%s%s", Path, DefaultVagrantfile)
	// check 是否存在
	if utils.FileExists(cfgpath) {
		logger.Exit(fmt.Sprintf("已存在相关配置文件: %v", Path))
	}
	if err = ioutil.WriteFile(cfgpath, []byte(vagranfile), 0644); err != nil {
		logger.Exit(fmt.Sprintf("write vagrantfile failed: %s", err))
	}
}

func (v *Vm) VmStartUP() {
	// TODO
	logger.Info("start up vm")
	vagrant, _ := vagrantutil.NewVagrant(Path)
	output, err := vagrant.Up()
	for line := range output {
		fmt.Println(line.Line)
	}
	if err != nil {
		vagrant.Destroy()
		utils.ErgoExit("启动虚拟机失败，清理失败数据")
	}
	if utils.String2Int(Instance) == 1 {
		logger.Info(fmt.Sprintf("ip: 11.11.11.%v, root/vagrant", 111))
	} else {
		logger.Info(fmt.Sprintf("ip: 11.11.11.%v-11.11.11.%v, root/vagrant", 111, 110+utils.String2Int(Instance)))
	}

	logger.Info(fmt.Sprintf("销毁方式: cd %s, vagrant destroy -f", Path))
}
