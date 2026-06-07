# AI服务设计文档

> 本文档用于指导 OCR、智能组题、自动批改、学情总结、教学评价等 AI 能力的工程落地。

## 1. AI能力范围

| 能力 | MVP 是否需要 | 说明 |
|---|---:|---|
| OCR 作业识别 | 是 | 识别学生拍照/扫描上传的答案 |
| 智能组题 | 是 | 根据课程、知识点、难度、题型推荐题目 |
| 自动批改 | 是 | 优先支持客观题、填空题、简单计算题 |
| 错题归因 | 是 | 将错误关联到题目、知识点、原因 |
| 学情总结 | 是 | 生成班级/学生文字总结 |
| 教学评价 | 后置 | 上传课件和录音后分析教学质量 |
| 复杂主观题评分 | 后置 | 论述题、画图题、实验题先人工处理 |

## 2. AI服务架构

建议后端单独封装 `AI Service`，业务系统不要直接调用第三方模型。

```text
前端
  ↓
业务后端 API
  ↓
AI Service Adapter
  ├─ OCR Provider
  ├─ Document Parse Provider（PDF/Word/扫描件解析，如 MinerU）
  ├─ LLM Provider
  ├─ Embedding / Rerank Provider（后续）
  └─ Cost & Log Tracker
  ↓
数据库：ai_call_logs / grading_records / learning_stats
```

## 3. AI调用日志表建议

建议新增表：`ai_call_logs`

| 字段 | 类型 | 说明 |
|---|---|---|
| id | bigint | 主键 |
| user_id | bigint | 调用用户 |
| school_id | bigint | 学校 |
| task_type | varchar | ocr / grade / recommend / summary / evaluation |
| provider | varchar | 模型供应商 |
| model | varchar | 模型名称 |
| input_tokens | int | 输入 token |
| output_tokens | int | 输出 token |
| file_count | int | 文件数 |
| cost_amount | decimal | 成本 |
| status | varchar | success / failed |
| error_message | text | 错误信息 |
| latency_ms | int | 耗时 |
| created_at | datetime | 创建时间 |

## 4. OCR / 文档解析设计

> OCR 主要解决学生作业图片、扫描件答案识别；文档解析主要解决 PDF/Word 题库、试卷、教材教辅资料的结构化提取。两者可以共用任务队列、日志和人工复核机制，但建议在服务层拆成不同 Provider，便于后续替换。

### 4.0 Provider 候选：MinerU

MinerU 可作为 **PDF/文档解析 Provider 候选方案**，优先用于题库导入和试卷资料结构化，不直接绑定为唯一方案。

适用场景：

- PDF 真题、教材教辅、扫描资料解析。
- 多栏排版、图片、表格、公式较多的资料预处理。
- 将 PDF 转为 Markdown / JSON / 图片资源，供后续 AI 切题、打标签和人工审核。
- 扫描版 PDF 需要结合 OCR 还原文本和版面结构。

不建议场景：

- 学生手写答案批改的实时 OCR 主流程，仍应优先选择手写识别效果更稳定的 OCR Provider。
- 对延迟要求很高的在线批改同步接口，应放入异步任务队列。

建议接入方式：

1. 上传 PDF/Word/图片文件后进入文件任务队列。
2. 按文件类型选择 Provider：图片/手写作业走 OCR Provider；PDF 题库/试卷走 Document Parse Provider。
3. MinerU 输出 Markdown、版面块、图片、表格、公式等中间结果。
4. LLM 基于中间结果提取题目块、答案、解析、知识点和来源信息。
5. 生成导入草稿，进入人工审核。
6. 审核通过后写入正式题库。

Provider 输出建议统一为：

```json
{
  "provider": "mineru",
  "sourceFileId": 10001,
  "pages": [
    {
      "pageNo": 1,
      "markdown": "...",
      "blocks": [
        {
          "type": "text/table/formula/image",
          "text": "...",
          "bbox": [10, 20, 300, 120],
          "confidence": 0.9
        }
      ]
    }
  ],
  "assets": [
    {
      "type": "image",
      "path": "assets/page-1-img-1.png"
    }
  ]
}
```

### 4.1 输入

- 图片：JPG、PNG、HEIC 转换后图片
- PDF：扫描件或拍照合成 PDF
- 元数据：学生 ID、作业 ID、页码、题号区域可选

### 4.2 输出

```json
{
  "pages": [
    {
      "pageNo": 1,
      "blocks": [
        {
          "type": "answer",
          "questionNo": "1",
          "text": "学生答案文本",
          "confidence": 0.92,
          "bbox": [10, 20, 300, 120]
        }
      ]
    }
  ]
}
```

### 4.3 失败兜底

- 识别置信度低于阈值时标记“需人工确认”。
- 图片模糊、角度异常、过暗时提示学生重新上传。
- 批量扫描无法识别学生姓名/学号时进入人工分配队列。

## 5. 智能组题设计

### 5.1 输入条件

- 课程体系：A-Level / IGCSE / AP / DSE / IB
- 年级
- 学科
- 考局
- 教材版本
- 知识点
- 教学模式：初次上课 / 复习课
- 难度分布
- 题型分布
- 真题隔离年限
- 题目数量

### 5.2 推荐逻辑

1. 先按硬条件过滤：课程、学科、考局、知识点、题型。
2. 根据教学模式调整排序：
   - 初次上课：教材/教辅基础题、难度 1-2 星优先。
   - 复习课：真题、综合题、难度 3-5 星优先。
3. 根据历史作业去重，避免近期重复。
4. 根据知识点频率和学生薄弱点加权。
5. 返回可解释推荐理由。

### 5.3 输出示例

```json
{
  "questions": [
    {
      "questionId": 1001,
      "reason": "匹配当前知识点，难度2星，适合初次上课巩固",
      "score": 0.91
    }
  ]
}
```

## 6. 自动批改设计

### 6.1 支持题型优先级

| 题型 | MVP 支持 | 处理方式 |
|---|---:|---|
| 选择题 | 是 | 精确匹配答案 |
| 判断题 | 是 | 精确匹配答案 |
| 填空题 | 是 | 标准化后匹配，支持同义/等价答案 |
| 计算题 | 是 | 过程可辅助判断，最终答案和关键步骤评分 |
| 简单简答题 | 有限支持 | LLM 辅助评分，教师复核 |
| 画图题 | 否 | 标记人工批改 |
| 复杂论述题 | 否 | 标记人工批改 |
| 实验题 | 否 | 标记人工批改 |

### 6.2 批改输出

```json
{
  "submissionId": 2001,
  "results": [
    {
      "questionId": 1001,
      "isCorrect": true,
      "score": 5,
      "maxScore": 5,
      "comment": "答案正确，步骤完整",
      "knowledgePointIds": [301],
      "needsManualReview": false,
      "confidence": 0.95
    }
  ]
}
```

### 6.3 人工复核规则

以下情况必须进入人工复核：

- 模型置信度低于 0.85
- 题型为画图、实验、复杂论述
- 学生答案缺失或 OCR 失败
- AI 给出的分数与规则引擎差异过大
- 学生或教师提出申诉

## 7. 学情总结设计

### 7.1 班级学情总结输入

- 班级平均分
- 作业完成率
- 逾期率
- 各知识点掌握率
- 各题型得分率
- 错题 Top N
- 与历史周期对比

### 7.2 输出结构

```json
{
  "summary": "本周班级整体完成率较高，但氧化还原相关知识点掌握不足。",
  "strengths": ["基础概念题得分稳定"],
  "weaknesses": ["计算题步骤不完整", "氧化还原方程配平错误较多"],
  "suggestions": ["下节课增加配平专项练习", "安排5道中等难度巩固题"]
}
```

## 8. 成本与限额控制

- 每次 AI 调用写入 `ai_call_logs`。
- 教师限额表 `teacher_quotas` 记录周期额度和已用量。
- 到达 80% 时提醒教师和管理员。
- 到达 100% 时禁止高成本 AI 调用或要求管理员审批。
- 批量任务要支持队列，避免并发成本失控。

## 9. 安全与合规

- AI 请求中尽量不要传不必要的学生身份信息。
- 可用学生匿名 ID 替代姓名。
- 上传文件设置有效期和访问权限。
- AI 输出不能直接作为最终成绩，MVP 建议保留教师复核机制。
- 所有 AI 结果应保存模型版本，方便追溯。

## 10. 待确认问题

1. 使用哪家 OCR 和 LLM 服务？
2. AI 批改准确率验收标准是多少？例如客观题 ≥ 99%，简单主观题 ≥ 90%。
3. AI 结果是否必须教师确认后才展示给学生？
4. 每位教师每月 AI 调用预算是多少？
5. 是否需要私有化部署模型？
6. PDF/文档解析是否采用 MinerU 作为 PoC 方案？若采用，需要确认部署方式、GPU/CPU 资源、解析耗时和准确率验收标准。
