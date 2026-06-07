# 数据库建表 SQL 草案

> 数据库类型建议：PostgreSQL 15+。如使用 MySQL，需要将 `jsonb`、`timestamptz` 等类型调整。  
> 本 SQL 为 MVP 草案，正式开发前需结合技术栈、租户策略、软删除策略再确认。

```sql
CREATE TABLE schools (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(128) NOT NULL,
  address VARCHAR(255),
  contact_name VARCHAR(64),
  contact_phone VARCHAR(32),
  course_systems JSONB DEFAULT '[]',
  logo_url TEXT,
  status VARCHAR(20) NOT NULL DEFAULT 'active',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  school_id BIGINT REFERENCES schools(id),
  username VARCHAR(80) NOT NULL UNIQUE,
  password_hash VARCHAR(255) NOT NULL,
  role VARCHAR(32) NOT NULL,
  real_name VARCHAR(80),
  phone VARCHAR(32),
  email VARCHAR(128),
  status VARCHAR(20) NOT NULL DEFAULT 'active',
  last_login_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX idx_users_school_role ON users(school_id, role, status);

CREATE TABLE teachers (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL UNIQUE REFERENCES users(id),
  teacher_no VARCHAR(64),
  subject VARCHAR(64),
  position VARCHAR(64),
  permission_profile JSONB DEFAULT '{}',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE classes (
  id BIGSERIAL PRIMARY KEY,
  school_id BIGINT NOT NULL REFERENCES schools(id),
  name VARCHAR(128) NOT NULL,
  grade VARCHAR(32),
  subject VARCHAR(64),
  course_system VARCHAR(32),
  exam_board VARCHAR(64),
  textbook_version VARCHAR(128),
  head_teacher_id BIGINT REFERENCES teachers(id),
  status VARCHAR(20) NOT NULL DEFAULT 'active',
  archived_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX idx_classes_school_subject ON classes(school_id, grade, subject, course_system, exam_board);

CREATE TABLE students (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL UNIQUE REFERENCES users(id),
  class_id BIGINT REFERENCES classes(id),
  student_no VARCHAR(64),
  grade VARCHAR(32),
  status VARCHAR(20) NOT NULL DEFAULT 'active',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX idx_students_class ON students(class_id, student_no);

CREATE TABLE parents (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL UNIQUE REFERENCES users(id),
  contact_phone VARCHAR(32),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE parent_students (
  id BIGSERIAL PRIMARY KEY,
  parent_id BIGINT NOT NULL REFERENCES parents(id),
  student_id BIGINT NOT NULL REFERENCES students(id),
  relation VARCHAR(32),
  UNIQUE(parent_id, student_id)
);

CREATE TABLE knowledge_points (
  id BIGSERIAL PRIMARY KEY,
  parent_id BIGINT REFERENCES knowledge_points(id),
  course_system VARCHAR(32),
  subject VARCHAR(64),
  name VARCHAR(128) NOT NULL,
  level SMALLINT NOT NULL DEFAULT 1,
  sort_no INT DEFAULT 0
);
CREATE INDEX idx_kp_parent ON knowledge_points(parent_id, course_system, subject);

CREATE TABLE syllabi (
  id BIGSERIAL PRIMARY KEY,
  course_system VARCHAR(32) NOT NULL,
  subject VARCHAR(64) NOT NULL,
  exam_board VARCHAR(64),
  version VARCHAR(64),
  title VARCHAR(255)
);

CREATE TABLE textbooks (
  id BIGSERIAL PRIMARY KEY,
  course_system VARCHAR(32),
  subject VARCHAR(64),
  name VARCHAR(255) NOT NULL,
  version VARCHAR(128),
  chapter VARCHAR(128),
  page_start INT,
  page_end INT
);

CREATE TABLE workbooks (
  id BIGSERIAL PRIMARY KEY,
  course_system VARCHAR(32),
  subject VARCHAR(64),
  name VARCHAR(255) NOT NULL,
  version VARCHAR(128),
  chapter VARCHAR(128),
  page_start INT,
  page_end INT
);

CREATE TABLE questions (
  id BIGSERIAL PRIMARY KEY,
  question_code VARCHAR(80),
  course_system VARCHAR(32) NOT NULL,
  grade VARCHAR(32),
  subject VARCHAR(64) NOT NULL,
  exam_board VARCHAR(64),
  knowledge_point_id BIGINT REFERENCES knowledge_points(id),
  syllabus_id BIGINT REFERENCES syllabi(id),
  textbook_id BIGINT REFERENCES textbooks(id),
  workbook_id BIGINT REFERENCES workbooks(id),
  question_type VARCHAR(32) NOT NULL,
  difficulty SMALLINT NOT NULL CHECK (difficulty BETWEEN 1 AND 5),
  source_type VARCHAR(32) NOT NULL,
  source_name VARCHAR(255),
  exam_year INT,
  exam_season VARCHAR(64),
  paper_code VARCHAR(64),
  question_no VARCHAR(64),
  stem TEXT NOT NULL,
  options JSONB DEFAULT '[]',
  answer TEXT NOT NULL,
  explanation TEXT,
  score NUMERIC(6,2) NOT NULL DEFAULT 1,
  tags JSONB DEFAULT '[]',
  copyright_source VARCHAR(255) NOT NULL,
  review_status VARCHAR(32) NOT NULL DEFAULT 'pending_review',
  created_by BIGINT REFERENCES users(id),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX idx_questions_filter ON questions(course_system, subject, exam_board, knowledge_point_id, difficulty, question_type, review_status);

CREATE TABLE assignments (
  id BIGSERIAL PRIMARY KEY,
  teacher_id BIGINT NOT NULL REFERENCES teachers(id),
  class_id BIGINT NOT NULL REFERENCES classes(id),
  title VARCHAR(255) NOT NULL,
  type VARCHAR(32) NOT NULL DEFAULT 'manual',
  mode VARCHAR(32),
  due_at TIMESTAMPTZ,
  late_rule VARCHAR(32) DEFAULT 'allow_late',
  publish_scope VARCHAR(32) DEFAULT 'class',
  status VARCHAR(32) NOT NULL DEFAULT 'draft',
  student_pdf_url TEXT,
  teacher_pdf_url TEXT,
  published_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX idx_assignments_class_status ON assignments(class_id, status, due_at);

CREATE TABLE assignment_items (
  id BIGSERIAL PRIMARY KEY,
  assignment_id BIGINT NOT NULL REFERENCES assignments(id) ON DELETE CASCADE,
  question_id BIGINT NOT NULL REFERENCES questions(id),
  sort_no INT NOT NULL DEFAULT 0,
  score NUMERIC(6,2) NOT NULL DEFAULT 1,
  UNIQUE(assignment_id, question_id)
);

CREATE TABLE submissions (
  id BIGSERIAL PRIMARY KEY,
  assignment_id BIGINT NOT NULL REFERENCES assignments(id),
  student_id BIGINT NOT NULL REFERENCES students(id),
  file_urls JSONB DEFAULT '[]',
  status VARCHAR(32) NOT NULL DEFAULT 'submitted',
  submitted_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  graded_at TIMESTAMPTZ,
  UNIQUE(assignment_id, student_id)
);
CREATE INDEX idx_submissions_assignment_status ON submissions(assignment_id, status);

CREATE TABLE grading_records (
  id BIGSERIAL PRIMARY KEY,
  submission_id BIGINT NOT NULL REFERENCES submissions(id) ON DELETE CASCADE,
  question_id BIGINT NOT NULL REFERENCES questions(id),
  teacher_id BIGINT REFERENCES teachers(id),
  ai_score NUMERIC(6,2),
  final_score NUMERIC(6,2),
  is_correct BOOLEAN,
  feedback TEXT,
  ai_confidence NUMERIC(5,4),
  needs_manual_review BOOLEAN DEFAULT false,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  UNIQUE(submission_id, question_id)
);

CREATE TABLE wrong_questions (
  id BIGSERIAL PRIMARY KEY,
  student_id BIGINT NOT NULL REFERENCES students(id),
  question_id BIGINT NOT NULL REFERENCES questions(id),
  submission_id BIGINT REFERENCES submissions(id),
  note TEXT,
  mastered_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX idx_wrong_student_mastered ON wrong_questions(student_id, mastered_at);

CREATE TABLE words (
  id BIGSERIAL PRIMARY KEY,
  knowledge_point_id BIGINT REFERENCES knowledge_points(id),
  word VARCHAR(128) NOT NULL,
  part_of_speech VARCHAR(32),
  meaning_zh TEXT,
  example_en TEXT,
  chapter VARCHAR(128)
);

CREATE TABLE learning_stats (
  id BIGSERIAL PRIMARY KEY,
  student_id BIGINT REFERENCES students(id),
  class_id BIGINT REFERENCES classes(id),
  knowledge_point_id BIGINT REFERENCES knowledge_points(id),
  stat_date DATE NOT NULL,
  score_rate NUMERIC(5,2),
  mastery_rate NUMERIC(5,2),
  assignment_count INT DEFAULT 0,
  wrong_count INT DEFAULT 0
);
CREATE INDEX idx_learning_stats_scope ON learning_stats(student_id, class_id, knowledge_point_id, stat_date);

CREATE TABLE teacher_quotas (
  id BIGSERIAL PRIMARY KEY,
  teacher_id BIGINT NOT NULL REFERENCES teachers(id),
  period_type VARCHAR(32) NOT NULL DEFAULT 'month',
  period_start DATE NOT NULL,
  period_end DATE NOT NULL,
  ai_calls_limit INT DEFAULT 0,
  ai_calls_used INT DEFAULT 0,
  assignment_limit INT DEFAULT 0,
  assignment_used INT DEFAULT 0,
  pdf_limit INT DEFAULT 0,
  pdf_used INT DEFAULT 0,
  UNIQUE(teacher_id, period_type, period_start)
);

CREATE TABLE ai_call_logs (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT REFERENCES users(id),
  school_id BIGINT REFERENCES schools(id),
  task_type VARCHAR(32) NOT NULL,
  provider VARCHAR(64),
  model VARCHAR(128),
  input_tokens INT DEFAULT 0,
  output_tokens INT DEFAULT 0,
  file_count INT DEFAULT 0,
  cost_amount NUMERIC(12,6) DEFAULT 0,
  status VARCHAR(32) NOT NULL,
  error_message TEXT,
  latency_ms INT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);


CREATE TABLE class_teachers (
  id BIGSERIAL PRIMARY KEY,
  class_id BIGINT NOT NULL REFERENCES classes(id) ON DELETE CASCADE,
  teacher_id BIGINT NOT NULL REFERENCES teachers(id) ON DELETE CASCADE,
  subject VARCHAR(64),
  role VARCHAR(32) DEFAULT 'subject_teacher',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  UNIQUE(class_id, teacher_id, subject)
);
CREATE INDEX idx_class_teachers_teacher ON class_teachers(teacher_id, class_id);

CREATE TABLE assignment_targets (
  id BIGSERIAL PRIMARY KEY,
  assignment_id BIGINT NOT NULL REFERENCES assignments(id) ON DELETE CASCADE,
  target_type VARCHAR(32) NOT NULL, -- class / group / student
  target_id BIGINT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX idx_assignment_targets_assignment ON assignment_targets(assignment_id, target_type, target_id);

CREATE TABLE files (
  id BIGSERIAL PRIMARY KEY,
  school_id BIGINT REFERENCES schools(id),
  owner_user_id BIGINT REFERENCES users(id),
  biz_type VARCHAR(64), -- submission / question_import / pdf / courseware / audio
  biz_id BIGINT,
  original_name VARCHAR(255),
  storage_key TEXT NOT NULL,
  file_url TEXT,
  mime_type VARCHAR(128),
  size_bytes BIGINT,
  status VARCHAR(32) NOT NULL DEFAULT 'uploaded',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX idx_files_biz ON files(biz_type, biz_id);
CREATE INDEX idx_files_owner ON files(owner_user_id, created_at);

CREATE TABLE notifications (
  id BIGSERIAL PRIMARY KEY,
  school_id BIGINT REFERENCES schools(id),
  user_id BIGINT NOT NULL REFERENCES users(id),
  type VARCHAR(64) NOT NULL,
  title VARCHAR(255) NOT NULL,
  content TEXT,
  biz_type VARCHAR(64),
  biz_id BIGINT,
  status VARCHAR(32) NOT NULL DEFAULT 'unread',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  read_at TIMESTAMPTZ
);
CREATE INDEX idx_notifications_user_status ON notifications(user_id, status, created_at);

CREATE TABLE teacher_quality_stats (
  id BIGSERIAL PRIMARY KEY,
  teacher_id BIGINT NOT NULL REFERENCES teachers(id),
  class_id BIGINT REFERENCES classes(id),
  stat_date DATE NOT NULL,
  period_type VARCHAR(32) NOT NULL DEFAULT 'week',
  grading_timely_rate NUMERIC(5,2),
  grading_completion_rate NUMERIC(5,2),
  report_generation_rate NUMERIC(5,2),
  student_score_improvement_rate NUMERIC(5,2),
  class_mastery_rate NUMERIC(5,2),
  ai_teaching_score NUMERIC(5,2),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX idx_teacher_quality_scope ON teacher_quality_stats(teacher_id, class_id, stat_date);

CREATE TABLE class_compare_stats (
  id BIGSERIAL PRIMARY KEY,
  class_id BIGINT NOT NULL REFERENCES classes(id),
  stat_date DATE NOT NULL,
  grade VARCHAR(32),
  subject VARCHAR(64),
  exam_board VARCHAR(64),
  avg_score NUMERIC(6,2),
  completion_rate NUMERIC(5,2),
  late_rate NUMERIC(5,2),
  excellent_rate NUMERIC(5,2),
  pass_rate NUMERIC(5,2),
  rank_no INT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX idx_class_compare_scope ON class_compare_stats(grade, subject, exam_board, stat_date);

CREATE TABLE operation_logs (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT REFERENCES users(id),
  school_id BIGINT REFERENCES schools(id),
  action VARCHAR(128) NOT NULL,
  target_type VARCHAR(64),
  target_id BIGINT,
  ip VARCHAR(64),
  user_agent TEXT,
  detail JSONB DEFAULT '{}',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE system_logs (
  id BIGSERIAL PRIMARY KEY,
  level VARCHAR(16) NOT NULL,
  module VARCHAR(64),
  message TEXT NOT NULL,
  detail JSONB DEFAULT '{}',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
```

## 后续建议新增表

- `class_teachers`：一个班级多个任课教师。
- `assignment_targets`：作业发布到指定学生/小组。
- `files`：统一文件元数据。
- `notifications`：站内消息通知。
