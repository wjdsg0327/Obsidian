# Apex Dashboard 黑曜主题设计

## 目标

为 Obsidian 插件 `apex-dashboard` 新增一个独立主题预设：

- 内部键名：`obsidian`
- 中文名称：`黑曜`
- 英文名称：`Obsidian`
- 主题始终保持暗色，不随 Obsidian 的明暗模式切换
- 仅加入主题列表，不自动替换用户当前的 `tundra` 主题

## 视觉方向

黑曜主题使用接近纯黑的背景与柔和紫色强调色，保持 Apex Dashboard
现有玻璃拟态结构：

- 页面背景：近黑渐变，并加入低透明度紫色环境光
- 卡片和侧栏：半透明深灰紫表面，配合轻微模糊与阴影
- 文字：高对比灰白正文、柔和灰紫次级文字
- 强调色：紫罗兰到浅紫渐变
- 边框：低透明度紫灰边框
- 交互：悬停、聚焦、进度条、复选框和链接统一使用紫色体系

视觉效果应清晰、柔和，适合长时间使用，同时与现有 `carbon`
主题的工业黑白风格保持明显区别。

## 实现范围

### JavaScript 注册

在编译后的 `main.js` 中完成以下注册：

1. 英文翻译表增加 `settings.styleObsidian: "Obsidian"`。
2. 中文翻译表增加 `settings.styleObsidian: "黑曜"`。
3. 设置页主题下拉框增加 `obsidian` 选项。
4. `cycle-theme` 命令的主题顺序增加 `obsidian`。

不修改默认主题，不修改现有用户设置。

### CSS 主题

在 `styles.css` 中增加
`.apex-dashboard-root[data-theme="obsidian"]` 主题块，定义完整的
`--db-*` 变量。该选择器不依赖 `.theme-light` 或 `.theme-dark`，
因此无论 Obsidian 当前处于何种模式，仪表盘都保持暗色。

针对黑曜主题补充必要的组件样式：

- 页面背景和环境光渐变
- 卡片、区块与侧栏玻璃表面
- 卡片悬停与输入框聚焦
- 新增区块按钮和拖放指示
- 弹窗表面、输入框及主要操作按钮

尽量复用现有组件变量，避免复制所有组件规则。

## 兼容与数据

- `data.json` 保持不变，当前 `stylePreset: "tundra"` 不变。
- 已有主题名称、键名和样式保持不变。
- 新主题可从设置下拉框选择，也可通过“切换到下一个主题”命令进入。
- 插件升级可能覆盖直接修改的编译产物；本次修改仅针对当前知识库中的已安装插件。

## 验证

1. 使用 `node --check main.js` 验证 JavaScript 语法。
2. 检查英文、中文、下拉框和循环命令四个注册点均包含 `obsidian`。
3. 检查 CSS 包含独立的 `data-theme="obsidian"` 变量和组件规则。
4. 检查 `data.json` 仍为 `stylePreset: "tundra"`。
5. 在 Obsidian 中重新加载插件后，确认设置列表出现“黑曜”，选择后仪表盘与弹窗均为暗黑紫色。

