// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package config

import (
	"fmt"
	"github.com/ysicing/ergo/pkg/logger"
	"github.com/ysicing/ergo/utils"
	"github.com/ysicing/go-utils/exfile"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Global GlobalConfig `yaml:"global"`
	Drone  DroneConfig  `yaml:"drone"`
	Debian DebianConfig `yaml:"debian"`
}

type GlobalConfig struct {
	Debug bool `yaml:"debug"`
}

type DebianConfig struct {
	InitImage       Image    `yaml:"init"`
	VagrantfilePath FilePath `yaml:"vagrantfile"`
}

type DroneConfig struct {
	Host  string `yaml:"host"`
	Token string `yaml:"token"`
}

type Image struct {
	Imagev1 string `yaml:"image"`
}

type FilePath struct {
	FilePathv1 string `yaml:"path"`
}

func exampleConfig() Config {
	return Config{
		Global: GlobalConfig{
			Debug: true,
		},
		Drone: DroneConfig{
			Host:  "http://drone.company.com",
			Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
		},
		Debian: DebianConfig{
			InitImage: Image{
				Imagev1: "registry.cn-beijing.aliyuncs.com/k7scn/ansible:1.0",
			},
			VagrantfilePath: FilePath{
				FilePathv1: fmt.Sprintf("%v/%v/Vagrantfile", utils.GetHome(), "vm"),
			},
		},
	}
}

func WriteDefaultCfg(path string) {
	cfg, _ := yaml.Marshal(exampleConfig())
	logger.Debug(fmt.Sprintf("[ergo] write default config(%v): \n%v", path, string(cfg)))
	exfile.WriteFile(path, string(cfg))
}
