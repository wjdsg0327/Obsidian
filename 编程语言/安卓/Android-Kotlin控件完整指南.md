# Android Studio Kotlin 控件完整指南

> 整理时间：2026-03-12  
> 适用于 Android 开发初学者和进阶者

---

## 📚 目录

1. [基础控件](#基础控件)
2. [布局容器](#布局容器)
3. [列表控件](#列表控件)
4. [Material Design 控件](#material-design-控件)
5. [高级控件](#高级控件)
6. [对话框](#对话框)
7. [常用属性速查表](#常用属性速查表)

---

## 基础控件

### 1. TextView - 文本显示

**说明**：用于显示文本内容，是最基础的控件之一。

**常用属性**：
- `text`：显示的文本内容
- `textSize`：文字大小（sp单位）
- `textColor`：文字颜色
- `textStyle`：文字样式（bold, italic, normal）
- `maxLines`：最大行数
- `ellipsize`：文字超出时的省略方式（end, start, middle, marquee）

**XML 示例**：
```xml
<TextView
    android:id="@+id/tvTitle"
    android:layout_width="wrap_content"
    android:layout_height="wrap_content"
    android:text="Hello Kotlin!"
    android:textSize="18sp"
    android:textColor="#333333"
    android:textStyle="bold" />
```

**Kotlin 代码示例**：
```kotlin
val tvTitle: TextView = findViewById(R.id.tvTitle)
tvTitle.text = "动态设置的文本"
tvTitle.setTextColor(Color.BLUE)
tvTitle.setTextSize(TypedValue.COMPLEX_UNIT_SP, 20f)

// 或者使用 ViewBinding
binding.tvTitle.text = "使用 ViewBinding"
```

---

### 2. EditText - 输入框

**说明**：用于用户输入文本，继承自 TextView。

**常用属性**：
- `hint`：输入提示文字
- `inputType`：输入类型（text, number, phone, textPassword, textEmailAddress）
- `maxLength`：最大输入长度
- `maxLines`：最大行数
- `singleLine`：是否单行输入

**XML 示例**：
```xml
<EditText
    android:id="@+id/etUsername"
    android:layout_width="match_parent"
    android:layout_height="wrap_content"
    android:hint="请输入用户名"
    android:inputType="text"
    android:maxLines="1" />

<EditText
    android:id="@+id/etPassword"
    android:layout_width="match_parent"
    android:layout_height="wrap_content"
    android:hint="请输入密码"
    android:inputType="textPassword" />
```

**Kotlin 代码示例**：
```kotlin
val etUsername: EditText = findViewById(R.id.etUsername)

// 获取输入内容
val username = etUsername.text.toString()

// 设置输入监听
etUsername.addTextChangedListener(object : TextWatcher {
    override fun beforeTextChanged(s: CharSequence?, start: Int, count: Int, after: Int) {}
    override fun onTextChanged(s: CharSequence?, start: Int, before: Int, count: Int) {
        // 实时监听输入变化
        Log.d("EditText", "当前输入: $s")
    }
    override fun afterTextChanged(s: Editable?) {}
})

// 设置输入过滤器（限制只能输入数字和字母）
etUsername.filters = arrayOf(InputFilter { source, start, end, dest, dstart, dend ->
    if (source.matches(Regex("[a-zA-Z0-9]*"))) source else ""
})
```

---

### 3. Button - 按钮

**说明**：用于用户点击交互，继承自 TextView。

**常用属性**：
- `text`：按钮文字
- `background`：背景（可设置为 selector 实现点击效果）
- `enabled`：是否可用

**XML 示例**：
```xml
<Button
    android:id="@+id/btnSubmit"
    android:layout_width="match_parent"
    android:layout_height="wrap_content"
    android:text="提交"
    android:textAllCaps="false" />
```

**Kotlin 代码示例**：
```kotlin
val btnSubmit: Button = findViewById(R.id.btnSubmit)

// 方式1：设置点击监听
btnSubmit.setOnClickListener {
    Toast.makeText(this, "按钮被点击", Toast.LENGTH_SHORT).show()
}

// 方式2：lambda 表达式
btnSubmit.setOnClickListener { view ->
    // 处理点击事件
    submitData()
}

// 禁用按钮
btnSubmit.isEnabled = false

// 启用按钮
btnSubmit.isEnabled = true
```

---

### 4. ImageView - 图片显示

**说明**：用于显示图片资源。

**常用属性**：
- `src`：图片资源
- `scaleType`：缩放类型（center, centerCrop, centerInside, fitXY, fitCenter, fitStart, fitEnd, matrix）
- `adjustViewBounds`：是否保持宽高比

**XML 示例**：
```xml
<ImageView
    android:id="@+id/ivAvatar"
    android:layout_width="100dp"
    android:layout_height="100dp"
    android:src="@drawable/ic_avatar"
    android:scaleType="centerCrop" />
```

**Kotlin 代码示例**：
```kotlin
val ivAvatar: ImageView = findViewById(R.id.ivAvatar)

// 设置资源图片
ivAvatar.setImageResource(R.drawable.ic_avatar)

// 使用 Glide 加载网络图片
Glide.with(this)
    .load("https://example.com/avatar.jpg")
    .placeholder(R.drawable.ic_placeholder)
    .error(R.drawable.ic_error)
    .circleCrop() // 圆形裁剪
    .into(ivAvatar)

// 使用 Coil 加载网络图片（Kotlin 推荐库）
ivAvatar.load("https://example.com/avatar.jpg") {
    crossfade(true)
    placeholder(R.drawable.ic_placeholder)
    transformations(CircleCropTransformation())
}
```

---

### 5. ImageButton - 图片按钮

**说明**：可点击的图片按钮，继承自 ImageView。

**XML 示例**：
```xml
<ImageButton
    android:id="@+id/ibPlay"
    android:layout_width="48dp"
    android:layout_height="48dp"
    android:src="@drawable/ic_play"
    android:background="?attr/selectableItemBackgroundBorderless"
    android:contentDescription="播放按钮" />
```

**Kotlin 代码示例**：
```kotlin
val ibPlay: ImageButton = findViewById(R.id.ibPlay)
ibPlay.setOnClickListener {
    // 播放/暂停逻辑
    togglePlayState()
}
```

---

### 6. CheckBox - 复选框

**说明**：用于多选场景，用户可选择多个选项。

**常用属性**：
- `checked`：是否选中
- `text`：复选框文字

**XML 示例**：
```xml
<CheckBox
    android:id="@+id/cbAgree"
    android:layout_width="wrap_content"
    android:layout_height="wrap_content"
    android:text="我同意用户协议" />

<CheckBox
    android:id="@+id/cbRemember"
    android:layout_width="wrap_content"
    android:layout_height="wrap_content"
    android:text="记住密码"
    android:checked="true" />
```

**Kotlin 代码示例**：
```kotlin
val cbAgree: CheckBox = findViewById(R.id.cbAgree)

// 获取选中状态
val isChecked = cbAgree.isChecked

// 设置选中状态
cbAgree.isChecked = true

// 监听选中状态变化
cbAgree.setOnCheckedChangeListener { buttonView, isChecked ->
    if (isChecked) {
        Toast.makeText(this, "已勾选", Toast.LENGTH_SHORT).show()
    } else {
        Toast.makeText(this, "已取消", Toast.LENGTH_SHORT).show()
    }
}
```

---

### 7. RadioButton & RadioGroup - 单选按钮

**说明**：RadioButton 用于单选场景，必须放在 RadioGroup 中使用。

**XML 示例**：
```xml
<RadioGroup
    android:id="@+id/rgGender"
    android:layout_width="match_parent"
    android:layout_height="wrap_content"
    android:orientation="vertical">

    <RadioButton
        android:id="@+id/rbMale"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="男"
        android:checked="true" />

    <RadioButton
        android:id="@+id/rbFemale"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="女" />
</RadioGroup>
```

**Kotlin 代码示例**：
```kotlin
val rgGender: RadioGroup = findViewById(R.id.rgGender)

// 方式1：获取选中的 RadioButton ID
val selectedId = rgGender.checkedRadioButtonId
val rbSelected: RadioButton = findViewById(selectedId)
val gender = rbSelected.text.toString()

// 方式2：监听选择变化
rgGender.setOnCheckedChangeListener { group, checkedId ->
    when (checkedId) {
        R.id.rbMale -> Log.d("RadioGroup", "选中：男")
        R.id.rbFemale -> Log.d("RadioGroup", "选中：女")
    }
}
```

---

### 8. Switch / ToggleButton - 开关

**说明**：用于开关切换，Switch 是 Material 风格，ToggleButton 是老式风格。

**XML 示例**：
```xml
<Switch
    android:id="@+id/swNotification"
    android:layout_width="wrap_content"
    android:layout_height="wrap_content"
    android:text="通知开关"
    android:checked="true" />

<ToggleButton
    android:id="@+id/tbWifi"
    android:layout_width="wrap_content"
    android:layout_height="wrap_content"
    android:textOn="开启"
    android:textOff="关闭" />
```

**Kotlin 代码示例**：
```kotlin
val swNotification: Switch = findViewById(R.id.swNotification)

// 获取开关状态
val isNotificationOn = swNotification.isChecked

// 设置开关状态
swNotification.isChecked = false

// 监听开关变化
swNotification.setOnCheckedChangeListener { buttonView, isChecked ->
    if (isChecked) {
        // 开启通知
        enableNotifications()
    } else {
        // 关闭通知
        disableNotifications()
    }
}
```

---

### 9. ProgressBar - 进度条

**说明**：用于显示加载进度。

**常用属性**：
- `style`：样式（horizontal 水平进度条，默认圆形）
- `max`：最大值
- `progress`：当前进度
- `indeterminate`：是否不确定进度

**XML 示例**：
```xml
<!-- 圆形进度条 -->
<ProgressBar
    android:id="@+id/pbLoading"
    android:layout_width="wrap_content"
    android:layout_height="wrap_content"
    android:visibility="gone" />

<!-- 水平进度条 -->
<ProgressBar
    android:id="@+id/pbDownload"
    style="@style/Widget.AppCompat.ProgressBar.Horizontal"
    android:layout_width="match_parent"
    android:layout_height="wrap_content"
    android:max="100"
    android:progress="30" />
```

**Kotlin 代码示例**：
```kotlin
val pbLoading: ProgressBar = findViewById(R.id.pbLoading)
val pbDownload: ProgressBar = findViewById(R.id.pbDownload)

// 显示/隐藏加载进度条
pbLoading.visibility = View.VISIBLE
pbLoading.visibility = View.GONE

// 设置水平进度条进度
pbDownload.progress = 50
pbDownload.max = 100

// 模拟下载进度
lifecycleScope.launch {
    for (i in 0..100) {
        delay(50)
        pbDownload.progress = i
    }
    Toast.makeText(this@MainActivity, "下载完成", Toast.LENGTH_SHORT).show()
}
```

---

### 10. SeekBar - 拖动条

**说明**：可拖动的进度条，常用于音量调节、亮度调节等。

**XML 示例**：
```xml
<SeekBar
    android:id="@+id/sbVolume"
    android:layout_width="match_parent"
    android:layout_height="wrap_content"
    android:max="100"
    android:progress="50" />
```

**Kotlin 代码示例**：
```kotlin
val sbVolume: SeekBar = findViewById(R.id.sbVolume)

// 设置进度
sbVolume.progress = 70

// 监听拖动
sbVolume.setOnSeekBarChangeListener(object : SeekBar.OnSeekBarChangeListener {
    override fun onProgressChanged(seekBar: SeekBar?, progress: Int, fromUser: Boolean) {
        // 进度变化
        Log.d("SeekBar", "当前进度: $progress")
    }

    override fun onStartTrackingTouch(seekBar: SeekBar?) {
        // 开始拖动
    }

    override fun onStopTrackingTouch(seekBar: SeekBar?) {
        // 停止拖动
        val volume = seekBar?.progress ?: 0
        setVolume(volume)
    }
})
```

---

### 11. RatingBar - 评分条

**说明**：用于星级评分。

**XML 示例**：
```xml
<RatingBar
    android:id="@+id/rbRating"
    android:layout_width="wrap_content"
    android:layout_height="wrap_content"
    android:numStars="5"
    android:rating="3.5"
    android:stepSize="0.5"
    android:isIndicator="false" />
```

**Kotlin 代码示例**：
```kotlin
val rbRating: RatingBar = findViewById(R.id.rbRating)

// 获取评分
val rating = rbRating.rating

// 设置评分
rbRating.rating = 4.5f

// 监听评分变化
rbRating.setOnRatingBarChangeListener { ratingBar, rating, fromUser ->
    Log.d("RatingBar", "新评分: $rating")
    submitRating(rating)
}
```

---

### 12. WebView - 网页视图

**说明**：用于在应用内显示网页内容。

**XML 示例**：
```xml
<WebView
    android:id="@+id/wvContent"
    android:layout_width="match_parent"
    android:layout_height="match_parent" />
```

**Kotlin 代码示例**：
```kotlin
val wvContent: WebView = findViewById(R.id.wvContent)

// 配置 WebView
wvContent.settings.javaScriptEnabled = true
wvContent.settings.domStorageEnabled = true
wvContent.settings.setSupportZoom(true)
wvContent.settings.builtInZoomControls = true

// 设置 WebViewClient（在当前 WebView 中打开链接）
wvContent.webViewClient = WebViewClient()

// 设置 WebChromeClient（处理标题、进度等）
wvContent.webChromeClient = object : WebChromeClient() {
    override fun onProgressChanged(view: WebView?, newProgress: Int) {
        // 显示加载进度
    }
}

// 加载网页
wvContent.loadUrl("https://www.baidu.com")

// 加载 HTML 内容
val htmlContent = "<html><body><h1>Hello WebView</h1></body></html>"
wvContent.loadDataWithBaseURL(null, htmlContent, "text/html", "UTF-8", null)

// 处理返回键
override fun onBackPressed() {
    if (wvContent.canGoBack()) {
        wvContent.goBack()
    } else {
        super.onBackPressed()
    }
}
```

---

## 布局容器

### 1. LinearLayout - 线性布局

**说明**：子控件按水平或垂直方向排列。

**常用属性**：
- `orientation`：排列方向（horizontal 水平，vertical 垂直）
- `gravity`：内部元素对齐方式
- `layout_weight`：权重，分配剩余空间

**XML 示例**：
```xml
<LinearLayout
    android:layout_width="match_parent"
    android:layout_height="wrap_content"
    android:orientation="horizontal"
    android:gravity="center_vertical">

    <TextView
        android:layout_width="0dp"
        android:layout_height="wrap_content"
        android:layout_weight="1"
        android:text="标题" />

    <Button
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="确定" />
</LinearLayout>
```

**Kotlin 代码创建**：
```kotlin
val linearLayout = LinearLayout(this).apply {
    orientation = LinearLayout.VERTICAL
    layoutParams = LinearLayout.LayoutParams(
        LinearLayout.LayoutParams.MATCH_PARENT,
        LinearLayout.LayoutParams.WRAP_CONTENT
    )
    gravity = Gravity.CENTER
}

// 添加子 View
val textView = TextView(this).apply {
    text = "动态添加的文本"
}
linearLayout.addView(textView)
```

---

### 2. ConstraintLayout - 约束布局

**说明**：功能最强大的布局，通过约束关系定位控件，推荐使用。

**常用属性**：
- `layout_constraintLeft_toLeftOf`：左边约束
- `layout_constraintRight_toRightOf`：右边约束
- `layout_constraintTop_toTopOf`：上边约束
- `layout_constraintBottom_toBottomOf`：下边约束
- `layout_constraintHorizontal_bias`：水平偏移（0-1）
- `layout_constraintVertical_bias`：垂直偏移（0-1）

**XML 示例**：
```xml
<androidx.constraintlayout.widget.ConstraintLayout
    android:layout_width="match_parent"
    android:layout_height="match_parent">

    <!-- 居中显示 -->
    <TextView
        android:id="@+id/tvTitle"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="居中标题"
        app:layout_constraintLeft_toLeftOf="parent"
        app:layout_constraintRight_toRightOf="parent"
        app:layout_constraintTop_toTopOf="parent"
        app:layout_constraintBottom_toBottomOf="parent" />

    <!-- 位于标题下方 -->
    <Button
        android:id="@+id/btnConfirm"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="确认"
        android:layout_marginTop="16dp"
        app:layout_constraintTop_toBottomOf="@id/tvTitle"
        app:layout_constraintLeft_toLeftOf="parent"
        app:layout_constraintRight_toRightOf="parent" />

</androidx.constraintlayout.widget.ConstraintLayout>
```

**Kotlin 代码创建约束**：
```kotlin
val constraintLayout = ConstraintLayout(this)

val textView = TextView(this).apply {
    id = View.generateViewId()
    text = "约束布局中的文本"
}

textView.layoutParams = ConstraintLayout.LayoutParams(
    ConstraintLayout.LayoutParams.WRAP_CONTENT,
    ConstraintLayout.LayoutParams.WRAP_CONTENT
).apply {
    // 约束到父布局中心
    leftToLeft = ConstraintLayout.LayoutParams.PARENT_ID
    rightToRight = ConstraintLayout.LayoutParams.PARENT_ID
    topToTop = ConstraintLayout.LayoutParams.PARENT_ID
    bottomToBottom = ConstraintLayout.LayoutParams.PARENT_ID
}

constraintLayout.addView(textView)
```

---

### 3. FrameLayout - 帧布局

**说明**：所有子控件默认叠加在左上角，适合单个控件或叠加效果。

**XML 示例**：
```xml
<FrameLayout
    android:layout_width="match_parent"
    android:layout_height="200dp">

    <ImageView
        android:layout_width="match_parent"
        android:layout_height="match_parent"
        android:src="@drawable/background"
        android:scaleType="centerCrop" />

    <TextView
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="叠加的文字"
        android:textColor="@android:color/white"
        android:layout_gravity="center" />
</FrameLayout>
```

---

### 4. RelativeLayout - 相对布局

**说明**：通过相对位置定位控件，较老的布局方式。

**XML 示例**：
```xml
<RelativeLayout
    android:layout_width="match_parent"
    android:layout_height="match_parent">

    <TextView
        android:id="@+id/tvTitle"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="标题"
        android:layout_centerInParent="true" />

    <Button
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="左边按钮"
        android:layout_toLeftOf="@id/tvTitle"
        android:layout_centerVertical="true" />

</RelativeLayout>
```

---

### 5. GridLayout - 网格布局

**说明**：将子控件按网格排列。

**XML 示例**：
```xml
<GridLayout
    android:layout_width="match_parent"
    android:layout_height="wrap_content"
    android:columnCount="3"
    android:rowCount="2">

    <Button android:text="1" />
    <Button android:text="2" />
    <Button android:text="3" />
    <Button android:text="4" />
    <Button android:text="5" />
    <Button android:text="6" />

</GridLayout>
```

---

## 列表控件

### 1. ListView - 列表视图

**说明**：用于显示垂直滚动列表，已被 RecyclerView 取代。

**XML 示例**：
```xml
<ListView
    android:id="@+id/lvItems"
    android:layout_width="match_parent"
    android:layout_height="match_parent"
    android:divider="@color/gray"
    android:dividerHeight="1dp" />
```

**Kotlin 代码示例**：
```kotlin
val lvItems: ListView = findViewById(R.id.lvItems)

// 数据源
val items = listOf("苹果", "香蕉", "橙子", "葡萄", "西瓜")

// 使用 ArrayAdapter
val adapter = ArrayAdapter(this, android.R.layout.simple_list_item_1, items)
lvItems.adapter = adapter

// 点击事件
lvItems.setOnItemClickListener { parent, view, position, id ->
    Toast.makeText(this, "点击了: ${items[position]}", Toast.LENGTH_SHORT).show()
}

// 长按事件
lvItems.setOnItemLongClickListener { parent, view, position, id ->
    Toast.makeText(this, "长按了: ${items[position]}", Toast.LENGTH_SHORT).show()
    true
}
```

---

### 2. RecyclerView - 可回收列表视图（推荐）

**说明**：功能强大的列表控件，支持多种布局类型，性能优秀，是 ListView 的替代品。

**依赖**：
```gradle
implementation "androidx.recyclerview:recyclerview:1.3.0"
```

**XML 示例**：
```xml
<androidx.recyclerview.widget.RecyclerView
    android:id="@+id/rvList"
    android:layout_width="match_parent"
    android:layout_height="match_parent" />
```

**Kotlin 代码示例（完整实现）**：

```kotlin
// 1. 数据类
data class User(
    val name: String,
    val avatar: String,
    val message: String
)

// 2. Adapter
class UserAdapter(private val userList: List<User>) : RecyclerView.Adapter<UserAdapter.UserViewHolder>() {

    // ViewHolder
    class UserViewHolder(itemView: View) : RecyclerView.ViewHolder(itemView) {
        val ivAvatar: ImageView = itemView.findViewById(R.id.ivAvatar)
        val tvName: TextView = itemView.findViewById(R.id.tvName)
        val tvMessage: TextView = itemView.findViewById(R.id.tvMessage)
    }

    // 创建 ViewHolder
    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): UserViewHolder {
        val view = LayoutInflater.from(parent.context)
            .inflate(R.layout.item_user, parent, false)
        return UserViewHolder(view)
    }

    // 绑定数据
    override fun onBindViewHolder(holder: UserViewHolder, position: Int) {
        val user = userList[position]
        holder.tvName.text = user.name
        holder.tvMessage.text = user.message
        // 使用 Glide 加载头像
        Glide.with(holder.itemView.context)
            .load(user.avatar)
            .into(holder.ivAvatar)
        
        // 点击事件
        holder.itemView.setOnClickListener {
            Log.d("RecyclerView", "点击了: ${user.name}")
        }
    }

    // 数据数量
    override fun getItemCount(): Int = userList.size
}

// 3. Activity 中使用
class MainActivity : AppCompatActivity() {
    
    private lateinit var rvList: RecyclerView
    private lateinit var userAdapter: UserAdapter
    
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
        
        rvList = findViewById(R.id.rvList)
        
        // 准备数据
        val users = listOf(
            User("张三", "https://example.com/avatar1.jpg", "你好！"),
            User("李四", "https://example.com/avatar2.jpg", "在吗？"),
            User("王五", "https://example.com/avatar3.jpg", "明天见")
        )
        
        // 创建 Adapter
        userAdapter = UserAdapter(users)
        
        // 设置 LayoutManager
        rvList.layoutManager = LinearLayoutManager(this) // 垂直列表
        // rvList.layoutManager = GridLayoutManager(this, 2) // 网格布局
        // rvList.layoutManager = StaggeredGridLayoutManager(2, StaggeredGridLayoutManager.VERTICAL) // 瀑布流
        
        // 设置 Adapter
        rvList.adapter = userAdapter
        
        // 添加分割线
        rvList.addItemDecoration(DividerItemDecoration(this, DividerItemDecoration.VERTICAL))
    }
}
```

**item_user.xml 布局文件**：
```xml
<?xml version="1.0" encoding="utf-8"?>
<androidx.constraintlayout.widget.ConstraintLayout
    xmlns:android="http://schemas.android.com/apk/res/android"
    xmlns:app="http://schemas.android.com/apk/res-auto"
    android:layout_width="match_parent"
    android:layout_height="wrap_content"
    android:padding="16dp">

    <ImageView
        android:id="@+id/ivAvatar"
        android:layout_width="50dp"
        android:layout_height="50dp"
        app:layout_constraintLeft_toLeftOf="parent"
        app:layout_constraintTop_toTopOf="parent" />

    <TextView
        android:id="@+id/tvName"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:textSize="16sp"
        android:textStyle="bold"
        android:layout_marginStart="12dp"
        app:layout_constraintLeft_toRightOf="@id/ivAvatar"
        app:layout_constraintTop_toTopOf="@id/ivAvatar" />

    <TextView
        android:id="@+id/tvMessage"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:textSize="14sp"
        android:layout_marginTop="4dp"
        app:layout_constraintLeft_toLeftOf="@id/tvName"
        app:layout_constraintTop_toBottomOf="@id/tvName" />

</androidx.constraintlayout.widget.ConstraintLayout>
```

---

### 3. Spinner - 下拉选择框

**说明**：下拉列表选择器。

**XML 示例**：
```xml
<Spinner
    android:id="@+id/spinner"
    android:layout_width="match_parent"
    android:layout_height="wrap_content" />
```

**Kotlin 代码示例**：
```kotlin
val spinner: Spinner = findViewById(R.id.spinner)

// 数据源
val cities = arrayOf("北京", "上海", "广州", "深圳")

// ArrayAdapter
val adapter = ArrayAdapter(this, android.R.layout.simple_spinner_item, cities)
adapter.setDropDownViewResource(android.R.layout.simple_spinner_dropdown_item)
spinner.adapter = adapter

// 选择监听
spinner.onItemSelectedListener = object : AdapterView.OnItemSelectedListener {
    override fun onItemSelected(parent: AdapterView<*>, view: View?, position: Int, id: Long) {
        val selectedCity = cities[position]
        Toast.makeText(this@MainActivity, "选中: $selectedCity", Toast.LENGTH_SHORT).show()
    }

    override fun onNothingSelected(parent: AdapterView<*>) {}
}
```

---

### 4. GridView - 网格视图

**说明**：网格布局展示数据，已被 RecyclerView 替代。

**XML 示例**：
```xml
<GridView
    android:id="@+id/gvImages"
    android:layout_width="match_parent"
    android:layout_height="match_parent"
    android:numColumns="3"
    android:horizontalSpacing="10dp"
    android:verticalSpacing="10dp" />
```

**Kotlin 代码示例**：
```kotlin
val gvImages: GridView = findViewById(R.id.gvImages)

// 数据源
val images = listOf(R.drawable.img1, R.drawable.img2, R.drawable.img3, R.drawable.img4)

// 使用自定义 Adapter
val adapter = ImageAdapter(this, images)
gvImages.adapter = adapter

// 点击事件
gvImages.onItemClickListener = AdapterView.OnItemClickListener { parent, view, position, id ->
    Toast.makeText(this, "点击了图片 $position", Toast.LENGTH_SHORT).show()
}
```

---

## Material Design 控件

### 1. TextInputLayout / TextInputEditText - 输入框容器

**说明**：Material 风格输入框，支持浮动标签、错误提示、密码可见切换。

**依赖**：
```gradle
implementation "com.google.android.material:material:1.9.0"
```

**XML 示例**：
```xml
<com.google.android.material.textfield.TextInputLayout
    android:id="@+id/tilUsername"
    android:layout_width="match_parent"
    android:layout_height="wrap_content"
    android:hint="用户名"
    app:startIconDrawable="@drawable/ic_person"
    app:endIconMode="clear_text"
    style="@style/Widget.MaterialComponents.TextInputLayout.OutlinedBox">

    <com.google.android.material.textfield.TextInputEditText
        android:id="@+id/etUsername"
        android:layout_width="match_parent"
        android:layout_height="wrap_content"
        android:inputType="text" />

</com.google.android.material.textfield.TextInputLayout>

<com.google.android.material.textfield.TextInputLayout
    android:id="@+id/tilPassword"
    android:layout_width="match_parent"
    android:layout_height="wrap_content"
    android:hint="密码"
    app:endIconMode="password_toggle"
    style="@style/Widget.MaterialComponents.TextInputLayout.OutlinedBox">

    <com.google.android.material.textfield.TextInputEditText
        android:id="@+id/etPassword"
        android:layout_width="match_parent"
        android:layout_height="wrap_content"
        android:inputType="textPassword" />

</com.google.android.material.textfield.TextInputLayout>
```

**Kotlin 代码示例**：
```kotlin
val tilUsername: TextInputLayout = findViewById(R.id.tilUsername)
val etUsername: TextInputEditText = findViewById(R.id.etUsername)

// 获取输入内容
val username = etUsername.text.toString()

// 显示错误
tilUsername.error = "用户名不能为空"

// 清除错误
tilUsername.error = null
tilUsername.isErrorEnabled = false

// 设置字符计数
tilUsername.counterEnabled = true
tilUsername.counterMaxLength = 20
```

---

### 2. FloatingActionButton - 浮动按钮

**说明**：Material Design 风格的悬浮按钮。

**XML 示例**：
```xml
<com.google.android.material.floatingactionbutton.FloatingActionButton
    android:id="@+id/fabAdd"
    android:layout_width="wrap_content"
    android:layout_height="wrap_content"
    android:src="@drawable/ic_add"
    android:contentDescription="添加"
    app:backgroundTint="@color/purple_500"
    app:tint="@android:color/white"
    app:fabSize="normal" />
```

**Kotlin 代码示例**：
```kotlin
val fabAdd: FloatingActionButton = findViewById(R.id.fabAdd)

fabAdd.setOnClickListener {
    // 添加操作
    showAddDialog()
}

// 动态隐藏/显示
fabAdd.hide()
fabAdd.show()
```

---

### 3. CardView - 卡片视图

**说明**：Material 风格卡片容器，带阴影和圆角。

**依赖**：
```gradle
implementation "androidx.cardview:cardview:1.0.0"
```

**XML 示例**：
```xml
<androidx.cardview.widget.CardView
    android:layout_width="match_parent"
    android:layout_height="wrap_content"
    app:cardCornerRadius="8dp"
    app:cardElevation="4dp"
    app:cardUseCompatPadding="true"
    app:contentPadding="16dp">

    <LinearLayout
        android:layout_width="match_parent"
        android:layout_height="wrap_content"
        android:orientation="vertical">

        <TextView
            android:layout_width="wrap_content"
            android:layout_height="wrap_content"
            android:text="卡片标题"
            android:textSize="18sp"
            android:textStyle="bold" />

        <TextView
            android:layout_width="wrap_content"
            android:layout_height="wrap_content"
            android:text="卡片内容描述"
            android:layout_marginTop="8dp" />

    </LinearLayout>

</androidx.cardview.widget.CardView>
```

---

### 4. BottomNavigationView - 底部导航栏

**说明**：底部导航栏，用于主要功能切换。

**XML 示例**：
```xml
<com.google.android.material.bottomnavigation.BottomNavigationView
    android:id="@+id/bottomNav"
    android:layout_width="match_parent"
    android:layout_height="wrap_content"
    app:menu="@menu/bottom_nav_menu"
    app:labelVisibilityMode="labeled" />
```

**menu/bottom_nav_menu.xml**：
```xml
<?xml version="1.0" encoding="utf-8"?>
<menu xmlns:android="http://schemas.android.com/apk/res/android">
    <item
        android:id="@+id/nav_home"
        android:icon="@drawable/ic_home"
        android:title="首页" />
    <item
        android:id="@+id/nav_message"
        android:icon="@drawable/ic_message"
        android:title="消息" />
    <item
        android:id="@+id/nav_profile"
        android:icon="@drawable/ic_person"
        android:title="我的" />
</menu>
```

**Kotlin 代码示例**：
```kotlin
val bottomNav: BottomNavigationView = findViewById(R.id.bottomNav)

bottomNav.setOnItemSelectedListener { item ->
    when (item.itemId) {
        R.id.nav_home -> {
            // 切换到首页
            true
        }
        R.id.nav_message -> {
            // 切换到消息
            true
        }
        R.id.nav_profile -> {
            // 切换到我的
            true
        }
        else -> false
    }
}

// 设置选中项
bottomNav.selectedItemId = R.id.nav_home

// 获取选中项
val selectedId = bottomNav.selectedItemId
```

---

### 5. TabLayout - 标签栏

**说明**：标签导航栏，常与 ViewPager2 配合使用。

**XML 示例**：
```xml
<com.google.android.material.tabs.TabLayout
    android:id="@+id/tabLayout"
    android:layout_width="match_parent"
    android:layout_height="wrap_content"
    app:tabMode="scrollable"
    app:tabIndicatorColor="@color/purple_500"
    app:tabTextColor="@color/gray"
    app:tabSelectedTextColor="@color/purple_500" />

<androidx.viewpager2.widget.ViewPager2
    android:id="@+id/viewPager"
    android:layout_width="match_parent"
    android:layout_height="match_parent" />
```

**Kotlin 代码示例**：
```kotlin
val tabLayout: TabLayout = findViewById(R.id.tabLayout)
val viewPager: ViewPager2 = findViewById(R.id.viewPager)

// 添加标签
val tabs = listOf("首页", "消息", "我的")
tabs.forEach { title ->
    tabLayout.addTab(tabLayout.newTab().setText(title))
}

// 设置 ViewPager
val adapter = ViewPagerAdapter(this)
viewPager.adapter = adapter

// 关联 TabLayout 和 ViewPager
TabLayoutMediator(tabLayout, viewPager) { tab, position ->
    tab.text = tabs[position]
}.attach()

// 标签选择监听
tabLayout.addOnTabSelectedListener(object : TabLayout.OnTabSelectedListener {
    override fun onTabSelected(tab: TabLayout.Tab?) {
        Log.d("TabLayout", "选中: ${tab?.text}")
    }

    override fun onTabUnselected(tab: TabLayout.Tab?) {}

    override fun onTabReselected(tab: TabLayout.Tab?) {}
})
```

---

### 6. Snackbar - 提示条

**说明**：Material 风格的底部提示条，可带操作按钮。

**Kotlin 代码示例**：
```kotlin
// 简单提示
Snackbar.make(view, "操作成功", Snackbar.LENGTH_SHORT).show()

// 带操作按钮
Snackbar.make(view, "文件已删除", Snackbar.LENGTH_LONG)
    .setAction("撤销") {
        // 撤销操作
        restoreFile()
    }
    .show()

// 自定义样式
val snackbar = Snackbar.make(view, "自定义样式", Snackbar.LENGTH_INDEFINITE)
snackbar.setAction("确定") { snackbar.dismiss() }
snackbar.setActionTextColor(Color.RED)
snackbar.view.setBackgroundColor(Color.BLUE)
snackbar.show()
```

---

## 高级控件

### 1. ViewPager2 - 滑动页面容器

**说明**：用于实现左右滑动的页面切换，替代老旧的 ViewPager。

**依赖**：
```gradle
implementation "androidx.viewpager2:viewpager2:1.0.0"
```

**XML 示例**：
```xml
<androidx.viewpager2.widget.ViewPager2
    android:id="@+id/viewPager"
    android:layout_width="match_parent"
    android:layout_height="match_parent" />
```

**Kotlin 代码示例**：
```kotlin
// Adapter
class ViewPagerAdapter(activity: FragmentActivity) : FragmentStateAdapter(activity) {
    override fun getItemCount(): Int = 3

    override fun createFragment(position: Int): Fragment {
        return when (position) {
            0 -> HomeFragment()
            1 -> MessageFragment()
            else -> ProfileFragment()
        }
    }
}

// 使用
val viewPager: ViewPager2 = findViewById(R.id.viewPager)
viewPager.adapter = ViewPagerAdapter(this)

// 监听页面切换
viewPager.registerOnPageChangeCallback(object : ViewPager2.OnPageChangeCallback() {
    override fun onPageSelected(position: Int) {
        Log.d("ViewPager2", "当前页面: $position")
    }
})

// 切换页面
viewPager.currentItem = 1

// 禁用滑动
viewPager.isUserInputEnabled = false
```

---

### 2. DrawerLayout - 抽屉布局

**说明**：侧滑菜单布局。

**XML 示例**：
```xml
<androidx.drawerlayout.widget.DrawerLayout
    android:id="@+id/drawerLayout"
    android:layout_width="match_parent"
    android:layout_height="match_parent">

    <!-- 主内容区域 -->
    <FrameLayout
        android:layout_width="match_parent"
        android:layout_height="match_parent">
        
        <!-- 主界面内容 -->
        
    </FrameLayout>

    <!-- 侧滑菜单 -->
    <com.google.android.material.navigation.NavigationView
        android:id="@+id/navView"
        android:layout_width="wrap_content"
        android:layout_height="match_parent"
        android:layout_gravity="start"
        app:menu="@menu/drawer_menu"
        app:headerLayout="@layout/nav_header" />

</androidx.drawerlayout.widget.DrawerLayout>
```

**Kotlin 代码示例**：
```kotlin
val drawerLayout: DrawerLayout = findViewById(R.id.drawerLayout)
val navView: NavigationView = findViewById(R.id.navView)

// 打开抽屉
drawerLayout.openDrawer(GravityCompat.START)

// 关闭抽屉
drawerLayout.closeDrawer(GravityCompat.START)

// 菜单项点击
navView.setNavigationItemSelectedListener { item ->
    when (item.itemId) {
        R.id.nav_home -> {
            // 首页
        }
        R.id.nav_settings -> {
            // 设置
        }
    }
    drawerLayout.closeDrawer(GravityCompat.START)
    true
}

// 处理返回键
override fun onBackPressed() {
    if (drawerLayout.isDrawerOpen(GravityCompat.START)) {
        drawerLayout.closeDrawer(GravityCompat.START)
    } else {
        super.onBackPressed()
    }
}
```

---

### 3. SwipeRefreshLayout - 下拉刷新

**说明**：下拉刷新容器。

**XML 示例**：
```xml
<androidx.swiperefreshlayout.widget.SwipeRefreshLayout
    android:id="@+id/swipeRefresh"
    android:layout_width="match_parent"
    android:layout_height="match_parent">

    <androidx.recyclerview.widget.RecyclerView
        android:id="@+id/rvList"
        android:layout_width="match_parent"
        android:layout_height="match_parent" />

</androidx.swiperefreshlayout.widget.SwipeRefreshLayout>
```

**Kotlin 代码示例**：
```kotlin
val swipeRefresh: SwipeRefreshLayout = findViewById(R.id.swipeRefresh)

// 设置刷新颜色
swipeRefresh.setColorSchemeResources(R.color.purple_500, R.color.teal_200)

// 下拉刷新监听
swipeRefresh.setOnRefreshListener {
    // 模拟刷新
    lifecycleScope.launch {
        delay(2000)
        // 刷新数据
        refreshData()
        // 停止刷新动画
        swipeRefresh.isRefreshing = false
    }
}

// 手动触发刷新
swipeRefresh.isRefreshing = true
```

---

### 4. NestedScrollView - 嵌套滚动视图

**说明**：支持嵌套滚动的 ScrollView，可与 CoordinatorLayout 配合使用。

**XML 示例**：
```xml
<androidx.core.widget.NestedScrollView
    android:layout_width="match_parent"
    android:layout_height="match_parent">

    <LinearLayout
        android:layout_width="match_parent"
        android:layout_height="wrap_content"
        android:orientation="vertical">
        
        <!-- 内容 -->
        
    </LinearLayout>

</androidx.core.widget.NestedScrollView>
```

---

### 5. SearchView - 搜索框

**说明**：搜索输入控件。

**XML 示例**：
```xml
<SearchView
    android:id="@+id/searchView"
    android:layout_width="match_parent"
    android:layout_height="wrap_content"
    android:iconifiedByDefault="false"
    android:queryHint="搜索..." />
```

**Kotlin 代码示例**：
```kotlin
val searchView: SearchView = findViewById(R.id.searchView)

// 搜索监听
searchView.setOnQueryTextListener(object : SearchView.OnQueryTextListener {
    override fun onQueryTextSubmit(query: String?): Boolean {
        // 提交搜索
        performSearch(query)
        return true
    }

    override fun onQueryTextChange(newText: String?): Boolean {
        // 实时搜索
        filterList(newText)
        return true
    }
})

// 设置搜索提示
searchView.queryHint = "请输入关键词"

// 获取搜索内容
val query = searchView.query.toString()
```

---

### 6. Toolbar - 工具栏

**说明**：替代 ActionBar 的工具栏，更灵活。

**XML 示例**：
```xml
<androidx.appcompat.widget.Toolbar
    android:id="@+id/toolbar"
    android:layout_width="match_parent"
    android:layout_height="?attr/actionBarSize"
    android:background="@color/purple_500"
    app:title="标题"
    app:subtitle="副标题"
    app:navigationIcon="@drawable/ic_back"
    app:titleTextColor="@android:color/white"
    app:subtitleTextColor="@android:color/white" />
```

**Kotlin 代码示例**：
```kotlin
val toolbar: Toolbar = findViewById(R.id.toolbar)

// 设置为 ActionBar
setSupportActionBar(toolbar)

// 设置标题
toolbar.title = "新标题"
toolbar.subtitle = "副标题"

// 导航按钮点击
toolbar.setNavigationOnClickListener {
    onBackPressed()
}

// 添加菜单
toolbar.inflateMenu(R.menu.toolbar_menu)

// 菜单点击
toolbar.setOnMenuItemClickListener { item ->
    when (item.itemId) {
        R.id.action_search -> {
            // 搜索
            true
        }
        R.id.action_settings -> {
            // 设置
            true
        }
        else -> false
    }
}
```

---

### 7. DatePicker / DatePickerDialog - 日期选择

**Kotlin 代码示例**：
```kotlin
// 使用 DatePickerDialog
val calendar = Calendar.getInstance()
val year = calendar.get(Calendar.YEAR)
val month = calendar.get(Calendar.MONTH)
val day = calendar.get(Calendar.DAY_OF_MONTH)

val datePickerDialog = DatePickerDialog(
    this,
    { _, selectedYear, selectedMonth, selectedDay ->
        val date = "$selectedYear-${selectedMonth + 1}-$selectedDay"
        Toast.makeText(this, "选中日期: $date", Toast.LENGTH_SHORT).show()
    },
    year, month, day
)
datePickerDialog.show()

// 设置日期范围
datePickerDialog.datePicker.minDate = calendar.timeInMillis // 最小日期
datePickerDialog.datePicker.maxDate = calendar.timeInMillis + 365L * 24 * 60 * 60 * 1000 // 最大日期
```

---

### 8. TimePicker / TimePickerDialog - 时间选择

**Kotlin 代码示例**：
```kotlin
val calendar = Calendar.getInstance()
val hour = calendar.get(Calendar.HOUR_OF_DAY)
val minute = calendar.get(Calendar.MINUTE)

val timePickerDialog = TimePickerDialog(
    this,
    { _, selectedHour, selectedMinute ->
        val time = String.format("%02d:%02d", selectedHour, selectedMinute)
        Toast.makeText(this, "选中时间: $time", Toast.LENGTH_SHORT).show()
    },
    hour, minute, true // 24小时制
)
timePickerDialog.show()
```

---

## 对话框

### 1. AlertDialog - 提示对话框

**Kotlin 代码示例**：
```kotlin
// 简单提示对话框
AlertDialog.Builder(this)
    .setTitle("提示")
    .setMessage("确定要删除吗？")
    .setPositiveButton("确定") { dialog, which ->
        // 确定操作
        deleteItem()
    }
    .setNegativeButton("取消", null)
    .show()

// 列表选择对话框
val items = arrayOf("选项1", "选项2", "选项3")
AlertDialog.Builder(this)
    .setTitle("请选择")
    .setItems(items) { dialog, which ->
        Toast.makeText(this, "选中: ${items[which]}", Toast.LENGTH_SHORT).show()
    }
    .show()

// 单选对话框
var selectedIndex = 0
AlertDialog.Builder(this)
    .setTitle("单选")
    .setSingleChoiceItems(items, selectedIndex) { dialog, which ->
        selectedIndex = which
    }
    .setPositiveButton("确定") { dialog, which ->
        Toast.makeText(this, "选中: ${items[selectedIndex]}", Toast.LENGTH_SHORT).show()
    }
    .show()

// 多选对话框
val selectedItems = booleanArrayOf(false, false, false)
AlertDialog.Builder(this)
    .setTitle("多选")
    .setMultiChoiceItems(items, selectedItems) { dialog, which, isChecked ->
        selectedItems[which] = isChecked
    }
    .setPositiveButton("确定") { dialog, which ->
        val selected = items.filterIndexed { index, _ -> selectedItems[index] }
        Toast.makeText(this, "选中: $selected", Toast.LENGTH_SHORT).show()
    }
    .show()

// 自定义布局对话框
val view = layoutInflater.inflate(R.layout.dialog_custom, null)
AlertDialog.Builder(this)
    .setView(view)
    .setPositiveButton("确定") { dialog, which ->
        // 处理自定义布局中的控件
    }
    .show()
```

---

### 2. ProgressDialog - 进度对话框（已废弃，建议自定义）

```kotlin
// 注意：ProgressDialog 已废弃，建议使用自定义布局或 ProgressBar
// 旧用法（不推荐）
val progressDialog = ProgressDialog(this).apply {
    setMessage("加载中...")
    setCancelable(false)
}
progressDialog.show()
progressDialog.dismiss()

// 推荐方案：使用自定义对话框
```

---

### 3. BottomSheetDialog - 底部滑出对话框

**依赖**：
```gradle
implementation "com.google.android.material:material:1.9.0"
```

**Kotlin 代码示例**：
```kotlin
val bottomSheetDialog = BottomSheetDialog(this)
val view = layoutInflater.inflate(R.layout.bottom_sheet_layout, null)
bottomSheetDialog.setContentView(view)
bottomSheetDialog.show()

// 处理底部对话框中的控件
view.findViewById<Button>(R.id.btnConfirm).setOnClickListener {
    // 处理点击
    bottomSheetDialog.dismiss()
}
```

**bottom_sheet_layout.xml**：
```xml
<?xml version="1.0" encoding="utf-8"?>
<LinearLayout
    xmlns:android="http://schemas.android.com/apk/res/android"
    android:layout_width="match_parent"
    android:layout_height="wrap_content"
    android:orientation="vertical"
    android:padding="16dp">

    <TextView
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="底部对话框"
        android:textSize="18sp"
        android:textStyle="bold" />

    <Button
        android:id="@+id/btnConfirm"
        android:layout_width="match_parent"
        android:layout_height="wrap_content"
        android:text="确认"
        android:layout_marginTop="16dp" />

</LinearLayout>
```

---

## 常用属性速查表

### 尺寸单位
- `dp`（density-independent pixels）：密度无关像素，用于控件尺寸
- `sp`（scale-independent pixels）：缩放无关像素，用于字体大小
- `px`：像素，一般不推荐使用

### 宽高属性
- `match_parent`：填满父容器
- `wrap_content`：根据内容自适应
- 具体数值：如 `100dp`

### 常用属性

| 属性 | 说明 | 示例 |
|------|------|------|
| `android:id` | 控件ID | `@+id/btnSubmit` |
| `android:layout_width` | 宽度 | `match_parent`, `wrap_content`, `100dp` |
| `android:layout_height` | 高度 | `match_parent`, `wrap_content`, `100dp` |
| `android:text` | 文本内容 | `"确定"` |
| `android:textSize` | 文字大小 | `16sp` |
| `android:textColor` | 文字颜色 | `#333333`, `@color/black` |
| `android:textStyle` | 文字样式 | `normal`, `bold`, `italic` |
| `android:background` | 背景 | `@color/white`, `@drawable/bg_button` |
| `android:src` | 图片资源 | `@drawable/ic_launcher` |
| `android:visibility` | 可见性 | `visible`, `invisible`, `gone` |
| `android:enabled` | 是否可用 | `true`, `false` |
| `android:padding` | 内边距 | `16dp` |
| `android:layout_margin` | 外边距 | `16dp` |
| `android:gravity` | 内容对齐 | `center`, `left`, `right` |
| `android:layout_gravity` | 控件对齐 | `center`, `left`, `right` |

### 可见性说明
- `visible`：可见（默认）
- `invisible`：不可见，但占用空间
- `gone`：不可见，不占用空间

**Kotlin 代码中设置**：
```kotlin
view.visibility = View.VISIBLE
view.visibility = View.INVISIBLE
view.visibility = View.GONE
```

---

## 常用第三方库推荐

### 图片加载
```gradle
// Glide
implementation "com.github.bumptech.glide:glide:4.15.1"

// Coil（Kotlin 推荐）
implementation "io.coil-kt:coil:2.4.0"
```

### 网络
```gradle
// Retrofit + OkHttp
implementation "com.squareup.retrofit2:retrofit:2.9.0"
implementation "com.squareup.retrofit2:converter-gson:2.9.0"
implementation "com.squareup.okhttp3:okhttp:4.11.0"
```

### 异步处理
```gradle
// Kotlin Coroutines
implementation "org.jetbrains.kotlinx:kotlinx-coroutines-android:1.7.1"
```

### JSON 解析
```gradle
// Gson
implementation "com.google.code.gson:gson:2.10.1"

// Moshi
implementation "com.squareup.moshi:moshi:1.15.0"
```

---

## 最佳实践

### 1. 使用 ViewBinding 替代 findViewById

```kotlin
// build.gradle
android {
    buildFeatures {
        viewBinding = true
    }
}

// Activity
class MainActivity : AppCompatActivity() {
    private lateinit var binding: ActivityMainBinding

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        binding = ActivityMainBinding.inflate(layoutInflater)
        setContentView(binding.root)
        
        binding.tvTitle.text = "Hello ViewBinding"
        binding.btnSubmit.setOnClickListener {
            // 处理点击
        }
    }
}
```

### 2. 使用 Data Binding（进阶）

```kotlin
// build.gradle
android {
    buildFeatures {
        dataBinding = true
    }
}

// XML
<layout xmlns:android="http://schemas.android.com/apk/res/android">
    <data>
        <variable
            name="user"
            type="com.example.User" />
    </data>
    
    <LinearLayout
        android:layout_width="match_parent"
        android:layout_height="match_parent"
        android:orientation="vertical">
        
        <TextView
            android:layout_width="wrap_content"
            android:layout_height="wrap_content"
            android:text="@{user.name}" />
            
    </LinearLayout>
</layout>
```

### 3. 使用扩展函数简化代码

```kotlin
// 扩展函数
fun View.show() {
    visibility = View.VISIBLE
}

fun View.hide() {
    visibility = View.GONE
}

fun View.invisible() {
    visibility = View.INVISIBLE
}

// 使用
binding.progressBar.show()
binding.progressBar.hide()
```

### 4. 避免内存泄漏

```kotlin
// 使用 lifecycleScope 替代 GlobalScope
lifecycleScope.launch {
    // 协程会随 Activity 销毁自动取消
    val data = repository.getData()
    updateUI(data)
}

// 在 Fragment 中使用 viewLifecycleOwner.lifecycleScope
viewLifecycleOwner.lifecycleScope.launch {
    // 协程会随 Fragment View 销毁自动取消
}
```

---

## 总结

本文档整理了 Android Studio Kotlin 开发中最常用的控件，包括：

✅ **基础控件**：TextView, EditText, Button, ImageView, CheckBox, RadioButton, Switch 等  
✅ **布局容器**：LinearLayout, ConstraintLayout, FrameLayout, RelativeLayout, GridLayout  
✅ **列表控件**：ListView, RecyclerView, Spinner, GridView  
✅ **Material Design**：TextInputLayout, FloatingActionButton, CardView, BottomNavigationView, TabLayout  
✅ **高级控件**：ViewPager2, DrawerLayout, SwipeRefreshLayout, Toolbar, DatePicker  
✅ **对话框**：AlertDialog, BottomSheetDialog  

建议：
1. 新项目优先使用 **ConstraintLayout** 和 **RecyclerView**
2. 使用 **ViewBinding** 或 **DataBinding** 提高开发效率
3. Material Design 控件提供更好的用户体验
4. 善用扩展函数简化代码

---

> 整理者：尘曦 🌅  
> 最后更新：2026-03-12
