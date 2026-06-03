# 云盘同步方案

## 简介
将 Obsidian vault 放在云盘同步文件夹中，利用云盘客户端自动同步。支持 OneDrive、Dropbox、坚果云、Google Drive 等。

---

## 优点
- ✅ **最简单**：无需额外配置，放文件夹里就行
- ✅ **自动同步**：云盘客户端自动处理
- ✅ **免费**：大部分云盘有免费额度
- ✅ **跨平台**：所有云盘都支持多平台
- ✅ **稳定可靠**：云盘厂商维护，稳定性好

## 缺点
- ❌ **冲突风险**：多设备同时编辑可能冲突
- ❌ **隐私性一般**：文件存储在第三方服务器
- ❌ **同步延迟**：不是实时同步，有几秒到几分钟延迟
- ❌ **空间限制**：免费额度有限
- ❌ **.obsidian 同步**：配置文件也会同步，可能有问题

---

## 各云盘对比

| 云盘 | 免费空间 | 国内可用 | WebDAV | 推荐度 |
|------|----------|----------|--------|--------|
| OneDrive | 5GB | ✅ | ❌ | ⭐⭐⭐⭐ |
| 坚果云 | 1GB/月上传 | ✅ | ✅ | ⭐⭐⭐⭐⭐ |
| Dropbox | 2GB | ⚠️ 需翻墙 | ❌ | ⭐⭐⭐ |
| Google Drive | 15GB | ⚠️ 需翻墙 | ❌ | ⭐⭐⭐ |
| iCloud | 5GB | ✅ | ❌ | ⭐⭐⭐ |
| 百度网盘 | 2TB | ✅ | ❌ | ⭐⭐ |

---

## 配置步骤（以坚果云为例）

### 1. 安装坚果云客户端
- 官网：https://www.jianguoyun.com/
- 下载对应系统版本并安装

### 2. 创建同步文件夹
```bash
# 在坚果云同步目录创建 vault 文件夹
mkdir -p ~/Nutstore/ObsidianVault
```

### 3. 移动 vault 到同步目录
```bash
# 移动现有 vault
mv /path/to/your/vault ~/Nutstore/ObsidianVault/

# 或新建 vault
obsidian ~/Nutstore/ObsidianVault/
```

### 4. 配置忽略规则（可选）
在坚果云设置中添加忽略规则：
```
.obsidian/workspace.json
.obsidian/workspace-mobile.json
.trash/
```

---

## 配置步骤（以 OneDrive 为例）

### 1. 确保 OneDrive 已登录
- Windows 10/11 自带 OneDrive
- 登录微软账号

### 2. 创建同步文件夹
```bash
# OneDrive 同步目录
mkdir -p ~/OneDrive/ObsidianVault
```

### 3. 移动 vault
```bash
mv /path/to/your/vault ~/OneDrive/ObsidianVault/
```

### 4. 等待同步
OneDrive 会自动同步文件，状态栏显示同步进度

---

## 配置步骤（以 Dropbox 为例）

### 1. 安装 Dropbox
- 官网：https://www.dropbox.com/
- 下载安装并登录

### 2. 创建同步文件夹
```bash
mkdir -p ~/Dropbox/ObsidianVault
```

### 3. 移动 vault
```bash
mv /path/to/your/vault ~/Dropbox/ObsidianVault/
```

---

## 冲突处理

### 冲突类型
- **文件冲突**：两边修改同一文件
- **文件夹冲突**：两边创建同名文件/文件夹

### Dropbox 冲突处理
Dropbox 会自动创建冲突副本：
```
file (conflicted copy 2024-01-01).md
```

### OneDrive 冲突处理
OneDrive 会保留两个版本，需要手动选择

### 坚果云冲突处理
坚果云会提示冲突，需要手动解决

### 预防冲突
1. **避免同时编辑**：确保同一时间只在一台设备编辑
2. **同步后再编辑**：切换设备前等待同步完成
3. **使用文件锁**：编辑时标记文件（如在文件名加 `[编辑中]`）

---

## 优化建议

### 1. 忽略工作区文件
在云盘设置中添加忽略规则：
```
.obsidian/workspace.json
.obsidian/workspace-mobile.json
```

### 2. 大文件处理
- 图片压缩后再放入 vault
- 大附件单独存储，用链接引用
- 考虑用图床服务

### 3. 定期备份
- 云盘不是备份，只是同步
- 定期导出 vault 备份
- 考虑用 Git 做版本控制

---

## 常见问题

### Q：同步慢怎么办？
A：检查网络，减少 vault 大小，关闭不必要的同步。

### Q：会丢失数据吗？
A：正常情况不会，但建议定期备份。

### Q：.obsidian 配置会同步吗？
A：会，这可能导致不同设备的插件配置冲突。

### Q：手机能用吗？
A：可以，安装对应云盘 App，用 Obsidian 打开同步文件夹。

---

## 适用场景
- 不想折腾，最省事
- 已有云盘账号
- 设备少（2-3台）
- 对隐私要求不高

---

*相关链接：*
- 坚果云：https://www.jianguoyun.com/
- OneDrive：https://onedrive.live.com/
- Dropbox：https://www.dropbox.com/
