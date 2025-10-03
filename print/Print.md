Go 语言的 `fmt` 包提供了多个打印函数（如 `fmt.Print`、`fmt.Println`、`fmt.Printf` 等），它们都支持**几乎所有内置类型和复合类型**的直接打印。Go 的设计原则是“开箱即用”，因此 `fmt` 包对类型有很强的自动识别和格式化能力。

---

### ✅ 一、`fmt.Print` 和 `fmt.Println` 支持的类型

这两个函数可以**直接打印任何类型**，无需格式化动词。它们会自动调用类型的默认格式输出。

#### 1. **基本数据类型**

| 类型 | 示例 | 输出 |
|------|------|------|
| `int`, `int8/16/32/64` | `42` | `42` |
| `uint`, `uint8/16/32/64` | `255` | `255` |
| `float32`, `float64` | `3.14` | `3.14` |
| `bool` | `true`, `false` | `true` |
| `string` | `"hello"` | `hello` |
| `complex64`, `complex128` | `1+2i` | `(1+2i)` |
| `byte`（= `uint8`） | `'A'` | `65`（数值）或 `%c` 打印字符 |
| `rune`（= `int32`） | `'世'` | `19990` 或 `%c` 打印字符 |

#### 2. **复合类型**

| 类型 | 示例 | 输出格式 |
|------|------|----------|
| **数组** `[3]int{1,2,3}` | `[1 2 3]` | 空格分隔，括号包围 |
| **切片** `[]int{1,2,3}` | `[1 2 3]` | 同数组 |
| **map** | `map[a:1 b:2]` | 按键排序或乱序输出 |
| **结构体** | `struct{A int}{A:1}` | `{1}` 或 `{A:1}`（取决于是否导出） |

示例：

```go
package main

import "fmt"

func main() {
    fmt.Println("hello")                    // string
    fmt.Println(42)                         // int
    fmt.Println(3.14)                       // float64
    fmt.Println(true)                       // bool
    fmt.Println([3]int{1, 2, 3})            // 数组: [1 2 3]
    fmt.Println([]int{4, 5, 6})             // 切片: [4 5 6]
    fmt.Println(map[string]int{"a": 1})     // map: map[a:1]
    
    type Person struct{ Name string; Age int }
    fmt.Println(Person{"Alice", 30})        // struct: {Alice 30}
}
```

---

### ✅ 二、`fmt.Printf` 支持更精细控制

`fmt.Printf` 使用格式动词（verbs）来控制输出格式，支持更多选项：

| 动词 | 说明 |
|------|------|
| `%v` | 值的默认格式（最常用） |
| `%+v` | 结构体字段名也打印：`{Name:Alice Age:30}` |
| `%#v` | Go 语法表示：`main.Person{Name:"Alice", Age:30}` |
| `%T` | 打印类型 |
| `%d` | 十进制整数 |
| `%f` | 浮点数 |
| `%s` | 字符串 |
| `%t` | 布尔值 |
| `%p` | 指针地址 |

示例：

```go
person := Person{"Bob", 25}
fmt.Printf("%v\n", person)   // {Bob 25}
fmt.Printf("%+v\n", person)  // {Name:Bob Age:25}
fmt.Printf("%#v\n", person)  // main.Person{Name:"Bob", Age:25}
fmt.Printf("%T\n", person)   // main.Person
```

---

### ✅ 三、自定义类型的打印

如果你定义了一个类型，`fmt` 也可以打印，但你可以通过实现以下接口来自定义输出：

- `String() string` 方法（实现 `fmt.Stringer` 接口）

```go
func (p Person) String() string {
    return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

fmt.Println(person) // 输出: Bob (25 years old)
```

---

### ✅ 四、指针和 nil

- 指针：`fmt.Print(&x)` 会输出地址（如 `0xc000012010`）
- `nil` 切片、map、指针：输出 `nil`

```go
var p *int
fmt.Println(p) // 输出: <nil>
```

---

### ✅ 总结：`fmt.Print` 系列函数支持的类型

| 类型 | 是否支持直接打印 | 说明 |
|------|------------------|------|
| 所有基本类型 | ✅ 是 | int, float, bool, string 等 |
| 数组、切片 | ✅ 是 | 输出 `[1 2 3]` 格式 |
| map | ✅ 是 | 输出 `map[key:value]` |
| struct | ✅ 是 | 默认 `{v1 v2}` 或 `{Field:v}` |
| 指针 | ✅ 是 | 输出内存地址 |
| channel, func | ✅ 是 | 输出 `0xc00...` 或 `0x10502e0` |
| interface{} | ✅ 是 | 输出其动态值 |
| nil 值 | ✅ 是 | 输出 `nil` |

---

### 📌 结论

> **`fmt.Print`, `fmt.Println`, `fmt.Printf` 几乎支持所有 Go 类型的直接打印**，无需额外处理。  
> 推荐使用 `%v` 查看值，`%+v` 查看结构体字段，`%#v` 查看完整 Go 语法表示。

非常适合调试和日志输出！



# fmt.Print, fmt.Println区别
fmt.Print：不换行、不自动加空格（但不同类型间有例外），适合精确输出。
fmt.Println：自动加空格、自动换行，适合快速调试和日志。