package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

// 实际应用场景示例

// 邮件模板示例
func emailTemplateExample() {
	fmt.Println("=== 邮件模板示例 ===")

	funcMap := template.FuncMap{
		"formatDate": func(t time.Time) string {
			return t.Format("2006年01月02日")
		},
		"currency": func(amount float64) string {
			return fmt.Sprintf("¥%.2f", amount)
		},
	}

	emailTemplate := `
主题: 订单确认 - {{.OrderID}}

亲爱的 {{.CustomerName}}，

感谢您的订单！以下是您的订单详情：

订单编号: {{.OrderID}}
下单时间: {{formatDate .OrderDate}}
订单状态: {{.Status}}

商品清单:
{{range .Items}}
- {{.Name}} x {{.Quantity}} = {{currency .Total}}
{{end}}

订单总额: {{currency .TotalAmount}}
{{if .DiscountAmount}}折扣金额: -{{currency .DiscountAmount}}{{end}}
实付金额: {{currency .FinalAmount}}

配送地址:
{{.ShippingAddress}}

预计送达时间: {{formatDate .EstimatedDelivery}}

如有任何问题，请联系我们的客服。

祝您购物愉快！
电商平台客服团队
`

	tmpl, err := template.New("email").Funcs(funcMap).Parse(emailTemplate)
	if err != nil {
		log.Printf("解析邮件模板失败: %v", err)
		return
	}

	// 订单数据
	orderData := struct {
		OrderID      string
		CustomerName string
		OrderDate    time.Time
		Status       string
		Items        []struct {
			Name     string
			Quantity int
			Total    float64
		}
		TotalAmount       float64
		DiscountAmount    float64
		FinalAmount       float64
		ShippingAddress   string
		EstimatedDelivery time.Time
	}{
		OrderID:      "ORD-2024-001",
		CustomerName: "张三",
		OrderDate:    time.Now(),
		Status:       "已确认",
		Items: []struct {
			Name     string
			Quantity int
			Total    float64
		}{
			{"智能手机", 1, 2999.00},
			{"手机壳", 2, 58.00},
			{"钢化膜", 1, 29.00},
		},
		TotalAmount:       3086.00,
		DiscountAmount:    100.00,
		FinalAmount:       2986.00,
		ShippingAddress:   "北京市朝阳区某某街道123号",
		EstimatedDelivery: time.Now().Add(48 * time.Hour),
	}

	err = tmpl.Execute(&bytes.Buffer{}, orderData)
	if err != nil {
		log.Printf("执行邮件模板失败: %v", err)
		return
	}

	// 输出到控制台
	tmpl.Execute(os.Stdout, orderData)
}

// 配置文件生成示例
func configTemplateExample() {
	fmt.Println("\n=== 配置文件生成示例 ===")

	funcMap := template.FuncMap{
		"join": strings.Join,
		"quote": func(s string) string {
			return fmt.Sprintf(`"%s"`, s)
		},
	}

	configTemplate := `# 应用配置文件
# 生成时间: {{.GeneratedAt.Format "2006-01-02 15:04:05"}}

[server]
host = {{quote .Server.Host}}
port = {{.Server.Port}}
debug = {{.Server.Debug}}

[database]
driver = {{quote .Database.Driver}}
host = {{quote .Database.Host}}
port = {{.Database.Port}}
name = {{quote .Database.Name}}
user = {{quote .Database.User}}
password = {{quote .Database.Password}}
max_connections = {{.Database.MaxConnections}}

[redis]
host = {{quote .Redis.Host}}
port = {{.Redis.Port}}
password = {{quote .Redis.Password}}
db = {{.Redis.DB}}

[logging]
level = {{quote .Logging.Level}}
file = {{quote .Logging.File}}
max_size = {{.Logging.MaxSize}}

[features]
enabled = [{{range $i, $feature := .Features.Enabled}}{{if $i}}, {{end}}{{quote $feature}}{{end}}]
disabled = [{{range $i, $feature := .Features.Disabled}}{{if $i}}, {{end}}{{quote $feature}}{{end}}]
`

	tmpl, err := template.New("config").Funcs(funcMap).Parse(configTemplate)
	if err != nil {
		log.Printf("解析配置模板失败: %v", err)
		return
	}

	configData := struct {
		GeneratedAt time.Time
		Server      struct {
			Host  string
			Port  int
			Debug bool
		}
		Database struct {
			Driver         string
			Host           string
			Port           int
			Name           string
			User           string
			Password       string
			MaxConnections int
		}
		Redis struct {
			Host     string
			Port     int
			Password string
			DB       int
		}
		Logging struct {
			Level   string
			File    string
			MaxSize int
		}
		Features struct {
			Enabled  []string
			Disabled []string
		}
	}{
		GeneratedAt: time.Now(),
		Server: struct {
			Host  string
			Port  int
			Debug bool
		}{
			Host:  "localhost",
			Port:  8080,
			Debug: true,
		},
		Database: struct {
			Driver         string
			Host           string
			Port           int
			Name           string
			User           string
			Password       string
			MaxConnections int
		}{
			Driver:         "mysql",
			Host:           "localhost",
			Port:           3306,
			Name:           "myapp",
			User:           "root",
			Password:       "password123",
			MaxConnections: 100,
		},
		Redis: struct {
			Host     string
			Port     int
			Password string
			DB       int
		}{
			Host:     "localhost",
			Port:     6379,
			Password: "",
			DB:       0,
		},
		Logging: struct {
			Level   string
			File    string
			MaxSize int
		}{
			Level:   "info",
			File:    "/var/log/myapp.log",
			MaxSize: 100,
		},
		Features: struct {
			Enabled  []string
			Disabled []string
		}{
			Enabled:  []string{"user_registration", "email_notifications", "api_v2"},
			Disabled: []string{"legacy_api", "debug_mode"},
		},
	}

	err = tmpl.Execute(os.Stdout, configData)
	if err != nil {
		log.Printf("执行配置模板失败: %v", err)
	}
}

// SQL 查询生成示例
func sqlTemplateExample() {
	fmt.Println("\n=== SQL 查询生成示例 ===")

	funcMap := template.FuncMap{
		"join": strings.Join,
		"quote": func(s string) string {
			return fmt.Sprintf("`%s`", s)
		},
		"sqlValue": func(v interface{}) string {
			switch val := v.(type) {
			case string:
				return fmt.Sprintf("'%s'", strings.ReplaceAll(val, "'", "''"))
			case int, int64, float64:
				return fmt.Sprintf("%v", val)
			case bool:
				if val {
					return "1"
				}
				return "0"
			default:
				return "NULL"
			}
		},
	}

	// SELECT 查询模板
	selectTemplate := `SELECT {{range $i, $col := .Columns}}{{if $i}}, {{end}}{{quote $col}}{{end}}
FROM {{quote .Table}}
{{if .Conditions}}WHERE {{range $i, $cond := .Conditions}}{{if $i}} AND {{end}}{{quote $cond.Column}} {{$cond.Operator}} {{sqlValue $cond.Value}}{{end}}{{end}}
{{if .OrderBy}}ORDER BY {{range $i, $order := .OrderBy}}{{if $i}}, {{end}}{{quote $order.Column}} {{$order.Direction}}{{end}}{{end}}
{{if .Limit}}LIMIT {{.Limit}}{{end}};`

	// INSERT 查询模板
	insertTemplate := `INSERT INTO {{quote .Table}} ({{range $i, $col := .Columns}}{{if $i}}, {{end}}{{quote $col}}{{end}})
VALUES {{range $i, $row := .Values}}{{if $i}}, {{end}}({{range $j, $val := $row}}{{if $j}}, {{end}}{{sqlValue $val}}{{end}}){{end}};`

	// UPDATE 查询模板
	updateTemplate := `UPDATE {{quote .Table}}
SET {{range $i, $set := .Sets}}{{if $i}}, {{end}}{{quote $set.Column}} = {{sqlValue $set.Value}}{{end}}
{{if .Conditions}}WHERE {{range $i, $cond := .Conditions}}{{if $i}} AND {{end}}{{quote $cond.Column}} {{$cond.Operator}} {{sqlValue $cond.Value}}{{end}}{{end}};`

	selectTmpl, _ := template.New("select").Funcs(funcMap).Parse(selectTemplate)
	insertTmpl, _ := template.New("insert").Funcs(funcMap).Parse(insertTemplate)
	updateTmpl, _ := template.New("update").Funcs(funcMap).Parse(updateTemplate)

	// SELECT 示例
	fmt.Println("SELECT 查询:")
	selectData := struct {
		Table      string
		Columns    []string
		Conditions []struct {
			Column   string
			Operator string
			Value    interface{}
		}
		OrderBy []struct {
			Column    string
			Direction string
		}
		Limit int
	}{
		Table:   "users",
		Columns: []string{"id", "name", "email", "created_at"},
		Conditions: []struct {
			Column   string
			Operator string
			Value    interface{}
		}{
			{"age", ">=", 18},
			{"status", "=", "active"},
		},
		OrderBy: []struct {
			Column    string
			Direction string
		}{
			{"created_at", "DESC"},
			{"name", "ASC"},
		},
		Limit: 10,
	}
	selectTmpl.Execute(os.Stdout, selectData)

	// INSERT 示例
	fmt.Println("\nINSERT 查询:")
	insertData := struct {
		Table   string
		Columns []string
		Values  [][]interface{}
	}{
		Table:   "users",
		Columns: []string{"name", "email", "age", "active"},
		Values: [][]interface{}{
			{"张三", "zhangsan@example.com", 25, true},
			{"李四", "lisi@example.com", 30, false},
		},
	}
	insertTmpl.Execute(os.Stdout, insertData)

	// UPDATE 示例
	fmt.Println("\nUPDATE 查询:")
	updateData := struct {
		Table string
		Sets  []struct {
			Column string
			Value  interface{}
		}
		Conditions []struct {
			Column   string
			Operator string
			Value    interface{}
		}
	}{
		Table: "users",
		Sets: []struct {
			Column string
			Value  interface{}
		}{
			{"last_login", time.Now().Format("2006-01-02 15:04:05")},
			{"login_count", 10},
		},
		Conditions: []struct {
			Column   string
			Operator string
			Value    interface{}
		}{
			{"id", "=", 1},
		},
	}
	updateTmpl.Execute(os.Stdout, updateData)
}

// 代码生成示例
func codeGenerationExample() {
	fmt.Println("\n=== 代码生成示例 ===")

	funcMap := template.FuncMap{
		"title": strings.Title,
		"lower": strings.ToLower,
		"join":  strings.Join,
	}

	// Go 结构体生成模板
	structTemplate := `// {{.Comment}}
type {{.Name}} struct {
{{range .Fields}}	{{.Name}} {{.Type}} {{if .Tag}}{{.Tag}}{{end}} {{if .Comment}}// {{.Comment}}{{end}}
{{end}}}

// New{{.Name}} 创建新的{{.Name}}实例
func New{{.Name}}({{range $i, $field := .Fields}}{{if $i}}, {{end}}{{lower $field.Name}} {{$field.Type}}{{end}}) *{{.Name}} {
	return &{{.Name}}{
{{range .Fields}}		{{.Name}}: {{lower .Name}},
{{end}}	}
}

{{range .Fields}}{{if eq .Type "string"}}
// Get{{.Name}} 获取{{.Name}}
func ({{lower $.Name}} *{{$.Name}}) Get{{.Name}}() {{.Type}} {
	return {{lower $.Name}}.{{.Name}}
}

// Set{{.Name}} 设置{{.Name}}
func ({{lower $.Name}} *{{$.Name}}) Set{{.Name}}({{lower .Name}} {{.Type}}) {
	{{lower $.Name}}.{{.Name}} = {{lower .Name}}
}
{{end}}{{end}}`

	tmpl, err := template.New("struct").Funcs(funcMap).Parse(structTemplate)
	if err != nil {
		log.Printf("解析结构体模板失败: %v", err)
		return
	}

	structData := struct {
		Name    string
		Comment string
		Fields  []struct {
			Name    string
			Type    string
			Tag     string
			Comment string
		}
	}{
		Name:    "User",
		Comment: "User 表示系统用户",
		Fields: []struct {
			Name    string
			Type    string
			Tag     string
			Comment string
		}{
			{"ID", "int64", "`json:\"id\" db:\"id\"`", "用户ID"},
			{"Name", "string", "`json:\"name\" db:\"name\"`", "用户姓名"},
			{"Email", "string", "`json:\"email\" db:\"email\"`", "邮箱地址"},
			{"CreatedAt", "time.Time", "`json:\"created_at\" db:\"created_at\"`", "创建时间"},
		},
	}

	err = tmpl.Execute(os.Stdout, structData)
	if err != nil {
		log.Printf("执行结构体模板失败: %v", err)
	}
}

// 运行所有实际应用示例
func runPracticalExamples() {
	emailTemplateExample()
	configTemplateExample()
	sqlTemplateExample()
	codeGenerationExample()
}
