# Remotely Save 插件方案

## 简介
Remotely Save 是 Obsidian 官方推荐的同步插件，支持 WebDAV、S3、OneDrive、Dropbox、Google Drive 等多种云存储。

---

## 优点
- ✅ **无需额外软件**：直接在 Obsidian 内配置
- ✅ **多云支持**：WebDAV、S3、OneDrive、Dropbox、Google Drive
- ✅ **配置简单**：图形界面配置，无需命令行
- ✅ **免费**：插件免费，配合免费云存储使用
- ✅ **跨平台**：支持 Windows、Mac、Linux、移动端
- ✅ **加密支持**：可选端到端加密

## 缺点
- ❌ **依赖第三方**：需要云存储服务
- ❌ **同步延迟**：非实时同步，需手动或定时触发
- ❌ **冲突处理弱**：冲突时可能覆盖，需小心
- ❌ **免费额度有限**：坚果云每月 1GB 上传流量

---

## 安装方法

1. 打开 Obsidian
2. 设置 → 第三方插件 → 浏览
3. 搜索 "Remotely Save"
4. 安装并启用

---

## 配置步骤（以坚果云 WebDAV 为例）

### 1. 获取坚果云 WebDAV 信息
1. 登录坚果云网页版
2. 右上角用户名 → 账户信息
3. 安全选项 → 第三方应用管理
4. 添加应用，获取：
   - 服务器地址
   - 用户名
   - 密码（应用专用密码）

### 2. 配置 Remotely Save
1. 设置 → Remotely Save
2. 选择 **WebDAV**
3. 填入信息：
   - **Address**：`https://dav.jianguoyun.com/dav/`
   - **Username**：你的坚果云邮箱
   - **Password**：应用专用密码
4. 点击 **Check** 测试连接
5. 点击 **Save** 保存

### 3. 配置同步选项
- **Auto Sync**：开启自动同步
- **Sync on Startup**：启动时同步
- **Sync Interval**：同步间隔（分钟）
- **Password for Encryption**：可选加密密码

---

## 其他云存储配置

### OneDrive
1. 选择 **OneDrive**
2. 点击 **Auth** 进行授权
3. 登录微软账号并授权

### Dropbox
1. 选择 **Dropbox**
2. 点击 **Auth** 进行授权
3. 登录 Dropbox 并授权

### Google Drive
1. 选择 **Google Drive**
2. 点击 **Auth** 进行授权
3. 登录 Google 账号并授权

### S3（自建或云服务）
1. 选择 **S3**
2. 填入：
   - **Endpoint**：S3 服务地址
   - **Access Key ID**
   - **Secret Access Key**
   - **Bucket**：存储桶名称

---

## 日常使用

### 手动同步
- 左侧边栏点击同步图标
- 或 `Ctrl+P` → "Remotely Save: Start Sync"

### 自动同步
开启 Auto Sync 后，插件会按设定间隔自动同步

### 同步状态
- ✅ 同步成功
- ⚠️ 有冲突
- ❌ 同步失败

---

## 冲突处理

### 冲突类型
- **同名冲突**：两边创建了同名文件
- **修改冲突**：两边修改了同一文件

### 处理方法
1. 插件会保留两个版本
2. 手动比较并合并
3. 删除不需要的版本

### 预防冲突
- 避免在多设备同时编辑同一文件
- 同步前先 pull 最新版本
- 使用插件的加密功能保护隐私

---

## 常见问题

### Q：坚果云流量不够用？
A：考虑付费升级，或用 OneDrive（5GB 免费）。

### Q：同步失败怎么办？
A：检查网络、账号密码、WebDAV 地址是否正确。

### Q：会丢失数据吗？
A：正常情况不会，但建议定期备份 vault。

### Q：支持端到端加密吗？
A：支持，在设置中开启并设置密码。

---

## 适用场景
- 不想安装额外软件
- 已有坚果云/OneDrive/Dropbox 账号
- 设备简单（主要是电脑）
- 对同步实时性要求不高

---

*相关链接：*
- 插件 GitHub：https://github.com/remotely-save/remotely-save
- 坚果云 WebDAV：https://www.jianguoyun.com/
