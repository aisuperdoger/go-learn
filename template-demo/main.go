package main

import (
	"fmt"
	htmlTemplate "html/template"
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

// 用户数据结构
type User struct {
	Name     string
	Age      int
	Email    string
	IsActive bool
	Tags     []string
	Profile  Profile
}

type Profile struct {
	Bio       string
	Website   string
	JoinDate  time.Time
	LastLogin *time.Time
}

// 产品数据结构
type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	InStock     bool
	Categories  []string
}

func main() {
	fmt.Println("=== Go text/template 库使用示例 ===\n")

	// 示例1: 基本模板使用
	fmt.Println("1. 基本模板使用:")
	basicTemplateExample()

	// 示例2: 使用 FuncMap 添加自定义函数
	fmt.Println("\n2. 使用 FuncMap 添加自定义函数:")
	funcMapExample()

	// 示例3: 条件和循环
	fmt.Println("\n3. 条件和循环:")
	conditionalAndLoopExample()

	// 示例4: 嵌套模板
	fmt.Println("\n4. 嵌套模板:")
	nestedTemplateExample()

	// 示例5: HTML 模板
	fmt.Println("\n5. HTML 模板:")
	htmlTemplateExample()

	// 示例6: 复杂的 FuncMap 示例
	fmt.Println("\n6. 复杂的 FuncMap 示例:")
	advancedFuncMapExample()

	// 示例7: 模板文件使用
	fmt.Println("\n7. 模板文件使用:")
	templateFileExample()

	// 示例8: 实际应用场景
	fmt.Println("\n8. 实际应用场景:")
	runPracticalExamples()
}

// 基本模板使用
func basicTemplateExample() {
	// 定义模板字符串
	tmplStr := `
用户信息:
姓名: {{.Name}}
年龄: {{.Age}}
邮箱: {{.Email}}
状态: {{if .IsActive}}活跃{{else}}非活跃{{end}}
`

	// 创建模板
	tmpl, err := template.New("user").Parse(tmplStr)
	if err != nil {
		log.Printf("解析模板失败: %v", err)
		return
	}

	// 准备数据
	user := User{
		Name:     "张三",
		Age:      25,
		Email:    "zhangsan@example.com",
		IsActive: true,
	}

	// 执行模板
	err = tmpl.Execute(os.Stdout, user)
	if err != nil {
		log.Printf("执行模板失败: %v", err)
	}
}

// 使用 FuncMap 添加自定义函数
func funcMapExample() {
	// 定义自定义函数映射
	funcMap := template.FuncMap{
		// 字符串处理函数
		"upper": strings.ToUpper,
		"lower": strings.ToLower,
		"title": strings.Title,

		// 数学函数
		"add": func(a, b int) int {
			return a + b
		},
		"multiply": func(a, b float64) float64 {
			return a * b
		},

		// 格式化函数
		"formatPrice": func(price float64) string {
			return fmt.Sprintf("¥%.2f", price)
		},

		// 时间格式化函数
		"formatDate": func(t time.Time) string {
			return t.Format("2006-01-02 15:04:05")
		},

		// 条件函数
		"isEven": func(n int) bool {
			return n%2 == 0
		},

		// 字符串切片连接
		"join": func(sep string, items []string) string {
			return strings.Join(items, sep)
		},
	}

	// 定义模板字符串
	tmplStr := `
产品信息:
名称: {{.Name | upper}}
价格: {{formatPrice .Price}}
描述: {{.Description | title}}
库存: {{if .InStock}}有货{{else}}缺货{{end}}
分类: {{join ", " .Categories}}
ID是否为偶数: {{isEven .ID}}
计算: 2 + 3 = {{add 2 3}}
价格乘以2: {{multiply .Price 2.0}}
`

	// 创建带有自定义函数的模板
	tmpl, err := template.New("product").Funcs(funcMap).Parse(tmplStr)
	if err != nil {
		log.Printf("解析模板失败: %v", err)
		return
	}

	// 准备数据
	product := Product{
		ID:          123,
		Name:        "智能手机",
		Price:       2999.99,
		Description: "latest smartphone with advanced features",
		InStock:     true,
		Categories:  []string{"电子产品", "手机", "智能设备"},
	}

	// 执行模板
	err = tmpl.Execute(os.Stdout, product)
	if err != nil {
		log.Printf("执行模板失败: %v", err)
	}
}

// 条件和循环示例
func conditionalAndLoopExample() {
	funcMap := template.FuncMap{
		"len": func(items []string) int {
			return len(items)
		},
		"index": func(i int) int {
			return i + 1
		},
	}

	tmplStr := `
用户详情:
姓名: {{.Name}}
{{if .IsActive -}}
状态: 活跃用户
{{- else -}}
状态: 非活跃用户
{{- end}}

{{if .Tags -}}
标签 (共{{len .Tags}}个):
{{- range $i, $tag := .Tags}}
  {{index $i}}. {{$tag}}
{{- end}}
{{- else}}
暂无标签
{{- end}}

{{with .Profile -}}
个人资料:
  简介: {{.Bio}}
  网站: {{.Website}}
  加入时间: {{.JoinDate.Format "2006-01-02"}}
  {{- if .LastLogin}}
  最后登录: {{.LastLogin.Format "2006-01-02 15:04:05"}}
  {{- else}}
  最后登录: 从未登录
  {{- end}}
{{- end}}
`

	tmpl, err := template.New("userDetail").Funcs(funcMap).Parse(tmplStr)
	if err != nil {
		log.Printf("解析模板失败: %v", err)
		return
	}

	lastLogin := time.Now().Add(-24 * time.Hour)
	user := User{
		Name:     "李四",
		Age:      30,
		Email:    "lisi@example.com",
		IsActive: true,
		Tags:     []string{"开发者", "Go语言", "后端"},
		Profile: Profile{
			Bio:       "资深Go开发工程师",
			Website:   "https://lisi.dev",
			JoinDate:  time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC),
			LastLogin: &lastLogin,
		},
	}

	err = tmpl.Execute(os.Stdout, user)
	if err != nil {
		log.Printf("执行模板失败: %v", err)
	}
}

// 嵌套模板示例
func nestedTemplateExample() {
	// 定义子模板
	headerTmpl := `{{define "header"}}
=== {{.Title}} ===
生成时间: {{.GeneratedAt.Format "2006-01-02 15:04:05"}}
{{end}}`

	footerTmpl := `{{define "footer"}}
---
版权所有 © 2024
{{end}}`

	// 主模板
	mainTmpl := `{{template "header" .}}

用户列表:
{{range .Users}}
- {{.Name}} ({{.Age}}岁) - {{.Email}}
{{end}}

{{template "footer" .}}`

	// 解析所有模板
	tmpl, err := template.New("main").Parse(headerTmpl)
	if err != nil {
		log.Printf("解析header模板失败: %v", err)
		return
	}

	tmpl, err = tmpl.Parse(footerTmpl)
	if err != nil {
		log.Printf("解析footer模板失败: %v", err)
		return
	}

	tmpl, err = tmpl.Parse(mainTmpl)
	if err != nil {
		log.Printf("解析main模板失败: %v", err)
		return
	}

	// 准备数据
	data := struct {
		Title       string
		GeneratedAt time.Time
		Users       []User
	}{
		Title:       "用户报告",
		GeneratedAt: time.Now(),
		Users: []User{
			{Name: "张三", Age: 25, Email: "zhangsan@example.com"},
			{Name: "李四", Age: 30, Email: "lisi@example.com"},
			{Name: "王五", Age: 28, Email: "wangwu@example.com"},
		},
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Printf("执行模板失败: %v", err)
	}
}

// HTML 模板示例
func htmlTemplateExample() {
	// 使用 html/template 包，自动转义HTML
	funcMap := htmlTemplate.FuncMap{
		"safeHTML": func(s string) htmlTemplate.HTML {
			return htmlTemplate.HTML(s)
		},
		"formatPrice": func(price float64) string {
			return fmt.Sprintf("¥%.2f", price)
		},
	}

	tmplStr := `<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
</head>
<body>
    <h1>{{.Title}}</h1>

    <div class="user-info">
        <h2>用户: {{.User.Name}}</h2>
        <p>邮箱: {{.User.Email}}</p>
        <p>状态: {{if .User.IsActive}}<span style="color: green;">活跃</span>{{else}}<span style="color: red;">非活跃</span>{{end}}</p>
    </div>

    <div class="content">
        <!-- 这会被自动转义 -->
        <p>原始内容: {{.RawContent}}</p>

        <!-- 这不会被转义 -->
        <p>安全HTML: {{safeHTML .SafeContent}}</p>
    </div>

    <script>
        console.log("用户名: {{.User.Name}}");
    </script>
</body>
</html>`

	tmpl, err := htmlTemplate.New("html").Funcs(funcMap).Parse(tmplStr)
	if err != nil {
		log.Printf("解析HTML模板失败: %v", err)
		return
	}

	data := struct {
		Title       string
		User        User
		RawContent  string
		SafeContent string
	}{
		Title: "用户信息页面",
		User: User{
			Name:     "张三",
			Email:    "zhangsan@example.com",
			IsActive: true,
		},
		RawContent:  "<script>alert('XSS攻击')</script>",
		SafeContent: "<strong>这是安全的HTML内容</strong>",
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Printf("执行HTML模板失败: %v", err)
	}
}

// 复杂的 FuncMap 示例
func advancedFuncMapExample() {
	funcMap := template.FuncMap{
		// 字符串处理
		"truncate": func(s string, length int) string {
			if len(s) <= length {
				return s
			}
			return s[:length] + "..."
		},

		// 数组/切片操作
		"first": func(items []string) string {
			if len(items) > 0 {
				return items[0]
			}
			return ""
		},
		"last": func(items []string) string {
			if len(items) > 0 {
				return items[len(items)-1]
			}
			return ""
		},
		"slice": func(items []string, start, end int) []string {
			if start < 0 || start >= len(items) {
				return []string{}
			}
			if end > len(items) {
				end = len(items)
			}
			return items[start:end]
		},

		// 条件函数
		"default": func(defaultValue, value interface{}) interface{} {
			if value == nil || value == "" {
				return defaultValue
			}
			return value
		},

		// 数学函数
		"max": func(a, b int) int {
			if a > b {
				return a
			}
			return b
		},
		"min": func(a, b int) int {
			if a < b {
				return a
			}
			return b
		},

		// 时间函数
		"timeAgo": func(t time.Time) string {
			duration := time.Since(t)
			if duration < time.Minute {
				return "刚刚"
			} else if duration < time.Hour {
				return fmt.Sprintf("%d分钟前", int(duration.Minutes()))
			} else if duration < 24*time.Hour {
				return fmt.Sprintf("%d小时前", int(duration.Hours()))
			} else {
				return fmt.Sprintf("%d天前", int(duration.Hours()/24))
			}
		},

		// 格式化函数
		"percentage": func(value, total float64) string {
			if total == 0 {
				return "0%"
			}
			return fmt.Sprintf("%.1f%%", (value/total)*100)
		},

		// 字符串连接函数
		"join": strings.Join,
	}

	tmplStr := `
高级函数示例:

用户: {{.Name}}
简介: {{truncate .Profile.Bio 20}}
网站: {{default "未设置" .Profile.Website}}

标签信息:
{{if .Tags -}}
第一个标签: {{first .Tags}}
最后一个标签: {{last .Tags}}
前两个标签: {{(slice .Tags 0 2) | join ", "}}
{{- end}}

时间信息:
加入时间: {{timeAgo .Profile.JoinDate}}
{{if .Profile.LastLogin -}}
最后登录: {{timeAgo .Profile.LastLogin}}
{{- end}}

统计信息:
年龄范围: {{min .Age 100}} - {{max .Age 18}}
活跃度: {{if .IsActive}}{{percentage 85.5 100}}{{else}}{{percentage 12.3 100}}{{end}}
`

	tmpl, err := template.New("advanced").Funcs(funcMap).Parse(tmplStr)
	if err != nil {
		log.Printf("解析高级模板失败: %v", err)
		return
	}

	lastLogin := time.Now().Add(-2 * time.Hour)
	user := User{
		Name:     "王五",
		Age:      28,
		IsActive: true,
		Tags:     []string{"Go", "Docker", "Kubernetes", "微服务", "云原生"},
		Profile: Profile{
			Bio:       "专注于云原生技术的全栈开发工程师，拥有丰富的容器化和微服务架构经验",
			Website:   "https://wangwu.dev",
			JoinDate:  time.Now().Add(-365 * 24 * time.Hour), // 一年前
			LastLogin: &lastLogin,
		},
	}

	err = tmpl.Execute(os.Stdout, user)
	if err != nil {
		log.Printf("执行高级模板失败: %v", err)
	}
}

// 模板文件使用示例
func templateFileExample() {
	// 创建模板文件
	createTemplateFiles()

	// 解析模板文件
	tmpl, err := template.ParseFiles("user.tmpl", "layout.tmpl")
	if err != nil {
		log.Printf("解析模板文件失败: %v", err)
		return
	}

	user := User{
		Name:     "赵六",
		Age:      32,
		Email:    "zhaoliu@example.com",
		IsActive: true,
		Tags:     []string{"架构师", "技术管理"},
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "layout.tmpl", user)
	if err != nil {
		log.Printf("执行模板文件失败: %v", err)
	}
}

// 创建示例模板文件
func createTemplateFiles() {
	// 创建用户模板文件
	userTmpl := `{{define "user-info"}}
<div class="user-card">
    <h3>{{.Name}}</h3>
    <p>年龄: {{.Age}}</p>
    <p>邮箱: {{.Email}}</p>
    <p>状态: {{if .IsActive}}活跃{{else}}非活跃{{end}}</p>
    {{if .Tags}}
    <div class="tags">
        {{range .Tags}}<span class="tag">{{.}}</span>{{end}}
    </div>
    {{end}}
</div>
{{end}}`

	err := os.WriteFile("user.tmpl", []byte(userTmpl), 0644)
	if err != nil {
		log.Printf("创建用户模板文件失败: %v", err)
	}

	// 创建布局模板文件
	layoutTmpl := `<!DOCTYPE html>
<html>
<head>
    <title>用户信息</title>
    <style>
        .user-card { border: 1px solid #ccc; padding: 10px; margin: 10px; }
        .tag { background: #007cba; color: white; padding: 2px 6px; margin: 2px; border-radius: 3px; }
    </style>
</head>
<body>
    <h1>用户管理系统</h1>
    {{template "user-info" .}}
</body>
</html>`

	err = os.WriteFile("layout.tmpl", []byte(layoutTmpl), 0644)
	if err != nil {
		log.Printf("创建布局模板文件失败: %v", err)
	}
}
