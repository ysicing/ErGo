## ergo

> 一个使用 Go 编写运维工具,尽量减少重复工作，同时降低维护脚本的成本

### 镜像使用

```bash
ysicing/ergo
```

### 二进制安装

可直接从 [release](https://github.com/ysicing/ergo/releases) 页下载预编译的二进制文件

### Mac OS安装

```bash
brew tap ysicing/tap
brew install ergo
```

## 命令支持 TODO

分类: 传统运维cli, 云原生运维cli, 云服务商cli

#### 传统运维cli

- [ ] debian系

```
# 新建debian vm
ergo vm new --mem 4096 --cpu 2 --num 2 --ip 10.0.0.0/24 # 内存，CPU，副本数, 默认IP端，建议使用默认的
# 初始化debian vm
ergo vm init --pass vagrant --ips 10.0.0.11 --ips 10.0.0.12
# 安装常用工具
ergo vm install --pass vagrant --ips 10.0.0.11 --ips 10.0.0.12 docker
# 执行shell
ergo vm exec --pass vagrant --ips 10.0.0.11 --ips 10.0.0.12 docker ps
```

#### 云原生运维cli

- [ ] 安装k8s 1.16.14

#### 云服务商cli

- [ ] 阿里云镜像仓库, ucloud镜像仓库