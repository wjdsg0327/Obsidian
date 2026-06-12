# Apex Dashboard 黑曜主题 Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 在 Apex Dashboard 中新增始终暗色的「黑曜 / Obsidian」主题预设，同时保留当前 `tundra` 设置和全部已有主题。

**Architecture:** 直接扩展已安装插件的编译产物。`main.js` 负责主题名称、设置下拉框和循环主题命令的注册，`styles.css` 通过独立的 `data-theme="obsidian"` 变量块与少量组件增强实现暗黑紫色外观；不引入新的运行时依赖或配置字段。

**Tech Stack:** Obsidian Plugin API 编译产物、JavaScript、CSS 自定义属性、PowerShell 静态断言、Node.js 语法检查

---

## 文件结构

- Modify: `.obsidian/plugins/apex-dashboard/main.js`：注册中英文主题名、下拉选项和循环顺序。
- Modify: `.obsidian/plugins/apex-dashboard/styles.css`：定义黑曜主题变量、背景及必要组件状态。
- Verify: `.obsidian/plugins/apex-dashboard/data.json`：确认当前主题仍为 `tundra`，不写入该文件。

### Task 1: 建立注册点失败检查

**Files:**
- Test: `.obsidian/plugins/apex-dashboard/main.js`
- Test: `.obsidian/plugins/apex-dashboard/styles.css`
- Verify: `.obsidian/plugins/apex-dashboard/data.json`

- [ ] **Step 1: 运行主题注册前置断言**

Run:

```powershell
$main = Get-Content -Raw '.obsidian/plugins/apex-dashboard/main.js'
$css = Get-Content -Raw '.obsidian/plugins/apex-dashboard/styles.css'
$data = Get-Content -Raw '.obsidian/plugins/apex-dashboard/data.json' | ConvertFrom-Json

$checks = [ordered]@{
  EnglishLabel = $main.Contains('"settings.styleObsidian":"Obsidian"')
  ChineseLabel = $main.Contains('"settings.styleObsidian":"\u9ED1\u66DC"')
  Dropdown = $main.Contains('obsidian:b("settings.styleObsidian")')
  CycleTheme = $main.Contains('"carbon","obsidian"]')
  CssTheme = $css.Contains('[data-theme="obsidian"]')
  CurrentThemePreserved = $data.stylePreset -eq 'tundra'
}

$checks.GetEnumerator() | ForEach-Object {
  '{0}: {1}' -f $_.Key, $_.Value
}
```

Expected:

```text
EnglishLabel: False
ChineseLabel: False
Dropdown: False
CycleTheme: False
CssTheme: False
CurrentThemePreserved: True
```

- [ ] **Step 2: 确认修改目标没有未提交的既有改动**

Run:

```powershell
git status --short -- `
  '.obsidian/plugins/apex-dashboard/main.js' `
  '.obsidian/plugins/apex-dashboard/styles.css' `
  '.obsidian/plugins/apex-dashboard/data.json'
```

Expected: no output. If output exists, inspect the diff and preserve the user's changes before proceeding.

### Task 2: 注册黑曜主题

**Files:**
- Modify: `.obsidian/plugins/apex-dashboard/main.js:6`
- Modify: `.obsidian/plugins/apex-dashboard/main.js:172`

- [ ] **Step 1: 增加英文翻译键**

In the English translation object, replace:

```javascript
"settings.styleCarbon":"Eclipse","settings.widgetTheme"
```

with:

```javascript
"settings.styleCarbon":"Eclipse","settings.styleObsidian":"Obsidian","settings.widgetTheme"
```

- [ ] **Step 2: 增加中文翻译键**

In the Chinese translation object, replace:

```javascript
"settings.styleCarbon":"\u65E5\u98DF","settings.widgetTheme"
```

with:

```javascript
"settings.styleCarbon":"\u65E5\u98DF","settings.styleObsidian":"\u9ED1\u66DC","settings.widgetTheme"
```

- [ ] **Step 3: 将黑曜加入设置下拉框**

Replace:

```javascript
carbon:b("settings.styleCarbon")
```

with:

```javascript
carbon:b("settings.styleCarbon"),obsidian:b("settings.styleObsidian")
```

Expected: the settings dropdown lists `黑曜` in Chinese and `Obsidian` in English.

- [ ] **Step 4: 将黑曜加入循环主题命令**

Replace:

```javascript
["earth","nordic","aurora","prism","island","tundra","blossom","matcha","lilac","haze","ember","jade","carbon"]
```

with:

```javascript
["earth","nordic","aurora","prism","island","tundra","blossom","matcha","lilac","haze","ember","jade","carbon","obsidian"]
```

- [ ] **Step 5: 验证 JavaScript 注册与语法**

Run:

```powershell
node --check '.obsidian/plugins/apex-dashboard/main.js'

$main = Get-Content -Raw '.obsidian/plugins/apex-dashboard/main.js'
@(
  $main.Contains('"settings.styleObsidian":"Obsidian"'),
  $main.Contains('"settings.styleObsidian":"\u9ED1\u66DC"'),
  $main.Contains('obsidian:b("settings.styleObsidian")'),
  $main.Contains('"carbon","obsidian"]')
) | ForEach-Object {
  if (-not $_) { throw 'Missing Obsidian theme registration' }
}
```

Expected: `node --check` exits with code 0 and the PowerShell command prints no error.

- [ ] **Step 6: Commit JavaScript registration**

```powershell
git add -- '.obsidian/plugins/apex-dashboard/main.js'
git commit -m 'feat: register Apex Dashboard Obsidian theme'
```

### Task 3: 实现黑曜主题 CSS

**Files:**
- Modify: `.obsidian/plugins/apex-dashboard/styles.css:1424`

- [ ] **Step 1: 在组件样式区之前添加主题变量块**

Insert immediately before the `COMPONENT STYLES` section:

```css
/* ----- 16. OBSIDIAN (黑曜) — Always-dark violet glass ----- */
.apex-dashboard-root[data-theme="obsidian"] {
	--db-bg: #09080d;
	--db-bg-section: rgba(20, 17, 28, 0.44);
	--db-bg-card: rgba(28, 24, 38, 0.72);
	--db-bg-card-hover: rgba(39, 32, 54, 0.86);
	--db-bg-sidebar: rgba(15, 13, 21, 0.78);
	--db-bg-input: rgba(22, 19, 30, 0.82);
	--db-bg-btn: rgba(139, 92, 246, 0.12);
	--db-bg-btn-hover: rgba(139, 92, 246, 0.22);
	--db-bg-banner: rgba(24, 20, 33, 0.70);
	--db-bg-overlay: linear-gradient(135deg, rgba(7, 5, 11, 0.76), rgba(50, 31, 78, 0.32));
	--db-bg-hover: rgba(167, 139, 250, 0.08);
	--db-bg-hover-strong: rgba(167, 139, 250, 0.14);
	--db-bg-add-section: rgba(139, 92, 246, 0.07);
	--db-bg-drop-indicator: rgba(167, 139, 250, 0.16);

	--db-border: rgba(196, 181, 253, 0.10);
	--db-border-card: rgba(196, 181, 253, 0.16);
	--db-border-section: rgba(196, 181, 253, 0.08);
	--db-border-sidebar: rgba(196, 181, 253, 0.10);
	--db-border-input: rgba(196, 181, 253, 0.14);
	--db-border-input-focus: #a78bfa;
	--db-border-btn: rgba(196, 181, 253, 0.18);
	--db-border-add-section: rgba(167, 139, 250, 0.18);

	--db-text: #f4f1fa;
	--db-text-muted: #b2aabe;
	--db-text-faint: #746c80;
	--db-text-inverse: #ffffff;
	--db-text-inverse-muted: rgba(255, 255, 255, 0.68);

	--db-accent: #a78bfa;
	--db-accent-light: #d8b4fe;
	--db-danger: #fb7185;

	--db-shadow-card: 0 8px 24px rgba(0, 0, 0, 0.30);
	--db-shadow-card-hover: 0 12px 32px rgba(0, 0, 0, 0.38), 0 0 20px rgba(139, 92, 246, 0.12);
	--db-backdrop-blur: blur(18px);

	--db-radius-sm: 7px;
	--db-radius-md: 11px;
	--db-radius-lg: 16px;

	--db-font: 'Inter', 'Noto Sans SC', var(--font-interface);
	--db-progress-from: #8b5cf6;
	--db-progress-to: #d8b4fe;
	--db-checkbox: #a78bfa;
	--db-quote-border: #8b5cf6;
	--db-link: #c4b5fd;

	background:
		radial-gradient(circle at 18% 0%, rgba(124, 58, 237, 0.19), transparent 38%),
		radial-gradient(circle at 88% 22%, rgba(168, 85, 247, 0.10), transparent 34%),
		linear-gradient(155deg, #111016 0%, #09080d 54%, #060508 100%);
	color-scheme: dark;
	transition: color 0.3s ease;
}
```

- [ ] **Step 2: 增加黑曜主题组件增强**

Append to the theme-specific enhancement area:

```css
.apex-dashboard-root[data-theme="obsidian"] .dashboard-section-row,
.apex-dashboard-root[data-theme="obsidian"] .dashboard-section,
.apex-dashboard-root[data-theme="obsidian"] .dashboard-card,
.apex-dashboard-root[data-theme="obsidian"] .dashboard-sidebar,
.apex-dashboard-root[data-theme="obsidian"] .dashboard-banner {
	backdrop-filter: var(--db-backdrop-blur);
	-webkit-backdrop-filter: var(--db-backdrop-blur);
}

.apex-dashboard-root[data-theme="obsidian"] .dashboard-card:hover {
	border-color: rgba(196, 181, 253, 0.28);
	box-shadow: var(--db-shadow-card-hover);
}

.apex-dashboard-root[data-theme="obsidian"] .dashboard-memo-textarea:focus,
.apex-dashboard-root[data-theme="obsidian"] .dashboard-modal-input:focus {
	border-color: var(--db-border-input-focus);
	box-shadow: 0 0 0 3px rgba(139, 92, 246, 0.14);
}

.apex-dashboard-root[data-theme="obsidian"] .dashboard-add-section {
	background: var(--db-bg-add-section);
	border-color: var(--db-border-add-section);
}

.apex-dashboard-root[data-theme="obsidian"] .dashboard-add-section:hover {
	background: var(--db-bg-btn-hover);
	border-color: rgba(196, 181, 253, 0.32);
}

[data-theme="obsidian"] .dashboard-modal {
	background: rgba(17, 14, 24, 0.96);
	border: 1px solid var(--db-border-card);
	box-shadow: 0 24px 64px rgba(0, 0, 0, 0.52), 0 0 28px rgba(124, 58, 237, 0.10);
	color: var(--db-text);
}

[data-theme="obsidian"] .dashboard-modal-actions .mod-cta {
	background: linear-gradient(135deg, #7c3aed, #a78bfa);
	border-color: rgba(216, 180, 254, 0.28);
	color: #ffffff;
}

[data-theme="obsidian"] .dashboard-modal-actions .mod-cta:hover {
	background: linear-gradient(135deg, #8b5cf6, #c084fc);
	box-shadow: 0 8px 24px rgba(124, 58, 237, 0.28);
}
```

- [ ] **Step 3: 运行 CSS 静态完整性检查**

Run:

```powershell
$css = Get-Content -Raw '.obsidian/plugins/apex-dashboard/styles.css'
$required = @(
  '.apex-dashboard-root[data-theme="obsidian"]',
  '--db-bg: #09080d;',
  '--db-accent: #a78bfa;',
  'color-scheme: dark;',
  '[data-theme="obsidian"] .dashboard-modal',
  '[data-theme="obsidian"] .dashboard-modal-actions .mod-cta'
)

foreach ($token in $required) {
  if (-not $css.Contains($token)) {
    throw "Missing CSS token: $token"
  }
}

$open = ([regex]::Matches($css, '\{')).Count
$close = ([regex]::Matches($css, '\}')).Count
if ($open -ne $close) {
  throw "Unbalanced CSS braces: $open opening, $close closing"
}
```

Expected: no output and exit code 0.

- [ ] **Step 4: Commit CSS implementation**

```powershell
git add -- '.obsidian/plugins/apex-dashboard/styles.css'
git commit -m 'feat: add Apex Dashboard Obsidian theme styles'
```

### Task 4: 完整验证

**Files:**
- Verify: `.obsidian/plugins/apex-dashboard/main.js`
- Verify: `.obsidian/plugins/apex-dashboard/styles.css`
- Verify: `.obsidian/plugins/apex-dashboard/data.json`

- [ ] **Step 1: 运行最终自动验证**

Run:

```powershell
node --check '.obsidian/plugins/apex-dashboard/main.js'

$main = Get-Content -Raw '.obsidian/plugins/apex-dashboard/main.js'
$css = Get-Content -Raw '.obsidian/plugins/apex-dashboard/styles.css'
$data = Get-Content -Raw '.obsidian/plugins/apex-dashboard/data.json' | ConvertFrom-Json

$checks = [ordered]@{
  EnglishLabel = $main.Contains('"settings.styleObsidian":"Obsidian"')
  ChineseLabel = $main.Contains('"settings.styleObsidian":"\u9ED1\u66DC"')
  Dropdown = $main.Contains('obsidian:b("settings.styleObsidian")')
  CycleTheme = $main.Contains('"carbon","obsidian"]')
  CssTheme = $css.Contains('.apex-dashboard-root[data-theme="obsidian"]')
  AlwaysDark = $css.Contains('color-scheme: dark;')
  ModalTheme = $css.Contains('[data-theme="obsidian"] .dashboard-modal')
  CurrentThemePreserved = $data.stylePreset -eq 'tundra'
}

$failed = $checks.GetEnumerator() | Where-Object { -not $_.Value }
$checks.GetEnumerator() | ForEach-Object {
  '{0}: {1}' -f $_.Key, $_.Value
}
if ($failed) {
  throw "Final verification failed: $($failed.Key -join ', ')"
}
```

Expected: all checks print `True`.

- [ ] **Step 2: 检查范围和空白错误**

Run:

```powershell
git diff --check HEAD~2..HEAD
git show --stat --oneline HEAD~2..HEAD
git status --short
```

Expected:

- `git diff --check` has no output.
- Only `main.js` and `styles.css` appear in the two implementation commits.
- Existing unrelated user changes may remain in `git status`; none are reverted or staged.

- [ ] **Step 3: Obsidian 手工验收**

1. 在 Obsidian 中关闭再启用 Apex Dashboard 插件，或重启 Obsidian。
2. 打开 `设置 → 第三方插件 → Apex Dashboard`。
3. 确认“样式”下拉框出现“黑曜”。
4. 选择“黑曜”，打开 Dashboard。
5. 在 Obsidian 浅色与深色模式之间切换，确认 Dashboard 都保持暗黑紫色。
6. 打开卡片编辑弹窗，确认弹窗、输入框聚焦和主按钮均使用黑曜配色。
7. 将主题切回“苔原”，确认原主题仍可正常使用。

