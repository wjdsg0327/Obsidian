# React + TypeScript

## 组件类型

### 函数组件

```typescript
// 方式一：React.FC（不推荐，有一些问题）
const Greet: React.FC<{ name: string }> = ({ name }) => {
  return <h1>Hello, {name}</h1>
}

// 方式二：直接类型（推荐）
interface GreetProps {
  name: string
}

function Greet({ name }: GreetProps) {
  return <h1>Hello, {name}</h1>
}
```

### 类组件

```typescript
interface Props {
  name: string
}

interface State {
  count: number
}

class Counter extends React.Component<Props, State> {
  state: State = {
    count: 0
  }

  increment = () => {
    this.setState({ count: this.state.count + 1 })
  }

  render() {
    return (
      <div>
        <p>{this.state.count}</p>
        <button onClick={this.increment}>+</button>
      </div>
    )
  }
}
```

---

## Props 类型

### 基础 Props

```typescript
interface UserCardProps {
  name: string
  age: number
  email?: string  // 可选
  readonly id: number  // 只读
}

function UserCard({ name, age, email, id }: UserCardProps) {
  return (
    <div>
      <h2>{name}</h2>
      <p>Age: {age}</p>
      {email && <p>Email: {email}</p>}
    </div>
  )
}
```

### children Props

```typescript
interface CardProps {
  children: React.ReactNode
}

function Card({ children }: CardProps) {
  return <div className="card">{children}</div>
}

// 使用
<Card>
  <h2>Title</h2>
  <p>Content</p>
</Card>
```

### 函数 Props

```typescript
interface ButtonProps {
  onClick: (event: React.MouseEvent<HTMLButtonElement>) => void
  children: React.ReactNode
}

function Button({ onClick, children }: ButtonProps) {
  return <button onClick={onClick}>{children}</button>
}
```

---

## Hooks 类型

### useState

```typescript
// 自动推断
const [count, setCount] = useState(0)  // number

// 显式类型
const [user, setUser] = useState<User | null>(null)

// 复杂对象
interface FormState {
  name: string
  email: string
}

const [form, setForm] = useState<FormState>({
  name: "",
  email: ""
})
```

### useRef

```typescript
// DOM 引用
const inputRef = useRef<HTMLInputElement>(null)

// 可变值
const countRef = useRef<number>(0)

// 使用
<input ref={inputRef} />
```

### useContext

```typescript
interface ThemeContextType {
  theme: "light" | "dark"
  toggleTheme: () => void
}

const ThemeContext = createContext<ThemeContextType | undefined>(undefined)

function useTheme() {
  const context = useContext(ThemeContext)
  if (!context) {
    throw new Error("useTheme must be used within ThemeProvider")
  }
  return context
}
```

### useReducer

```typescript
interface State {
  count: number
  loading: boolean
}

type Action =
  | { type: "increment" }
  | { type: "decrement" }
  | { type: "setLoading"; payload: boolean }

function reducer(state: State, action: Action): State {
  switch (action.type) {
    case "increment":
      return { ...state, count: state.count + 1 }
    case "decrement":
      return { ...state, count: state.count - 1 }
    case "setLoading":
      return { ...state, loading: action.payload }
  }
}

const [state, dispatch] = useReducer(reducer, {
  count: 0,
  loading: false
})
```

---

## 事件处理

### 常见事件类型

```typescript
// 鼠标事件
function handleClick(event: React.MouseEvent<HTMLButtonElement>) {
  console.log(event.clientX, event.clientY)
}

// 输入事件
function handleChange(event: React.ChangeEvent<HTMLInputElement>) {
  console.log(event.target.value)
}

// 表单事件
function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
  event.preventDefault()
}

// 键盘事件
function handleKeyDown(event: React.KeyboardEvent<HTMLInputElement>) {
  if (event.key === "Enter") {
    // ...
  }
}
```

### 事件处理器

```typescript
interface SearchProps {
  onSearch: (query: string) => void
}

function Search({ onSearch }: SearchProps) {
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    onSearch(e.target.value)
  }

  return <input onChange={handleChange} />
}
```

---

## 泛型组件

### 泛型列表

```typescript
interface ListProps<T> {
  items: T[]
  renderItem: (item: T) => React.ReactNode
  keyExtractor: (item: T) => string
}

function List<T>({ items, renderItem, keyExtractor }: ListProps<T>) {
  return (
    <ul>
      {items.map((item) => (
        <li key={keyExtractor(item)}>{renderItem(item)}</li>
      ))}
    </ul>
  )
}

// 使用
<List
  items={users}
  renderItem={(user) => <span>{user.name}</span>}
  keyExtractor={(user) => user.id.toString()}
/>
```

### 泛型选择器

```typescript
interface SelectProps<T> {
  options: T[]
  value: T
  onChange: (value: T) => void
  getLabel: (option: T) => string
  getKey: (option: T) => string
}

function Select<T>({
  options,
  value,
  onChange,
  getLabel,
  getKey
}: SelectProps<T>) {
  return (
    <select
      value={getKey(value)}
      onChange={(e) => {
        const selected = options.find(
          (opt) => getKey(opt) === e.target.value
        )
        if (selected) onChange(selected)
      }}
    >
      {options.map((option) => (
        <option key={getKey(option)} value={getKey(option)}>
          {getLabel(option)}
        </option>
      ))}
    </select>
  )
}
```

---

## 自定义 Hook 类型

```typescript
// 返回类型
function useLocalStorage<T>(key: string, initialValue: T) {
  const [value, setValue] = useState<T>(() => {
    const item = localStorage.getItem(key)
    return item ? JSON.parse(item) : initialValue
  })

  useEffect(() => {
    localStorage.setItem(key, JSON.stringify(value))
  }, [key, value])

  return [value, setValue] as const
}

// 使用
const [name, setName] = useLocalStorage("name", "")
```

---

## 高级模式

### 渲染 Props

```typescript
interface DataFetcherProps<T> {
  url: string
  children: (data: T | null, loading: boolean) => React.ReactNode
}

function DataFetcher<T>({ url, children }: DataFetcherProps<T>) {
  const [data, setData] = useState<T | null>(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    fetch(url)
      .then((res) => res.json())
      .then((data) => {
        setData(data)
        setLoading(false)
      })
  }, [url])

  return <>{children(data, loading)}</>
}

// 使用
<DataFetcher<User[]> url="/api/users">
  {(users, loading) => {
    if (loading) return <div>Loading...</div>
    return <UserList users={users!} />
  }}
</DataFetcher>
```

### 转发 Ref

```typescript
interface InputProps {
  label: string
  value: string
  onChange: (value: string) => void
}

const Input = React.forwardRef<HTMLInputElement, InputProps>(
  ({ label, value, onChange }, ref) => {
    return (
      <div>
        <label>{label}</label>
        <input
          ref={ref}
          value={value}
          onChange={(e) => onChange(e.target.value)}
        />
      </div>
    )
  }
)
```

---

## 实战示例

### 表单组件

```typescript
interface FormData {
  name: string
  email: string
  password: string
}

interface FormProps {
  onSubmit: (data: FormData) => void
  initialValues?: Partial<FormData>
}

function Form({ onSubmit, initialValues = {} }: FormProps) {
  const [form, setForm] = useState<FormData>({
    name: "",
    email: "",
    password: "",
    ...initialValues
  })

  const handleChange = (field: keyof FormData) => (
    e: React.ChangeEvent<HTMLInputElement>
  ) => {
    setForm((prev) => ({ ...prev, [field]: e.target.value }))
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    onSubmit(form)
  }

  return (
    <form onSubmit={handleSubmit}>
      <input value={form.name} onChange={handleChange("name")} />
      <input value={form.email} onChange={handleChange("email")} />
      <input
        type="password"
        value={form.password}
        onChange={handleChange("password")}
      />
      <button type="submit">Submit</button>
    </form>
  )
}
```

---

*下一节：[11-Vue3+TypeScript](./11-Vue3+TypeScript.md)*
