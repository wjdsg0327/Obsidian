# Vue3 + TypeScript

## 组件类型

### defineComponent

```typescript
import { defineComponent } from "vue"

export default defineComponent({
  props: {
    name: {
      type: String,
      required: true
    },
    age: {
      type: Number,
      default: 0
    }
  },
  setup(props) {
    // props 有完整的类型推断
    console.log(props.name)  // string
    console.log(props.age)   // number
  }
})
```

### script setup（推荐）

```vue
<script setup lang="ts">
// Props 类型
interface Props {
  name: string
  age?: number
}

const props = defineProps<Props>()

// 带默认值的 Props
const propsWithDefaults = withDefaults(defineProps<Props>(), {
  age: 0
})
</script>
```

---

## Props 类型

### 基础 Props

```vue
<script setup lang="ts">
interface UserProps {
  id: number
  name: string
  email?: string
  role: "admin" | "user" | "guest"
}

const props = defineProps<UserProps>()
</script>

<template>
  <div>
    <h2>{{ name }}</h2>
    <p>{{ email }}</p>
  </div>
</template>
```

### 复杂 Props

```vue
<script setup lang="ts">
interface User {
  id: number
  name: string
}

interface Props {
  users: User[]
  selectedId: number | null
  onSelect: (user: User) => void
  renderHeader?: () => string
}

const props = defineProps<Props>()
</script>
```

### Props 验证

```vue
<script setup lang="ts">
import type { PropType } from "vue"

interface User {
  id: number
  name: string
}

const props = defineProps({
  user: {
    type: Object as PropType<User>,
    required: true
  },
  status: {
    type: String as () => "active" | "inactive",
    default: "active"
  }
})
</script>
```

---

## Emits 类型

### 基础 Emits

```vue
<script setup lang="ts">
// 定义 emits 类型
const emit = defineEmits<{
  (e: "update", value: string): void
  (e: "delete", id: number): void
  (e: "submit"): void
}>()

// 使用
function handleUpdate(value: string) {
  emit("update", value)
}
</script>
```

### 对象语法（Vue 3.3+）

```vue
<script setup lang="ts">
const emit = defineEmits<{
  update: [value: string]
  delete: [id: number]
  submit: []
}>()
</script>
```

---

## Ref 类型

### 基础 Ref

```vue
<script setup lang="ts">
import { ref } from "vue"

// 自动推断
const count = ref(0)  // Ref<number>

// 显式类型
const name = ref<string>("张三")

// 复杂类型
interface User {
  id: number
  name: string
}

const user = ref<User | null>(null)

// DOM 引用
const inputRef = ref<HTMLInputElement | null>(null)
</script>

<template>
  <input ref="inputRef" />
</template>
```

### 复杂 Ref

```vue
<script setup lang="ts">
import { ref } from "vue"

interface FormState {
  step: number
  data: {
    name: string
    email: string
  }
  errors: Record<string, string>
}

const form = ref<FormState>({
  step: 1,
  data: { name: "", email: "" },
  errors: {}
})
</script>
```

---

## Reactive 类型

```vue
<script setup lang="ts">
import { reactive } from "vue"

interface FormState {
  name: string
  email: string
  loading: boolean
}

const form = reactive<FormState>({
  name: "",
  email: "",
  loading: false
})

// 直接修改
form.name = "张三"
form.loading = true
</script>
```

---

## Computed 类型

```vue
<script setup lang="ts">
import { computed, ref } from "vue"

const items = ref([1, 2, 3, 4, 5])

// 自动推断返回类型
const doubled = computed(() => items.value.map((n) => n * 2))

// 显式类型
const sum = computed<number>(() => {
  return items.value.reduce((acc, n) => acc + n, 0)
})

// 可写 computed
const fullName = computed({
  get: () => `${firstName.value} ${lastName.value}`,
  set: (value: string) => {
    const [first, last] = value.split(" ")
    firstName.value = first
    lastName.value = last
  }
})
</script>
```

---

## 函数类型

```vue
<script setup lang="ts">
// 普通函数
function handleClick(event: MouseEvent) {
  console.log(event.clientX, event.clientY)
}

// 异步函数
async function fetchUser(id: number): Promise<User> {
  const response = await fetch(`/api/users/${id}`)
  return response.json()
}

// 事件处理器
function handleSubmit(e: Event) {
  e.preventDefault()
  const form = e.target as HTMLFormElement
  const formData = new FormData(form)
}
</script>

<template>
  <button @click="handleClick">Click</button>
  <form @submit="handleSubmit">...</form>
</template>
```

---

## 组合式函数（Composables）

### 基础组合式函数

```typescript
// composables/useCounter.ts
import { ref, computed } from "vue"

export function useCounter(initialValue = 0) {
  const count = ref(initialValue)
  const doubled = computed(() => count.value * 2)

  function increment() {
    count.value++
  }

  function decrement() {
    count.value--
  }

  return {
    count,
    doubled,
    increment,
    decrement
  }
}
```

### 带类型的组合式函数

```typescript
// composables/useFetch.ts
import { ref, watchEffect } from "vue"

interface UseFetchReturn<T> {
  data: Ref<T | null>
  loading: Ref<boolean>
  error: Ref<Error | null>
  refetch: () => Promise<void>
}

export function useFetch<T>(url: string): UseFetchReturn<T> {
  const data = ref<T | null>(null)
  const loading = ref(true)
  const error = ref<Error | null>(null)

  async function fetchData() {
    loading.value = true
    error.value = null
    try {
      const response = await fetch(url)
      if (!response.ok) throw new Error(response.statusText)
      data.value = await response.json()
    } catch (e) {
      error.value = e as Error
    } finally {
      loading.value = false
    }
  }

  watchEffect(() => {
    fetchData()
  })

  return { data, loading, error, refetch: fetchData }
}
```

### 使用组合式函数

```vue
<script setup lang="ts">
import { useFetch } from "@/composables/useFetch"

interface User {
  id: number
  name: string
}

const { data: users, loading, error } = useFetch<User[]>("/api/users")
</script>

<template>
  <div v-if="loading">Loading...</div>
  <div v-else-if="error">{{ error.message }}</div>
  <ul v-else>
    <li v-for="user in users" :key="user.id">{{ user.name }}</li>
  </ul>
</template>
```

---

## provide/inject 类型

```typescript
// 父组件
import { provide, ref } from "vue"

interface ThemeContext {
  theme: "light" | "dark"
  toggleTheme: () => void
}

const theme = ref<"light" | "dark">("light")

provide<ThemeContext>("theme", {
  theme: theme.value,
  toggleTheme: () => {
    theme.value = theme.value === "light" ? "dark" : "light"
  }
})
```

```typescript
// 子组件
import { inject } from "vue"

interface ThemeContext {
  theme: "light" | "dark"
  toggleTheme: () => void
}

const theme = inject<ThemeContext>("theme")
```

---

## 模板引用类型

```vue
<script setup lang="ts">
import { ref, onMounted } from "vue"

// DOM 元素引用
const inputRef = ref<HTMLInputElement | null>(null)

// 组件引用
const childRef = ref<InstanceType<typeof ChildComponent> | null>(null)

onMounted(() => {
  inputRef.value?.focus()
  childRef.value?.someMethod()
})
</script>

<template>
  <input ref="inputRef" />
  <ChildComponent ref="childRef" />
</template>
```

---

## Pinia 类型

### Store 类型

```typescript
// stores/user.ts
import { defineStore } from "pinia"

interface UserState {
  id: number | null
  name: string
  email: string
  loading: boolean
}

export const useUserStore = defineStore("user", {
  state: (): UserState => ({
    id: null,
    name: "",
    email: "",
    loading: false
  }),

  getters: {
    isLoggedIn: (state) => state.id !== null,
    displayName: (state) => state.name || state.email
  },

  actions: {
    async login(email: string, password: string) {
      this.loading = true
      try {
        const user = await api.login(email, password)
        this.id = user.id
        this.name = user.name
        this.email = user.email
      } finally {
        this.loading = false
      }
    },

    logout() {
      this.id = null
      this.name = ""
      this.email = ""
    }
  }
})
```

### Setup Store（推荐）

```typescript
// stores/counter.ts
import { ref, computed } from "vue"
import { defineStore } from "pinia"

export const useCounterStore = defineStore("counter", () => {
  const count = ref(0)
  const doubled = computed(() => count.value * 2)

  function increment() {
    count.value++
  }

  function decrement() {
    count.value--
  }

  return { count, doubled, increment, decrement }
})
```

---

## Vue Router 类型

```typescript
// router/index.ts
import { createRouter, createWebHistory } from "vue-router"

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      component: () => import("@/views/Home.vue")
    },
    {
      path: "/user/:id",
      component: () => import("@/views/User.vue"),
      props: true
    }
  ]
})

export default router
```

```vue
<script setup lang="ts">
import { useRoute, useRouter } from "vue-router"

const route = useRoute()
const router = useRouter()

// 获取路由参数
const userId = route.params.id as string

// 编程式导航
router.push({ name: "user", params: { id: "123" } })
</script>
```

---

## 实战示例

### 表单组件

```vue
<script setup lang="ts">
import { reactive, ref } from "vue"

interface FormData {
  username: string
  password: string
  remember: boolean
}

interface FormErrors {
  username?: string
  password?: string
}

const form = reactive<FormData>({
  username: "",
  password: "",
  remember: false
})

const errors = ref<FormErrors>({})
const loading = ref(false)

function validate(): boolean {
  errors.value = {}
  if (!form.username) {
    errors.value.username = "请输入用户名"
  }
  if (!form.password) {
    errors.value.password = "请输入密码"
  }
  return Object.keys(errors.value).length === 0
}

async function handleSubmit() {
  if (!validate()) return
  loading.value = true
  try {
    await login(form.username, form.password)
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <form @submit.prevent="handleSubmit">
    <div>
      <input v-model="form.username" placeholder="用户名" />
      <span v-if="errors.username">{{ errors.username }}</span>
    </div>
    <div>
      <input v-model="form.password" type="password" placeholder="密码" />
      <span v-if="errors.password">{{ errors.password }}</span>
    </div>
    <label>
      <input v-model="form.remember" type="checkbox" />
      记住我
    </label>
    <button :disabled="loading">登录</button>
  </form>
</template>
```

---

*下一节：[12-常见问题与技巧](./12-常见问题与技巧.md)*
