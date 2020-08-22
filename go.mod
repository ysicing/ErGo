module github.com/ysicing/ergo

go 1.14

require (
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/aliyun/aliyun-oss-go-sdk v2.1.4+incompatible
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/cuisongliu/sshcmd v1.5.2
	github.com/drone/drone-go v1.3.1
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/hashicorp/go-version v1.2.1 // indirect
	github.com/koding/logging v0.0.0-20160720134017-8b5a689ed69b // indirect
	github.com/koding/vagrantutil v0.0.0-20180710063911-70827343f116
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/mapstructure v1.3.3 // indirect
	github.com/pelletier/go-toml v1.8.0 // indirect
	github.com/prometheus/client_golang v1.7.1
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/shirou/gopsutil v2.20.7+incompatible
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/afero v1.3.4 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.7.1
	github.com/wangle201210/githubapi v0.0.0-20200804144924-cde7bbdc36ab
	github.com/wonderivan/logger v1.0.0
	github.com/ysicing/ginmid v0.2.1
	github.com/ysicing/go-utils v0.3.1
	go.uber.org/zap v1.15.0
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sys v0.0.0-20200821140526-fda516888d29 // indirect
	golang.org/x/text v0.3.3 // indirect
	gopkg.in/ini.v1 v1.60.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.3.0
	k8s.io/klog v1.0.0
)

replace github.com/cuisongliu/sshcmd v1.5.2 => github.com/kunnos/sshcmd v1.6.0
