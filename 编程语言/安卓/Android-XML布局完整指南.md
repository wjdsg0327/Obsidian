# Android Studio Kotlin XML 布局完整指南

> 整理时间：2026-03-12  
> 适用于 Android 开发初学者和进阶者

---

## 📚 目录

1. [XML 布局基础](#xml-布局基础)
2. [布局容器详解](#布局容器详解)
   - [LinearLayout 线性布局](#1-linearlayout---线性布局)
   - [ConstraintLayout 约束布局](#2-constraintlayout---约束布局推荐)
   - [FrameLayout 帧布局](#3-framelayout---帧布局)
   - [RelativeLayout 相对布局](#4-relativelayout---相对布局)
   - [GridLayout 网格布局](#5-gridlayout---网格布局)
   - [TableLayout 表格布局](#6-tablelayout---表格布局)
3. [常用布局属性](#常用布局属性)
4. [布局优化技巧](#布局优化技巧)
5. [实战案例](#实战案例)

---

## XML 布局基础

### 什么是 XML 布局？

XML 布局是 Android 用于定义用户界面的声明式方式。每个 XML 文件对应一个视图层级结构。

### 布局文件位置

```
app/src/main/res/layout/
├── activity_main.xml      # Activity 布局
├── fragment_home.xml      # Fragment 布局
├── item_user.xml          # RecyclerView 列表项
├── dialog_custom.xml      # 自定义对话框
└── layout_error.xml       # 其他布局
```

### 布局文件基本结构

```xml
<?xml version="1.0" encoding="utf-8"?>
<!-- 根元素：布局容器 -->
<LinearLayout xmlns:android="http://schemas.android.com/apk/res/android"
    xmlns:app="http://schemas.android.com/apk/res-auto"
    xmlns:tools="http://schemas.android.com/tools"
    android:layout_width="match_parent"
    android:layout_height="match_parent"
    android:orientation="vertical"
    tools:context=".MainActivity">

    <!-- 子控件 -->
    <TextView
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="Hello World!" />

</LinearLayout>
```

### 命名空间说明

| 命名空间 | 作用 | 示例 |
|---------|------|------|
| `xmlns:android` | Android 标准属性 | `android:layout_width` |
| `xmlns:app` | 自定义属性/第三方库属性 | `app:layout_constraintLeft_toLeftOf` |
| `xmlns:tools` | 预览工具属性（不打包进 APK） | `tools:text="预览文字"` |

---

## 布局容器详解

---

## 1. LinearLayout - 线性布局

### 说明

子控件按**水平**或**垂直**方向依次排列，是最简单常用的布局之一。

### 核心属性

| 属性 | 值 | 说明 |
|------|-----|------|
| `android:orientation` | `horizontal` / `vertical` | 排列方向 |
| `android:gravity` | 多种对齐方式 | 内部子控件的对齐方式 |
| `android:layout_gravity` | 多种对齐方式 | 子控件自身在父容器中的对齐 |
| `android:layout_weight` | 数值 | 权重，分配剩余空间 |
| `android:weightSum` | 数值 | 权重总和（父容器） |

### 基础示例

```xml
<?xml version="1.0" encoding="utf-8"?>
<!-- 垂直线性布局 -->
<LinearLayout xmlns:android="http://schemas.android.com/apk/res/android"
    android:layout_width="match_parent"
    android:layout_height="match_parent"
    android:orientation="vertical"
    android:gravity="center"
    android:padding="16dp">

    <TextView
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="标题"
        android:textSize="24sp" />

    <TextView
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="副标题"
        android:layout_marginTop="8dp" />

    <Button
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="确定"
        android:layout_marginTop="16dp" />

</LinearLayout>
```

### 权重（weight）使用

```xml
<?xml version="1.0" encoding="utf-8"?>
<LinearLayout xmlns:android="http://schemas.android.com/apk/res/android"
    android:layout_width="match_parent"
    android:layout_height="wrap_content"
    android:orientation="horizontal">

    <!-- 权重分配：1:2:1 -->
    <Button
        android:layout_width="0dp"
        android:layout_height="wrap_content"
        android:layout_weight="1"
        android:text="按钮1" />

    <Button
        android:layout_width="0dp"
        android:layout_height="wrap_content"
        android:layout_weight="2"
        android:text="按钮2" />

    <Button
        android:layout_width="0dp"
        android:layout_height="wrap_content"
        android:layout_weight="1"
        android:text="按钮3" />

</LinearLayout>
```

### gravity 对齐方式

```xml
<!-- gravity 可组合使用 -->
<LinearLayout
    android:layout_width="match_parent"
    android:layout_height="200dp"
    android:orientation="vertical"
    android:gravity="center_horizontal|bottom"
    android:background="#E0E0E0">

    <TextView
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="水平居中 + 底部对齐" />

</LinearLayout>
```

**gravity 常用值**：
- `center`：居中
- `center_horizontal`：水平居中
- `center_vertical`：垂直居中
- `left` / `right` / `top` / `bottom`
- `start` / `end`（支持 RTL 布局）
- 可用 `|` 组合：`center_horizontal|bottom`

### layout_gravity 使用

```xml
<LinearLayout
    android:layout_width="match_parent"
    android:layout_height="200dp"
    android:orientation="horizontal"
    android:background="#E0E0E0">

    <!-- 子控件使用 layout_gravity 自身对齐 -->
    <TextView
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="顶部"
        android:layout_gravity="top"
        android:background="#FFCDD2" />

    <TextView
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="居中"
        android:layout_gravity="center_vertical"
        android:background="#C8E6C9" />

    <TextView
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="底部"
        android:layout_gravity="bottom"
        android:background="#BBDEFB" />

</LinearLayout>
```

### 实战：登录表单

```xml
<?xml version="1.0" encoding="utf-8"?>
<LinearLayout xmlns:android="http://schemas.android.com/apk/res/android"
    android:layout_width="match_parent"
    android:layout_height="match_parent"
    android:orientation="vertical"
    android:padding="24dp"
    android:gravity="center">

    <!-- Logo -->
    <ImageView
        android:layout_width="80dp"
        android:layout_height="80dp"
        android:src="@drawable/ic_logo"
        android:layout_marginBottom="32dp" />

    <!-- 用户名 -->
    <EditText
        android:id="@+id/etUsername"
        android:layout_width="match_parent"
        android:layout_height="wrap_content"
        android:hint="用户名"
        android:inputType="text"
        android:layout_marginBottom="16dp" />

    <!-- 密码 -->
    <EditText
        android:id="@+id/etPassword"
        android:layout_width="match_parent"
        android:layout_height="wrap_content"
        android:hint="密码"
        android:inputType="textPassword"
        android:layout_marginBottom="24dp" />

    <!-- 登录按钮 -->
    <Button
        android:id="@+id/btnLogin"
        android:layout_width="match_parent"
        android:layout_height="wrap_content"
        android:text="登录"
        android:layout_marginBottom="16dp" />

    <!-- 底部链接 -->
    <LinearLayout
        android:layout_width="match_parent"
        android:layout_height="wrap_content"
        android:orientation="horizontal"
        android:gravity="center">

        <TextView
            android:layout_width="wrap_content"
            android:layout_height="wrap_content"
            android:text="还没有账号？" />

        <TextView
            android:id="@+id/tvRegister"
            android:layout_width="wrap_content"
            android:layout_height="wrap_content"
            android:text="立即注册"
            android:textColor="#2196F3"
            android:layout_marginStart="4dp" />

    </LinearLayout>

</LinearLayout>
```

---

## 2. ConstraintLayout - 约束布局（推荐）

### 说明

ConstraintLayout 是 Android 推荐的布局容器，通过约束关系定位控件，性能优秀，能替代大多数嵌套布局。

### 依赖

```gradle
implementation "androidx.constraintlayout:constraintlayout:2.1.4"
```

### 核心概念

**约束**：控件的某一边与另一个控件（或父容器）的某一边建立关系。

```
控件A.左边 → 控件B.右边
控件A.上边 → 控件B.下边
控件A.左边 → 父容器.左边
```

### 核心属性

| 属性 | 说明 |
|------|------|
| `app:layout_constraintLeft_toLeftOf` | 左边约束到目标的左边 |
| `app:layout_constraintLeft_toRightOf` | 左边约束到目标的右边 |
| `app:layout_constraintRight_toLeftOf` | 右边约束到目标的左边 |
| `app:layout_constraintRight_toRightOf` | 右边约束到目标的右边 |
| `app:layout_constraintTop_toTopOf` | 上边约束到目标的上边 |
| `app:layout_constraintTop_toBottomOf` | 上边约束到目标的下边 |
| `app:layout_constraintBottom_toTopOf` | 下边约束到目标的上边 |
| `app:layout_constraintBottom_toBottomOf` | 下边约束到目标的下边 |
| `app:layout_constraintStart_toStartOf` | 起始边约束（支持 RTL） |
| `app:layout_constraintEnd_toEndOf` | 结束边约束（支持 RTL） |

**目标值**：
- `parent`：父容器
- `@id/viewId`：指定控件 ID

### 基础示例

#### 居中显示

```xml
<?xml version="1.0" encoding="utf-8"?>
<androidx.constraintlayout.widget.ConstraintLayout
    xmlns:android="http://schemas.android.com/apk/res/android"
    xmlns:app="http://schemas.android.com/apk/res-auto"
    android:layout_width="match_parent"
    android:layout_height="match_parent">

    <!-- 居中：左右上下都约束到父容器 -->
    <TextView
        android:id="@+id/tvTitle"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="居中显示"
        android:textSize="24sp"
        app:layout_constraintLeft_toLeftOf="parent"
        app:layout_constraintRight_toRightOf="parent"
        app:layout_constraintTop_toTopOf="parent"
        app:layout_constraintBottom_toBottomOf="parent" />

</androidx.constraintlayout.widget.ConstraintLayout>
```

#### 相对定位

```xml
<?xml version="1.0" encoding="utf-8"?>
<androidx.constraintlayout.widget.ConstraintLayout
    xmlns:android="http://schemas.android.com/apk/res/android"
    xmlns:app="http://schemas.android.com/apk/res-auto"
    android:layout_width="match_parent"
    android:layout_height="match_parent"
    android:padding="16dp">

    <!-- 头像 -->
    <ImageView
        android:id="@+id/ivAvatar"
        android:layout_width="60dp"
        android:layout_height="60dp"
        android:src="@drawable/ic_avatar"
        app:layout_constraintLeft_toLeftOf="parent"
        app:layout_constraintTop_toTopOf="parent" />

    <!-- 用户名：在头像右侧 -->
    <TextView
        android:id="@+id/tvName"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="用户名"
        android:textSize="16sp"
        android:textStyle="bold"
        android:layout_marginStart="12dp"
        app:layout_constraintLeft_toRightOf="@id/ivAvatar"
        app:layout_constraintTop_toTopOf="@id/ivAvatar" />

    <!-- 消息：在用户名下方 -->
    <TextView
        android:id="@+id/tvMessage"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="这是消息内容..."
        android:textSize="14sp"
        android:layout_marginTop="4dp"
        app:layout_constraintLeft_toLeftOf="@id/tvName"
        app:layout_constraintTop_toBottomOf="@id/tvName" />

    <!-- 时间：右上角 -->
    <TextView
        android:id="@+id/tvTime"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="12:30"
        android:textSize="12sp"
        android:textColor="#999999"
        app:layout_constraintRight_toRightOf="parent"
        app:layout_constraintTop_toTopOf="@id/ivAvatar" />

</androidx.constraintlayout.widget.ConstraintLayout>
```

### Bias 偏移

当控件两边都有约束时，可以设置偏移比例（默认 0.5 居中）。

```xml
<TextView
    android:layout_width="wrap_content"
    android:layout_height="wrap_content"
    android:text="偏移 30%"
    app:layout_constraintLeft_toLeftOf="parent"
    app:layout_constraintRight_toRightOf="parent"
    app:layout_constraintHorizontal_bias="0.3" />

<TextView
    android:layout_width="wrap_content"
    android:layout_height="wrap_content"
    android:text="偏移 70%"
    app:layout_constraintTop_toTopOf="parent"
    app:layout_constraintBottom_toBottomOf="parent"
    app:layout_constraintVertical_bias="0.7" />
```

### 边距属性

```xml
<!-- 普通边距 -->
android:layout_margin="16dp"
android:layout_marginTop="8dp"
android:layout_marginStart="12dp"
android:layout_marginEnd="12dp"

<!-- 约束边距：当目标可见性为 gone 时生效 -->
app:layout_goneMarginTop="16dp"
app:layout_goneMarginStart="16dp"
```

### 宽高比例（Dimension Ratio）

```xml
<!-- 16:9 比例 -->
<ImageView
    android:layout_width="0dp"
    android:layout_height="0dp"
    android:src="@drawable/image"
    app:layout_constraintLeft_toLeftOf="parent"
    app:layout_constraintRight_toRightOf="parent"
    app:layout_constraintDimensionRatio="16:9" />

<!-- 宽度固定，高度按比例 -->
<ImageView
    android:layout_width="200dp"
    android:layout_height="0dp"
    app:layout_constraintDimensionRatio="1:1" />  <!-- 正方形 -->
```

### 链（Chain）

多个控件在同一方向形成链，可控制它们的分布方式。

```xml
<!-- 水平链：三个按钮均分 -->
<androidx.constraintlayout.widget.ConstraintLayout
    android:layout_width="match_parent"
    android:layout_height="wrap_content">

    <Button
        android:id="@+id/btn1"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="按钮1"
        app:layout_constraintLeft_toLeftOf="parent"
        app:layout_constraintRight_toLeftOf="@id/btn2"
        app:layout_constraintHorizontal_chainStyle="spread" />

    <Button
        android:id="@+id/btn2"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="按钮2"
        app:layout_constraintLeft_toRightOf="@id/btn1"
        app:layout_constraintRight_toLeftOf="@id/btn3" />

    <Button
        android:id="@+id/btn3"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="按钮3"
        app:layout_constraintLeft_toRightOf="@id/btn2"
        app:layout_constraintRight_toRightOf="parent" />

</androidx.constraintlayout.widget.ConstraintLayout>
```

**链样式**：
- `spread`：均匀分布（默认）
- `spread_inside`：两端贴边，中间均匀分布
- `packed`：紧凑排列

### Guideline 辅助线

不可见的辅助线，用于定位参考。

```xml
<androidx.constraintlayout.widget.ConstraintLayout
    android:layout_width="match_parent"
    android:layout_height="match_parent">

    <!-- 垂直辅助线：距离左边 30% -->
    <androidx.constraintlayout.widget.Guideline
        android:id="@+id/guideline"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:orientation="vertical"
        app:layout_constraintGuide_percent="0.3" />

    <!-- 也可以用固定 dp -->
    <androidx.constraintlayout.widget.Guideline
        android:id="@+id/guideline2"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:orientation="vertical"
        app:layout_constraintGuide_begin="100dp" />

    <!-- 使用辅助线约束 -->
    <TextView
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="参考辅助线定位"
        app:layout_constraintLeft_toLeftOf="@id/guideline"
        app:layout_constraintTop_toTopOf="parent" />

</androidx.constraintlayout.widget.ConstraintLayout>
```

### Barrier 屏障

根据多个控件的位置动态创建边界。

```xml
<androidx.constraintlayout.widget.ConstraintLayout
    android:layout_width="match_parent"
    android:layout_height="wrap_content">

    <TextView
        android:id="@+id/tvLabel1"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="标签1："
        app:layout_constraintLeft_toLeftOf="parent"
        app:layout_constraintTop_toTopOf="parent" />

    <TextView
        android:id="@+id/tvLabel2"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="较长的标签名："
        android:layout_marginTop="8dp"
        app:layout_constraintLeft_toLeftOf="parent"
        app:layout_constraintTop_toBottomOf="@id/tvLabel1" />

    <!-- 屏障：根据最右边的标签创建边界 -->
    <androidx.constraintlayout.widget.Barrier
        android:id="@+id/barrier"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        app:barrierDirection="end"
        app:constraint_referenced_ids="tvLabel1,tvLabel2" />

    <!-- 输入框：始终在屏障右侧 -->
    <EditText
        android:id="@+id/etInput"
        android:layout_width="0dp"
        android:layout_height="wrap_content"
        android:layout_marginStart="8dp"
        app:layout_constraintLeft_toRightOf="@id/barrier"
        app:layout_constraintRight_toRightOf="parent"
        app:layout_constraintTop_toTopOf="parent" />

</androidx.constraintlayout.widget.ConstraintLayout>
```

### 实战：个人中心页面

```xml
<?xml version="1.0" encoding="utf-8"?>
<androidx.constraintlayout.widget.ConstraintLayout
    xmlns:android="http://schemas.android.com/apk/res/android"
    xmlns:app="http://schemas.android.com/apk/res-auto"
    android:layout_width="match_parent"
    android:layout_height="match_parent"
    android:background="#F5F5F5">

    <!-- 顶部背景 -->
    <View
        android:id="@+id/viewHeader"
        android:layout_width="0dp"
        android:layout_height="200dp"
        android:background="#6200EE"
        app:layout_constraintLeft_toLeftOf="parent"
        app:layout_constraintRight_toRightOf="parent"
        app:layout_constraintTop_toTopOf="parent" />

    <!-- 头像 -->
    <ImageView
        android:id="@+id/ivAvatar"
        android:layout_width="80dp"
        android:layout_height="80dp"
        android:src="@drawable/ic_avatar"
        android:layout_marginTop="120dp"
        app:layout_constraintLeft_toLeftOf="parent"
        app:layout_constraintRight_toRightOf="parent"
        app:layout_constraintTop_toTopOf="parent" />

    <!-- 用户名 -->
    <TextView
        android:id="@+id/tvName"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="用户名"
        android:textSize="20sp"
        android:textStyle="bold"
        android:layout_marginTop="12dp"
        app:layout_constraintLeft_toLeftOf="parent"
        app:layout_constraintRight_toRightOf="parent"
        app:layout_constraintTop_toBottomOf="@id/ivAvatar" />

    <!-- 个人简介 -->
    <TextView
        android:id="@+id/tvBio"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="这个人很懒，什么都没写"
        android:textSize="14sp"
        android:textColor="#666666"
        android:layout_marginTop="4dp"
        app:layout_constraintLeft_toLeftOf="parent"
        app:layout_constraintRight_toRightOf="parent"
        app:layout_constraintTop_toBottomOf="@id/tvName" />

    <!-- 统计信息 -->
    <TextView
        android:id="@+id/tvPosts"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="128"
        android:textSize="18sp"
        android:textStyle="bold"
        android:layout_marginTop="24dp"
        app:layout_constraintLeft_toLeftOf="parent"
        app:layout_constraintRight_toLeftOf="@id/tvFollowers"
        app:layout_constraintTop_toBottomOf="@id/tvBio" />

    <TextView
        android:id="@+id/tvPostsLabel"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="帖子"
        android:textSize="12sp"
        android:textColor="#999999"
        app:layout_constraintLeft_toLeftOf="@id/tvPosts"
        app:layout_constraintRight_toRightOf="@id/tvPosts"
        app:layout_constraintTop_toBottomOf="@id/tvPosts" />

    <TextView
        android:id="@+id/tvFollowers"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="1.2k"
        android:textSize="18sp"
        android:textStyle="bold"
        app:layout_constraintLeft_toRightOf="@id/tvPosts"
        app:layout_constraintRight_toLeftOf="@id/tvFollowing"
        app:layout_constraintTop_toTopOf="@id/tvPosts" />

    <TextView
        android:id="@+id/tvFollowersLabel"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="粉丝"
        android:textSize="12sp"
        android:textColor="#999999"
        app:layout_constraintLeft_toLeftOf="@id/tvFollowers"
        app:layout_constraintRight_toRightOf="@id/tvFollowers"
        app:layout_constraintTop_toBottomOf="@id/tvFollowers" />

    <TextView
        android:id="@+id/tvFollowing"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="256"
        android:textSize="18sp"
        android:textStyle="bold"
        app:layout_constraintLeft_toRightOf="@id/tvFollowers"
        app:layout_constraintRight_toRightOf="parent"
        app:layout_constraintTop_toTopOf="@id/tvPosts" />

    <TextView
        android:id="@+id/tvFollowingLabel"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="关注"
        android:textSize="12sp"
        android:textColor="#999999"
        app:layout_constraintLeft_toLeftOf="@id/tvFollowing"
        app:layout_constraintRight_toRightOf="@id/tvFollowing"
        app:layout_constraintTop_toBottomOf="@id/tvFollowing" />

    <!-- 编辑资料按钮 -->
    <Button
        android:id="@+id/btnEdit"
        android:layout_width="0dp"
        android:layout_height="wrap_content"
        android:text="编辑资料"
        android:layout_marginHorizontal="24dp"
        android:layout_marginTop="24dp"
        app:layout_constraintLeft_toLeftOf="parent"
        app:layout_constraintRight_toRightOf="parent"
        app:layout_constraintTop_toBottomOf="@id/tvPostsLabel" />

</androidx.constraintlayout.widget.ConstraintLayout>
```

---

## 3. FrameLayout - 帧布局

### 说明

所有子控件默认叠加在左上角，适合单一控件或叠加效果（如图片上放文字）。

### 核心属性

| 属性 | 说明 |
|------|------|
| `android:foreground` | 前景图 |
| `android:foregroundGravity` | 前景位置 |

### 基础示例

```xml
<?xml version="1.0" encoding="utf-8"?>
<FrameLayout xmlns:android="http://schemas.android.com/apk/res/android"
    android:layout_width="match_parent"
    android:layout_height="200dp">

    <!-- 底层：背景图 -->
    <ImageView
        android:layout_width="match_parent"
        android:layout_height="match_parent"
        android:src="@drawable/background"
        android:scaleType="centerCrop" />

    <!-- 中层：半透明遮罩 -->
    <View
        android:layout_width="match_parent"
        android:layout_height="match_parent"
        android:background="#80000000" />

    <!-- 顶层：文字 -->
    <TextView
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="叠加文字"
        android:textColor="@android:color/white"
        android:textSize="24sp"
        android:layout_gravity="center" />

</FrameLayout>
```

### 实战：带标签的图片

```xml
<?xml version="1.0" encoding="utf-8"?>
<FrameLayout xmlns:android="http://schemas.android.com/apk/res/android"
    android:layout_width="match_parent"
    android:layout_height="200dp">

    <!-- 图片 -->
    <ImageView
        android:id="@+id/ivCover"
        android:layout_width="match_parent"
        android:layout_height="match_parent"
        android:src="@drawable/cover"
        android:scaleType="centerCrop" />

    <!-- 标签：右上角 -->
    <TextView
        android:id="@+id/tvTag"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="NEW"
        android:textColor="@android:color/white"
        android:background="#FF5722"
        android:paddingHorizontal="8dp"
        android:paddingVertical="4dp"
        android:layout_gravity="top|end"
        android:layout_margin="8dp" />

    <!-- 标题：底部 -->
    <TextView
        android:id="@+id/tvTitle"
        android:layout_width="match_parent"
        android:layout_height="wrap_content"
        android:text="文章标题"
        android:textColor="@android:color/white"
        android:textSize="18sp"
        android:background="#80000000"
        android:padding="12dp"
        android:layout_gravity="bottom" />

</FrameLayout>
```

### 实战：加载中遮罩

```xml
<?xml version="1.0" encoding="utf-8"?>
<FrameLayout xmlns:android="http://schemas.android.com/apk/res/android"
    android:layout_width="match_parent"
    android:layout_height="match_parent">

    <!-- 主内容 -->
    <LinearLayout
        android:layout_width="match_parent"
        android:layout_height="match_parent"
        android:orientation="vertical">

        <!-- 页面内容 -->
        <TextView
            android:layout_width="wrap_content"
            android:layout_height="wrap_content"
            android:text="页面内容" />

    </LinearLayout>

    <!-- 加载遮罩 -->
    <FrameLayout
        android:id="@+id/loadingOverlay"
        android:layout_width="match_parent"
        android:layout_height="match_parent"
        android:background="#80000000"
        android:visibility="gone"
        android:clickable="true"
        android:focusable="true">

        <ProgressBar
            android:layout_width="wrap_content"
            android:layout_height="wrap_content"
            android:layout_gravity="center" />

    </FrameLayout>

</FrameLayout>
```

---

## 4. RelativeLayout - 相对布局

### 说明

通过控件之间的相对位置定位，较老的布局方式，已被 ConstraintLayout 取代。

### 核心属性

| 属性 | 说明 |
|------|------|
| `android:layout_centerInParent` | 相对于父容器居中 |
| `android:layout_centerHorizontal` | 水平居中 |
| `android:layout_centerVertical` | 垂直居中 |
| `android:layout_above` | 在某控件上方 |
| `android:layout_below` | 在某控件下方 |
| `android:layout_toLeftOf` | 在某控件左边 |
| `android:layout_toRightOf` | 在某控件右边 |
| `android:layout_alignParentTop` | 与父容器上边对齐 |
| `android:layout_alignParentBottom` | 与父容器下边对齐 |
| `android:layout_alignParentLeft` | 与父容器左边对齐 |
| `android:layout_alignParentRight` | 与父容器右边对齐 |
| `android:layout_alignTop` | 与某控件上边对齐 |
| `android:layout_alignBottom` | 与某控件下边对齐 |
| `android:layout_alignLeft` | 与某控件左边对齐 |
| `android:layout_alignRight` | 与某控件右边对齐 |

### 基础示例

```xml
<?xml version="1.0" encoding="utf-8"?>
<RelativeLayout xmlns:android="http://schemas.android.com/apk/res/android"
    android:layout_width="match_parent"
    android:layout_height="match_parent">

    <!-- 居中的标题 -->
    <TextView
        android:id="@+id/tvTitle"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="标题"
        android:textSize="24sp"
        android:layout_centerInParent="true" />

    <!-- 标题左边 -->
    <Button
        android:id="@+id/btnLeft"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="左"
        android:layout_toLeftOf="@id/tvTitle"
        android:layout_centerVertical="true"
        android:layout_marginEnd="16dp" />

    <!-- 标题右边 -->
    <Button
        android:id="@+id/btnRight"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="右"
        android:layout_toRightOf="@id/tvTitle"
        android:layout_centerVertical="true"
        android:layout_marginStart="16dp" />

    <!-- 底部按钮 -->
    <Button
        android:id="@+id/btnBottom"
        android:layout_width="match_parent"
        android:layout_height="wrap_content"
        android:text="底部按钮"
        android:layout_alignParentBottom="true"
        android:layout_margin="16dp" />

</RelativeLayout>
```

---

## 5. GridLayout - 网格布局

### 说明

将子控件按网格排列，适合计算器、键盘等场景。

### 核心属性

| 属性 | 说明 |
|------|------|
| `android:columnCount` | 列数 |
| `android:rowCount` | 行数 |
| `android:orientation` | 排列方向 |
| `android:layout_column` | 指定列位置 |
| `android:layout_row` | 指定行位置 |
| `android:layout_columnSpan` | 跨列数 |
| `android:layout_rowSpan` | 跨行数 |
| `android:layout_gravity` | 填充方式（fill） |

### 基础示例

```xml
<?xml version="1.0" encoding="utf-8"?>
<GridLayout xmlns:android="http://schemas.android.com/apk/res/android"
    android:layout_width="match_parent"
    android:layout_height="wrap_content"
    android:columnCount="4"
    android:rowCount="5"
    android:padding="8dp">

    <!-- 第一行：跨 4 列 -->
    <EditText
        android:id="@+id/etDisplay"
        android:layout_width="0dp"
        android:layout_height="wrap_content"
        android:layout_columnSpan="4"
        android:layout_gravity="fill_horizontal"
        android:gravity="end"
        android:textSize="32sp"
        android:text="0" />

    <!-- 第二行 -->
    <Button android:text="C" />
    <Button android:text="±" />
    <Button android:text="%" />
    <Button android:text="÷" />

    <!-- 第三行 -->
    <Button android:text="7" />
    <Button android:text="8" />
    <Button android:text="9" />
    <Button android:text="×" />

    <!-- 第四行 -->
    <Button android:text="4" />
    <Button android:text="5" />
    <Button android:text="6" />
    <Button android:text="-" />

    <!-- 第五行 -->
    <Button android:text="1" />
    <Button android:text="2" />
    <Button android:text="3" />
    <Button android:text="+" />

    <!-- 第六行 -->
    <Button
        android:text="0"
        android:layout_columnSpan="2"
        android:layout_gravity="fill_horizontal" />
    <Button android:text="." />
    <Button android:text="=" />

</GridLayout>
```

---

## 6. TableLayout - 表格布局

### 说明

表格形式的布局，使用 TableRow 定义行。

### 核心属性

| 属性 | 说明 |
|------|------|
| `android:stretchColumns` | 可拉伸的列（索引从 0 开始） |
| `android:shrinkColumns` | 可收缩的列 |
| `android:collapseColumns` | 隐藏的列 |
| `android:layout_span` | 跨列数 |

### 基础示例

```xml
<?xml version="1.0" encoding="utf-8"?>
<TableLayout xmlns:android="http://schemas.android.com/apk/res/android"
    android:layout_width="match_parent"
    android:layout_height="wrap_content"
    android:stretchColumns="1,2"
    android:padding="8dp">

    <!-- 表头 -->
    <TableRow android:background="#E0E0E0">
        <TextView
            android:text="序号"
            android:padding="8dp"
            android:textStyle="bold" />
        <TextView
            android:text="姓名"
            android:padding="8dp"
            android:textStyle="bold" />
        <TextView
            android:text="分数"
            android:padding="8dp"
            android:textStyle="bold" />
    </TableRow>

    <!-- 数据行 -->
    <TableRow>
        <TextView
            android:text="1"
            android:padding="8dp" />
        <TextView
            android:text="张三"
            android:padding="8dp" />
        <TextView
            android:text="95"
            android:padding="8dp" />
    </TableRow>

    <TableRow>
        <TextView
            android:text="2"
            android:padding="8dp" />
        <TextView
            android:text="李四"
            android:padding="8dp" />
        <TextView
            android:text="88"
            android:padding="8dp" />
    </TableRow>

    <!-- 跨列 -->
    <TableRow>
        <TextView
            android:text="平均分：91.5"
            android:padding="8dp"
            android:layout_span="3"
            android:gravity="center" />
    </TableRow>

</TableLayout>
```

---

## 常用布局属性

### 尺寸单位

| 单位 | 说明 | 使用场景 |
|------|------|----------|
| `dp` | 密度无关像素 | 控件尺寸、边距 |
| `sp` | 缩放无关像素 | 字体大小 |
| `px` | 像素 | 一般不使用 |
| `in` | 英寸 | 特殊场景 |
| `mm` | 毫米 | 特殊场景 |

### 宽高值

| 值 | 说明 |
|------|------|
| `match_parent` | 填满父容器 |
| `wrap_content` | 根据内容自适应 |
| `具体数值` | 如 `100dp`、`200px` |
| `0dp` | 配合 weight 或约束使用 |

### 内外边距

```xml
<!-- 内边距 padding -->
android:padding="16dp"           <!-- 四边 -->
android:paddingTop="8dp"         <!-- 上 -->
android:paddingBottom="8dp"      <!-- 下 -->
android:paddingLeft="8dp"        <!-- 左 -->
android:paddingRight="8dp"       <!-- 右 -->
android:paddingStart="8dp"       <!-- 起始（支持 RTL） -->
android:paddingEnd="8dp"         <!-- 结束（支持 RTL） -->
android:paddingHorizontal="8dp"  <!-- 水平 -->
android:paddingVertical="8dp"    <!-- 垂直 -->

<!-- 外边距 margin -->
android:layout_margin="16dp"
android:layout_marginTop="8dp"
android:layout_marginBottom="8dp"
android:layout_marginLeft="8dp"
android:layout_marginRight="8dp"
android:layout_marginStart="8dp"
android:layout_marginEnd="8dp"
android:layout_marginHorizontal="8dp"
android:layout_marginVertical="8dp"
```

### 可见性

```xml
android:visibility="visible"    <!-- 可见（默认） -->
android:visibility="invisible"  <!-- 不可见，占用空间 -->
android:visibility="gone"       <!-- 不可见，不占用空间 -->
```

**Kotlin 代码控制**：
```kotlin
view.visibility = View.VISIBLE
view.visibility = View.INVISIBLE
view.visibility = View.GONE
```

### 背景

```xml
<!-- 纯色背景 -->
android:background="#FF5722"
android:background="@color/red"

<!-- 图片背景 -->
android:background="@drawable/bg_image"

<!-- 圆角背景（drawable） -->
android:background="@drawable/bg_rounded"

<!-- 透明背景 -->
android:background="@null"
android:background="#00000000"
```

**res/drawable/bg_rounded.xml**：
```xml
<?xml version="1.0" encoding="utf-8"?>
<shape xmlns:android="http://schemas.android.com/apk/res/android">
    <solid android:color="#FFFFFF" />
    <corners android:radius="8dp" />
    <stroke
        android:width="1dp"
        android:color="#E0E0E0" />
</shape>
```

---

## 布局优化技巧

### 1. 减少布局层级

❌ **不推荐**：嵌套过多
```xml
<LinearLayout>
    <LinearLayout>
        <LinearLayout>
            <TextView />
        </LinearLayout>
    </LinearLayout>
</LinearLayout>
```

✅ **推荐**：使用 ConstraintLayout 扁平化
```xml
<ConstraintLayout>
    <TextView />
</ConstraintLayout>
```

### 2. 使用 merge 标签

当布局作为子布局被 include 时，可以用 `<merge>` 消除多余层级。

**res/layout/item_buttons.xml**：
```xml
<?xml version="1.0" encoding="utf-8"?>
<merge xmlns:android="http://schemas.android.com/apk/res/android">
    <Button
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="确定" />
    <Button
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="取消" />
</merge>
```

### 3. 使用 include 复用布局

```xml
<LinearLayout xmlns:android="http://schemas.android.com/apk/res/android"
    android:layout_width="match_parent"
    android:layout_height="match_parent"
    android:orientation="vertical">

    <!-- 引入公共标题栏 -->
    <include layout="@layout/layout_toolbar" />

    <!-- 主内容 -->
    <FrameLayout
        android:id="@+id/container"
        android:layout_width="match_parent"
        android:layout_height="match_parent" />

</LinearLayout>
```

### 4. 使用 ViewStub 延迟加载

```xml
<ViewStub
    android:id="@+id/vsEmpty"
    android:layout_width="match_parent"
    android:layout_height="match_parent"
    android:layout="@layout/layout_empty"
    android:inflatedId="@+id/emptyView" />
```

**Kotlin 代码**：
```kotlin
// 需要时才加载
val viewStub = findViewById<ViewStub>(R.id.vsEmpty)
val emptyView = viewStub.inflate()

// 或者设置为可见时自动加载
viewStub.visibility = View.VISIBLE
```

### 5. 使用 Space 占位

```xml
<LinearLayout
    android:layout_width="match_parent"
    android:layout_height="wrap_content"
    android:orientation="horizontal">

    <TextView
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="左边" />

    <!-- 占位，更轻量 -->
    <Space
        android:layout_width="0dp"
        android:layout_height="0dp"
        android:layout_weight="1" />

    <TextView
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="右边" />

</LinearLayout>
```

---

## 实战案例

### 案例 1：列表项布局

```xml
<?xml version="1.0" encoding="utf-8"?>
<!-- RecyclerView 列表项 -->
<androidx.constraintlayout.widget.ConstraintLayout
    xmlns:android="http://schemas.android.com/apk/res/android"
    xmlns:app="http://schemas.android.com/apk/res-auto"
    xmlns:tools="http://schemas.android.com/tools"
    android:layout_width="match_parent"
    android:layout_height="wrap_content"
    android:padding="16dp">

    <!-- 头像 -->
    <ImageView
        android:id="@+id/ivAvatar"
        android:layout_width="48dp"
        android:layout_height="48dp"
        android:src="@drawable/ic_avatar"
        app:layout_constraintLeft_toLeftOf="parent"
        app:layout_constraintTop_toTopOf="parent"
        tools:src="@tools:sample/avatars" />

    <!-- 用户名 -->
    <TextView
        android:id="@+id/tvName"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:textSize="16sp"
        android:textStyle="bold"
        android:layout_marginStart="12dp"
        app:layout_constraintLeft_toRightOf="@id/ivAvatar"
        app:layout_constraintTop_toTopOf="@id/ivAvatar"
        tools:text="用户名" />

    <!-- 时间 -->
    <TextView
        android:id="@+id/tvTime"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:textSize="12sp"
        android:textColor="#999999"
        android:layout_marginTop="2dp"
        app:layout_constraintLeft_toLeftOf="@id/tvName"
        app:layout_constraintTop_toBottomOf="@id/tvName"
        tools:text="3分钟前" />

    <!-- 内容 -->
    <TextView
        android:id="@+id/tvContent"
        android:layout_width="0dp"
        android:layout_height="wrap_content"
        android:textSize="14sp"
        android:layout_marginTop="8dp"
        android:layout_marginStart="60dp"
        app:layout_constraintLeft_toLeftOf="parent"
        app:layout_constraintRight_toRightOf="parent"
        app:layout_constraintTop_toBottomOf="@id/ivAvatar"
        tools:text="这是消息内容，可以很长很长很长很长很长很长很长很长" />

    <!-- 分割线 -->
    <View
        android:layout_width="0dp"
        android:layout_height="1dp"
        android:background="#E0E0E0"
        android:layout_marginTop="16dp"
        app:layout_constraintLeft_toLeftOf="parent"
        app:layout_constraintRight_toRightOf="parent"
        app:layout_constraintTop_toBottomOf="@id/tvContent" />

</androidx.constraintlayout.widget.ConstraintLayout>
```

### 案例 2：登录页面

```xml
<?xml version="1.0" encoding="utf-8"?>
<ScrollView xmlns:android="http://schemas.android.com/apk/res/android"
    xmlns:app="http://schemas.android.com/apk/res-auto"
    android:layout_width="match_parent"
    android:layout_height="match_parent"
    android:fillViewport="true">

    <androidx.constraintlayout.widget.ConstraintLayout
        android:layout_width="match_parent"
        android:layout_height="wrap_content"
        android:padding="24dp">

        <!-- Logo -->
        <ImageView
            android:id="@+id/ivLogo"
            android:layout_width="80dp"
            android:layout_height="80dp"
            android:src="@drawable/ic_logo"
            android:layout_marginTop="48dp"
            app:layout_constraintLeft_toLeftOf="parent"
            app:layout_constraintRight_toRightOf="parent"
            app:layout_constraintTop_toTopOf="parent" />

        <!-- 欢迎文字 -->
        <TextView
            android:id="@+id/tvWelcome"
            android:layout_width="wrap_content"
            android:layout_height="wrap_content"
            android:text="欢迎登录"
            android:textSize="24sp"
            android:textStyle="bold"
            android:layout_marginTop="24dp"
            app:layout_constraintLeft_toLeftOf="parent"
            app:layout_constraintRight_toRightOf="parent"
            app:layout_constraintTop_toBottomOf="@id/ivLogo" />

        <!-- 用户名输入框 -->
        <com.google.android.material.textfield.TextInputLayout
            android:id="@+id/tilUsername"
            android:layout_width="0dp"
            android:layout_height="wrap_content"
            android:hint="用户名"
            android:layout_marginTop="32dp"
            app:startIconDrawable="@drawable/ic_person"
            style="@style/Widget.MaterialComponents.TextInputLayout.OutlinedBox"
            app:layout_constraintLeft_toLeftOf="parent"
            app:layout_constraintRight_toRightOf="parent"
            app:layout_constraintTop_toBottomOf="@id/tvWelcome">

            <com.google.android.material.textfield.TextInputEditText
                android:id="@+id/etUsername"
                android:layout_width="match_parent"
                android:layout_height="wrap_content"
                android:inputType="text" />

        </com.google.android.material.textfield.TextInputLayout>

        <!-- 密码输入框 -->
        <com.google.android.material.textfield.TextInputLayout
            android:id="@+id/tilPassword"
            android:layout_width="0dp"
            android:layout_height="wrap_content"
            android:hint="密码"
            android:layout_marginTop="16dp"
            app:startIconDrawable="@drawable/ic_lock"
            app:endIconMode="password_toggle"
            style="@style/Widget.MaterialComponents.TextInputLayout.OutlinedBox"
            app:layout_constraintLeft_toLeftOf="parent"
            app:layout_constraintRight_toRightOf="parent"
            app:layout_constraintTop_toBottomOf="@id/tilUsername">

            <com.google.android.material.textfield.TextInputEditText
                android:id="@+id/etPassword"
                android:layout_width="match_parent"
                android:layout_height="wrap_content"
                android:inputType="textPassword" />

        </com.google.android.material.textfield.TextInputLayout>

        <!-- 忘记密码 -->
        <TextView
            android:id="@+id/tvForgotPassword"
            android:layout_width="wrap_content"
            android:layout_height="wrap_content"
            android:text="忘记密码？"
            android:textColor="#2196F3"
            android:layout_marginTop="8dp"
            app:layout_constraintRight_toRightOf="parent"
            app:layout_constraintTop_toBottomOf="@id/tilPassword" />

        <!-- 登录按钮 -->
        <Button
            android:id="@+id/btnLogin"
            android:layout_width="0dp"
            android:layout_height="48dp"
            android:text="登录"
            android:textSize="16sp"
            android:layout_marginTop="24dp"
            app:layout_constraintLeft_toLeftOf="parent"
            app:layout_constraintRight_toRightOf="parent"
            app:layout_constraintTop_toBottomOf="@id/tvForgotPassword" />

        <!-- 分割线 -->
        <LinearLayout
            android:id="@+id/llDivider"
            android:layout_width="0dp"
            android:layout_height="wrap_content"
            android:orientation="horizontal"
            android:gravity="center"
            android:layout_marginTop="32dp"
            app:layout_constraintLeft_toLeftOf="parent"
            app:layout_constraintRight_toRightOf="parent"
            app:layout_constraintTop_toBottomOf="@id/btnLogin">

            <View
                android:layout_width="0dp"
                android:layout_height="1dp"
                android:layout_weight="1"
                android:background="#E0E0E0" />

            <TextView
                android:layout_width="wrap_content"
                android:layout_height="wrap_content"
                android:text="或"
                android:textColor="#999999"
                android:layout_marginHorizontal="16dp" />

            <View
                android:layout_width="0dp"
                android:layout_height="1dp"
                android:layout_weight="1"
                android:background="#E0E0E0" />

        </LinearLayout>

        <!-- 第三方登录 -->
        <LinearLayout
            android:id="@+id/llSocialLogin"
            android:layout_width="wrap_content"
            android:layout_height="wrap_content"
            android:orientation="horizontal"
            android:layout_marginTop="24dp"
            app:layout_constraintLeft_toLeftOf="parent"
            app:layout_constraintRight_toRightOf="parent"
            app:layout_constraintTop_toBottomOf="@id/llDivider">

            <ImageButton
                android:id="@+id/btnWechat"
                android:layout_width="48dp"
                android:layout_height="48dp"
                android:src="@drawable/ic_wechat"
                android:background="?attr/selectableItemBackgroundBorderless" />

            <ImageButton
                android:id="@+id/btnQQ"
                android:layout_width="48dp"
                android:layout_height="48dp"
                android:src="@drawable/ic_qq"
                android:background="?attr/selectableItemBackgroundBorderless"
                android:layout_marginStart="24dp" />

        </LinearLayout>

        <!-- 注册提示 -->
        <LinearLayout
            android:layout_width="wrap_content"
            android:layout_height="wrap_content"
            android:orientation="horizontal"
            android:layout_marginBottom="24dp"
            app:layout_constraintLeft_toLeftOf="parent"
            app:layout_constraintRight_toRightOf="parent"
            app:layout_constraintBottom_toBottomOf="parent">

            <TextView
                android:layout_width="wrap_content"
                android:layout_height="wrap_content"
                android:text="还没有账号？" />

            <TextView
                android:id="@+id/tvRegister"
                android:layout_width="wrap_content"
                android:layout_height="wrap_content"
                android:text="立即注册"
                android:textColor="#2196F3"
                android:layout_marginStart="4dp" />

        </LinearLayout>

    </androidx.constraintlayout.widget.ConstraintLayout>

</ScrollView>
```

### 案例 3：底部对话框

```xml
<?xml version="1.0" encoding="utf-8"?>
<!-- 底部弹出对话框布局 -->
<LinearLayout xmlns:android="http://schemas.android.com/apk/res/android"
    android:layout_width="match_parent"
    android:layout_height="wrap_content"
    android:orientation="vertical"
    android:background="@drawable/bg_bottom_sheet"
    android:padding="16dp">

    <!-- 顶部拖动条 -->
    <View
        android:layout_width="40dp"
        android:layout_height="4dp"
        android:background="#E0E0E0"
        android:layout_gravity="center_horizontal"
        android:layout_marginBottom="16dp" />

    <!-- 标题 -->
    <TextView
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="选择操作"
        android:textSize="18sp"
        android:textStyle="bold"
        android:layout_marginBottom="16dp" />

    <!-- 拍照 -->
    <LinearLayout
        android:id="@+id/llCamera"
        android:layout_width="match_parent"
        android:layout_height="wrap_content"
        android:orientation="horizontal"
        android:paddingVertical="12dp"
        android:background="?attr/selectableItemBackground">

        <ImageView
            android:layout_width="24dp"
            android:layout_height="24dp"
            android:src="@drawable/ic_camera" />

        <TextView
            android:layout_width="wrap_content"
            android:layout_height="wrap_content"
            android:text="拍照"
            android:textSize="16sp"
            android:layout_marginStart="16dp" />

    </LinearLayout>

    <!-- 从相册选择 -->
    <LinearLayout
        android:id="@+id/llGallery"
        android:layout_width="match_parent"
        android:layout_height="wrap_content"
        android:orientation="horizontal"
        android:paddingVertical="12dp"
        android:background="?attr/selectableItemBackground">

        <ImageView
            android:layout_width="24dp"
            android:layout_height="24dp"
            android:src="@drawable/ic_gallery" />

        <TextView
            android:layout_width="wrap_content"
            android:layout_height="wrap_content"
            android:text="从相册选择"
            android:textSize="16sp"
            android:layout_marginStart="16dp" />

    </LinearLayout>

    <!-- 分割线 -->
    <View
        android:layout_width="match_parent"
        android:layout_height="1dp"
        android:background="#E0E0E0"
        android:layout_marginVertical="8dp" />

    <!-- 取消 -->
    <TextView
        android:id="@+id/tvCancel"
        android:layout_width="match_parent"
        android:layout_height="wrap_content"
        android:text="取消"
        android:textSize="16sp"
        android:textColor="#999999"
        android:gravity="center"
        android:paddingVertical="12dp"
        android:background="?attr/selectableItemBackground" />

</LinearLayout>
```

---

## 总结

### 布局选择指南

| 场景 | 推荐布局 |
|------|----------|
| 复杂界面、扁平化 | **ConstraintLayout** |
| 简单垂直/水平排列 | LinearLayout |
| 叠加效果（图片+文字） | FrameLayout |
| 网格排列（计算器、键盘） | GridLayout |
| 表格数据 | TableLayout |

### 最佳实践

1. ✅ **优先使用 ConstraintLayout**，减少布局层级
2. ✅ 使用 `tools:` 属性预览，不打包进 APK
3. ✅ 使用 `<include>` 复用公共布局
4. ✅ 使用 `<ViewStub>` 延迟加载不常用布局
5. ✅ 使用 `0dp + weight` 或 `ConstraintLayout` 实现自适应
6. ❌ 避免过度嵌套 LinearLayout
7. ❌ 避免使用 RelativeLayout（已被 ConstraintLayout 取代）

---

> 整理者：尘曦 🌅  
> 最后更新：2026-03-12
