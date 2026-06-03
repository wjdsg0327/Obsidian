# CentOS 7.9 安装 Docker

## 适用环境

- CentOS 7.9 64 位
- 建议使用 `root` 用户，或在命令前加 `sudo`

## 1. 卸载旧版本（可选）

如果系统之前安装过旧版 Docker，建议先卸载，避免冲突。新机器可跳过。

```bash
sudo yum remove docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-engine
```

## 2. 安装依赖

```bash
sudo yum install -y yum-utils
```

## 3. 配置 Docker CE 镜像仓库

使用阿里云 Docker CE 镜像源：

```bash
sudo yum-config-manager \
    --add-repo \
    http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
```

## 4. 安装 Docker Engine

```bash
sudo yum makecache fast
sudo yum install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

## 5. 启动并设置开机自启

```bash
sudo systemctl start docker
sudo systemctl enable docker
```

## 6. 验证安装

```bash
docker version
sudo docker run hello-world
```

## 7. 配置镜像加速器

创建或编辑 `/etc/docker/daemon.json`：

```bash
sudo mkdir -p /etc/docker
sudo vi /etc/docker/daemon.json
```

写入：

```json
{
  "registry-mirrors": [
    "https://registry.docker-cn.com",
    "http://hub-mirror.c.163.com",
    "https://docker.mirrors.ustc.edu.cn"
  ]
}
```

重新加载并重启 Docker：

```bash
sudo systemctl daemon-reload
sudo systemctl restart docker
```

## 备注

CentOS 7 已结束生命周期，不建议作为长期生产环境 Docker 宿主机。新环境优先考虑 Ubuntu 24.04 LTS、Debian 12 或 Rocky Linux / AlmaLinux。
