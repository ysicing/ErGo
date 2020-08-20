// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package cloud

type CloudMeta struct {
	Key string
	Secret string
	Region string
}

type Cloud interface {
	OSSupload()
}

func NewCloud(t string, cm CloudMeta) Cloud  {
	switch t {
	case "aliyun":
		return &AliCloud{Cloud: cm}
	default:
		return &AliCloud{Cloud: cm}
	}
}