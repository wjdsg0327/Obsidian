# scrcpy 命令完整指南

> scrcpy 是一个通过 USB/TCP/IP 连接 Android 设备进行屏幕镜像和控制的命令行工具

---

## 目录

- [连接相关](#连接相关)
- [视频相关](#视频相关)
- [音频相关](#音频相关)
- [录制相关](#录制相关)
- [控制相关](#控制相关)
- [相机模式](#相机模式)
- [窗口相关](#窗口相关)
- [其他选项](#其他选项)
- [常用场景示例](#常用场景示例)

---

## 连接相关

adb tcpip 5555

| 命令                    | 说明               | 示例                                |
| --------------------- | ---------------- | --------------------------------- |
| `-s, --serial=SERIAL` | 指定设备序列号          | `scrcpy -s 0123456789abcdef`      |
| `-d, --select-usb`    | 选择 USB 连接的设备     | `scrcpy -d`                       |
| `-e, --select-tcpip`  | 选择 TCP/IP 连接的设备  | `scrcpy -e`                       |
| `--tcpip[=IP[:PORT]]` | 自动配置 TCP/IP 连接   | `scrcpy --tcpip=192.168.1.1:5555` |
| `--force-adb-forward` | 强制使用 adb forward | `scrcpy --force-adb-forward`      |

**TCP/IP 无线连接示例：**
```bash
# 通过 USB 自动配置无线连接
scrcpy --tcpip

# 直接连接已知 IP
scrcpy --tcpip=192.168.1.1
scrcpy --tcpip=192.168.1.1:5555
```

---

## 视频相关

| 命令 | 说明 | 示例 |
|------|------|------|
| `-b, --video-bit-rate=VALUE` | 视频比特率 (默认 8M) | `scrcpy -b 4M` |
| `--video-source=SOURCE` | 视频源: display/camera | `scrcpy --video-source=camera` |
| `-m, --max-size=VALUE` | 最大分辨率 (默认 0=无限制) | `scrcpy -m 1024` |
| `--max-fps=VALUE` | 最大帧率 | `scrcpy --max-fps=30` |
| `--video-buffer=MS` | 视频缓冲大小 (毫秒) | `scrcpy --video-buffer=200` |
| `--video-codec=CODEC` | 视频编码: h264/h265/av1 | `scrcpy --video-codec=h265` |
| `--video-encoder=NAME` | 指定视频编码器 | `scrcpy --video-encoder=OMX.qcom.video.encoder.avc` |
| `--no-video` | 禁用视频 | `scrcpy --no-video` |
| `--display=ID` | 指定显示器 ID | `scrcpy --display=1` |
| `--display-buffer=MS` | 显示缓冲 | `scrcpy --display-buffer=50` |
| `--v4l2-sink=DEVICE` | V4L2 输出设备 (Linux) | `scrcpy --v4l2-sink=/dev/video0` |

**分辨率示例：**
```bash
# 限制最大分辨率
scrcpy -m 1920

# 限制帧率
scrcpy --max-fps=60

# 高质量镜像
scrcpy -b 16M -m 2048 --max-fps=60
```

---

## 音频相关

| 命令 | 说明 | 示例 |
|------|------|------|
| `--no-audio` | 禁用音频 | `scrcpy --no-audio` |
| `--audio-source=SOURCE` | 音频源 | `scrcpy --audio-source=mic` |
| `--audio-codec=CODEC` | 音频编码: opus/aac/flac/raw | `scrcpy --audio-codec=aac` |
| `--audio-bit-rate=VALUE` | 音频比特率 (默认 128K) | `scrcpy --audio-bit-rate=64K` |
| `--audio-buffer=MS` | 音频缓冲 (默认 50ms) | `scrcpy --audio-buffer=100` |
| `--audio-encoder=NAME` | 指定音频编码器 | `scrcpy --audio-encoder=c2.android.opus.encoder` |
| `--audio-dup` | 音频双工 (设备+电脑同时播放) | `scrcpy --audio-dup` |
| `--require-audio` | 音频失败时退出 | `scrcpy --require-audio` |

**音频源选项：**
- `output` (默认) - 设备音频输出
- `playback` - 播放音频 (Android 13+)
- `mic` - 麦克风
- `mic-unprocessed` - 原始麦克风
- `mic-camcorder` - 摄像模式麦克风
- `mic-voice-recognition` - 语音识别模式
- `mic-voice-communication` - 通话模式
- `voice-call` - 通话音频

**音频示例：**
```bash
# 仅播放音频
scrcpy --no-video --no-control

# 使用麦克风
scrcpy --audio-source=mic

# 设备和电脑同时播放音频
scrcpy --audio-dup

# 高质量音频
scrcpy --audio-codec=flac --audio-bit-rate=256K
```

---

## 录制相关

| 命令 | 说明 | 示例 |
|------|------|------|
| `-r, --record=FILE` | 录制到文件 | `scrcpy -r video.mp4` |
| `--record-format=FORMAT` | 录制格式: mp4/mkv/opus/aac/flac | `scrcpy --record-format=mkv` |
| `--no-playback` | 录制时不播放 | `scrcpy -r video.mp4 --no-playback` |

**录制示例：**
```bash
# 录制屏幕
scrcpy -r video.mp4

# 录制时不显示窗口
scrcpy -r video.mp4 --no-playback

# 录制麦克风
scrcpy --audio-source=mic --no-video --no-playback -r audio.opus

# 高质量录制
scrcpy -b 16M --video-codec=h265 -r video.mkv
```

---

## 控制相关

| 命令 | 说明 | 示例 |
|------|------|------|
| `-n, --no-control` | 禁用控制 (只读镜像) | `scrcpy -n` |
| `--no-window` | 无窗口模式 | `scrcpy --no-window` |
| `--keyboard=MODE` | 键盘模式: disabled/sdk/aoa/uhid | `scrcpy --keyboard=uhid` |
| `--mouse=MODE` | 鼠标模式: disabled/sdk/aoa/uhid | `scrcpy --mouse=uhid` |
| `--power-off-on-close` | 关闭时关闭屏幕 | `scrcpy --power-off-on-close` |
| `--display-power-on` | 启动时开启屏幕 | `scrcpy --display-power-on` |
| `--clipboard-autosync` | 剪贴板自动同步 | `scrcpy --clipboard-autosync` |
| `--shortcut-mod=KEY` | 快捷键修饰键 | `scrcpy --shortcut-mod=alt` |

**控制示例：**
```bash
# 只读镜像 (不能操作)
scrcpy -n

# 无窗口模式 (仅录制或音频)
scrcpy --no-window

# 关闭时关闭设备屏幕
scrcpy --power-off-on-close
```

---

## 相机模式

| 命令 | 说明 | 示例 |
|------|------|------|
| `--video-source=camera` | 使用相机作为视频源 | `scrcpy --video-source=camera` |
| `--camera-id=ID` | 指定相机 ID | `scrcpy --camera-id=0` |
| `--camera-facing=FACING` | 相机方向: front/back | `scrcpy --camera-facing=front` |
| `--camera-size=WIDTHxHEIGHT` | 相机分辨率 | `scrcpy --camera-size=1920x1080` |
| `--camera-fps=FPS` | 相机帧率 | `scrcpy --camera-fps=30` |
| `--list-cameras` | 列出可用相机 | `scrcpy --list-cameras` |
| `--list-camera-sizes` | 列出相机支持的分辨率 | `scrcpy --list-camera-sizes` |

**相机示例：**
```bash
# 列出相机
scrcpy --list-cameras

# 使用前置摄像头
scrcpy --video-source=camera --camera-facing=front

# 指定分辨率和帧率
scrcpy --video-source=camera --camera-size=1920x1080 --camera-fps=30
```

---

## 窗口相关

| 命令 | 说明 | 示例 |
|------|------|------|
| `-f, --fullscreen` | 全屏模式 | `scrcpy -f` |
| `--always-on-top` | 窗口置顶 | `scrcpy --always-on-top` |
| `--no-window-title` | 隐藏窗口标题 | `scrcpy --no-window-title` |
| `--window-title=TEXT` | 自定义窗口标题 | `scrcpy --window-title=MyPhone` |
| `--window-x=X` | 窗口 X 位置 | `scrcpy --window-x=100` |
| `--window-y=Y` | 窗口 Y 位置 | `scrcpy --window-y=100` |
| `--window-width=WIDTH` | 窗口宽度 | `scrcpy --window-width=800` |
| `--window-height=HEIGHT` | 窗口高度 | `scrcpy --window-height=600` |
| `--window-borderless` | 无边框窗口 | `scrcpy --window-borderless` |
| `--rotation=VALUE` | 旋转: 0/90/180/270 | `scrcpy --rotation=90` |
| `--orientation=VALUE` | 方向: vertical/horizontal | `scrcpy --orientation=horizontal` |
| `--aspect-ratio=RATIO` | 宽高比 | `scrcpy --aspect-ratio=16:9` |
| `--hidpi` | HiDPI 模式 | `scrcpy --hidpi` |

**窗口示例：**
```bash
# 全屏
scrcpy -f

# 置顶小窗口
scrcpy --always-on-top -m 1024

# 指定位置和大小
scrcpy --window-x=0 --window-y=0 --window-width=400 --window-height=800

# 横屏模式
scrcpy --rotation=90
```

---

## 其他选项

| 命令 | 说明 | 示例 |
|------|------|------|
| `-v, --version` | 显示版本 | `scrcpy -v` |
| `-h, --help` | 显示帮助 | `scrcpy -h` |
| `--list-encoders` | 列出编码器 | `scrcpy --list-encoders` |
| `--list-displays` | 列出显示器 | `scrcpy --list-displays` |
| `--verbosity=VALUE` | 日志级别: verbose/debug/info/warn/error | `scrcpy --verbosity=debug` |
| `--no-cleanup` | 退出时不清理 | `scrcpy --no-cleanup` |
| `--print-fps` | 打印帧率 | `scrcpy --print-fps` |
| `--no-key-repeat` | 禁用按键重复 | `scrcpy --no-key-repeat` |
| `--forward-all-clicks` | 转发所有点击 | `scrcpy --forward-all-clicks` |
| `--legacy-paste` | 传统粘贴方式 | `scrcpy --legacy-paste` |

---

## 常用场景示例

### 基础使用
```bash
# 最简单的连接
scrcpy

# 无线连接
scrcpy --tcpip=192.168.1.1
```

### 低配置设备
```bash
# 降低分辨率和比特率
scrcpy -m 1024 -b 2M --max-fps=30
```

### 高质量镜像
```bash
# 高分辨率 + 高比特率
scrcpy -m 2048 -b 16M --video-codec=h265
```

### 录制教程
```bash
# 高质量录制
scrcpy -r tutorial.mp4 -b 8M --video-codec=h265
```

### 后台录制
```bash
# 无窗口录制
scrcpy -r video.mp4 --no-window
```

### 仅音频播放
```bash
# 把手机当音箱
scrcpy --no-video
```

### 只读监控
```bash
# 仅查看不能操作
scrcpy -n --always-on-top
```

### 开发调试
```bash
# 显示帧率和调试信息
scrcpy --print-fps --verbosity=debug
```

### 远程会议共享
```bash
# 高帧率 + 音频双工
scrcpy --audio-dup --max-fps=60 -m 1280
```

---

## 快捷键参考

| 快捷键 | 功能 |
|--------|------|
| `Ctrl+h` | 主屏幕 |
| `Ctrl+b` | 返回 |
| `Ctrl+s` | 多任务 |
| `Ctrl+p` | 电源键 |
| `Ctrl+o` | 关闭屏幕 |
| `Ctrl+r` | 旋转屏幕 |
| `Ctrl+n` | 展开通知栏 |
| `Ctrl+Shift+n` | 折叠通知栏 |
| `Ctrl+c` | 复制到剪贴板 |
| `Ctrl+v` | 粘贴 |
| `Ctrl+x` | 剪切 |
| `Ctrl++/-` | 调整音量 |
| `Ctrl+m` | 静音 |
| `Alt+h` | 显示帮助 |
| `Alt+f` | 全屏切换 |
| `Alt+p` | 录屏暂停/继续 |

---

> 文档整理时间: 2026-03-14
> 基于 scrcpy v2.x 版本
