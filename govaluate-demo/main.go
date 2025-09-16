package main

import (
	"fmt"
	"log"
	"math"

	"github.com/Knetic/govaluate"
)

func main() {
	fmt.Println("=== govaluate 使用示例 ===")

	// 示例1: 基本数学表达式
	fmt.Println("\n1. 基本数学表达式:")
	basicMathExample()

	// 示例2: 变量替换
	fmt.Println("\n2. 变量替换:")
	variableExample()

	// 示例3: 逻辑表达式
	fmt.Println("\n3. 逻辑表达式:")
	logicalExample()

	// 示例4: 字符串操作
	fmt.Println("\n4. 字符串操作:")
	stringExample()

	// 示例5: 条件表达式
	fmt.Println("\n5. 条件表达式:")
	conditionalExample()

	// 示例6: 函数调用
	fmt.Println("\n6. 自定义函数:")
	functionExample()

	// 示例7: 表达式变量分析
	fmt.Println("\n7. 表达式变量分析:")
	variableAnalysisExample()

	// 示例8: expr.Vars() 高级用法
	fmt.Println("\n8. expr.Vars() 高级用法:")
	demonstrateVarsUsage()

	// 示例9: 复杂业务场景
	fmt.Println("\n9. 复杂业务场景:")
	businessExample()
}

// 基本数学表达式
func basicMathExample() {
	// 定义数学函数
	mathFunctions := map[string]govaluate.ExpressionFunction{
		"sqrt": func(args ...interface{}) (interface{}, error) {
			if len(args) != 1 {
				return nil, fmt.Errorf("sqrt函数需要1个参数")
			}
			if val, ok := args[0].(float64); ok {
				return math.Sqrt(val), nil
			}
			return nil, fmt.Errorf("参数必须是数字")
		},
		"abs": func(args ...interface{}) (interface{}, error) {
			if len(args) != 1 {
				return nil, fmt.Errorf("abs函数需要1个参数")
			}
			if val, ok := args[0].(float64); ok {
				return math.Abs(val), nil
			}
			return nil, fmt.Errorf("参数必须是数字")
		},
	}

	expressions := []string{
		"2 + 3 * 4",
		"(10 + 5) / 3",
		"2 ** 3",   // 幂运算
		"10 % 3",   // 取模
		"sqrt(16)", // 平方根
		"abs(-5)",  // 绝对值
	}

	for _, expr := range expressions {
		var expression *govaluate.EvaluableExpression
		var err error

		// 对于包含自定义函数的表达式，使用带函数的构造器
		if expr == "sqrt(16)" || expr == "abs(-5)" {
			expression, err = govaluate.NewEvaluableExpressionWithFunctions(expr, mathFunctions)
		} else {
			expression, err = govaluate.NewEvaluableExpression(expr)
		}

		if err != nil {
			log.Printf("解析表达式失败 '%s': %v", expr, err)
			continue
		}

		result, err := expression.Evaluate(nil)
		if err != nil {
			log.Printf("计算表达式失败 '%s': %v", expr, err)
			continue
		}

		fmt.Printf("  %s = %v\n", expr, result)
	}
}

// 变量替换示例
func variableExample() {
	expression, err := govaluate.NewEvaluableExpression("price * quantity * (1 - discount)")
	if err != nil {
		log.Printf("解析表达式失败: %v", err)
		return
	}

	// 定义变量
	parameters := map[string]interface{}{
		"price":    100.0,
		"quantity": 3,
		"discount": 0.1, // 10% 折扣
	}

	result, err := expression.Evaluate(parameters)
	if err != nil {
		log.Printf("计算失败: %v", err)
		return
	}

	fmt.Printf("  价格: %.2f, 数量: %d, 折扣: %.1f%%\n",
		parameters["price"], parameters["quantity"], parameters["discount"].(float64)*100)
	fmt.Printf("  总价: %.2f\n", result)
}

// 逻辑表达式示例
func logicalExample() {
	expressions := []string{
		"age >= 18",
		"score >= 60 && score <= 100",
		"status == 'active' || status == 'pending'",
		"!(banned == true)",
	}

	testCases := []map[string]interface{}{
		{"age": 20, "score": 85, "status": "active", "banned": false},
		{"age": 16, "score": 45, "status": "inactive", "banned": true},
	}

	for i, testCase := range testCases {
		fmt.Printf("  测试用例 %d: %v\n", i+1, testCase)

		for _, expr := range expressions {
			expression, err := govaluate.NewEvaluableExpression(expr)
			if err != nil {
				continue
			}

			result, err := expression.Evaluate(testCase)
			if err != nil {
				continue
			}

			fmt.Printf("    %s = %v\n", expr, result)
		}
		fmt.Println()
	}
}

// 字符串操作示例
func stringExample() {
	// 定义字符串函数
	stringFunctions := map[string]govaluate.ExpressionFunction{
		"len": func(args ...interface{}) (interface{}, error) {
			if len(args) != 1 {
				return nil, fmt.Errorf("len函数需要1个参数")
			}
			if str, ok := args[0].(string); ok {
				return float64(len(str)), nil
			}
			return nil, fmt.Errorf("参数必须是字符串")
		},
		"substr": func(args ...interface{}) (interface{}, error) {
			if len(args) != 3 {
				return nil, fmt.Errorf("substr函数需要3个参数")
			}
			str, ok1 := args[0].(string)
			start, ok2 := args[1].(float64)
			length, ok3 := args[2].(float64)

			if !ok1 || !ok2 || !ok3 {
				return nil, fmt.Errorf("参数类型错误")
			}

			startInt := int(start)
			lengthInt := int(length)

			if startInt < 0 || startInt >= len(str) {
				return "", nil
			}

			end := startInt + lengthInt
			if end > len(str) {
				end = len(str)
			}

			return str[startInt:end], nil
		},
	}

	expressions := []string{
		"name + ' ' + surname",
		"len(message) > 10",
		"substr(email, 0, 5)",
	}

	parameters := map[string]interface{}{
		"name":    "张",
		"surname": "三",
		"message": "Hello, World!",
		"email":   "user@example.com",
	}

	for _, expr := range expressions {
		var expression *govaluate.EvaluableExpression
		var err error

		// 对于包含自定义函数的表达式，使用带函数的构造器
		if expr == "len(message) > 10" || expr == "substr(email, 0, 5)" {
			expression, err = govaluate.NewEvaluableExpressionWithFunctions(expr, stringFunctions)
		} else {
			expression, err = govaluate.NewEvaluableExpression(expr)
		}

		if err != nil {
			log.Printf("解析表达式失败 '%s': %v", expr, err)
			continue
		}

		result, err := expression.Evaluate(parameters)
		if err != nil {
			log.Printf("计算表达式失败 '%s': %v", expr, err)
			continue
		}

		fmt.Printf("  %s = %v\n", expr, result)
	}
}

// 条件表达式示例
func conditionalExample() {
	expression, err := govaluate.NewEvaluableExpression(
		"score >= 90 ? 'A' : (score >= 80 ? 'B' : (score >= 70 ? 'C' : (score >= 60 ? 'D' : 'F')))")
	if err != nil {
		log.Printf("解析表达式失败: %v", err)
		return
	}

	scores := []int{95, 85, 75, 65, 55}

	for _, score := range scores {
		parameters := map[string]interface{}{
			"score": score,
		}

		result, err := expression.Evaluate(parameters)
		if err != nil {
			log.Printf("计算失败: %v", err)
			continue
		}

		fmt.Printf("  分数: %d, 等级: %s\n", score, result)
	}
}

// 自定义函数示例
func functionExample() {
	// 定义自定义函数
	functions := map[string]govaluate.ExpressionFunction{
		"max": func(args ...interface{}) (interface{}, error) {
			if len(args) != 2 {
				return nil, fmt.Errorf("max函数需要2个参数")
			}

			a, ok1 := args[0].(float64)
			b, ok2 := args[1].(float64)

			if !ok1 || !ok2 {
				return nil, fmt.Errorf("参数必须是数字")
			}

			if a > b {
				return a, nil
			}
			return b, nil
		},
		"min": func(args ...interface{}) (interface{}, error) {
			if len(args) != 2 {
				return nil, fmt.Errorf("min函数需要2个参数")
			}

			a, ok1 := args[0].(float64)
			b, ok2 := args[1].(float64)

			if !ok1 || !ok2 {
				return nil, fmt.Errorf("参数必须是数字")
			}

			if a < b {
				return a, nil
			}
			return b, nil
		},
	}

	expression, err := govaluate.NewEvaluableExpressionWithFunctions("max(a, b) + min(c, d)", functions)
	if err != nil {
		log.Printf("解析表达式失败: %v", err)
		return
	}

	parameters := map[string]interface{}{
		"a": 10.0,
		"b": 20.0,
		"c": 5.0,
		"d": 15.0,
	}

	result, err := expression.Evaluate(parameters)
	if err != nil {
		log.Printf("计算失败: %v", err)
		return
	}

	fmt.Printf("  max(%.1f, %.1f) + min(%.1f, %.1f) = %.1f\n",
		parameters["a"], parameters["b"], parameters["c"], parameters["d"], result)
}

// 表达式变量分析示例
func variableAnalysisExample() {
	// 测试不同复杂度的表达式
	testExpressions := []string{
		"a + b",
		"price * quantity * (1 - discount)",
		"(userAge >= 18) && (accountStatus == 'active') && !suspended",
		"basePrice + tax - couponDiscount + (weight > freeShippingLimit ? shippingCost : 0)",
		"sqrt(x * x + y * y)",
		"firstName + ' ' + lastName",
	}

	for i, exprStr := range testExpressions {
		fmt.Printf("  表达式 %d: %s\n", i+1, exprStr)

		// 解析表达式
		expression, err := govaluate.NewEvaluableExpression(exprStr)
		if err != nil {
			fmt.Printf("    解析失败: %v\n", err)
			continue
		}

		// 获取表达式中的所有变量
		variables := expression.Vars()

		fmt.Printf("    包含变量: ")
		if len(variables) == 0 {
			fmt.Printf("无变量\n")
		} else {
			for j, varName := range variables {
				if j > 0 {
					fmt.Printf(", ")
				}
				fmt.Printf("%s", varName)
			}
			fmt.Printf(" (共 %d 个)\n", len(variables))
		}

		// 演示如何根据变量动态构建参数
		fmt.Printf("    需要提供的参数: ")
		if len(variables) == 0 {
			fmt.Printf("无需参数\n")
		} else {
			for j, varName := range variables {
				if j > 0 {
					fmt.Printf(", ")
				}
				fmt.Printf("'%s': <值>", varName)
			}
			fmt.Println()
		}

		// 示例：为第一个表达式提供实际参数并计算
		if i == 0 && len(variables) > 0 {
			sampleParams := map[string]interface{}{
				"a": 10.0,
				"b": 20.0,
			}

			result, err := expression.Evaluate(sampleParams)
			if err != nil {
				fmt.Printf("    计算失败: %v\n", err)
			} else {
				fmt.Printf("    示例计算: a=10, b=20 => 结果=%v\n", result)
			}
		}

		fmt.Println()
	}

	// 实际应用场景：动态表单验证
	fmt.Println("  实际应用场景 - 动态表单验证:")
	validationRules := []string{
		"age >= 18 && age <= 120",
		"email != '' && len(email) > 5",
		"password != '' && len(password) >= 8",
		"confirmPassword == password",
		"salary > 0 && salary <= 1000000",
	}

	for _, rule := range validationRules {
		expression, err := govaluate.NewEvaluableExpression(rule)
		if err != nil {
			continue
		}

		variables := expression.Vars()
		fmt.Printf("    验证规则: %s\n", rule)
		fmt.Printf("      需要字段: %v\n", variables)
	}

	// 实际应用场景：配置驱动的业务规则
	fmt.Println("\n  实际应用场景 - 配置驱动的业务规则:")
	businessRules := map[string]string{
		"VIP客户判断": "totalSpent > 10000 && memberYears >= 2",
		"免运费条件":   "orderAmount >= 99 || memberLevel == 'gold'",
		"折扣计算":    "isNewCustomer ? 0.1 : (memberLevel == 'silver' ? 0.05 : 0)",
		"库存预警":    "currentStock <= minStock && dailySales > 0",
	}

	for ruleName, ruleExpr := range businessRules {
		expression, err := govaluate.NewEvaluableExpression(ruleExpr)
		if err != nil {
			continue
		}

		variables := expression.Vars()
		fmt.Printf("    %s: %s\n", ruleName, ruleExpr)
		fmt.Printf("      依赖数据: %v\n", variables)
	}
}

// 复杂业务场景示例
func businessExample() {
	// 电商价格计算规则
	priceExpression, err := govaluate.NewEvaluableExpression(
		"basePrice * quantity * (1 - memberDiscount) * (1 - couponDiscount) + (quantity > freeShippingThreshold ? 0 : shippingFee)")
	if err != nil {
		log.Printf("解析价格表达式失败: %v", err)
		return
	}

	// 用户权限检查规则
	permissionExpression, err := govaluate.NewEvaluableExpression(
		"(userRole == 'admin' || userRole == 'manager') && accountStatus == 'active' && !suspended")
	if err != nil {
		log.Printf("解析权限表达式失败: %v", err)
		return
	}

	// 测试数据
	orderData := map[string]interface{}{
		"basePrice":             99.99,
		"quantity":              2,
		"memberDiscount":        0.05, // 会员5%折扣
		"couponDiscount":        0.10, // 优惠券10%折扣
		"freeShippingThreshold": 3,    // 满3件免运费
		"shippingFee":           15.0,
	}

	userData := map[string]interface{}{
		"userRole":      "manager",
		"accountStatus": "active",
		"suspended":     false,
	}

	// 计算价格
	totalPrice, err := priceExpression.Evaluate(orderData)
	if err != nil {
		log.Printf("价格计算失败: %v", err)
		return
	}

	// 检查权限
	hasPermission, err := permissionExpression.Evaluate(userData)
	if err != nil {
		log.Printf("权限检查失败: %v", err)
		return
	}

	fmt.Printf("  订单信息:\n")
	fmt.Printf("    商品单价: %.2f\n", orderData["basePrice"])
	fmt.Printf("    购买数量: %d\n", orderData["quantity"])
	fmt.Printf("    会员折扣: %.1f%%\n", orderData["memberDiscount"].(float64)*100)
	fmt.Printf("    优惠券折扣: %.1f%%\n", orderData["couponDiscount"].(float64)*100)
	fmt.Printf("    运费: %.2f\n", orderData["shippingFee"])
	fmt.Printf("    总价: %.2f\n", totalPrice)

	fmt.Printf("  用户权限:\n")
	fmt.Printf("    角色: %s\n", userData["userRole"])
	fmt.Printf("    账户状态: %s\n", userData["accountStatus"])
	fmt.Printf("    是否有权限: %v\n", hasPermission)
}
