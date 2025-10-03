是的，Go 的 `fmt` 包提供了多个打印相关的函数，除了 `fmt.Print`、`fmt.Println` 之外，**`fmt.Sprintf` 是非常重要的一员**，它用于**格式化并返回字符串，而不是直接输出到控制台**。

---

### ✅ `fmt` 包中常用的打印/格式化函数

| 函数 | 作用 | 输出目标 |
|------|------|----------|
| `fmt.Print` | 打印内容，不换行 | 标准输出（终端） |
| `fmt.Println` | 打印内容，自动换行 | 标准输出 |
| `fmt.Printf` | 格式化打印，支持 `%v`、`%s` 等动词 | 标准输出 |
| `fmt.Sprintf` | 格式化并**返回字符串** | 返回 `string` |
| `fmt.Fprint` | 打印到任意 `io.Writer`（如文件） | 指定的写入目标 |
| `fmt.Fprintln` | 类似 `Println`，但输出到 `io.Writer` | 指定的写入目标 |
| `fmt.Fprintf` | 格式化打印到 `io.Writer` | 指定的写入目标 |

---

### 📌 1. `fmt.Sprintf`：格式化并返回字符串

**用途**：将变量格式化为字符串，用于拼接、日志、返回值等场景。

```go
name := "Alice"
age := 30

// 使用 Sprintf 返回字符串，而不是打印
greeting := fmt.Sprintf("Hello, %s! You are %d years old.", name, age)

fmt.Println(greeting)
// 输出：Hello, Alice! You are 30 years old.
```

> ✅ 返回类型是 `string`，非常适合用于：
> - 构造 SQL 语句
> - 生成日志消息
> - 拼接 URL
> - 返回错误信息等

---

### 📌 2. `fmt.Printf`：带格式的打印

**用途**：在终端直接打印格式化内容。

```go
fmt.Printf("Name: %s, Age: %d\n", name, age)
// 输出：Name: Alice, Age: 30
```

区别：
- `fmt.Printf`：直接输出到屏幕
- `fmt.Sprintf`：返回字符串，由你决定怎么用

---

### 📌 3. `fmt.Fprint` 系列：输出到文件或其他目标

这些函数可以将内容写入**文件、网络连接、缓冲区**等实现了 `io.Writer` 接口的目标。

#### 示例：写入文件

```go
file, _ := os.Create("output.txt")
defer file.Close()

fmt.Fprint(file, "Hello, File!\n")
fmt.Fprintf(file, "User: %s, Score: %d\n", "Bob", 95)
fmt.Fprintln(file, "Done")
```

这会把内容写入 `output.txt` 文件。

#### 常见 `io.Writer` 目标：
- `*os.File`（文件）
- `*bytes.Buffer`（内存缓冲区）
- `http.ResponseWriter`（HTTP 响应）
- `net.Conn`（网络连接）

---
