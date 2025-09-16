package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/Knetic/govaluate"
)

// ExpressionAnalyzer 表达式分析器
type ExpressionAnalyzer struct {
	expression *govaluate.EvaluableExpression
	rawExpr    string
}

// NewExpressionAnalyzer 创建表达式分析器
func NewExpressionAnalyzer(expr string) (*ExpressionAnalyzer, error) {
	evaluable, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		return nil, err
	}
	
	return &ExpressionAnalyzer{
		expression: evaluable,
		rawExpr:    expr,
	}, nil
}

// GetUniqueVars 获取去重后的变量列表
func (ea *ExpressionAnalyzer) GetUniqueVars() []string {
	vars := ea.expression.Vars()
	uniqueVars := make(map[string]bool)
	
	for _, v := range vars {
		uniqueVars[v] = true
	}
	
	result := make([]string, 0, len(uniqueVars))
	for v := range uniqueVars {
		result = append(result, v)
	}
	
	sort.Strings(result) // 排序便于查看
	return result
}

// GetVarCount 获取变量使用次数统计
func (ea *ExpressionAnalyzer) GetVarCount() map[string]int {
	vars := ea.expression.Vars()
	count := make(map[string]int)
	
	for _, v := range vars {
		count[v]++
	}
	
	return count
}

// ValidateParameters 验证参数是否完整
func (ea *ExpressionAnalyzer) ValidateParameters(params map[string]interface{}) []string {
	requiredVars := ea.GetUniqueVars()
	missing := make([]string, 0)
	
	for _, varName := range requiredVars {
		if _, exists := params[varName]; !exists {
			missing = append(missing, varName)
		}
	}
	
	return missing
}

// GenerateParameterTemplate 生成参数模板
func (ea *ExpressionAnalyzer) GenerateParameterTemplate() map[string]interface{} {
	vars := ea.GetUniqueVars()
	template := make(map[string]interface{})
	
	for _, varName := range vars {
		// 根据变量名推测类型
		if strings.Contains(strings.ToLower(varName), "price") ||
		   strings.Contains(strings.ToLower(varName), "amount") ||
		   strings.Contains(strings.ToLower(varName), "cost") {
			template[varName] = 0.0 // 价格相关用浮点数
		} else if strings.Contains(strings.ToLower(varName), "count") ||
				  strings.Contains(strings.ToLower(varName), "quantity") ||
				  strings.Contains(strings.ToLower(varName), "age") {
			template[varName] = 0 // 数量相关用整数
		} else if strings.Contains(strings.ToLower(varName), "status") ||
				  strings.Contains(strings.ToLower(varName), "level") ||
				  strings.Contains(strings.ToLower(varName), "name") {
			template[varName] = "" // 状态相关用字符串
		} else if strings.Contains(strings.ToLower(varName), "is") ||
				  strings.Contains(strings.ToLower(varName), "has") ||
				  strings.Contains(strings.ToLower(varName), "enabled") {
			template[varName] = false // 布尔相关用布尔值
		} else {
			template[varName] = nil // 默认为 nil
		}
	}
	
	return template
}

// PrintAnalysis 打印分析结果
func (ea *ExpressionAnalyzer) PrintAnalysis() {
	fmt.Printf("表达式: %s\n", ea.rawExpr)
	
	// 基本信息
	allVars := ea.expression.Vars()
	uniqueVars := ea.GetUniqueVars()
	varCount := ea.GetVarCount()
	
	fmt.Printf("  总变量引用次数: %d\n", len(allVars))
	fmt.Printf("  唯一变量数量: %d\n", len(uniqueVars))
	
	// 变量列表
	fmt.Printf("  变量列表: %v\n", uniqueVars)
	
	// 使用频率
	fmt.Printf("  变量使用频率:\n")
	for _, varName := range uniqueVars {
		fmt.Printf("    %s: %d 次\n", varName, varCount[varName])
	}
	
	// 参数模板
	template := ea.GenerateParameterTemplate()
	fmt.Printf("  建议参数模板:\n")
	for _, varName := range uniqueVars {
		fmt.Printf("    \"%s\": %v,\n", varName, template[varName])
	}
	
	fmt.Println()
}

func demonstrateVarsUsage() {
	fmt.Println("=== expr.Vars() 高级用法演示 ===\n")
	
	// 测试表达式
	expressions := []string{
		"price * quantity",
		"userAge >= 18 && userAge <= 65",
		"basePrice + tax - discount + (weight > 5 ? shippingCost : 0)",
		"isVIP && (totalSpent > 1000 || memberYears >= 2)",
		"firstName + ' ' + lastName + ' (' + email + ')'",
	}
	
	analyzers := make([]*ExpressionAnalyzer, 0, len(expressions))
	
	// 创建分析器
	for _, expr := range expressions {
		analyzer, err := NewExpressionAnalyzer(expr)
		if err != nil {
			log.Printf("创建分析器失败 '%s': %v", expr, err)
			continue
		}
		analyzers = append(analyzers, analyzer)
	}
	
	// 分析每个表达式
	for _, analyzer := range analyzers {
		analyzer.PrintAnalysis()
	}
	
	// 演示参数验证
	fmt.Println("=== 参数验证演示 ===")
	if len(analyzers) > 0 {
		analyzer := analyzers[0] // 使用第一个表达式
		
		// 完整参数
		completeParams := map[string]interface{}{
			"price":    99.99,
			"quantity": 2,
		}
		
		// 不完整参数
		incompleteParams := map[string]interface{}{
			"price": 99.99,
			// 缺少 quantity
		}
		
		fmt.Printf("表达式: %s\n", analyzer.rawExpr)
		
		missing1 := analyzer.ValidateParameters(completeParams)
		fmt.Printf("完整参数验证 - 缺失变量: %v\n", missing1)
		
		missing2 := analyzer.ValidateParameters(incompleteParams)
		fmt.Printf("不完整参数验证 - 缺失变量: %v\n", missing2)
	}
	
	// 演示动态表单生成
	fmt.Println("\n=== 动态表单生成演示 ===")
	formValidationRules := map[string]string{
		"用户注册": "age >= 18 && email != '' && password != '' && confirmPassword == password",
		"订单提交": "productId > 0 && quantity > 0 && shippingAddress != ''",
		"会员升级": "currentLevel != 'gold' && totalSpent >= upgradeThreshold",
	}
	
	for formName, rule := range formValidationRules {
		analyzer, err := NewExpressionAnalyzer(rule)
		if err != nil {
			continue
		}
		
		fmt.Printf("%s表单:\n", formName)
		fmt.Printf("  验证规则: %s\n", rule)
		fmt.Printf("  需要字段: %v\n", analyzer.GetUniqueVars())
		
		template := analyzer.GenerateParameterTemplate()
		fmt.Printf("  字段类型推测:\n")
		for _, varName := range analyzer.GetUniqueVars() {
			fmt.Printf("    %s: %T\n", varName, template[varName])
		}
		fmt.Println()
	}
}

// 如果直接运行此文件
func init() {
	// 这个函数可以被 main.go 调用，也可以独立运行
}

// 独立运行的 main 函数（当直接运行此文件时）
// 注释掉，因为一个包只能有一个 main 函数
/*
func main() {
	demonstrateVarsUsage()
}
*/
