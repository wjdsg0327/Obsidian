# LM Eval / lm-evaluation-harness 使用教程与实战案例

整理时间：2026-06-03  
项目：EleutherAI `lm-evaluation-harness`  
官方仓库：https://github.com/EleutherAI/lm-evaluation-harness

---

## 1. LM Eval 是什么

LM Eval 通常指 EleutherAI 的 `lm-evaluation-harness`，命令行工具叫 `lm-eval` 或 `lm_eval`。它是目前很常用的大模型评测框架，Hugging Face Open LLM Leaderboard 也用过它做后端。

它适合做这些事：

- 用标准 benchmark 评估模型能力，例如 `mmlu`、`hellaswag`、`arc_easy`、`arc_challenge`、`gsm8k`、`ifeval`。
- 评估本地 Hugging Face 模型。
- 评估 vLLM / TGI / OpenAI-compatible API 部署出来的模型。
- 写自己的 YAML 任务，用私有数据集做评测。
- 保存模型输出，分析错题和失败案例。

核心思路很简单：

```text
模型 + 任务集 + 提示模板 + 指标 = 评测结果
```

---

## 2. 安装

现在官方版本的安装方式有一点变化：基础包不一定自带 `torch` / `transformers`，建议按后端安装 extras。

### 2.1 基础安装

```bash
pip install lm_eval
```

### 2.2 Hugging Face 本地模型评测

```bash
pip install "lm_eval[hf]"
```

### 2.3 vLLM 后端

```bash
pip install "lm_eval[vllm]"
```

### 2.4 API 模型 / OpenAI 兼容接口

```bash
pip install "lm_eval[api]"
```

### 2.5 一次装多个后端

```bash
pip install "lm_eval[hf,vllm,api]"
```

### 2.6 从源码安装

```bash
git clone --depth 1 https://github.com/EleutherAI/lm-evaluation-harness
cd lm-evaluation-harness
pip install -e .
```

---

## 3. CLI 基础命令

新版 CLI 主要有三个子命令：

```bash
lm-eval run       # 跑评测
lm-eval ls        # 查看任务、分组、标签
lm-eval validate  # 校验任务配置
```

旧写法仍兼容，例如：

```bash
lm_eval --model hf --tasks hellaswag
```

但推荐新写法：

```bash
lm-eval run --model hf --tasks hellaswag
```

查看可用任务：

```bash
lm-eval ls tasks
lm-eval ls groups
lm-eval ls subtasks
lm-eval ls tags
```

---

## 4. 最小实战：评测 GPT-2 的 HellaSwag

```bash
lm-eval run \
  --model hf \
  --model_args pretrained=gpt2 \
  --tasks hellaswag \
  --device cuda:0 \
  --batch_size 8
```

如果没有 GPU：

```bash
lm-eval run \
  --model hf \
  --model_args pretrained=gpt2 \
  --tasks hellaswag \
  --device cpu \
  --batch_size 1
```

只想试跑 10 条，避免一上来跑很久：

```bash
lm-eval run \
  --model hf \
  --model_args pretrained=gpt2 \
  --tasks hellaswag \
  --limit 10 \
  --device cuda:0
```

注意：`--limit` 只适合调试，不适合正式汇报结果。

---

## 5. 常用参数解释

### 5.1 模型相关

```bash
--model hf
```

指定模型后端。常见值：

- `hf`：Hugging Face transformers 本地加载。
- `vllm`：用 vLLM 加载模型。
- `local-completions`：OpenAI-compatible completions API。
- `local-chat-completions`：OpenAI-compatible chat completions API。

```bash
--model_args pretrained=Qwen/Qwen2.5-7B-Instruct,dtype=bfloat16
```

传模型参数。常见参数：

- `pretrained=模型名或本地路径`
- `dtype=float16` / `dtype=bfloat16` / `dtype=float32`
- `revision=分支或 checkpoint`
- `trust_remote_code=True`
- `parallelize=True`

### 5.2 任务相关

```bash
--tasks hellaswag,arc_easy,gsm8k
```

可以传单个任务、多个任务、任务组。

```bash
--num_fewshot 5
```

few-shot 示例数量。

```bash
--apply_chat_template
```

对 instruct/chat 模型很重要，会用 tokenizer 的 chat template 包装 prompt。

```bash
--fewshot_as_multiturn
```

把 few-shot 示例按多轮对话形式组织，适合 chat 模型。

### 5.3 性能相关

```bash
--batch_size auto
--batch_size auto:4
```

自动寻找合适 batch size。`auto:4` 表示过程中重新估计 4 次。

```bash
--use_cache ./cache/lm_eval_cache_
```

缓存模型响应，重复实验会省时间。

### 5.4 输出相关

```bash
--output_path ./results/
--log_samples
```

保存结果和每条样本的输入输出，方便分析错题。

```bash
--write_out
```

打印前几条 prompt，调试模板时很好用。

```bash
--show_config
```

显示实际任务配置。

---

## 6. 实战案例一：评测本地 Hugging Face 模型

以 Qwen2.5-7B-Instruct 为例：

```bash
lm-eval run \
  --model hf \
  --model_args pretrained=Qwen/Qwen2.5-7B-Instruct,dtype=bfloat16,trust_remote_code=True \
  --tasks arc_easy,arc_challenge,hellaswag \
  --device cuda:0 \
  --batch_size auto \
  --output_path ./results/qwen2.5-7b/ \
  --log_samples
```

如果是 instruct/chat 模型，并且任务是生成类，建议加：

```bash
--apply_chat_template
```

完整例子：

```bash
lm-eval run \
  --model hf \
  --model_args pretrained=Qwen/Qwen2.5-7B-Instruct,dtype=bfloat16,trust_remote_code=True \
  --tasks gsm8k \
  --apply_chat_template \
  --device cuda:0 \
  --batch_size auto \
  --output_path ./results/qwen2.5-7b-gsm8k/ \
  --log_samples
```

---

## 7. 实战案例二：用 vLLM 直接跑评测

如果装了 vLLM 后端：

```bash
pip install "lm_eval[vllm]"
```

评测：

```bash
lm-eval run \
  --model vllm \
  --model_args pretrained=Qwen/Qwen2.5-7B-Instruct,dtype=bfloat16,tensor_parallel_size=1,trust_remote_code=True \
  --tasks gsm8k,ifeval \
  --apply_chat_template \
  --batch_size auto \
  --output_path ./results/qwen2.5-7b-vllm/ \
  --log_samples
```

vLLM 适合大模型、高吞吐评测。GPU 显存够时，比纯 HF backend 更适合批量跑。

---

## 8. 实战案例三：评测 OpenAI-compatible API

这个最适合生产部署验证：模型已经用 vLLM、TGI、SGLang、Ollama 或其它服务暴露成 OpenAI 风格接口。

### 8.1 启动 vLLM OpenAI API 服务

示例：

```bash
vllm serve Qwen/Qwen2.5-7B-Instruct \
  --host 0.0.0.0 \
  --port 8000 \
  --dtype bfloat16 \
  --trust-remote-code
```

### 8.2 用 lm-eval 跑 chat completions

```bash
lm-eval run \
  --model local-chat-completions \
  --tasks gsm8k_cot,ifeval \
  --model_args model=Qwen/Qwen2.5-7B-Instruct,base_url=http://localhost:8000/v1/chat/completions,num_concurrent=16,max_retries=3,tokenized_requests=False \
  --apply_chat_template \
  --fewshot_as_multiturn \
  --output_path ./results/qwen-api/ \
  --log_samples
```

重点参数：

- `model=`：模型名，通常和服务端模型名一致。
- `base_url=`：注意 chat 接口通常是 `/v1/chat/completions`。
- `num_concurrent=`：并发数。太大会打爆服务，建议从 4、8、16 逐步加。
- `tokenized_requests=False`：chat API 常用。
- `--apply_chat_template`：让 prompt 更贴合 chat 模型格式。
- `--fewshot_as_multiturn`：few-shot 按多轮消息组织。

### 8.3 completions 接口

如果服务提供的是 completions API：

```bash
lm-eval run \
  --model local-completions \
  --tasks hellaswag \
  --model_args model=your-model,base_url=http://localhost:8000/v1/completions,num_concurrent=16,max_retries=3 \
  --output_path ./results/local-completions/ \
  --log_samples
```

注意：有些 benchmark 需要 loglikelihood，API 后端未必支持；chat completions 更适合生成式任务，例如 `gsm8k_cot`、`ifeval`。

---

## 9. 实战案例四：评测 DeepSeek-R1 / Qwen3 这类 thinking 模型

一些模型会输出 `<think>...</think>` 这类思考过程。评测时通常只想拿最终答案算分，需要配置 `think_end_token`。

vLLM / SGLang 示例：

```bash
lm-eval run \
  --model vllm \
  --model_args pretrained=Qwen/Qwen3-32B,enable_thinking=True,think_end_token="</think>" \
  --tasks gsm8k \
  --apply_chat_template
```

HF backend 可以用 token id：

```bash
lm-eval run \
  --model hf \
  --model_args pretrained=Qwen/Qwen3-32B,enable_thinking=True,think_end_token=200008 \
  --tasks gsm8k \
  --apply_chat_template
```

注意：thinking 模式通常只适合生成类任务，不适合 loglikelihood 类任务。

---

## 10. 实战案例五：自定义中文选择题任务

假设有一个本地 JSONL 文件：

`data/chinese_exam.jsonl`

```jsonl
{"question":"Python 中列表追加元素的方法是？","A":"push","B":"append","C":"add","D":"insertEnd","answer":"B"}
{"question":"Linux 查看当前目录文件的命令是？","A":"pwd","B":"cd","C":"ls","D":"mkdir","answer":"C"}
```

创建目录：

```bash
mkdir -p custom_tasks/chinese_exam
```

创建 `custom_tasks/chinese_exam/chinese_exam.yaml`：

```yaml
task: chinese_exam
dataset_path: json
dataset_kwargs:
  data_files:
    test: data/chinese_exam.jsonl
test_split: test
output_type: multiple_choice
doc_to_text: "题目：{{question}}\nA. {{A}}\nB. {{B}}\nC. {{C}}\nD. {{D}}\n答案："
doc_to_choice: ["A", "B", "C", "D"]
doc_to_target: "{{ ['A', 'B', 'C', 'D'].index(answer) }}"
metric_list:
  - metric: acc
    aggregation: mean
    higher_is_better: true
metadata:
  version: 1.0
```

校验任务：

```bash
lm-eval validate \
  --tasks chinese_exam \
  --include_path ./custom_tasks
```

运行评测：

```bash
lm-eval run \
  --model hf \
  --model_args pretrained=Qwen/Qwen2.5-7B-Instruct,dtype=bfloat16,trust_remote_code=True \
  --tasks chinese_exam \
  --include_path ./custom_tasks \
  --apply_chat_template \
  --device cuda:0 \
  --batch_size auto \
  --output_path ./results/chinese_exam/ \
  --log_samples \
  --write_out
```

`--write_out` 很重要：先看实际 prompt 长什么样，确认没有模板错误。

---

## 11. 自定义生成式任务：问答 / 摘要 / 数学

如果数据是：

```jsonl
{"question":"小明有3个苹果，又买了5个，一共有几个？","answer":"8"}
```

YAML 可以这样写：

```yaml
task: chinese_math_qa
dataset_path: json
dataset_kwargs:
  data_files:
    test: data/chinese_math_qa.jsonl
test_split: test
output_type: generate_until
doc_to_text: "请解答下面的问题，只输出最终答案。\n问题：{{question}}\n答案："
doc_to_target: "{{answer}}"
generation_kwargs:
  until:
    - "\n"
  do_sample: false
  temperature: 0.0
metric_list:
  - metric: exact_match
    aggregation: mean
    higher_is_better: true
filter_list:
  - name: strict-match
    filter:
      - function: regex
        regex_pattern: "(-?[0-9]+(?:\\.[0-9]+)?)"
      - function: take_first
metadata:
  version: 1.0
```

运行：

```bash
lm-eval run \
  --model local-chat-completions \
  --tasks chinese_math_qa \
  --include_path ./custom_tasks \
  --model_args model=Qwen/Qwen2.5-7B-Instruct,base_url=http://localhost:8000/v1/chat/completions,num_concurrent=8,max_retries=3,tokenized_requests=False \
  --apply_chat_template \
  --output_path ./results/chinese_math_qa/ \
  --log_samples
```

---

## 12. 结果怎么看

命令跑完后会输出类似表格：

```text
| Tasks     | Version | Filter | n-shot | Metric | Value | Stderr |
| hellaswag | ...     | none   | 0      | acc    | 0.xxx | ...    |
| hellaswag | ...     | none   | 0      | acc_norm | 0.xxx | ...  |
```

常见指标：

- `acc`：准确率。
- `acc_norm`：归一化准确率，常用于多选题，减少长选项/短选项造成的偏差。
- `exact_match`：生成文本和标准答案完全匹配。
- `f1`：常见于问答/抽取类。
- `perplexity` / `word_perplexity`：困惑度，越低越好。

如果用了：

```bash
--output_path ./results/xxx --log_samples
```

会保存详细 JSON，里面能看到每条样本的：

- prompt
- gold answer
- model output
- metric result
- filter 后的答案

这是做误差分析最有用的文件。

---

## 13. 推荐评测组合

### 13.1 快速 sanity check

```bash
--tasks hellaswag,arc_easy --limit 20
```

用于确认流程没问题。

### 13.2 通用英文能力

```bash
--tasks hellaswag,arc_easy,arc_challenge,mmlu
```

### 13.3 数学推理

```bash
--tasks gsm8k,gsm8k_cot
```

### 13.4 指令遵循

```bash
--tasks ifeval
```

### 13.5 中文能力

LM Eval 官方内置中文任务随版本变化，建议先查：

```bash
lm-eval ls tasks | grep -i chinese
lm-eval ls tasks | grep -i cmmlu
lm-eval ls tasks | grep -i ceval
```

如果没有合适任务，就用自定义 YAML 跑自己的中文数据集。

---

## 14. 正式评测建议流程

### 第一步：确认模型能跑

```bash
lm-eval run \
  --model hf \
  --model_args pretrained=你的模型 \
  --tasks hellaswag \
  --limit 5 \
  --write_out
```

### 第二步：确认 prompt 格式

对 chat 模型加：

```bash
--apply_chat_template --write_out
```

人工看几条 prompt，确认没有奇怪格式。

### 第三步：小样本试跑

```bash
--limit 100 --log_samples --output_path ./results/debug/
```

### 第四步：正式跑全量

去掉 `--limit`，固定：

- 模型版本
- lm-eval 版本
- 数据集版本
- prompt 配置
- 随机种子

### 第五步：做错题分析

看 `--log_samples` 生成的文件，按任务分类分析失败原因。

---

## 15. 常见坑

### 15.1 `--limit` 结果不能当正式成绩

`--limit` 只截取部分样本，适合调试，不适合对外汇报。

### 15.2 chat 模型没加 chat template

Instruct 模型如果直接用普通文本 prompt，结果可能明显偏低。优先试：

```bash
--apply_chat_template
```

### 15.3 API 后端不支持 loglikelihood

有些任务依赖 loglikelihood，比如多选题常见的比较选项概率。Chat API 不一定支持。遇到报错时，换生成式任务，或者用 `hf` / `vllm` 后端直接加载模型。

### 15.4 batch size 太大 OOM

先用：

```bash
--batch_size auto
```

或手动降到 1、2、4。

### 15.5 自定义任务答案格式不稳定

生成式任务建议加 filter，比如 regex 抽数字、抽选项字母，否则模型多输出一句解释就可能 exact match 失败。

### 15.6 没保存 samples，后面无法分析

正式跑建议总是加：

```bash
--output_path ./results/xxx --log_samples
```

### 15.7 模型输出思考链影响打分

Qwen3 / DeepSeek-R1 类模型需要处理 `think_end_token`，否则 metric 可能拿到一堆思考文本。

---

## 16. 一个推荐目录结构

```text
lm-eval-project/
  data/
    chinese_exam.jsonl
    chinese_math_qa.jsonl
  custom_tasks/
    chinese_exam/
      chinese_exam.yaml
    chinese_math_qa/
      chinese_math_qa.yaml
  results/
    qwen2.5-7b/
    qwen-api/
  scripts/
    run_eval_hf.sh
    run_eval_api.sh
```

---

## 17. 可直接改用的脚本

### 17.1 `scripts/run_eval_hf.sh`

```bash
#!/usr/bin/env bash
set -euo pipefail

MODEL="Qwen/Qwen2.5-7B-Instruct"
TASKS="gsm8k,ifeval"
OUT="./results/qwen2.5-7b-hf"

lm-eval run \
  --model hf \
  --model_args pretrained=${MODEL},dtype=bfloat16,trust_remote_code=True \
  --tasks ${TASKS} \
  --apply_chat_template \
  --device cuda:0 \
  --batch_size auto \
  --output_path ${OUT} \
  --log_samples \
  --show_config
```

### 17.2 `scripts/run_eval_api.sh`

```bash
#!/usr/bin/env bash
set -euo pipefail

MODEL="Qwen/Qwen2.5-7B-Instruct"
BASE_URL="http://localhost:8000/v1/chat/completions"
TASKS="gsm8k_cot,ifeval"
OUT="./results/qwen2.5-7b-api"

lm-eval run \
  --model local-chat-completions \
  --tasks ${TASKS} \
  --model_args model=${MODEL},base_url=${BASE_URL},num_concurrent=8,max_retries=3,tokenized_requests=False \
  --apply_chat_template \
  --fewshot_as_multiturn \
  --output_path ${OUT} \
  --log_samples \
  --show_config
```

---

## 18. 资料来源

- 官方仓库：https://github.com/EleutherAI/lm-evaluation-harness
- 官方 CLI 文档：https://github.com/EleutherAI/lm-evaluation-harness/blob/main/docs/interface.md
- 官方任务 YAML 配置文档：https://github.com/EleutherAI/lm-evaluation-harness/blob/main/docs/task_guide.md
- vLLM / TGI + OpenAI API 实战参考：https://www.philschmid.de/evaluate-llms-with-lm-eval-and-tgi-vllm

---

## 19. 老王实战建议

如果你只是想快速把自己的模型跑起来，我建议按这个顺序：

1. 先用 `--limit 5 --write_out` 跑通。
2. 本地模型优先用 `hf`，大模型/生产部署优先用 `vllm` 或 `local-chat-completions`。
3. Instruct/chat 模型优先加 `--apply_chat_template`。
4. 正式评测一定加 `--output_path` 和 `--log_samples`。
5. 私有中文任务直接写 YAML，自定义数据用 JSONL 最省事。

最短可用命令：

```bash
lm-eval run \
  --model hf \
  --model_args pretrained=Qwen/Qwen2.5-7B-Instruct,dtype=bfloat16,trust_remote_code=True \
  --tasks gsm8k,ifeval \
  --apply_chat_template \
  --device cuda:0 \
  --batch_size auto \
  --output_path ./results/qwen-eval/ \
  --log_samples
```
