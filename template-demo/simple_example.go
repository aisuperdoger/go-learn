package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

// 简单的 text/template 和 FuncMap 使用示例

func simpleExample() {
	fmt.Println("=== 简单的 text/template 和 FuncMap 示例 ===\n")

	// 1. 基本模板使用
	fmt.Println("1. 基本模板:")
	basicExample()

	// 2. FuncMap 使用
	fmt.Println("\n2. FuncMap 自定义函数:")
	funcMapBasicExample()

	// 3. 实用的 FuncMap 函数
	fmt.Println("\n3. 实用的 FuncMap 函数:")
	practicalFuncMapExample()
}

// 基本模板示例
func basicExample() {
	// 定义模板字符串
	tmplStr := `
姓名: {{.Name}}
年龄: {{.Age}}
邮箱: {{.Email}}
{{if .IsActive}}状态: 活跃{{else}}状态: 非活跃{{end}}
`

	// 创建并解析模板
	tmpl, err := template.New("basic").Parse(tmplStr)
	if err != nil {
		log.Printf("解析模板失败: %v", err)
		return
	}

	// 准备数据
	data := struct {
		Name     string
		Age      int
		Email    string
		IsActive bool
	}{
		Name:     "张三",
		Age:      25,
		Email:    "zhangsan@example.com",
		IsActive: true,
	}

	// 执行模板
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Printf("执行模板失败: %v", err)
	}
}

// FuncMap 基本使用示例
func funcMapBasicExample() {
	// 定义自定义函数映射
	funcMap := template.FuncMap{
		// 字符串处理
		"upper": strings.ToUpper,
		"lower": strings.ToLower,

		// 数学运算
		"add": func(a, b int) int {
			return a + b
		},
		"multiply": func(a, b float64) float64 {
			return a * b
		},

		// 格式化
		"formatPrice": func(price float64) string {
			return fmt.Sprintf("¥%.2f", price)
		},
	}

	// 定义模板字符串
	tmplStr := `
产品信息:
名称: {{.Name | upper}}
价格: {{formatPrice .Price}}
折扣价: {{formatPrice (multiply .Price 0.8)}}
总数: {{add .Count 10}}
`

	// 创建带有自定义函数的模板
	tmpl, err := template.New("product").Funcs(funcMap).Parse(tmplStr)
	if err != nil {
		log.Printf("解析模板失败: %v", err)
		return
	}

	// 准备数据
	data := struct {
		Name  string
		Price float64
		Count int
	}{
		Name:  "智能手机",
		Price: 2999.99,
		Count: 5,
	}

	// 执行模板
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Printf("执行模板失败: %v", err)
	}
}

// 实用的 FuncMap 函数示例
func practicalFuncMapExample() {
	// 定义实用的自定义函数
	funcMap := template.FuncMap{
		// 默认值函数
		"default": func(defaultValue, value interface{}) interface{} {
			if value == nil || value == "" {
				return defaultValue
			}
			return value
		},

		// 字符串截断
		"truncate": func(s string, length int) string {
			if len(s) <= length {
				return s
			}
			return s[:length] + "..."
		},

		// 时间格式化
		"formatTime": func(t time.Time, layout string) string {
			return t.Format(layout)
		},

		// 数组长度
		"len": func(items []string) int {
			return len(items)
		},

		// 字符串连接
		"join": strings.Join,

		// 条件判断
		"eq": func(a, b interface{}) bool {
			return a == b
		},
		"gt": func(a, b int) bool {
			return a > b
		},
	}

	// 定义模板字符串
	tmplStr := `
用户资料:
姓名: {{.Name}}
简介: {{truncate .Bio 30}}
网站: {{default "未设置" .Website}}
注册时间: {{formatTime .RegisterTime "2006-01-02"}}
标签数量: {{len .Tags}}
{{if gt (len .Tags) 0}}标签: {{join .Tags ", "}}{{end}}
{{if eq .Role "admin"}}权限: 管理员{{else}}权限: 普通用户{{end}}
`

	// 创建模板
	tmpl, err := template.New("profile").Funcs(funcMap).Parse(tmplStr)
	if err != nil {
		log.Printf("解析模板失败: %v", err)
		return
	}

	// 准备数据
	data := struct {
		Name         string
		Bio          string
		Website      string
		RegisterTime time.Time
		Tags         []string
		Role         string
	}{
		Name:         "李四",
		Bio:          "这是一个很长的个人简介，展示如何使用truncate函数来截断长文本内容",
		Website:      "", // 空值，会使用默认值
		RegisterTime: time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC),
		Tags:         []string{"Go", "Docker", "Kubernetes"},
		Role:         "admin",
	}

	// 执行模板
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Printf("执行模板失败: %v", err)
	}
}

// 如果直接运行这个文件
func main() {
	simpleExample()
}
