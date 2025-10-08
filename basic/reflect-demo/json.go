package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 定义一个示例结构体
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {
	// 2. 原始的JSON字符串
	jsonData := `{"name": "Alice", "age": 25}`

	// 3. 创建一个空的Person结构体实例，用于接收解析后的数据
	var p Person

	// 4. 调用json.Unmarshal进行转换
	//    函数内部会使用反射来解析jsonData并填充结构体p
	err := json.Unmarshal([]byte(jsonData), &p) // 注意：必须传递指针&p，否则无法修改原结构体
	if err != nil {
		panic(err)
	}

	// 5. 打印转换结果
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age) // 输出: Name: Alice, Age: 25

	// --- 下面部分模拟json.Unmarshal内部大致如何使用反射 ---
	fmt.Println("\n模拟反射过程:")

	// 获取结构体实例的反射值对象（必须可寻址，因此传入指针后调用Elem()）
	structValue := reflect.ValueOf(&p).Elem()
	// 获取结构体类型信息
	structType := structValue.Type()

	// 为了演示，我们再次解析JSON到一个map，模拟json包解析JSON字符串后的中间状态
	var rawData map[string]interface{}
	json.Unmarshal([]byte(jsonData), &rawData)

	// 遍历结构体的所有字段
	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)   // 获取第i个字段的reflect.Value
		fieldType := structType.Field(i) // 获取第i个字段的reflect.StructField（包含元数据）

		// 获取该字段在JSON中对应的键名（从json标签中提取，如果没有标签则使用字段名）
		jsonKey := fieldType.Tag.Get("json")
		if jsonKey == "" {
			jsonKey = fieldType.Name
		}

		// 从原始数据map中查找该键对应的值
		if valueFromJSON, exists := rawData[jsonKey]; exists {
			fmt.Printf("字段: %s (JSON键: '%s'), 原始值: %v (类型: %T)\n", fieldType.Name, jsonKey, valueFromJSON, valueFromJSON)

			// 反射设置字段值是一个复杂的过程，这里简化演示其核心：类型判断与转换
			// json.Unmarshal内部会根据结构体字段的类型，将interface{}类型的值转换为正确的类型并设置
			switch field.Kind() {
			case reflect.String:
				// 如果JSON中的值是float64（JSON数字的默认Go类型），而我们需要string，需要转换
				if str, ok := valueFromJSON.(string); ok {
					field.SetString(str)
				}
			case reflect.Int, reflect.Int64, reflect.Int32:
				// JSON数字默认被解析为float64，需要转换为int
				if num, ok := valueFromJSON.(float64); ok {
					field.SetInt(int64(num)) // 使用SetInt方法设置整型值
				}
			// ... 处理其他各种类型（Bool, Slice, Struct等）的逻辑会更加复杂
			}
		}
	}

	fmt.Printf("通过模拟反射设置后: Name: %s, Age: %d\n", p.Name, p.Age)
}