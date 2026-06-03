# Ubuntu 24.04 安装 Docker

## 适用环境

- Ubuntu 24.04 64 位
- 需要具有 `sudo` 权限的用户

## 1. 卸载旧版本（可选）

如果之前安装过 Docker，先卸载旧包避免冲突：

```bash
sudo apt-get remove docker docker-engine docker.io containerd runc
```

## 2. 更新软件源并安装依赖

```bash
sudo apt-get update
sudo apt-get install -y ca-certificates curl gnupg
```

## 3. 添加 Docker GPG 密钥（阿里云源）

```bash
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc
```

## 4. 添加 Docker APT 仓库

```bash
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://mirrors.aliyun.com/docker-ce/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
```

## 5. 安装 Docker 核心组件

```bash
sudo apt-get update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

## 6. 启动、开机自启并验证

```bash
sudo systemctl enable --now docker
sudo docker run hello-world
```

如果终端打印 `Hello from Docker!`，说明安装成功。

## 7. 免 sudo 运行 Docker（可选）

```bash
sudo usermod -aG docker $USER
```

执行后需要注销并重新登录，或重开终端，权限才会生效。

## 8. 配置镜像加速器

创建或编辑 `/etc/docker/daemon.json`：

```bash
sudo mkdir -p /etc/docker
sudo nano /etc/docker/daemon.json
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

重启 Docker：

```bash
sudo systemctl daemon-reload
sudo systemctl restart docker
```

## 备注

Ubuntu 24.04 LTS 是现代化 Docker 宿主机的好选择，默认内核较新，长期维护到 2029 年。
