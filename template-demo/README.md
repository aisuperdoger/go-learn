# Go text/template 库使用示例

本示例详细展示了Go语言中`text/template`库的使用方法，特别是`template.FuncMap`的各种用法。

## 什么是 text/template？

`text/template`是Go标准库中的模板引擎，用于生成文本输出。它支持：
- 变量替换
- 条件判断
- 循环遍历
- 自定义函数
- 嵌套模板
- 管道操作

## template.FuncMap 详解

`template.FuncMap`是一个映射类型，用于向模板注册自定义函数：

```go
type FuncMap map[string]interface{}
```

### 基本用法

```go
funcMap := template.FuncMap{
    "upper": strings.ToUpper,
    "add": func(a, b int) int {
        return a + b
    },
}

tmpl := template.New("example").Funcs(funcMap)
```

## 示例说明

### 1. 基本模板使用
展示最基本的模板语法：
- `{{.Field}}` - 访问字段
- `{{if .Condition}}...{{else}}...{{end}}` - 条件判断

### 2. FuncMap 自定义函数
演示如何添加各种类型的自定义函数：
- 字符串处理函数：`upper`, `lower`, `title`
- 数学计算函数：`add`, `multiply`
- 格式化函数：`formatPrice`, `formatDate`
- 条件判断函数：`isEven`
- 数组操作函数：`join`

### 3. 条件和循环
展示模板中的控制结构：
- `{{if}}...{{else}}...{{end}}` - 条件判断
- `{{range}}...{{end}}` - 循环遍历
- `{{with}}...{{end}}` - 上下文切换
- `{{- }}` - 去除空白字符

### 4. 嵌套模板
演示如何使用子模板：
- `{{define "name"}}...{{end}}` - 定义子模板
- `{{template "name" .}}` - 调用子模板

### 5. HTML 模板
展示`html/template`包的特性：
- 自动HTML转义
- `template.HTML`类型绕过转义
- JavaScript上下文中的安全处理

### 6. 高级 FuncMap 示例
展示更复杂的自定义函数：
- `truncate` - 字符串截断
- `first`, `last`, `slice` - 数组操作
- `default` - 默认值处理
- `max`, `min` - 数学函数
- `timeAgo` - 时间格式化
- `percentage` - 百分比计算

### 7. 模板文件使用
演示如何使用外部模板文件：
- `template.ParseFiles()` - 解析文件
- `ExecuteTemplate()` - 执行指定模板

## 常用 FuncMap 函数类型

### 字符串处理
```go
"upper":    strings.ToUpper,
"lower":    strings.ToLower,
"title":    strings.Title,
"trim":     strings.TrimSpace,
"replace":  strings.Replace,
```

### 数学运算
```go
"add": func(a, b int) int { return a + b },
"sub": func(a, b int) int { return a - b },
"mul": func(a, b int) int { return a * b },
"div": func(a, b int) int { return a / b },
```

### 时间处理
```go
"now": time.Now,
"formatTime": func(t time.Time, layout string) string {
    return t.Format(layout)
},
"timeAgo": func(t time.Time) string {
    return time.Since(t).String()
},
```

### 条件判断
```go
"eq":  func(a, b interface{}) bool { return a == b },
"ne":  func(a, b interface{}) bool { return a != b },
"lt":  func(a, b int) bool { return a < b },
"gt":  func(a, b int) bool { return a > b },
"and": func(a, b bool) bool { return a && b },
"or":  func(a, b bool) bool { return a || b },
```

### 数组/切片操作
```go
"len": func(items []interface{}) int { return len(items) },
"first": func(items []interface{}) interface{} {
    if len(items) > 0 { return items[0] }
    return nil
},
"join": strings.Join,
```

## 模板语法要点

### 变量访问
- `{{.}}` - 当前上下文
- `{{.Field}}` - 访问字段
- `{{.Method}}` - 调用方法
- `{{$var := .Field}}` - 定义变量

### 管道操作
```go
{{.Name | upper | truncate 10}}
{{.Price | formatPrice}}
```

### 条件判断
```go
{{if .IsActive}}
    活跃用户
{{else if .IsPending}}
    待审核用户
{{else}}
    非活跃用户
{{end}}
```

### 循环遍历
```go
{{range $index, $item := .Items}}
    {{$index}}: {{$item.Name}}
{{end}}

{{range .Items}}
    {{.Name}}
{{else}}
    没有项目
{{end}}
```

### 上下文切换
```go
{{with .Profile}}
    姓名: {{.Name}}
    年龄: {{.Age}}
{{end}}
```

## 最佳实践

1. **函数命名**：使用清晰、简洁的函数名
2. **错误处理**：自定义函数应该返回错误
3. **类型安全**：在函数中进行类型检查
4. **性能考虑**：避免在模板中进行复杂计算
5. **安全性**：使用`html/template`处理HTML内容

## 文件说明

- `main.go` - 完整的示例程序，包含所有功能演示
- `simple_example.go` - 简化的示例，适合快速学习
- `practical_examples.go` - 实际应用场景示例
- `user.tmpl` 和 `layout.tmpl` - 模板文件示例
- `README.md` - 详细说明文档

## 运行示例

### 运行完整示例
```bash
cd template-demo
go run main.go
```

### 运行简单示例
```bash
cd template-demo
go run simple_example.go
```

### 运行单个文件
```bash
cd template-demo
go run main.go practical_examples.go
```

## 输出说明

程序会依次执行所有示例，展示：
1. 基本模板渲染
2. 自定义函数的使用
3. 条件和循环控制
4. 嵌套模板结构
5. HTML模板的安全特性
6. 高级函数的实际应用
7. 外部模板文件的使用

每个示例都会在控制台输出相应的结果，帮助理解不同功能的使用方法。

## 扩展阅读

- [Go官方文档 - text/template](https://pkg.go.dev/text/template)
- [Go官方文档 - html/template](https://pkg.go.dev/html/template)
- [模板语法参考](https://golang.org/pkg/text/template/#hdr-Text_and_spaces)
