# 国际学校智能教学管理平台 VibeCoding 开发文档

> 用途：给 VibeCoding / AI 编程工具作为产品需求、技术约束和开发任务说明。  
> 当前版本：MVP 开发版  
> 核心目标：完成“题库导入与 AI 结构化、教师智能布置作业、学生提交、AI 批改、教师复核、学生和家长查看学情”的闭环。

## 0. 已确认口径

以下为当前已确认的产品与技术口径，开发时按此执行。

| 项目 | 已确认方案 |
|---|---|
| 学校范围 | 一家学校使用，系统按单校部署设计；数据库保留 `school_id` 和超级管理员能力，方便后续扩展 |
| 管理员层级 | 需要区分超级管理员和学校管理员 |
| 家长端 | MVP 做基础只读版，支持通过绑定码或学校审核绑定学生、查看作业、查看学情 |
| 课程体系 | 默认预置常见国际课程体系，学校管理员可在后台维护 |
| 学科 | 数学、物理、化学、英语 |
| 知识点维护 | 管理员维护，教师可提交新增建议 |
| 题目知识点 | 1 个主知识点，最多 3 个辅助知识点 |
| AI 模型 | OpenAI 兼容模型，通过 Base URL、API Key、Model 配置 |
| 题库入库 | AI 提取后必须人工审核 |
| 批改展示 | AI 批改结果必须教师复核后，学生和家长才能查看 |
| 判断题 | 走客观题规则批改 |
| 论述题 | 每题必须有评分标准，AI 参考评分标准批改，教师必须复核 |
| 文件存储 | 本地服务器存储 |
| 通知 | MVP 只做站内通知 |
| 语言与主题 | 需要中文、英文切换；需要主题切换 |
| 订正/申诉/错题重做 | MVP 需要支持 |
| 学生端设备 | 手机/平板优先，电脑也能用 |
| AI 调用计费 | 需要记录每次 AI 调用的归属教师、学校、任务类型和费用；管理员可设置默认额度和单个教师额度 |

## 1. 产品定位

建设一个面向单所国际学校的智能教学管理平台，服务四类用户：

| 角色 | 核心目标 |
|---|---|
| 超级管理员/学校管理员 | 管理学校、用户、班级、题库、知识点、审核和统计 |
| 教师 | 管理班级、上传讲义、智能组题、布置作业、复核批改、查看学情 |
| 学生 | 加入多个班级、查看作业、在线答题、上传图片答案、查看反馈和错题 |
| 家长 | 绑定学生、查看学生作业完成情况、成绩反馈和学习报告 |

平台必须支持国际课程场景下的题目导入、公式识别、知识点标签、主客观题批改、订正申诉、错题重做和学情分析。

## 2. 默认技术栈

如果没有特别指定，按以下技术栈实现：

| 层级 | 技术 |
|---|---|
| 前端 | Vue 3 + Vite + TypeScript |
| UI | Ant Design Vue |
| 前端路由 | Vue Router |
| 前端状态管理 | Pinia |
| HTTP 客户端 | Axios |
| 国际化 | Vue I18n |
| 图表 | ECharts |
| 富文本/公式展示 | Markdown 渲染 + KaTeX/MathJax |
| 后端 | Java 21 + Spring Boot 3.x |
| 构建工具 | Maven |
| 数据库 | MySQL 8 |
| ORM | MyBatis-Plus |
| 数据库迁移 | Flyway |
| 缓存 | Redis + Redisson |
| 异步任务 | Spring Async + ThreadPoolTaskExecutor，复杂任务可接入 RabbitMQ/RocketMQ |
| 定时任务 | Spring Scheduler，后续可接入 XXL-JOB |
| 文件存储 | 本地服务器存储 |
| 鉴权 | Spring Security + JWT Access Token + Refresh Token |
| 权限 | RBAC 角色权限 + 数据范围校验，方法级权限使用 Spring Security 注解 |
| 接口文档 | SpringDoc OpenAPI 3 |
| 参数校验 | Jakarta Validation |
| 对象映射 | MapStruct |
| 简化代码 | Lombok |
| Excel 导入导出 | EasyExcel |
| AI 接入 | OpenAI 兼容接口，统一 Java AI Adapter，支持多模态模型、文本模型、Embedding 模型 |
| 日志 | Logback + JSON 结构化日志 |
| 测试 | JUnit 5 + Mockito + Testcontainers |
| 部署 | Docker Compose |

前后端分离开发。前端使用 Vue 3 + Vite，后端使用 Spring Boot 单体模块化架构；MVP 不拆微服务，但代码边界要按业务模块拆清楚，方便后续拆分。

### 2.1 前端工程约定

前端采用 Vue 3 + Vite + TypeScript + Ant Design Vue 的主流后台管理技术栈。

推荐目录结构：

```text
src
├─ api              Axios 请求封装，按业务模块拆分
├─ assets           静态资源
├─ components       通用组件
├─ composables      组合式函数
├─ layouts          管理端、教师端、学生端、家长端布局
├─ locales          中文、英文语言包
├─ router           Vue Router 路由和权限守卫
├─ stores           Pinia 状态管理
├─ styles           全局样式、主题变量、移动端适配
├─ utils            工具函数
└─ views            页面，按角色和业务模块组织
```

前端基础规范：

- 使用 Vue 3 Composition API。
- 所有接口请求统一从 `src/api` 发起，Axios 拦截器自动携带 Token。
- 路由守卫根据当前用户角色和后端返回的权限控制页面访问。
- Pinia 保存当前用户、Token、语言、主题、通知数量、教师 AI 用量等全局状态。
- Ant Design Vue 作为主要组件库，后台表格、筛选、弹窗、步骤条优先使用组件库能力。
- 学生端页面按手机/平板优先设计，教师端和管理员端按桌面优先设计，同时保持响应式可用。
- 中英文切换使用 Vue I18n，语言选择写入 `user_preferences`。
- 主题切换使用 Ant Design Vue Token/CSS 变量，支持浅色、深色，预留学校自定义主题。
- 数学公式展示使用 KaTeX 或 MathJax，题干和解析中的公式必须正确渲染。
- 学情图表使用 ECharts，封装通用图表组件。

### 2.2 后端工程约定

后端采用 Spring Boot + MyBatis-Plus 的主流企业项目结构。

推荐包结构：

```text
com.school.ai
├─ common          通用响应、异常、分页、常量、工具类
├─ config          Spring Security、Redis、文件存储、OpenAPI、线程池配置
├─ auth            登录、刷新 Token、当前用户
├─ user            用户、教师、学生、家长
├─ school          学校、班级、班级教师、班级学生
├─ question        题库、知识点、题库导入、知识点建议
├─ assignment      作业、提交、答案、订正、申诉、错题重做
├─ grading         规则批改、AI 批改、教师复核
├─ analytics       学情统计、月度报告
├─ ai              OpenAI 兼容模型 Adapter、用量、计费、额度
├─ file            本地文件上传、下载、权限校验
├─ notification    站内通知
└─ system          操作日志、系统配置、主题、国际化
```

后端分层要求：

- `controller`：只处理 HTTP 入参、鉴权上下文和响应包装。
- `service`：写业务规则、权限数据范围校验、事务边界。
- `mapper`：MyBatis-Plus Mapper，复杂查询写 XML。
- `entity`：数据库实体。
- `dto`：请求和响应对象。
- `converter`：MapStruct 转换 entity、dto、vo。
- `job`：AI 解析、AI 批改、月度报告等异步任务。

基础规范：

- 所有接口统一返回 `{ code, message, data }`。
- 数据库写操作使用短事务；文件上传、AI 解析、AI 批改等长耗时流程必须通过任务表和状态机衔接，不能把外部调用包在一个数据库事务里。
- 所有列表接口支持分页、排序和筛选。
- 所有上传文件先落本地磁盘临时区，再写 `files` 表并移动到正式目录；失败时必须清理临时文件或标记为 `FAILED`。
- 所有 AI 调用必须经过 `ai` 模块，不能在业务模块里直接调模型。
- 所有异步任务必须记录状态、错误信息、重试次数、开始时间和结束时间，方便前端展示和后台排查。
- Flyway 管理数据库建表和变更脚本，禁止手工改库后不留迁移文件。
- MySQL 8 使用 `utf8mb4` 字符集和 `utf8mb4_0900_ai_ci` 排序规则；结构化扩展字段优先使用 MySQL `JSON` 类型。
- SpringDoc 自动生成 OpenAPI 文档，方便前端和 VibeCoding 对接。

## 3. MVP 范围

MVP 必须打通以下闭环：

1. 管理员创建学校、教师、学生、家长、班级。
2. 管理员维护知识点树和题库。
3. 管理员或教师上传 PDF、Word、TXT、图片等资料，AI 自动提取题型、题干、选项、答案、解析、评分标准、知识点标签。
4. 管理员审核系统题目，教师审核或管理自己的个人题目。
5. 教师上传讲义，系统自动识别知识点，并从题库推荐相关题目。
6. 教师设置题型和题量，保存个人默认配置。
7. 教师发布作业，学生收到通知。
8. 学生进入“我的作业”，完成选择题、多选题、判断题、填空题、简答题、论述题。
9. 学生可直接输入答案，也可上传拍照图片作为答案。
10. 提交后系统进行批改：客观题规则批改，主观题和图片答案交给多模态 AI 批改。
11. AI 批改完成后通知教师复核。
12. 教师复核并确认分数后，通知学生和家长查看结果。
13. 班级管理支持添加、删除、编辑学生。
14. 学生可以加入多个班级。
15. 系统生成学生月度做题报告和基础薄弱点分析。
16. 家长绑定学生后查看学习情况。
17. 学生可以对错题进行订正、申诉和重做。
18. 系统支持中文、英文语言切换和主题切换。
19. 系统记录每位教师的 AI 调用用量、费用、剩余额度，并允许管理员设置默认额度和单个教师额度。

## 4. 角色与权限

### 4.1 角色

- `SUPER_ADMIN`：平台超级管理员，可管理所有学校和全局题库。
- `SCHOOL_ADMIN`：学校管理员，只能管理本校数据。
- `TEACHER`：教师，只能管理自己任课班级、自己上传的题目和作业。
- `STUDENT`：学生，只能查看和提交自己的作业。
- `PARENT`：家长，只能查看已绑定学生的数据。

### 4.2 权限原则

- 所有接口必须同时校验角色权限和数据范围。
- 不允许只靠前端隐藏菜单控制权限。
- 教师访问学生时，必须校验 `class_teachers` 和 `class_students` 关系。
- 家长访问学生时，必须校验 `parent_students` 绑定关系。
- 题库导入、题目审核、作业发布、批改改分、导出数据等关键操作写入操作日志。
- 单校部署下也必须保留 `school_id` 字段，方便未来扩展和数据隔离。
- 教师发起 AI 题库提取、智能组题、AI 批改、AI 学情总结前，必须检查教师 AI 额度。

## 5. 核心业务流程

### 5.1 题库导入流程

1. 管理员或教师上传文件。
2. 系统保存文件元数据到 `files`。
3. 创建 `question_import_jobs` 任务。
4. 后端队列调用 AI 多模态解析。
5. AI 输出结构化题目草稿。
6. 系统按知识点树进行标签匹配，每题设置 1 个主知识点，最多 3 个辅助知识点。
7. 管理员或教师进入导入预览页人工确认。
8. 审核通过后写入 `questions`、`question_knowledge_points`、`question_options`。

管理员上传的题目来源为 `SYSTEM`，教师上传的题目来源为 `TEACHER`，并记录 `owner_teacher_id`。

### 5.2 智能布置作业流程

1. 教师选择班级。
2. 教师上传讲义或选择知识点。
3. AI 识别讲义覆盖的知识点。
4. 系统读取教师上次题型题量配置作为默认值。
5. 教师设置题型数量：选择题、判断题、多选题、填空题、简答题、论述题。
6. 系统根据知识点、题型、难度、来源、近期重复情况推荐题目。
7. 教师可替换、删除、排序题目。
8. 教师设置截止时间、发布范围和说明。
9. 发布后创建作业、通知学生。
10. 课程体系、题型、难度、知识点等筛选项来自学校管理员维护的后台字典。

### 5.3 学生作答流程

1. 学生登录后进入“我的作业”。
2. 未完成作业显示作业标题、班级、教师、截止时间、题目数量。
3. 学生打开作业开始答题。
4. 客观题使用选项控件。
5. 填空题、简答题、论述题支持文本输入和图片上传。
6. 学生可保存草稿。
7. 提交后状态变为 `SUBMITTED`，教师收到通知。
8. 学生查看批改结果后，可提交订正、发起申诉或进入错题重做。

### 5.4 批改与复核流程

1. 学生提交后创建批改任务。
2. 选择题、多选题、判断题由规则引擎自动批改。
3. 填空题、简答题、论述题、图片答案由多模态 AI 参考标准答案和评分标准批改。
4. AI 输出得分、评语、扣分点、置信度和是否需要人工复核。
5. 批改完成后通知教师。
6. 教师在批改详情页查看学生答案、标准答案、AI 评分和证据。
7. 教师可修改分数、评语、复核状态。
8. 教师确认后该学生的提交状态变为 `REVIEWED`；作业整体状态保持 `PUBLISHED`，直到教师或管理员关闭作业。
9. 系统通知学生和家长查看结果。
10. 若学生发起申诉，教师收到申诉通知并可重新复核。

### 5.5 学情分析流程

1. 每次作业复核完成后，更新学生题目表现。
2. 按知识点、题型、时间、班级维度统计正确率、得分率、完成率。
3. 每月生成学生 AI 学情报告。
4. 报告内容包括：本月完成情况、优势知识点、薄弱知识点、常错题型、建议练习方向。
5. 教师可查看班级和学生报告，家长可查看绑定学生报告。
6. 学情报告支持中文和英文展示，AI 总结根据当前界面语言生成或翻译。

## 6. 数据模型设计

### 6.1 用户与组织

- `schools`：学校
- `users`：统一账号表
- `teachers`：教师扩展信息
- `students`：学生扩展信息
- `parents`：家长扩展信息
- `parent_students`：家长学生绑定
- `classes`：班级
- `class_teachers`：班级教师关系
- `class_students`：班级学生关系，支持一个学生加入多个班级
- `user_preferences`：用户偏好，保存语言、主题等设置

### 6.2 题库与知识点

- `knowledge_points`：知识点树
- `knowledge_point_suggestions`：教师提交的知识点新增建议
- `questions`：题目主表
- `question_options`：选择题/多选题选项
- `question_knowledge_points`：题目知识点多对多
- `question_import_jobs`：题库导入任务
- `question_import_drafts`：AI 提取后的题目草稿

题目字段建议：

| 字段 | 说明 |
|---|---|
| `id` | 题目 ID |
| `school_id` | 所属学校，系统题可为空或为平台级 |
| `source_type` | `SYSTEM` / `TEACHER` |
| `owner_teacher_id` | 教师题目所属教师 |
| `subject` | 学科 |
| `course_system` | 课程体系 |
| `exam_board` | 考局 |
| `question_type` | `SINGLE_CHOICE` / `MULTIPLE_CHOICE` / `TRUE_FALSE` / `FILL_BLANK` / `SHORT_ANSWER` / `ESSAY` |
| `stem` | 题干 |
| `answer` | 标准答案 |
| `grading_rubric` | 评分标准 |
| `analysis` | 解析 |
| `difficulty` | 难度 |
| `status` | `DRAFT` / `PENDING_REVIEW` / `APPROVED` / `REJECTED` / `ARCHIVED` |
| `default_score` | 默认分值，可选；布置作业时可在 `assignment_items` 中覆盖 |
| `formula_latex_json` | 公式 LaTeX 数组，可选，例如 `["x^2+1"]` |
| `created_by` | 创建人 |

`question_knowledge_points` 必须包含 `relation_type` 字段：

| 字段 | 说明 |
|---|---|
| `question_id` | 题目 ID |
| `knowledge_point_id` | 知识点 ID |
| `relation_type` | `PRIMARY` / `SECONDARY` |
| `sort_order` | 辅助知识点排序 |

校验规则：

- 每道题必须且只能有 1 个 `PRIMARY` 主知识点。
- 每道题最多有 3 个 `SECONDARY` 辅助知识点。
- 教师不能直接新增正式知识点，只能写入 `knowledge_point_suggestions` 等管理员审核。

### 6.3 作业与提交

- `assignments`：作业
- `assignment_items`：作业题目
- `assignment_targets`：发布范围，全班、分组或指定学生
- `student_answers`：学生单题答案
- `submissions`：作业提交记录
- `grading_tasks`：批改任务，记录 AI/规则批改状态、重试次数和错误信息
- `grading_records`：单题批改记录
- `teacher_assignment_preferences`：教师上次题型题量配置
- `corrections`：学生订正记录
- `grading_appeals`：学生申诉记录
- `wrong_question_attempts`：错题重做记录

作业状态：

| 状态 | 含义 |
|---|---|
| `DRAFT` | 草稿 |
| `PUBLISHED` | 已发布 |
| `CLOSED` | 已关闭 |
| `ARCHIVED` | 已归档 |

作业整体状态只描述教师发布和关闭层面的生命周期。学生是否提交、批改中、待复核、已复核，必须以 `submissions` 和 `grading_tasks` 为准；同一份作业下不同学生可以处于不同状态。

`assignment_items` 建议包含以下关键字段：

| 字段 | 说明 |
|---|---|
| `assignment_id` | 作业 ID |
| `question_id` | 题目 ID |
| `score` | 本次作业中该题分值 |
| `sort_order` | 题目排序 |
| `required` | 是否必答 |

提交状态：

| 状态 | 含义 |
|---|---|
| `NOT_STARTED` | 未开始 |
| `IN_PROGRESS` | 作答中 |
| `SUBMITTED` | 已提交 |
| `AI_GRADING` | AI 批改中 |
| `WAIT_TEACHER_REVIEW` | 待教师复核 |
| `REVIEWED` | 已复核 |
| `RETURNED` | 打回重做 |

批改任务状态：

| 状态 | 含义 |
|---|---|
| `PENDING` | 待处理 |
| `RUNNING` | 批改中 |
| `SUCCEEDED` | 批改完成 |
| `FAILED` | 批改失败，可由教师或管理员重试 |
| `CANCELLED` | 已取消 |

申诉状态：

| 状态 | 含义 |
|---|---|
| `OPEN` | 学生已提交申诉 |
| `IN_REVIEW` | 教师处理中 |
| `ACCEPTED` | 申诉通过，分数或评语已调整 |
| `REJECTED` | 申诉驳回 |
| `CLOSED` | 申诉关闭 |

### 6.4 通知、文件、日志

- `files`：上传文件元数据
- `notifications`：站内通知
- `operation_logs`：操作日志
- `ai_call_logs`：AI 调用日志
- `ai_quota_defaults`：AI 默认额度配置，管理员设置全校教师默认值
- `teacher_ai_quotas`：教师 AI 额度配置和周期用量
- `ai_billing_records`：AI 调用计费明细，按学校、归属教师、任务、模型、token、金额记录
- `learning_reports`：学情报告
- `wrong_questions`：错题
- `theme_configs`：主题配置，可保存浅色、深色或学校自定义主题
- `i18n_messages`：可选，保存后台可维护的界面文案

### 6.5 AI 用量与计费

AI 调用计费需要支持管理员设置默认额度和单个教师额度。每次 AI 调用都必须明确费用归属，避免系统任务、管理员任务和教师任务混在一起。

`ai_quota_defaults` 字段建议：

| 字段 | 说明 |
|---|---|
| `school_id` | 学校 ID |
| `period_type` | `MONTHLY` / `TERM` |
| `default_quota_amount` | 默认金额额度 |
| `default_token_quota` | 默认 Token 额度，可选 |
| `currency` | 币种，例如 `CNY` / `USD` |
| `warning_threshold_percent` | 预警阈值，例如 80 |

`teacher_ai_quotas` 字段建议：

| 字段 | 说明 |
|---|---|
| `teacher_id` | 教师 ID |
| `period_start` | 额度周期开始 |
| `period_end` | 额度周期结束 |
| `quota_amount` | 当前周期金额额度 |
| `used_amount` | 当前周期已用金额 |
| `quota_tokens` | 当前周期 Token 额度，可选 |
| `used_tokens` | 当前周期已用 Token |
| `status` | `NORMAL` / `WARNING` / `EXCEEDED` / `DISABLED` |

`ai_call_logs` 和 `ai_billing_records` 必须包含以下归属字段：

| 字段 | 说明 |
|---|---|
| `school_id` | 学校 ID |
| `owner_teacher_id` | 费用归属教师，可为空 |
| `trigger_user_id` | 实际发起用户 |
| `task_type` | `QUESTION_IMPORT` / `HANDOUT_ANALYSIS` / `QUESTION_RECOMMEND` / `GRADING` / `LEARNING_REPORT` |
| `related_id` | 关联任务或业务对象 ID |
| `model_name` | 模型名称 |
| `input_tokens` | 输入 Token |
| `output_tokens` | 输出 Token |
| `image_count` | 图片数量，可选 |
| `file_count` | 文件数量，可选 |
| `amount` | 服务端计算费用 |
| `status` | `PENDING` / `SUCCEEDED` / `FAILED` |

计费规则：

- 每次 AI 调用必须写入 `ai_call_logs` 和 `ai_billing_records`。
- 教师主动发起的题库导入、讲义分析、智能组题，费用归属该教师。
- 学生提交后触发的 AI 批改，费用默认归属发布作业的教师；如果作业有多个教师，归属作业创建人。
- 管理员发起的系统题库导入，`owner_teacher_id` 为空，只记入学校级 AI 日志，不扣教师额度。
- 月度学情报告如果由教师手动生成，费用归属教师；如果由系统定时批量生成，费用归属学校公共任务，不扣教师额度。
- AI 调用完成后更新对应 `teacher_ai_quotas.used_amount` 和 `teacher_ai_quotas.used_tokens`；无归属教师的调用只写学校级日志和计费明细。
- 教师达到预警阈值时，通知教师和管理员。
- 教师超过额度后，默认禁止继续发起题库导入、讲义分析、智能组题、手动生成报告等高成本主动操作，管理员可手动调高额度或重置周期。
- 学生已经提交的作业不能因教师额度不足而阻止学生提交；如果批改时额度不足，提交状态进入 `WAIT_TEACHER_REVIEW` 或批改任务进入 `FAILED`，提示教师/管理员处理额度或改为人工批改。
- 管理员可以设置全校默认额度，也可以覆盖单个教师额度。

## 7. AI 能力设计

### 7.1 AI 总原则

- 所有 AI 能力通过后端 `AIService` 调用，不允许前端直接调用模型。
- AI Provider 使用 OpenAI 兼容协议，配置项包含 `OPENAI_COMPATIBLE_BASE_URL`、`OPENAI_COMPATIBLE_API_KEY`、`OPENAI_COMPATIBLE_MODEL`。
- PDF、Word、TXT 优先进行文件文本提取；扫描 PDF、图片和拍照答案优先进行 OCR/版面分析，再交给 AI 结构化和纠错。
- 多模态 AI 用于处理低质量图片、复杂版面、公式识别、学生拍照答案和 OCR 结果不可靠的场景。
- AI 输出必须结构化 JSON，并保存原始输出和模型版本。
- AI 输出 JSON 必须做 schema 校验；校验失败时任务标记为 `FAILED`，保留原始输出，允许人工修正或重试。
- AI 结果不能直接作为最终成绩，必须经过教师复核后展示给学生和家长。
- AI 学情总结需要支持中文和英文输出，默认跟随用户当前语言。

### 7.2 题库提取 AI 输出格式

```json
{
  "questions": [
    {
      "questionType": "SINGLE_CHOICE",
      "stem": "题干，公式尽量使用 LaTeX 表达",
      "options": [
        { "label": "A", "content": "选项 A" }
      ],
      "answer": "A",
      "analysis": "解析",
      "gradingRubric": "评分标准",
      "primaryKnowledgePointName": "Quadratic equations",
      "secondaryKnowledgePointNames": ["Solving equations", "Algebraic manipulation"],
      "difficulty": 2,
      "formulaLatex": ["x=\\frac{-b\\pm\\sqrt{b^2-4ac}}{2a}"],
      "confidence": 0.92
    }
  ],
  "warnings": ["第 3 题公式识别置信度较低，请人工复核"]
}
```

### 7.3 智能组题 AI 输入

```json
{
  "teacherId": 1,
  "classId": 10,
  "subject": "Mathematics",
  "courseSystem": "IGCSE",
  "knowledgePointIds": [101, 102],
  "questionTypeCounts": {
    "SINGLE_CHOICE": 5,
    "TRUE_FALSE": 2,
    "MULTIPLE_CHOICE": 2,
    "FILL_BLANK": 3,
    "SHORT_ANSWER": 2,
    "ESSAY": 1
  },
  "difficultyRange": [1, 3],
  "excludeRecentlyUsed": true
}
```

### 7.4 AI 批改输出格式

```json
{
  "submissionId": 10001,
  "results": [
    {
      "questionId": 20001,
      "score": 4,
      "maxScore": 5,
      "isCorrect": false,
      "comment": "最终答案正确，但关键步骤缺失，扣 1 分。",
      "deductions": [
        { "reason": "缺少公式推导过程", "points": 1 }
      ],
      "confidence": 0.88,
      "needsManualReview": true
    }
  ],
  "overallComment": "整体完成较好，但步骤表达需要加强。"
}
```

### 7.5 必须人工复核的情况

- AI 置信度低于 0.85。
- 学生上传图片模糊或无法判断。
- 论述题、长篇开放题。
- AI 输出 JSON 不合法。
- AI 评分超过题目满分或低于 0。
- 教师设置该作业为全部人工复核。

## 8. 页面与路由

### 8.1 公共页面

- `/login`：登录
- `/profile`：个人中心
- `/notifications`：通知中心

### 8.2 管理员端

- `/admin/dashboard`：管理看板
- `/admin/schools`：学校管理
- `/admin/users`：用户管理
- `/admin/classes`：班级管理
- `/admin/questions`：系统题库
- `/admin/question-imports`：题库导入任务
- `/admin/knowledge-points`：知识点管理
- `/admin/ai-logs`：AI 调用日志
- `/admin/ai-quotas`：教师 AI 额度和计费管理
- `/admin/operation-logs`：操作日志

### 8.3 教师端

- `/teacher/dashboard`：教师首页
- `/teacher/classes`：我的班级
- `/teacher/classes/:id/students`：班级学生管理
- `/teacher/question-bank`：个人题库
- `/teacher/question-imports`：上传题目
- `/teacher/assignments/new`：布置作业
- `/teacher/assignments/:id`：作业详情
- `/teacher/grading`：批改作业
- `/teacher/grading/:submissionId`：批改详情
- `/teacher/analytics/classes/:id`：班级学情
- `/teacher/analytics/students/:id`：学生报告

### 8.4 学生端

- `/student/dashboard`：学生首页
- `/student/classes`：我的班级
- `/student/assignments`：我的作业
- `/student/assignments/:id/do`：开始答题
- `/student/submissions/:id/result`：批改反馈
- `/student/wrong-questions`：错题集
- `/student/analytics`：我的学情

### 8.5 家长端

- `/parent/dashboard`：家长首页
- `/parent/children`：绑定学生
- `/parent/children/:studentId/assignments`：学生作业情况
- `/parent/children/:studentId/reports`：学生学情报告

## 9. API 设计

统一前缀：`/api/v1`

### 9.1 鉴权

| 方法 | 路径 | 说明 |
|---|---|---|
| POST | `/auth/login` | 登录 |
| POST | `/auth/refresh` | 刷新 Token |
| POST | `/auth/logout` | 退出 |
| GET | `/auth/me` | 当前用户 |

### 9.2 用户与班级

| 方法 | 路径 | 说明 |
|---|---|---|
| GET | `/users` | 用户列表 |
| POST | `/users` | 创建用户 |
| PATCH | `/users/:id` | 编辑用户 |
| PATCH | `/users/:id/status` | 启用/禁用 |
| POST | `/users/import` | 批量导入 |
| GET | `/classes` | 班级列表 |
| POST | `/classes` | 创建班级 |
| PATCH | `/classes/:id` | 编辑班级 |
| POST | `/classes/:id/students` | 添加学生到班级 |
| DELETE | `/classes/:id/students/:studentId` | 从班级移除学生 |
| POST | `/classes/:id/teachers` | 添加任课教师 |

### 9.3 题库

| 方法 | 路径 | 说明 |
|---|---|---|
| GET | `/questions` | 题目筛选 |
| POST | `/questions` | 手动新增题目 |
| PATCH | `/questions/:id` | 编辑题目 |
| POST | `/questions/:id/review` | 审核题目 |
| POST | `/question-imports` | 创建导入任务 |
| GET | `/question-imports` | 导入任务列表 |
| GET | `/question-imports/:id/drafts` | 查看 AI 提取草稿 |
| POST | `/question-imports/:id/approve` | 确认入库 |
| GET | `/knowledge-points` | 知识点树 |
| POST | `/knowledge-points` | 创建知识点 |
| POST | `/knowledge-point-suggestions` | 教师提交知识点新增建议 |
| GET | `/knowledge-point-suggestions` | 管理员查看建议 |
| POST | `/knowledge-point-suggestions/:id/review` | 管理员审核建议 |

### 9.4 作业

| 方法 | 路径 | 说明 |
|---|---|---|
| GET | `/assignments` | 作业列表 |
| POST | `/assignments` | 创建作业草稿 |
| POST | `/assignments/recommend` | 智能推荐题目 |
| PATCH | `/assignments/:id` | 修改作业 |
| POST | `/assignments/:id/publish` | 发布作业 |
| GET | `/assignments/:id/submissions` | 提交列表 |
| GET | `/teacher/preferences/assignment` | 获取教师上次题型配置 |
| PUT | `/teacher/preferences/assignment` | 保存教师题型配置 |

### 9.5 学生作答与批改

| 方法 | 路径 | 说明 |
|---|---|---|
| GET | `/student/assignments` | 学生作业 |
| GET | `/student/assignments/:id` | 作业详情 |
| POST | `/submissions` | 创建/提交作答 |
| PUT | `/submissions/:id/answers` | 保存答案草稿 |
| POST | `/submissions/:id/files` | 上传图片答案 |
| POST | `/submissions/:id/submit` | 最终提交，提交成功后自动创建批改任务 |
| GET | `/submissions/:id/grading` | 查看批改结果 |
| GET | `/grading-tasks/:id` | 查看批改任务状态 |
| POST | `/grading-tasks/:id/retry` | 教师或管理员重试失败的批改任务 |
| PATCH | `/grading-records/:id` | 教师复核单题 |
| POST | `/submissions/:id/review` | 教师完成复核 |
| POST | `/submissions/:id/corrections` | 学生提交订正 |
| POST | `/grading-records/:id/appeals` | 学生对单题发起申诉 |
| PATCH | `/grading-appeals/:id` | 教师处理申诉 |
| POST | `/wrong-questions/:id/attempts` | 学生错题重做 |

### 9.6 学情与家长

| 方法 | 路径 | 说明 |
|---|---|---|
| GET | `/analytics/classes/:id` | 班级学情 |
| GET | `/analytics/students/:id` | 学生学情 |
| POST | `/analytics/students/:id/monthly-report` | 生成月度报告 |
| GET | `/parent/children` | 家长绑定学生列表 |
| POST | `/parent/children/bind` | 使用绑定码申请绑定学生 |
| DELETE | `/parent/children/:studentId` | 解绑学生 |
| GET | `/parent/children/:studentId/report` | 查看学生报告 |
| GET | `/settings/preferences` | 获取语言、主题等个人偏好 |
| PUT | `/settings/preferences` | 保存语言、主题等个人偏好 |

家长绑定规则：

- MVP 使用学生绑定码作为默认绑定方式，绑定码由学校管理员或班主任生成。
- 绑定码必须有有效期和使用次数限制，使用后写入操作日志。
- 学校可配置是否需要管理员审核；如开启审核，绑定申请先进入待审核状态，审核通过后家长才能查看学生数据。
- 家长只能查看已绑定学生的作业结果和学情报告，不能查看教师内部批注、其他学生数据或班级整体数据。

### 9.7 AI 用量与额度

| 方法 | 路径 | 说明 |
|---|---|---|
| GET | `/admin/ai-quota-defaults` | 查看全校教师默认 AI 额度 |
| PUT | `/admin/ai-quota-defaults` | 设置全校教师默认 AI 额度 |
| GET | `/admin/teacher-ai-quotas` | 查看教师 AI 用量和额度列表 |
| PATCH | `/admin/teacher-ai-quotas/:teacherId` | 设置单个教师 AI 额度 |
| POST | `/admin/teacher-ai-quotas/:teacherId/reset` | 重置单个教师当前周期用量 |
| GET | `/admin/ai-billing-records` | 查看 AI 计费明细 |
| GET | `/teacher/ai-usage` | 教师查看自己的 AI 用量和剩余额度 |

## 10. 前端交互要求

### 10.1 教师布置作业

- 使用步骤式表单。
- 第一步选择班级，第二步上传讲义或选择知识点，第三步设置题型题量，第四步预览题目，第五步发布。
- 题型题量输入项必须默认读取教师上一次配置。
- 推荐题目列表允许替换、删除、排序。
- 显示每道题的题型、难度、知识点、来源、AI 推荐理由。
- 发起智能组题前显示本次预计消耗和当前剩余额度。

### 10.2 学生答题

- 学生端移动优先，桌面可用。
- 客观题使用单选、多选、判断控件。
- 主观题提供文本框和图片上传。
- 支持保存草稿。
- 提交前显示未答题提醒。
- 图片上传支持预览、删除、重新上传。
- 批改结果页提供订正、申诉、错题重做入口。
- 订正需要保留订正内容、提交时间、教师反馈。
- 申诉需要选择具体题目并填写申诉理由。

### 10.3 教师批改

- 左侧显示学生答案文本和图片。
- 右侧显示标准答案、评分标准、AI 建议得分、AI 评语。
- 教师可以快速通过，也可以改分和改评语。
- 显示 AI 置信度和需要复核原因。

### 10.4 学情报告

- 学生报告包含完成率、平均得分率、知识点掌握、错题统计、AI 总结。
- 家长端只读，不展示教师内部批注。
- 教师端可查看更详细的题目级别错误分析。
- 报告支持中文和英文显示。

### 10.5 语言与主题

- 顶部导航或个人中心提供语言切换：中文、English。
- 顶部导航或个人中心提供主题切换：浅色、深色，后续可扩展学校自定义主题。
- 用户选择语言和主题后写入 `user_preferences`。
- 未登录状态默认中文和浅色主题。

### 10.6 管理员 AI 额度管理

- 管理员可设置全校教师默认 AI 额度、周期类型和预警阈值。
- 管理员可查看教师列表中的本期已用金额、已用 Token、剩余额度、状态。
- 管理员可单独调整某个教师的额度。
- 管理员可查看每次 AI 调用的任务类型、模型、Token、费用、状态和错误信息。
- 教师端首页显示本期 AI 已用量和剩余额度。

## 11. 安全与合规

- 密码使用 Argon2id 或 bcrypt。
- 文件下载必须鉴权，不能公开暴露真实路径。
- AI 请求尽量使用匿名学生 ID，避免传递不必要的姓名、手机号等个人信息。
- 所有上传文件限制类型和大小。
- 所有写操作必须校验用户角色和数据范围。
- 批改改分、导出、删除、审核必须写操作日志。
- 防止学生访问其他学生作业、防止家长访问未绑定学生数据。
- 家长绑定学生必须通过绑定码或学校审核，不允许仅凭学生 ID、姓名或手机号直接绑定。
- 本地服务器文件存储必须按学校、业务类型和日期分目录，不允许前端直接访问真实文件路径。
- OpenAI 兼容模型的 API Key 只能保存在服务端环境变量或密钥配置中。
- AI 计费金额由服务端根据模型单价和 token/图片/文件调用量计算，前端不能提交最终费用。

## 12. 开发里程碑

### 12.1 第一阶段：基础框架

- 初始化前端 Vue 3 + Vite + TypeScript + Ant Design Vue 项目。
- 初始化后端 Spring Boot + MyBatis-Plus 项目。
- 配置前端 Vue Router、Pinia、Axios、Vue I18n、主题切换和权限路由。
- 配置 Maven、Flyway、SpringDoc、统一异常、统一响应、分页、日志。
- 完成登录、JWT、RBAC。
- 完成学校、用户、班级、班级学生关系。
- 完成文件上传基础能力。
- 完成中英文语言切换和主题切换基础框架。

### 12.2 第二阶段：题库与 AI 导入

- 完成知识点树。
- 完成教师知识点新增建议和管理员审核。
- 完成题库 CRUD。
- 完成题库文件上传和 AI 提取任务。
- 完成文件文本提取、OCR/版面分析和 AI 结构化的基础任务链路。
- 完成题目草稿审核入库。

### 12.3 第三阶段：作业闭环

- 完成教师布置作业。
- 完成题型题量偏好保存。
- 完成学生作答和提交。
- 完成通知。

### 12.4 第四阶段：批改与复核

- 完成客观题规则批改。
- 完成主观题多模态 AI 批改。
- 完成批改任务状态、失败重试和人工接管。
- 完成教师复核。
- 完成学生查看反馈。
- 完成学生订正、申诉和错题重做。

### 12.5 第五阶段：学情与家长

- 完成班级学情。
- 完成学生月度报告。
- 完成家长绑定码、可选审核和只读查看。
- 完成基础看板和日志。

### 12.6 第六阶段：AI 用量与额度

- 完成 AI 调用日志。
- 完成 AI 计费明细。
- 完成全校默认教师额度设置。
- 完成单个教师额度设置和周期用量统计。
- 完成额度预警、超额限制和站内通知。

## 13. 验收标准

1. 管理员可以创建教师、学生、家长和班级。
2. 一个学生可以加入多个班级。
3. 管理员可以上传题库文件，AI 能生成结构化题目草稿。
4. 每道题可以绑定多个知识点。
5. 每道题必须有 1 个主知识点，最多 3 个辅助知识点。
6. 公式内容可保存为 LaTeX 或可读数学表达。
7. 管理员审核后题目才进入系统题库。
8. 教师上传的题目归属教师个人题库。
9. 教师只能提交知识点新增建议，管理员审核后才进入正式知识点树。
10. 教师上传讲义后，系统能识别知识点并推荐题目。
11. 教师题型题量设置能自动保存为下次默认值。
12. 学生收到作业通知并能在线作答。
13. 学生可上传图片作为主观题答案。
14. 判断题、选择题、多选题能按规则自动批改。
15. 填空题、简答题、论述题和图片答案能调用 OpenAI 兼容多模态模型批改。
16. 论述题必须有评分标准，AI 必须参考评分标准输出建议分和评语。
17. 教师能复核 AI 批改结果并改分。
18. 教师复核后学生和家长才能看到最终结果。
19. 学生能提交订正、发起申诉、进行错题重做。
20. 家长通过绑定码或学校审核绑定学生后，才能查看学习情况。
21. 系统能生成学生月度 AI 学情分析。
22. 系统支持中文、英文切换和主题切换。
23. 管理员能设置全校教师默认 AI 额度。
24. 管理员能单独设置每个教师的 AI 额度。
25. 系统能记录每次 AI 调用的学校、归属教师、任务类型、Token、费用、状态和错误信息。
26. 教师达到预警阈值时，教师和管理员能收到站内通知。
27. 教师超过额度后，系统能限制高成本 AI 操作。
28. 所有核心接口有权限校验。
29. 关键操作有日志。
30. AI 调用有日志、状态和错误记录。
31. 作业整体状态与学生提交状态分离，同一作业下不同学生可以处于不同批改进度。
32. 批改任务失败后，教师或管理员可以查看失败原因并重试，必要时可以人工批改。

## 14. 给 VibeCoding 的执行提示

请按本文档开发一个可运行的 MVP。优先实现完整业务闭环，不要先做营销首页。

开发时遵守以下顺序：

1. 先建数据库 schema 和权限模型。
2. 再搭建 Vue 3 前端基础工程和 Spring Boot 后端基础工程。
3. 再实现登录和不同角色工作台。
4. 再实现题库、知识点和文件上传。
5. 再实现 AI Adapter 的 mock 版本，保证前后端流程跑通。
6. 再实现 AI 调用日志、额度和计费统计。
7. 再接入真实多模态 AI。
8. 最后完善学情、家长端和统计看板。

AI Adapter 初期可以提供 mock 返回，结构必须与本文档一致，后续替换真实模型时不影响业务代码。Mock 阶段也要模拟 token、费用和额度扣减，避免真实模型接入后再补计费逻辑。
